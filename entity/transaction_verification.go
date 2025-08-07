package entity

import "time"

type VerificationTransactionDonation struct {
	ID                    uint64 `gorm:"primaryKey"`
	TransactionDonationID uint64 `gorm:"not null"`
	VerifiedBy            uint64 `gorm:"not null"`
	Note                  string `gorm:"type:text"`
	Status                string `gorm:"type:enum('pending', 'verified', 'rejected');not null"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
