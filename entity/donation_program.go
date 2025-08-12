package entity

import "time"

type DonationProgram struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement"`
	UserID        uint64    `gorm:"not null"`
	CategoryID    uint64    `gorm:"not null"`
	Creator       string    `gorm:"type:varchar(100);not null"`
	Title         string    `gorm:"type:varchar(255);not null"`
	Description   string    `gorm:"type:text;not null"`
	TargetAmount  float64   `gorm:"type:decimal(15,2);not null"`
	CurrentAmount float64   `gorm:"type:decimal(15,2);default:0"`
	RekeningInfo  string    `gorm:"type:varchar(255);not null"`
	IsPersonal    bool      `gorm:"not null;default:false"`
	Status        string    `gorm:"type:enum('soon','ongoing','over','success');default:'soon';not null"`
	StartDate     time.Time `gorm:"not null"`
	EndDate       time.Time `gorm:"not null"`
	CoverImage    string    `gorm:"type:varchar(255);not null"`
	ContactInfo   string    `gorm:"type:varchar(100);not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type DonationProgramRequest struct {
	ID            uint64  `gorm:"primaryKey;autoIncrement"`
	UserID        uint64  `gorm:"not null"`
	CategoryID    uint64  `gorm:"not null"`
	ProgramID     *uint64 `gorm:"default:null"` // pointer supaya bisa nil
	StatusRequest string  `gorm:"type:enum('pending','approved','rejected');default:'pending';not null"`
	KTPPhoto      string  `gorm:"type:varchar(255)"`
	SelfiePhoto   string  `gorm:"type:varchar(255)"`
	LegalDoc      string  `gorm:"type:varchar(255)"`
	AdminNotes    string  `gorm:"type:text"`
	CreatedAt     time.Time
	UpdatedAt     time.Time

	DonationProgram *DonationProgram `gorm:"foreignKey:ProgramID;references:ID;constraint:OnDelete:SET NULL"`
}

type DonationProgramRequestCreate struct {
	CategoryID   uint64    `json:"category_id" binding:"required"`
	Creator      string    `json:"creator" binding:"required"`
	Title        string    `json:"title" binding:"required"`
	Description  string    `json:"description" binding:"required"`
	IsPersonal   bool      `json:"is_personal"`
	StartDate    time.Time `json:"start_date" binding:"required"`
	EndDate      time.Time `json:"end_date" binding:"required"`
	CoverImage   string    `json:"cover_image" binding:"required"`
	TargetAmount float64   `json:"target_amount" binding:"required"`
	RekeningInfo string    `json:"rekening_info" binding:"required"`
	KTPPhoto     string    `json:"ktp_photo,omitempty"`
	SelfiePhoto  string    `json:"selfie_photo,omitempty"`
	LegalDoc     string    `json:"legal_doc,omitempty"`
	AdminNotes   string    `json:"admin_notes,omitempty"`
	ContactInfo  string    `json:"contact_info" binding:"required"`
}

type DonationProgramListItem struct {
	ID            uint64
	CategoryID    uint64
	Title         string
	Description   string
	TargetAmount  float64
	CurrentAmount float64
	Creator       string
	StartDate     time.Time
	EndDate       time.Time
	CoverImage    string
	Status        string
	StatusRequest string
}
