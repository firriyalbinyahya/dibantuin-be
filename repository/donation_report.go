package repository

import (
	"dibantuin-be/entity"

	"gorm.io/gorm"
)

type DonationReportRepository struct {
	DB *gorm.DB
}

func NewDonationReportRepository(db *gorm.DB) *DonationReportRepository {
	return &DonationReportRepository{DB: db}
}

func (drr *DonationReportRepository) Create(report *entity.DonationReport) error {
	return drr.DB.Create(&report).Error
}

func (drr *DonationReportRepository) GetDonationReportByID(id uint) (*entity.DonationReport, error) {
	var report entity.DonationReport
	err := drr.DB.First(&report, id).Error
	return &report, err
}

func (drr *DonationReportRepository) Update(report *entity.DonationReport) (*entity.DonationReport, error) {
	err := drr.DB.Save(&report).Error
	return report, err
}

func (drr *DonationReportRepository) Delete(id uint) error {
	return drr.DB.Delete(&entity.DonationReport{}, id).Error
}

func (drr *DonationReportRepository) GetDonationReports(programID uint64, limit, page int) (*[]entity.DonationReport, int64, error) {
	var reports []entity.DonationReport
	var total int64
	offset := (page - 1) * limit

	if err := drr.DB.Model(&entity.DonationReport{}).Where("program_id = ?", programID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := drr.DB.Where("program_id = ?", programID).Order("created_at DESC").Limit(limit).Offset(offset).Find(&reports).Error; err != nil {
		return nil, 0, err
	}

	return &reports, total, nil
}
