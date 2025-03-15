package model

import (
	"time"

	"github.com/google/uuid"
)

type Job struct {
	ID          uuid.UUID `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	CompanyID   uuid.UUID `json:"company_id" gorm:"column:company_id;type:uuid;not null"`
	Company     Company   `json:"-" gorm:"foreignKey:CompanyID"`
	Title       string    `json:"title" gorm:"column:title;type:varchar;size:100;not null"`
	Description string    `json:"description" gorm:"column:description;type:text;not null"`
	Location    string    `json:"location" gorm:"column:location;type:varchar;size:100"`
	JobType     string    `json:"job_type" gorm:"column:job_type;type:varchar;size:50"` // Full-time, Part-time, Contract
	SalaryMin   float64   `json:"salary_min" gorm:"column:salary_min;type:decimal(10,2)"`
	SalaryMax   float64   `json:"salary_max" gorm:"column:salary_max;type:decimal(10,2)"`
	Skills      []Skill   `json:"skills" gorm:"many2many:job_skills"`
	IsActive    bool      `json:"is_active" gorm:"column:is_active;type:boolean;default:true"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp"`
}
