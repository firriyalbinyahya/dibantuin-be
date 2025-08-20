package service

import (
	"dibantuin-be/entity"
	"dibantuin-be/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{UserRepository: repository}
}

func (us *UserService) GetUserByID(userID uint64) (*entity.UserDetailResponse, error) {
	user, err := us.UserRepository.GetByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	response := &entity.UserDetailResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}

	return response, nil
}

func (us *UserService) UpdateUser(userID uint64, req *entity.UserUpdate) error {
	user, err := us.UserRepository.GetByID(userID)
	if err != nil {
		return errors.New("user not found")
	}

	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		req.Password = string(hashedPassword)
	}

	err = us.UserRepository.Update(user, req)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) DeleteUser(actingUserID uint64, targetUserID uint64, actingUserRole string) error {
	if actingUserID != targetUserID && actingUserRole != "admin" {
		return errors.New("unauthorized to delete this user")
	}

	err := us.UserRepository.Delete(targetUserID)
	if err != nil {
		return errors.New("user not found")
	}
	return nil
}
