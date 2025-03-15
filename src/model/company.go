package model

import (
	"time"

	"github.com/google/uuid"
)

type Company struct {
	ID             uuid.UUID `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID         uuid.UUID `json:"user_id" gorm:"column:user_id;type:uuid;not null"`
	User           User      `json:"-" gorm:"foreignKey:UserID"`
	CompanyName    string    `json:"company_name" gorm:"column:company_name;type:varchar;size:100;not null"`
	Industry       string    `json:"industry" gorm:"column:industry;type:varchar;size:50"`
	CompanySize    string    `json:"company_size" gorm:"column:company_size;type:varchar;size:20"`
	CompanyWebsite string    `json:"company_website" gorm:"column:company_website;type:varchar;size:255"`
	CompanyLogo    string    `json:"company_logo" gorm:"column:company_logo;type:varchar;size:255"`
	CreatedAt      time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:timestamp"`
}
