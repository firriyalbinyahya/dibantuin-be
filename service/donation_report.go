package service

import (
	"dibantuin-be/entity"
	"dibantuin-be/repository"
	"errors"
)

type DonationReportService struct {
	DonationReportRepository *repository.DonationReportRepository
}

func NewDonationReportService(repository *repository.DonationReportRepository) *DonationReportService {
	return &DonationReportService{DonationReportRepository: repository}
}

func (drs *DonationReportService) CreateDonationReport(userID uint64, programID uint64, request *entity.DonationReportRequest) (*entity.DonationReport, error) {

	report := &entity.DonationReport{
		UserID:      userID,
		ProgramID:   programID,
		Title:       request.Title,
		Description: request.Description,
		ReportPhoto: request.ReportPhoto,
	}

	if err := drs.DonationReportRepository.Create(report); err != nil {
		return nil, errors.New("failed to create donation report")
	}

	return report, nil

}
