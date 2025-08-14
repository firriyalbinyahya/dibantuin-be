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

func (ulr *UserLogRepository) UserLogs(userID uint64, limit, page int) (*[]entity.UserLog, int64, error) {
	var logs []entity.UserLog
	var total int64
	offset := (page - 1) * limit

	if err := ulr.DB.Model(&entity.UserLog{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := ulr.DB.Where("user_id = ?", userID).Order("created_at DESC").Limit(limit).Offset(offset).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	return &logs, total, nil
}
