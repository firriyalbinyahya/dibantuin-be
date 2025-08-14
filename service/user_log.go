package service

import (
	"dibantuin-be/entity"
	"dibantuin-be/repository"
	"errors"
	"math"
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

func (uls *UserLogService) GetUserLogs(userID uint64, limit, page int) (*entity.PaginatedUserLogs, error) {
	logs, totalItems, err := uls.Repository.UserLogs(userID, limit, page)
	if err != nil {
		return nil, errors.New("failed to get data of user logs")
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))
	return &entity.PaginatedUserLogs{
		Items:       *logs,
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: page,
	}, nil
}
