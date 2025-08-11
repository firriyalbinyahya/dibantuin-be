package service

import (
	"dibantuin-be/entity"
	"dibantuin-be/repository"
)

type UserLogService struct {
	Repository *repository.UserLogRepository
}

func NewUserLogService(repository *repository.UserLogRepository) *UserLogService {
	return &UserLogService{Repository: repository}
}

func (uls *UserLogService) LogUserAction(userID uint64, actionType, targetTable string, targetID uint64, description string) error {
	log := &entity.UserLog{
		UserID:      userID,
		ActionType:  actionType,
		TargetTable: targetTable,
		TargetID:    targetID,
		Description: description,
	}

	return uls.Repository.Create(log)
}
