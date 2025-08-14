package repository

import (
	"dibantuin-be/entity"

	"gorm.io/gorm"
)

type DonationTransactionRepository struct {
	DB *gorm.DB
}

func NewDonationTransactionRepository(db *gorm.DB) *DonationTransactionRepository {
	return &DonationTransactionRepository{DB: db}
}

func (dtr *DonationTransactionRepository) CreateMoneyDonationTransaction(donationTransaction *entity.MoneyTransactionDonation) (*entity.MoneyTransactionDonation, error) {
	err := dtr.DB.Create(donationTransaction).Error
	if err != nil {
		return nil, err
	}
	return donationTransaction, nil
}

func (dtr *DonationTransactionRepository) GetDonationTransactionById(id uint64) (*entity.MoneyTransactionDonation, error) {
	var request entity.MoneyTransactionDonation
	err := dtr.DB.Where("id = ?", id).First(&request).Error
	if err != nil {
		return nil, err
	}

	return &request, nil
}

func (dtr *DonationTransactionRepository) UpdateStatusDonationTransactionById(id uint64, status string) error {
	return dtr.DB.Model(&entity.MoneyTransactionDonation{}).Where("id = ?", id).Update("donation_status", status).Error
}

func (dtr *DonationTransactionRepository) IncreaseCurrentAmount(programID uint64, amount float64) error {
	return dtr.DB.Model(&entity.DonationProgram{}).
		Where("id = ?", programID).
		UpdateColumn("current_amount", gorm.Expr("current_amount + ?", amount)).
		Error
}

func (dtr *DonationTransactionRepository) ListDonationTransactions(userID *uint64, status, search string, limit, page int) (*[]entity.DonationTransactionListItem, int64, error) {
	var transactions []entity.DonationTransactionListItem
	var totalItems int64
	offset := (page - 1) * limit

	query := dtr.DB.Table("money_transaction_donations as mtd").
		Select("mtd.id, dp.title, dp.creator, u.name, mtd.amount, mtd.donation_status, mtd.donors_name, mtd.donation_message").
		Joins("left join donation_programs as dp on mtd.program_id = dp.id").
		Joins("left join users as u on mtd.user_id = u.id")

	if userID != nil {
		query = query.Where("mtd.user_id = ?", *userID)
	}

	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where(
			"dp.title LIKE ? OR u.name LIKE ? OR mtd.donors_name LIKE ? OR mtd.donation_message LIKE ?",
			searchPattern, searchPattern, searchPattern, searchPattern,
		)
	}

	if status != "" {
		query = query.Where("mtd.donation_status = ?", status)
	}

	err := query.Count(&totalItems).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Limit(limit).Offset(offset).Order("mtd.created_at desc").Scan(&transactions).Error
	if err != nil {
		return nil, 0, err
	}

	return &transactions, totalItems, nil
}
