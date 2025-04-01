package repository

import (
	"github.com/danielmoisa/envoy/src/model"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CandidatesRepository struct {
	logger *zap.SugaredLogger
	db     *gorm.DB
}

func NewCandidatesRepository(logger *zap.SugaredLogger, db *gorm.DB) *CandidatesRepository {
	return &CandidatesRepository{
		logger: logger,
		db:     db,
	}
}

func (r *CandidatesRepository) GetAll() ([]*model.Candidate, error) {
	var candidates []*model.Candidate
	if err := r.db.Find(&candidates).Error; err != nil {
		r.logger.Errorw("failed to retrieve candidates", "error", err)
		return nil, err
	}
	return candidates, nil
}

func (r *CandidatesRepository) GetByID(id uuid.UUID) (*model.Candidate, error) {
	var candidate *model.Candidate
	if err := r.db.Where("id = ?", id).First(&candidate).Error; err != nil {
		r.logger.Errorw("failed to retrieve candidate", "id", id, "error", err)
		return &model.Candidate{}, err
	}
	return candidate, nil
}

func (r *CandidatesRepository) Create(newCandidate *model.Candidate) (*model.Candidate, error) {
	candidate := &model.Candidate{
		UserID:      newCandidate.UserID,
		Title:       newCandidate.Title,
		Summary:     newCandidate.Summary,
		Experience:  newCandidate.Experience,
		IsAvailable: newCandidate.IsAvailable,
	}

	if err := r.db.Create(candidate).Error; err != nil {
		r.logger.Errorw("failed to create candidate", "error", err)
		return candidate, err
	}
	return candidate, nil
}

func (r *CandidatesRepository) Update(updatedCandidate *model.Candidate) (*model.Candidate, error) {
	candidate := &model.Candidate{
		UserID:      updatedCandidate.UserID,
		Title:       updatedCandidate.Title,
		Summary:     updatedCandidate.Summary,
		Experience:  updatedCandidate.Experience,
		IsAvailable: updatedCandidate.IsAvailable,
	}

	if err := r.db.Where("id = ?", updatedCandidate.ID).Updates(candidate).Error; err != nil {
		r.logger.Errorw("failed to update candidate", "id", updatedCandidate.ID, "error", err)
		return candidate, err
	}
	return candidate, nil
}

func (r *CandidatesRepository) Delete(id uuid.UUID) error {
	if err := r.db.Where("id = ?", id).Delete(&model.Candidate{}).Error; err != nil {
		r.logger.Errorw("failed to delete candidate", "id", id, "error", err)
		return err
	}
	return nil
}
