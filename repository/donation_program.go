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

func (dpr *DonationProgramRepository) ListDonationPrograms(statusRequest, search string, limit, page int, categoryID uint64) (*[]entity.DonationProgramListItem, int64, error) {
	var programs []entity.DonationProgramListItem
	var total int64

	query := dpr.DB.Table("donation_programs dp").
		Select(`dp.id, dp.category_id, dp.title, dp.description, dp.target_amount, dp.current_amount,
	dp.creator, dp.start_date, dp.end_date, dp.cover_image, dp.status,
	dpr.status_request`).
		Joins("LEFT JOIN donation_program_requests dpr ON dp.id = dpr.program_id")

	if statusRequest != "" {
		query = query.Where("dpr.status_request = ?", statusRequest)
	}

	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("dp.title LIKE ? OR dp.description LIKE ?", searchPattern, searchPattern)
	}

	if categoryID != 0 {
		query = query.Where("dp.category_id = ?", categoryID)
	}

	query.Count(&total)

	offset := (page - 1) * limit
	err := query.Limit(limit).Offset(offset).Find(&programs).Error
	if err != nil {
		return nil, 0, err
	}

	return &programs, total, nil
}
