package service

import (
	"dibantuin-be/entity"
	"dibantuin-be/repository"
	"errors"
	"math"
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

func (drs *DonationReportService) UpdateDonationReport(id uint, updated *entity.DonationReportRequest) (*entity.DonationReport, error) {
	report, err := drs.DonationReportRepository.GetDonationReportByID(id)
	if err != nil {
		return nil, errors.New("report id was not found")
	}

	report.Title = updated.Title
	report.ReportPhoto = updated.ReportPhoto
	report.Description = updated.Description

	err = drs.DonationReportRepository.Update(report)
	if err != nil {
		return nil, errors.New("failed to update donation report")
	}

	return report, nil
}

func (drs *DonationReportService) GetDonationReports(programID uint64, limit, page int) (*entity.PaginatedDonationReport, error) {
	reports, totalItems, err := drs.DonationReportRepository.GetDonationReports(programID, limit, page)
	if err != nil {
		return nil, errors.New("failed to get data of donation reports")
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))
	return &entity.PaginatedDonationReport{
		Items:       *reports,
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: page,
	}, nil
}

func (drs *DonationReportService) DeleteDonationReport(id uint) error {
	return drs.DonationReportRepository.Delete(id)
}
