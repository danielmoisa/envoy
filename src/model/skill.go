package model

import (
	"time"

	"github.com/google/uuid"
)

type Skill struct {
	ID        uuid.UUID `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string    `json:"name" gorm:"column:name;type:varchar;size:50;uniqueIndex"`
	Category  string    `json:"category" gorm:"column:category;type:varchar;size:50"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp"`
}
