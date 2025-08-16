package entity

import "time"

type DonationProgram struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        uint64    `gorm:"not null" json:"user_id"`
	CategoryID    uint64    `gorm:"not null" json:"category_id"`
	Creator       string    `gorm:"type:varchar(100);not null" json:"creator"`
	Title         string    `gorm:"type:varchar(255);not null" json:"title"`
	Description   string    `gorm:"type:text;not null" json:"description"`
	TargetAmount  float64   `gorm:"type:decimal(15,2);not null" json:"target_amount"`
	CurrentAmount float64   `gorm:"type:decimal(15,2);default:0" json:"current_amount"`
	RekeningInfo  string    `gorm:"type:varchar(255);not null" json:"rekening_info"`
	IsPersonal    bool      `gorm:"not null;default:false" json:"is_personal"`
	StartDate     time.Time `gorm:"not null" json:"start_date"`
	EndDate       time.Time `gorm:"not null" json:"end_date"`
	CoverImage    string    `gorm:"type:varchar(255);not null" json:"cover_image"`
	ContactInfo   string    `gorm:"type:varchar(100);not null" json:"contact_info"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	// Relasi
	Category               Category                 `gorm:"foreignKey:CategoryID;references:ID" json:"category"`
	DonationProgramRequest []DonationProgramRequest `gorm:"foreignKey:ProgramID;references:ID" json:"donation_program_request,omitempty"`
}

type DonationProgramRequest struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        uint64    `gorm:"not null" json:"user_id"`
	CategoryID    uint64    `gorm:"not null" json:"category_id"`
	ProgramID     *uint64   `gorm:"default:null" json:"program_id"` // pointer supaya bisa nil
	StatusRequest string    `gorm:"type:enum('pending','approved','rejected');default:'pending';not null" json:"status_request"`
	KTPPhoto      string    `gorm:"type:varchar(255)" json:"ktp_photo"`
	SelfiePhoto   string    `gorm:"type:varchar(255)" json:"selfie_photo"`
	LegalDoc      string    `gorm:"type:varchar(255)" json:"legal_doc"`
	AdminNotes    string    `gorm:"type:text" json:"admin_notes"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	DonationProgram *DonationProgram `gorm:"foreignKey:ProgramID;references:ID;constraint:OnDelete:SET NULL" json:"-"`
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
	ID            uint64    `json:"id"`
	CategoryID    uint64    `json:"category_id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	TargetAmount  float64   `json:"target_amount"`
	CurrentAmount float64   `json:"current_amount"`
	Creator       string    `json:"creator"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	CoverImage    string    `json:"cover_image"`
	StatusRequest string    `json:"status_request"`
}
