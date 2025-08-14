package entity

import "time"

type MoneyTransactionDonation struct {
	ID              uint64  `gorm:"primaryKey"`
	ProgramID       uint64  `gorm:"not null"`
	UserID          uint64  `gorm:"not null"`
	Amount          float64 `gorm:"type:decimal(15,2);not null"`
	DonationStatus  string  `gorm:"type:enum('pending', 'success', 'failed');not null"`
	DonorsName      string  `gorm:"type:varchar(100);not null"`
	DonationMessage string  `gorm:"type:text"`
	DonationPhoto   string  `gorm:"type:varchar(255);not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type MoneyTransactionDonationRequest struct {
	ProgramID       uint64  `json:"program_id" binding:"required"`
	Amount          float64 `json:"amount" binding:"required"`
	DonorsName      string  `json:"donors_name,omitempty"`
	DonationMessage string  `json:"donation_message,omitempty"`
	DonationPhoto   string  `json:"donation_photo" binding:"required"`
}

type DonationTransactionListItem struct {
	ID              uint64  `json:"id"`
	Title           string  `json:"title"`
	Creator         string  `json:"creator"`
	Name            string  `json:"name"`
	Amount          float64 `json:"amount"`
	DonationStatus  string  `json:"donation_status"`
	DonorsName      string  `json:"donors_name"`
	DonationMessage string  `json:"donation_message"`
}

type PaginatedDonationTransactions struct {
	Items       []DonationTransactionListItem `json:"items"`
	TotalItems  int64                         `json:"total_items"`
	TotalPages  int                           `json:"total_pages"`
	CurrentPage int                           `json:"current_page"`
}
