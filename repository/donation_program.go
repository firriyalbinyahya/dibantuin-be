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

func (dpr *DonationProgramRepository) CreateDonationProgram(donationProgram *entity.DonationProgram) error {
	return dpr.DB.Create(donationProgram).Error
}

func (dpr *DonationProgramRepository) CreateDonationProgramRequest(donationProgramRequest *entity.DonationProgramRequest) error {
	return dpr.DB.Create(donationProgramRequest).Error
}

func (dpr *DonationProgramRepository) GetDonationProgramRequestById(id uint64) (*entity.DonationProgramRequest, error) {
	var request entity.DonationProgramRequest
	err := dpr.DB.Where("id = ?", id).First(&request).Error
	if err != nil {
		return nil, err
	}

	return &request, nil
}

func (dpr *DonationProgramRepository) UpdateStatusDonationProgramRequestById(id uint64, status string) error {
	return dpr.DB.Model(&entity.DonationProgramRequest{}).Where("id = ?", id).Update("status_request", status).Error
}
