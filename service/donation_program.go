package service

import (
	"dibantuin-be/entity"
	"dibantuin-be/repository"
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type DonationProgramService struct {
	DB                            *gorm.DB
	DonationProgramRepository     *repository.DonationProgramRepository
	VerificationProgramRepository *repository.VerificationProgramRepository
	UserLogService                *UserLogService
}

func NewDonationProgramService(db *gorm.DB, donationProgramRepository *repository.DonationProgramRepository,
	verificationProgramRepository *repository.VerificationProgramRepository, userLogService *UserLogService) *DonationProgramService {
	return &DonationProgramService{
		DB:                            db,
		DonationProgramRepository:     donationProgramRepository,
		VerificationProgramRepository: verificationProgramRepository,
		UserLogService:                userLogService,
	}
}

func (dps *DonationProgramService) CreateRequest(req *entity.DonationProgramRequestCreate, userID uint64) error {
	if req.TargetAmount <= 0 {
		return errors.New("target amount must be greater than zero")
	}

	if time.Now().After(req.EndDate) {
		return errors.New("end date must be in the future")
	}

	return dps.DB.Transaction(func(tx *gorm.DB) error {
		donationRepo := repository.NewDonationProgramRepository(tx)

		newDonationProgram := &entity.DonationProgram{
			UserID:        userID,
			CategoryID:    req.CategoryID,
			Creator:       req.Creator,
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

		if err := donationRepo.CreateDonationProgram(newDonationProgram); err != nil {
			return err
		}

		newDonationProgramRequest := &entity.DonationProgramRequest{
			UserID:      userID,
			CategoryID:  req.CategoryID,
			KTPPhoto:    req.KTPPhoto,
			SelfiePhoto: req.SelfiePhoto,
			LegalDoc:    req.LegalDoc,
			AdminNotes:  req.AdminNotes,
			ProgramID:   &newDonationProgram.ID,
		}

		if err := donationRepo.CreateDonationProgramRequest(newDonationProgramRequest); err != nil {
			return err
		}

		// log user action
		desc := "request to create donation program"
		if err := dps.UserLogService.LogUserAction(userID, "REQUEST_PROGRAM", "donation_program_requests", newDonationProgramRequest.ID, desc); err != nil {
			log.Printf("Failed to create user log: %v", err)
		}

		return nil
	})
}

func (dps *DonationProgramService) VerifyProgram(donationProgramRequestID uint64, adminID uint64,
	status string, note string) error {
	_, err := dps.DonationProgramRepository.GetDonationProgramRequestById(donationProgramRequestID)
	if err != nil {
		return err
	}

	err = dps.DonationProgramRepository.UpdateStatusDonationProgramRequestById(donationProgramRequestID, status)
	if err != nil {
		return err
	}

	verificationProgram := &entity.VerificationProgram{
		ProgramRequestID: donationProgramRequestID,
		VerifiedBy:       adminID,
		Note:             note,
		Status:           status,
	}

	err = dps.VerificationProgramRepository.CreateVerificationProgram(verificationProgram)
	if err != nil {
		return err
	}

	// log user action
	desc := fmt.Sprintf("Admin Verify Donation Program, Status Program Become %s", status)
	if err := dps.UserLogService.LogUserAction(adminID, "VERIFY_PROGRAM", "donation_program_requests", donationProgramRequestID, desc); err != nil {
		log.Printf("Failed to create user log: %v", err)
	}

	return nil
}

func (dps *DonationProgramService) ListDonationPrograms(statusRequest, search string, limit, page int, categoryID uint64) (*[]entity.DonationProgramListItem, int64, error) {
	return dps.DonationProgramRepository.ListDonationPrograms(statusRequest, search, limit, page, categoryID)
}
