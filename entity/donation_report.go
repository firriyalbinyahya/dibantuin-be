package entity

import "time"

type DonationReport struct {
	ID          uint64 `gorm:"primaryKey"`
	UserID      uint64 `gorm:"not null"`
	ProgramID   uint64 `gorm:"not null"`
	Title       string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:text;not null"`
	ReportPhoto string `gorm:"type:varchar(255)"`
	CreatedAt   time.Time
}

type DonationReportRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ReportPhoto string `json:"report_photo"`
}
