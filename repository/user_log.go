package repository

import (
	"dibantuin-be/entity"

	"gorm.io/gorm"
)

type UserLogRepository struct {
	DB *gorm.DB
}

func NewUserLogRepository(db *gorm.DB) *UserLogRepository {
	return &UserLogRepository{DB: db}
}

func (ulr *UserLogRepository) Create(log *entity.UserLog) error {
	return ulr.DB.Create(log).Error
}
