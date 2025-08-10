package service

import (
	"dibantuin-be/config/redis"
	"dibantuin-be/entity"
	"dibantuin-be/repository"
	"dibantuin-be/utils/auth"
	"dibantuin-be/utils/response"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repository *repository.UserRepository
}

func NewAuthService(repository *repository.UserRepository) *AuthService {
	return &AuthService{Repository: repository}
}

func (as *AuthService) Register(req *entity.Register) error {
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
		Role:     "user",
	}

	return as.Repository.Create(newUser)
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

	return nil
}
