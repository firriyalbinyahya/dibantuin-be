package repository

import (
	"dibantuin-be/entity"

	"gorm.io/gorm"
)

type DonationProgramRepository struct {
	DB *gorm.DB
}

func NewDonationProgramRepository(db *gorm.DB) *DonationProgramRepository {
	return &DonationProgramRepository{DB: db}
}

func (dpr *DonationProgramRepository) Create(donationProgram *entity.DonationProgram) error {
	return dpr.DB.Create(donationProgram).Error
}
