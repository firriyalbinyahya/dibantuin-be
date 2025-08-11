package service

import (
	"dibantuin-be/entity"
	"dibantuin-be/repository"
	"errors"
	"time"
)

type DonationProgramService struct {
	Repository *repository.DonationProgramRepository
}

func NewDonationProgramService(repository *repository.DonationProgramRepository) *DonationProgramService {
	return &DonationProgramService{Repository: repository}
}

func (dps *DonationProgramService) CreateRequest(req *entity.DonationProgramRequestCreate, userID uint64) (*entity.DonationProgramRequest, error) {
	if req.TargetAmount <= 0 {
		return nil, errors.New("target amount must be greater than zero")
	}

	if time.Now().After(req.EndDate) {
		return nil, errors.New("end date must be in the future")
	}

	newDonationProgram := &entity.DonationProgram{
		UserID:        userID,
		CategoryID:    req.CategoryID,
		Title:         req.Title,
		Description:   req.Description,
		TargetAmount:  req.TargetAmount,
		CurrentAmount: 0,
		RekeningInfo:  req.RekeningInfo,
		IsPersonal:    req.IsPersonal,
		Status:        "soon",
		StartDate:     req.StartDate,
		EndDate:       req.EndDate,
		CoverImage:    req.CoverImage,
		ContactInfo:   req.ContactInfo,
	}

	newDonationProgramRequest := &entity.DonationProgramRequest{
		UserID:      userID,
		CategoryID:  req.CategoryID,
		KTPPhoto:    req.KTPPhoto,
		SelfiePhoto: req.SelfiePhoto,
		LegalDoc:    req.LegalDoc,
		AdminNotes:  req.AdminNotes,
	}

	err = dps.Repository.Create()
	if err != nil {
		return err
	}
}
