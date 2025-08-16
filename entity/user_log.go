package entity

import "time"

type UserLog struct {
	ID          uint64    `gorm:"primaryKey" json:"id"`
	UserID      uint64    `gorm:"not null" json:"user_id"`
	ActionType  string    `gorm:"type:enum('LOGIN','LOGOUT', 'CREATE_USER', 'REQUEST_PROGRAM','UPDATE_PROGRAM','DELETE_PROGRAM','SUBMIT_VERIFICATION','DONATE_TRANSACTION','UPDATE_PROFILE','UPLOAD_REPORT','VERIFY_PROGRAM','REJECT_VERIFICATION');not null" json:"action_type"`
	TargetTable string    `gorm:"type:varchar(100)" json:"target_table"`
	TargetID    uint64    `json:"target_id"`
	Description string    `gorm:"type:text" json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type PaginatedUserLogs struct {
	Items       []UserLog `json:"items"`
	TotalItems  int64     `json:"total_items"`
	TotalPages  int       `json:"total_pages"`
	CurrentPage int       `json:"current_page"`
}
