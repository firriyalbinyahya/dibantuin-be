package service

import (
	"dibantuin-be/entity"
	"dibantuin-be/repository"
	"errors"
	"fmt"
	"log"
	"math"
	"time"

	"gorm.io/gorm"
)

type DonationTransactionService struct {
	DB                                *gorm.DB
	DonationTransactionRepository     *repository.DonationTransactionRepository
	TransactionVerificationRepository *repository.TransactionVerificationRepository
	UserLogService                    *UserLogService
}

func NewDonationTransactionService(db *gorm.DB, donationTransactionRepository *repository.DonationTransactionRepository,
	transactionVerification *repository.TransactionVerificationRepository,
	userLogService *UserLogService) *DonationTransactionService {
	return &DonationTransactionService{DB: db, DonationTransactionRepository: donationTransactionRepository,
		TransactionVerificationRepository: transactionVerification,
		UserLogService:                    userLogService}
}

func (dts *DonationTransactionService) CreateMoneyDonationTransaction(req *entity.MoneyTransactionDonationRequest, userID uint64) (*entity.MoneyTransactionDonation, error) {

	if req.Amount < 1000 {
		return nil, errors.New("minimum donation is 1000")
	}

	if req.DonorsName == "" {
		req.DonorsName = "Anonymous"
	}

	if len(req.DonationMessage) > 500 {
		return nil, errors.New("donation message must be at most 500 characters")
	}

	donationProgramRepo := repository.NewDonationProgramRepository(dts.DB)
	program, err := donationProgramRepo.GetDonationProgramById(req.ProgramID)
	if err != nil {
		return nil, errors.New("failed to get donation program id")
	}

	if time.Now().After(program.EndDate) {
		return nil, errors.New("donation program has ended")
	}

	programRequest, err := donationProgramRepo.GetProgramRequestByProgramID(program.ID)
	if err != nil {
		return nil, errors.New("failed to get donation program request by program id")
	}

	if programRequest.StatusRequest != "approved" {
		return nil, errors.New("failed! Donation program status not approved yet")
	}

	newMoneyDonationTransaction := &entity.MoneyTransactionDonation{
		ProgramID:       req.ProgramID,
		UserID:          userID,
		Amount:          req.Amount,
		DonationStatus:  "pending",
		DonorsName:      req.DonorsName,
		DonationMessage: req.DonationMessage,
		DonationPhoto:   req.DonationPhoto,
	}

	transaction, err := dts.DonationTransactionRepository.CreateMoneyDonationTransaction(newMoneyDonationTransaction)
	if err != nil {
		return nil, err
	}

	// log user action
	desc := fmt.Sprintf("Success donate transaction as %s", req.DonorsName)
	if err := dts.UserLogService.LogUserAction(userID, "DONATE_TRANSACTION", "money_transaction_donations", newMoneyDonationTransaction.ID, desc); err != nil {
		log.Printf("Failed to create user log: %v", err)
	}

	return transaction, nil
}

func (dts *DonationTransactionService) VerifyDonationTransaction(donationTransactionID uint64, adminID uint64,
	status string, note string) error {
	transaction, err := dts.DonationTransactionRepository.GetDonationTransactionById(donationTransactionID)
	if err != nil {
		return errors.New("transaction was not found")
	}
	var verificationTransactionID uint64

	err = dts.DB.Transaction(func(tx *gorm.DB) error {
		donationTransactionRepo := repository.NewDonationTransactionRepository(tx)
		transactionVerificationRepo := repository.NewTransactionVerificationRepository(tx)
		err := donationTransactionRepo.UpdateStatusDonationTransactionById(donationTransactionID, status)
		if err != nil {
			return errors.New("failed to update donation status")
		}

		if status == "success" {
			err = donationTransactionRepo.IncreaseCurrentAmount(transaction.ProgramID, transaction.Amount)
			if err != nil {
				return errors.New("failed to update program current amount")
			}
		}

		verificationTransaction := &entity.VerificationTransactionDonation{
			TransactionDonationID: donationTransactionID,
			VerifiedBy:            adminID,
			Note:                  note,
			Status:                status,
		}

		err = transactionVerificationRepo.CreateTransactionVerification(verificationTransaction)
		if err != nil {
			return err
		}

		verificationTransactionID = verificationTransaction.ID

		return nil
	})

	if err != nil {
		return err
	}

	var log_status string
	if status == "success" {
		log_status = "SUBMIT_VERIFICATION"
	} else {
		log_status = "REJECT_VERIFICATION"
	}
	// log user action
	desc := fmt.Sprintf("Admin Verify Donation Transaction, Status Transaction Become %s", status)
	if err := dts.UserLogService.LogUserAction(adminID, log_status, "verification_transaction_donations", verificationTransactionID, desc); err != nil {
		log.Printf("Failed to create user log: %v", err)
	}

	return nil
}

func (dts *DonationTransactionService) ListDonationTransactions(userID *uint64, donationStatus, search string, limit, page int) (*entity.PaginatedDonationTransactions, error) {
	items, totalItems, err := dts.DonationTransactionRepository.ListDonationTransactions(userID, donationStatus, search, limit, page)
	if err != nil {
		return nil, errors.New("failed to get list of transactions")
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	return &entity.PaginatedDonationTransactions{
		Items:       *items,
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: page,
	}, nil
}
