package repository

import (
	"dibantuin-be/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ar *UserRepository) GetByEmail(email string) (data *entity.User, err error) {
	err = ar.DB.Where("email = ?", email).First(&data).Error
	return
}

func (ar *UserRepository) Create(user *entity.User) error {
	return ar.DB.Create(user).Error
}

func (ar *UserRepository) CountAdmins() (int64, error) {
	var count int64
	if err := ar.DB.Model(&entity.User{}).
		Where("role = ?", "admin").
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (ar *UserRepository) GetByID(userID uint64) (*entity.User, error) {
	var user entity.User
	if err := ar.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ar *UserRepository) Update(user *entity.User, req *entity.UserUpdate) error {
	if err := ar.DB.Model(user).Updates(req).Error; err != nil {
		return err
	}
	return nil
}

func (ar *UserRepository) Delete(userID uint64) error {
	if err := ar.DB.Delete(&entity.User{}, userID).Error; err != nil {
		return err
	}
	return nil
}
