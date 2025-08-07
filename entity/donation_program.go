package entity

import "time"

type DonationProgram struct {
	ID            uint64    `gorm:"primaryKey"`
	UserID        uint64    `gorm:"not null"`
	CategoryID    uint64    `gorm:"not null"`
	Creator       string    `gorm:"type:varchar(100);not null"`
	Title         string    `gorm:"type:varchar(255);not null"`
	Description   string    `gorm:"type:text;not null"`
	TargetAmount  float64   `gorm:"type:decimal(15,2);not null"`
	CurrentAmount float64   `gorm:"type:decimal(15,2);not null"`
	RekeningInfo  string    `gorm:"type:varchar(255);not null"`
	IsPersonal    bool      `gorm:"not null"`
	Status        string    `gorm:"type:enum('soon','ongoing','over','success');not null"`
	EndDate       time.Time `gorm:"not null"`
	CoverImage    string    `gorm:"type:varchar(255);not null"`
	ContactInfo   string    `gorm:"type:varchar(100);not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
