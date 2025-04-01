package repository

import (
	"time"

	"github.com/danielmoisa/envoy/src/model"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ApplicationsRepository struct {
	logger *zap.SugaredLogger
	db     *gorm.DB
}

func NewApplicationsRepository(logger *zap.SugaredLogger, db *gorm.DB) *ApplicationsRepository {
	return &ApplicationsRepository{
		logger: logger,
		db:     db,
	}
}

// GetAll retrieves all applications
func (r *ApplicationsRepository) GetAll() ([]*model.Application, error) {
	var resources []*model.Application
	if err := r.db.Find(&resources).Error; err != nil {
		r.logger.Errorw("failed to retrieve applications", "error", err)
		return nil, err
	}
	return resources, nil
}

// GetByID retrieves a single application by ID
func (r *ApplicationsRepository) GetByID(id uuid.UUID) (*model.Application, error) {
	var application *model.Application
	if err := r.db.Where("id = ?", id).First(&application).Error; err != nil {
		r.logger.Errorw("failed to retrieve application", "id", id, "error", err)
		return &model.Application{}, err
	}
	return application, nil
}

// Create creates a new application
func (r *ApplicationsRepository) Create(newApplication *model.Application) (*model.Application, error) {
	application := &model.Application{
		JobID:       newApplication.JobID,
		CandidateID: newApplication.CandidateID,
		Status:      newApplication.Status,
		CoverLetter: newApplication.CoverLetter,
		AppliedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := r.db.Create(application).Error; err != nil {
		r.logger.Errorw("failed to create application", "error", err)
		return application, err
	}
	return application, nil
}

// Update updates an existing application
func (r *ApplicationsRepository) Update(updatedApplication *model.Application) (*model.Application, error) {
	application := &model.Application{
		Status:      updatedApplication.Status,
		CoverLetter: updatedApplication.CoverLetter,
		UpdatedAt:   time.Now(),
	}
	if err := r.db.Where("id = ?", updatedApplication.ID).Updates(application).Error; err != nil {
		r.logger.Errorw("failed to update application", "id", updatedApplication.ID, "error", err)
		return application, err
	}
	return application, nil
}

// Delete deletes a application by ID
func (r *ApplicationsRepository) Delete(id uuid.UUID) error {
	if err := r.db.Where("id = ?", id).Delete(&model.Application{}).Error; err != nil {
		r.logger.Errorw("failed to delete application", "id", id, "error", err)
		return err
	}
	return nil
}
