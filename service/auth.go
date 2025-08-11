package service

import (
	"dibantuin-be/config/redis"
	"dibantuin-be/entity"
	"dibantuin-be/repository"
	"dibantuin-be/utils/auth"
	"dibantuin-be/utils/response"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repository     *repository.UserRepository
	UserLogService *UserLogService
}

func NewAuthService(repository *repository.UserRepository, userLogService *UserLogService) *AuthService {
	return &AuthService{Repository: repository, UserLogService: userLogService}
}

func (as *AuthService) CreateUser(req *entity.Register, role string, actorUserID uint64) error {
	// mengecek apakah email sudah ada
	_, err := as.Repository.GetByEmail(req.Email)
	if err == nil {
		return errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     role,
	}

	err = as.Repository.Create(newUser)
	if err != nil {
		return err
	}

	if role == "user" {
		actorUserID = newUser.ID
	}

	// Log sukses menambah user
	desc := fmt.Sprintf("%s user registered with email %s", role, newUser.Email)
	err = as.UserLogService.LogUserAction(actorUserID, "CREATE_USER", "users", newUser.ID, desc)
	if err != nil {
		log.Printf("Failed to create user log: %v", err)
	}

	return nil
}

func (as *AuthService) Register(req *entity.Register) error {
	return as.CreateUser(req, "user", 0)
}

func (as *AuthService) CreateaAdmin(req *entity.Register, fromAPIKey bool, role string, idAdmin uint64) error {
	if fromAPIKey {
		count, err := as.Repository.CountAdmins()
		if err != nil {
			return err
		}
		if count > 0 {
			return errors.New("you have to ask another admin to create an admin account")
		}
	} else {
		if role != "admin" {
			return errors.New("only admins can create another admin")
		}
	}

	return as.CreateUser(req, "admin", idAdmin)
}

func (as *AuthService) Login(req *entity.Login) (*entity.UserLoginResponse, error) {
	user, err := as.Repository.GetByEmail(req.Email)
	if err != nil {
		return &entity.UserLoginResponse{}, response.NewCustomError(http.StatusUnauthorized, "Email or password information does not match", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Printf("Error compare password: %v", err)
		return &entity.UserLoginResponse{}, response.NewCustomError(http.StatusUnauthorized, "Email or password information does not match", err)
	}

	// generate token
	accessToken, accessExp, err := auth.GenerateToken(user, false)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshExp, err := auth.GenerateToken(user, true)
	if err != nil {
		return nil, err
	}

	// menyimpan refresh token di redis
	err = redis.RedisClient.Set(redis.Ctx,
		"refresh_token:"+string(rune(user.ID)),
		refreshToken,
		time.Until(*refreshExp)).Err()
	if err != nil {
		return nil, err
	}

	//user log login
	desc := fmt.Sprintf("%s login with email %s", user.Name, user.Email)
	err = as.UserLogService.LogUserAction(user.ID, "LOGIN", "users", user.ID, desc)
	if err != nil {
		log.Printf("Failed to create user log: %v", err)
	}

	return &entity.UserLoginResponse{
		Name:           user.Name,
		Email:          user.Email,
		Role:           user.Role,
		AccessToken:    accessToken,
		AccessExpired:  accessExp,
		RefreshToken:   refreshToken,
		RefreshExpired: refreshExp,
	}, nil
}

func (as *AuthService) RefreshToken(refreshToken string) (*entity.UserLoginResponse, error) {
	user, err := auth.VerifyToken(refreshToken, true)
	if err != nil {
		log.Printf("Gagal memverifikasi refresh token: %v", err)
		return nil, response.NewCustomError(http.StatusInternalServerError, "Gagal memverifikasi token", nil)
	}

	// mengecek token di redis
	savedToken, err := redis.RedisClient.Get(redis.Ctx, "refresh_token:"+string(rune(user.ID))).Result()
	if err != nil || savedToken != refreshToken {
		if err != nil {
			log.Printf("Gagal memverifikasi refresh token: %v", err)
			return nil, response.NewCustomError(http.StatusInternalServerError, "Refresh token invalid", nil)
		}
	}

	newAccessToken, accessExp, err := auth.GenerateToken(user, false)
	if err != nil {
		return nil, err
	}

	return &entity.UserLoginResponse{
		AccessToken:    newAccessToken,
		AccessExpired:  accessExp,
		RefreshToken:   refreshToken,
		RefreshExpired: nil,
	}, nil
}

func (as *AuthService) Logout(userID uint64) error {
	userIDStr := strconv.FormatUint(userID, 10)

	// menghapus refresh token di redis
	err := redis.RedisClient.Del(redis.Ctx, "refresh_token:"+userIDStr).Err()
	if err != nil {
		return err
	}

	//user log logout
	desc := fmt.Sprintf("user id : %s logout", userIDStr)
	err = as.UserLogService.LogUserAction(userID, "LOGOUT", "users", userID, desc)
	if err != nil {
		log.Printf("Failed to create user log: %v", err)
	}

	return nil
}
