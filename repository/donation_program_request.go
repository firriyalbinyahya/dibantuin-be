package repository

import (
	"dibantuin-be/entity"

	"gorm.io/gorm"
)

type DonationProgramRequestRepository struct {
	DB *gorm.DB
}

func NewDonationProgramRequestRepository(db *gorm.DB) *DonationProgramRequestRepository {
	return &DonationProgramRequestRepository{DB: db}
}

func (dpr *DonationProgramRequestRepository) Create(donationProgramRequest *entity.DonationProgramRequest) error {
	return dpr.DB.Create(donationProgramRequest).Error
}
