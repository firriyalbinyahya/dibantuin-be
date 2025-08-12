package repository

import (
	"dibantuin-be/entity"

	"gorm.io/gorm"
)

type VerificationProgramRepository struct {
	DB *gorm.DB
}

func NewVerificationProgramRepository(db *gorm.DB) *VerificationProgramRepository {
	return &VerificationProgramRepository{DB: db}
}

func (vpr *VerificationProgramRepository) CreateVerificationProgram(verification *entity.VerificationProgram) error {
	return vpr.DB.Create(verification).Error
}
