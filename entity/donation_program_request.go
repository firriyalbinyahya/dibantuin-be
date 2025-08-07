package entity

import "time"

type DonationProgramRequest struct {
	ID           uint64    `gorm:"primaryKey"`
	UserID       uint64    `gorm:"not null"`
	CategoryID   uint64    `gorm:"not null"`
	Creator      string    `gorm:"type:varchar(100);not null"`
	Title        string    `gorm:"type:varchar(255);not null"`
	Description  string    `gorm:"type:text;not null"`
	IsPersonal   bool      `gorm:"not null"`
	Status       string    `gorm:"type:enum('pending', 'approved', 'rejected');not null"`
	EndDate      time.Time `gorm:"not null"`
	CoverImage   string    `gorm:"type:varchar(255);not null"`
	TargetAmount float64   `gorm:"type:decimal(15,2);not null"`
	RekeningInfo string    `gorm:"type:varchar(255);not null"`
	KTPPhoto     string    `gorm:"type:varchar(255)"`
	SelfiePhoto  string    `gorm:"type:varchar(255)"`
	LegalDoc     string    `gorm:"type:varchar(255)"`
	AdminNotes   string    `gorm:"type:text"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
