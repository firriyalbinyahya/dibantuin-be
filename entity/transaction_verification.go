package entity

import "time"

type VerificationTransactionDonation struct {
	ID                    uint64 `gorm:"primaryKey"`
	TransactionDonationID uint64 `gorm:"not null"`
	VerifiedBy            uint64 `gorm:"not null"`
	Note                  string `gorm:"type:text"`
	Status                string `gorm:"type:enum('pending', 'success', 'failed');not null"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

type VerificationTransactionRequest struct {
	Note   string `json:"note" binding:"omitempty"`
	Status string `json:"status" binding:"required,oneof=pending success failed"`
}
