package entity

import "time"

type VerificationProgram struct {
	ID               uint64 `gorm:"primaryKey"`
	ProgramRequestID uint64 `gorm:"not null"`
	VerifiedBy       uint64 `gorm:"not null"`
	Note             string `gorm:"type:text"`
	Status           string `gorm:"type:enum('pending', 'approved', 'rejected');not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type VerificationProgramRequest struct {
	Note   string `json:"note" binding:"omitempty"`
	Status string `json:"status" binding:"required,oneof=pending approved rejected"`
}
