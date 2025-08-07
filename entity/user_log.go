package entity

import "time"

type UserLog struct {
	ID          uint64 `gorm:"primaryKey"`
	UserID      uint64 `gorm:"not null"`
	ActionType  string `gorm:"type:enum('LOGIN','LOGOUT','REQUEST_PROGRAM','UPDATE_PROGRAM','DELETE_PROGRAM','SUBMIT_VERIFICATION','DONATE_TRANSACTION','UPDATE_PROFILE','UPLOAD_REPORT','VERIFY_PROGRAM','REJECT_VERIFICATION');not null"`
	TargetTable string `gorm:"type:varchar(100)"`
	TargetID    uint64
	Description string `gorm:"type:text"`
	CreatedAt   time.Time
}
