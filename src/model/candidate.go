package model

import (
	"time"

	"github.com/google/uuid"
)

type Candidate struct {
	ID          uuid.UUID `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID      uuid.UUID `json:"user_id" gorm:"column:user_id;type:uuid;not null"`
	User        User      `json:"-" gorm:"foreignKey:UserID"`
	Title       string    `json:"title" gorm:"column:title;type:varchar;size:100"`
	Summary     string    `json:"summary" gorm:"column:summary;type:text"`
	Experience  int       `json:"experience" gorm:"column:experience;type:int"` // Years of experience
	Skills      []Skill   `json:"skills" gorm:"many2many:candidate_skills"`
	IsAvailable bool      `json:"is_available" gorm:"column:is_available;type:boolean;default:true"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp"`
}
