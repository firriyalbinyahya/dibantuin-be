package entity

import "time"

type MoneyTransactionDonation struct {
	ID              uint64  `gorm:"primaryKey"`
	ProgramID       uint64  `gorm:"not null"`
	UserID          uint64  `gorm:"not null"`
	Amount          float64 `gorm:"type:decimal(15,2);not null"`
	DonationStatus  string  `gorm:"type:enum('pending', 'success', 'failed');not null"`
	DonorName       string  `gorm:"type:varchar(100);not null"`
	DonationMessage string  `gorm:"type:text"`
	DonationPhoto   string  `gorm:"type:varchar(255);not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
