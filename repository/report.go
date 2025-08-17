package repository

import (
	"dibantuin-be/entity"
	"errors"
	"time"

	"gorm.io/gorm"
)

type ReportRepository struct {
	DB *gorm.DB
}

func NewReportRepository(db *gorm.DB) *ReportRepository {
	return &ReportRepository{DB: db}
}

func (rr *ReportRepository) GetDonationSummary(programID uint64, startDate, endDate time.Time) (*entity.DonationSummary, error) {
	var summary entity.DonationSummary
	err := rr.DB.Model(&entity.MoneyTransactionDonation{}).
		Select("COUNT(DISTINCT user_id) as donors_in_range, SUM(amount) as total_donations_in_range, AVG(amount) as avg_donation_in_range").
		Where("program_id = ? AND created_at BETWEEN ? AND ?", programID, startDate, endDate).
		Find(&summary).Error

	return &summary, err
}

func (rr *ReportRepository) GetAggregatedDonations(programID uint64, startDate, endDate time.Time, aggType string) ([]entity.AggregatedDonation, error) {
	var results []entity.AggregatedDonation
	var groupBy string

	switch aggType {
	case "daily":
		groupBy = "DATE_FORMAT(created_at, '%Y-%m-%d')"
	case "monthly":
		groupBy = "DATE_FORMAT(created_at, '%Y-%m')"
	default:
		return nil, errors.New("invalid aggregation type")
	}

	err := rr.DB.Model(&entity.MoneyTransactionDonation{}).
		Select(groupBy+" as date_label, SUM(amount) as amount").
		Where("program_id = ? AND donation_status = 'success' AND created_at BETWEEN ? AND ?", programID, startDate, endDate).
		Group(groupBy).
		Order(groupBy + " ASC").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
}

func (rr *ReportRepository) GetTopDonors(programID uint64, startDate, endDate time.Time, limit int) ([]entity.Donors, error) {
	var donors []entity.Donors
	err := rr.DB.Model(&entity.MoneyTransactionDonation{}).
		Select("donors_name, SUM(amount) as total_amount").
		Where("program_id = ? AND created_at BETWEEN ? AND ?", programID, startDate, endDate).
		Group("donors_name").
		Order("total_amount DESC").
		Limit(limit).
		Find(&donors).Error

	return donors, err
}

func (rr *ReportRepository) GetTotalUniqueDonors(programID uint64) (int64, error) {
	var totalDonors int64
	err := rr.DB.Model(&entity.MoneyTransactionDonation{}).
		Where("program_id = ? AND donation_status = 'success'", programID).
		Distinct("user_id").
		Count(&totalDonors).Error

	if err != nil {
		return 0, err
	}
	return totalDonors, nil
}

func (r *DonationProgramRepository) GetProgramStatusSummary() (*entity.ProgramStatusSummary, error) {
	var summary entity.ProgramStatusSummary

	err := r.DB.Model(&entity.DonationProgram{}).
		Select("COUNT(*) as total_programs, "+
			"SUM(CASE WHEN end_date > NOW() THEN 1 ELSE 0 END) as active_programs, "+
			"SUM(CASE WHEN current_amount >= target_amount THEN 1 ELSE 0 END) as completed_programs, "+
			"SUM(CASE WHEN end_date < NOW() AND current_amount < target_amount THEN 1 ELSE 0 END) as failed_programs").
		Joins("JOIN donation_program_requests ON donation_program_requests.program_id = donation_programs.id").
		Where("donation_program_requests.status_request = ?", "approved").
		Scan(&summary).Error

	if err != nil {
		return nil, err
	}
	return &summary, nil
}

func (r *DonationTransactionRepository) GetGlobalDonationSummary() (*entity.GlobalDonationSummary, error) {
	var summary entity.GlobalDonationSummary

	err := r.DB.Model(&entity.MoneyTransactionDonation{}).
		Select("SUM(amount) as total_donations, COUNT(DISTINCT user_id) as unique_donors").
		Where("donation_status = 'success'").
		Scan(&summary).Error

	if err != nil {
		return nil, err
	}
	return &summary, nil
}
