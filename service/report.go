package service

import (
	"dibantuin-be/entity"
	"dibantuin-be/repository"
	"dibantuin-be/utils/math"
	"errors"
	"time"
)

type ReportService struct {
	DonationProgramRepository     *repository.DonationProgramRepository
	DonationTransactionRepository *repository.DonationTransactionRepository
	ReportRepository              *repository.ReportRepository
}

func NewReportService(donationProgramRepository *repository.DonationProgramRepository, donationTransactionRepository *repository.DonationTransactionRepository, reportRepository *repository.ReportRepository) *ReportService {
	return &ReportService{DonationProgramRepository: donationProgramRepository, DonationTransactionRepository: donationTransactionRepository, ReportRepository: reportRepository}
}

func (rs *ReportService) GetDonationProgramReport(programID uint64, userID uint64, startDate, endDate time.Time) (*entity.ProgramReport, error) {
	program, err := rs.DonationProgramRepository.GetDonationProgramWithoutRequestById(programID)
	if err != nil {
		return nil, errors.New("failed to get donation program by program id")
	}

	if program.UserID != userID {
		return nil, errors.New("you don't have access to this program")
	}

	if startDate.IsZero() && endDate.IsZero() {
		startDate = program.StartDate
		endDate = time.Now()
	}

	var aggType string
	durationDays := endDate.Sub(startDate).Hours() / 24
	if durationDays <= 30 {
		aggType = "daily"
	} else {
		aggType = "monthly"
	}

	// Ambil data dari repo
	summary, err := rs.ReportRepository.GetDonationSummary(programID, startDate, endDate)
	if err != nil {
		return nil, errors.New("failed to get donation summary")
	}

	// Ambil data donasi teragregasi
	donationsByDate, err := rs.ReportRepository.GetAggregatedDonations(programID, startDate, endDate, aggType)
	if err != nil {
		return nil, errors.New("failed to get aggregated donations")
	}

	// Ambil top donors
	topDonors, err := rs.ReportRepository.GetTopDonors(programID, startDate, endDate, 10)
	if err != nil {
		return nil, errors.New("failed to get top donors")
	}

	totalDonors, err := rs.ReportRepository.GetTotalUniqueDonors(programID)
	if err != nil {
		return nil, errors.New("failed to get total unique donors")
	}

	progress := 0.0
	if program.TargetAmount > 0 {
		progress = math.RoundToTwoDecimalPlaces((program.CurrentAmount / program.TargetAmount) * 100)
	}

	// Perhitungan sisa hari dan progres
	remainingDays := int(time.Until(program.EndDate).Hours() / 24)
	if remainingDays < 0 {
		remainingDays = 0
	}

	// Susun laporan
	report := &entity.ProgramReport{
		ProgramID:       program.ID,
		Title:           program.Title,
		TargetAmount:    program.TargetAmount,
		CurrentAmount:   program.CurrentAmount,
		TotalDonors:     int(totalDonors),
		ProgressPercent: progress,
		DonationsByDate: donationsByDate,
		TopDonors:       topDonors,
		RemainingDays:   remainingDays,
	}

	report.DonationsSummary.TotalDonationsInRange = summary.TotalDonationsInRange
	report.DonationsSummary.DonorsInRange = summary.DonorsInRange
	report.DonationsSummary.AvgDonationInRange = math.RoundToTwoDecimalPlaces(summary.AvgDonationInRange)

	return report, nil
}

func (rs *ReportService) GetGlobalReport() (*entity.GlobalReport, error) {
	// Ambil ringkasan status program
	programSummary, err := rs.DonationProgramRepository.GetProgramStatusSummary()
	if err != nil {
		return nil, errors.New("failed to get program status summary")
	}

	// Ambil ringkasan donasi global
	donationSummary, err := rs.DonationTransactionRepository.GetGlobalDonationSummary()
	if err != nil {
		return nil, errors.New("failed to get global donation summary")
	}

	report := &entity.GlobalReport{
		TotalDonations:    float64(donationSummary.TotalDonations),
		TotalPrograms:     programSummary.TotalPrograms,
		ActivePrograms:    programSummary.ActivePrograms,
		CompletedPrograms: programSummary.CompletedPrograms,
		FailedPrograms:    programSummary.FailedPrograms,
		UniqueDonors:      donationSummary.UniqueDonors,
	}

	return report, nil
}
