package service

import (
	"dibantuin-be/entity"
	"dibantuin-be/repository"
	"errors"
	"fmt"
	"log"
	"math"
)

type DonationTransactionService struct {
	DonationTransactionRepository     *repository.DonationTransactionRepository
	TransactionVerificationRepository *repository.TransactionVerificationRepository
	UserLogService                    *UserLogService
}

func NewDonationTransactionService(donationTransactionRepository *repository.DonationTransactionRepository,
	transactionVerification *repository.TransactionVerificationRepository,
	userLogService *UserLogService) *DonationTransactionService {
	return &DonationTransactionService{DonationTransactionRepository: donationTransactionRepository,
		TransactionVerificationRepository: transactionVerification,
		UserLogService:                    userLogService}
}

func (dts *DonationTransactionService) CreateMoneyDonationTransaction(req *entity.MoneyTransactionDonationRequest, userID uint64) (*entity.MoneyTransactionDonation, error) {
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
	_, err := dts.DonationTransactionRepository.GetDonationTransactionById(donationTransactionID)
	if err != nil {
		return errors.New("transaction was not found")
	}

	err = dts.DonationTransactionRepository.UpdateStatusDonationTransactionById(donationTransactionID, status)
	if err != nil {
		return errors.New("failed to update donation status")
	}

	verificationTransaction := &entity.VerificationTransactionDonation{
		TransactionDonationID: donationTransactionID,
		VerifiedBy:            adminID,
		Note:                  note,
		Status:                status,
	}

	err = dts.TransactionVerificationRepository.CreateTransactionVerification(verificationTransaction)
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
	if err := dts.UserLogService.LogUserAction(adminID, log_status, "verification_transaction_donations", verificationTransaction.ID, desc); err != nil {
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
