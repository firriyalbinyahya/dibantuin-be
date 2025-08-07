package entity

import "time"

type User struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(50);unique;not null"`
	Password  string `gorm:"type:varchar(100);not null"`
	Role      string `gorm:"type:enum('admin', 'user');not null"`
	CreatedAt time.Time
}
