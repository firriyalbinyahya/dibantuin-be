package entity

import "time"

type Category struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100);not null"`
	Slug      string `gorm:"type:varchar(100);unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
