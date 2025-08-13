package repository

import (
	"dibantuin-be/entity"

	"gorm.io/gorm"
)

type TransactionVerificationRepository struct {
	DB *gorm.DB
}

func NewTransactionVerificationRepository(db *gorm.DB) *TransactionVerificationRepository {
	return &TransactionVerificationRepository{DB: db}
}

func (tvr *TransactionVerificationRepository) CreateTransactionVerification(verification *entity.VerificationTransactionDonation) error {
	return tvr.DB.Create(verification).Error
}
