package entity

type DailySummary struct {
	Date   string  `json:"date"`
	Amount float64 `json:"amount"`
}

type DonationSummary struct {
	TotalDonationsInRange float64 `json:"total_donations_in_range"`
	DonorsInRange         int     `json:"donors_in_range"`
	AvgDonationInRange    float64 `json:"avg_donation_in_range"`
}

type Donors struct {
	DonorsName  string  `json:"donors_name"`
	TotalAmount float64 `json:"total_amount"`
}

type AggregatedDonation struct {
	DateLabel string
	Amount    float64
}

type ProgramReport struct {
	ProgramID        uint64               `json:"program_id"`
	Title            string               `json:"title"`
	TargetAmount     float64              `json:"target_amount"`
	CurrentAmount    float64              `json:"current_amount"`
	ProgressPercent  float64              `json:"progress_percent"`
	TotalDonors      int                  `json:"total_donors"`
	RemainingDays    int                  `json:"remaining_days"`
	DonationsSummary DonationSummary      `json:"donations_summary"`
	DonationsByDate  []AggregatedDonation `json:"donations_by_date"`
	TopDonors        []Donors             `json:"top_donors"`
}

// global report
type ProgramStatusSummary struct {
	TotalPrograms     int64 `json:"total_programs"`
	ActivePrograms    int64 `json:"active_programs"`
	CompletedPrograms int64 `json:"completed_programs"`
	FailedPrograms    int64 `json:"failed_programs"`
}

type GlobalDonationSummary struct {
	TotalDonations float64 `json:"total_donations"`
	UniqueDonors   int64   `json:"unique_donors"`
}

type GlobalReport struct {
	TotalDonations    float64 `json:"total_donations"`
	TotalPrograms     int64   `json:"total_programs"`
	ActivePrograms    int64   `json:"active_programs"`
	CompletedPrograms int64   `json:"completed_programs"`
	FailedPrograms    int64   `json:"failed_programs"`
	UniqueDonors      int64   `json:"unique_donors"`
}
