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

func (dps *DonationProgramService) UpdateProgram(programID uint64, req *entity.DonationProgramUpdateRequest, userID uint64) error {
	program, programRequest, err := dps.DonationProgramRepository.FindProgramAndRequest(programID)
	if err != nil {
		return err
	}

	if program.UserID != userID {
		return errors.New("unauthorized access")
	}

	if req.TargetAmount > 0 {
		if req.TargetAmount < program.CurrentAmount {
			return errors.New("target amount cannot be less than current amount")
		}
		program.TargetAmount = req.TargetAmount
	}

	if !req.EndDate.IsZero() {
		if time.Now().After(req.EndDate) {
			return errors.New("end date must be in the future")
		}
		if req.EndDate.Before(program.StartDate) {
			return errors.New("end date must be after start date")
		}
		program.EndDate = req.EndDate
	}

	return dps.DB.Transaction(func(tx *gorm.DB) error {
		donationRepo := repository.NewDonationProgramRepository(tx)

		if programRequest.StatusRequest == "approved" {
			programRequest.StatusRequest = "pending"

			if err := donationRepo.UpdateDonationProgramRequest(programRequest); err != nil {
				return err
			}

			// Log aksi
			dps.UserLogService.LogUserAction(userID, "UPDATE_PROGRAM", "donation_program_requests", programRequest.ID, "request to update approved donation program")
		} else {
			dps.UserLogService.LogUserAction(userID, "UPDATE_PROGRAM", "donation_programs", program.ID, "update pending/rejected donation program")
		}

		if err := donationRepo.UpdateDonationProgramFromStruct(program.ID, req); err != nil {
			return err
		}

		if err := donationRepo.UpdateDonationProgramRequestFromStruct(programRequest.ID, req); err != nil {
			return err
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

func (dps *DonationProgramService) GetDonationProgramDetail(id uint64) (*entity.DonationProgram, error) {
	return dps.DonationProgramRepository.GetDonationProgramById(id)
}

func (dps *DonationProgramService) GetDonationProgramDetailWithoutRequest(id uint64) (*entity.DonationProgram, error) {
	return dps.DonationProgramRepository.GetDonationProgramWithoutRequestById(id)
}

func (dps *DonationProgramService) DeleteProgram(programID uint64, userID uint64, userRole string) error {
	program, err := dps.DonationProgramRepository.GetDonationProgramWithoutRequestById(programID)
	if err != nil {
		return errors.New("program not found")
	}
	if program.UserID != userID && userRole != "admin" {
		return errors.New("unauthorized access")
	}

	if err := dps.DonationProgramRepository.DeleteProgram(programID); err != nil {
		return err
	}

	//Log aksi
	dps.UserLogService.LogUserAction(userID, "DELETE_PROGRAM", "donation_programs", programID, "deleted donation program")

	return nil
}
