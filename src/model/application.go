package model

import (
	"time"

	"github.com/google/uuid"
)

type ApplicationStatus string

const (
	StatusApplied   ApplicationStatus = "applied"
	StatusScreening ApplicationStatus = "screening"
	StatusInterview ApplicationStatus = "interview"
	StatusOffer     ApplicationStatus = "offer"
	StatusRejected  ApplicationStatus = "rejected"
	StatusWithdrawn ApplicationStatus = "withdrawn"
)

type Application struct {
	ID          uuid.UUID         `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	JobID       uuid.UUID         `json:"job_id" gorm:"column:job_id;type:uuid;not null"`
	Job         Job               `json:"-" gorm:"foreignKey:JobID"`
	CandidateID uuid.UUID         `json:"candidate_id" gorm:"column:candidate_id;type:uuid;not null"`
	Candidate   Candidate         `json:"-" gorm:"foreignKey:CandidateID"`
	Status      ApplicationStatus `json:"status" gorm:"column:status;type:varchar;size:20;default:'applied'"`
	CoverLetter string            `json:"cover_letter" gorm:"column:cover_letter;type:text"`
	AppliedAt   time.Time         `json:"applied_at" gorm:"column:applied_at;type:timestamp;not null"`
	UpdatedAt   time.Time         `gorm:"column:updated_at;type:timestamp"`
}
