package repository

import (
	"time"

	"github.com/danielmoisa/envoy/src/model"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type JobsRepository struct {
	logger *zap.SugaredLogger
	db     *gorm.DB
}

func NewJobsRepository(logger *zap.SugaredLogger, db *gorm.DB) *JobsRepository {
	return &JobsRepository{
		logger: logger,
		db:     db,
	}
}

// GetAll retrieves all jobs
func (r *JobsRepository) GetAll() ([]*model.Job, error) {
	var resources []*model.Job
	if err := r.db.Find(&resources).Error; err != nil {
		r.logger.Errorw("failed to retrieve jobs", "error", err)
		return nil, err
	}
	return resources, nil
}

// GetByID retrieves a single job by ID
func (r *JobsRepository) GetByID(id uuid.UUID) (*model.Job, error) {
	var job *model.Job
	if err := r.db.Where("id = ?", id).First(&job).Error; err != nil {
		r.logger.Errorw("failed to retrieve job", "id", id, "error", err)
		return &model.Job{}, err
	}
	return job, nil
}

// Create creates a new job
func (r *JobsRepository) Create(title, description, location, jobType string, salaryMin, salaryMax float64, companyID uuid.UUID) (*model.Job, error) {
	job := &model.Job{
		Title:       title,
		Description: description,
		Location:    location,
		JobType:     jobType,
		SalaryMin:   salaryMin,
		SalaryMax:   salaryMax,
		CompanyID:   companyID,
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := r.db.Create(job).Error; err != nil {
		r.logger.Errorw("failed to create job", "error", err)
		return job, err
	}
	return job, nil
}

// Update updates an existing job
func (r *JobsRepository) Update(id uuid.UUID, title, description, location, jobType string, salaryMin, salaryMax float64) (*model.Job, error) {
	job := &model.Job{
		Title:       title,
		Description: description,
		Location:    location,
		JobType:     jobType,
		SalaryMin:   salaryMin,
		SalaryMax:   salaryMax,
		UpdatedAt:   time.Now(),
	}
	if err := r.db.Where("id = ?", id).Updates(job).Error; err != nil {
		r.logger.Errorw("failed to update job", "id", id, "error", err)
		return job, err
	}
	return job, nil
}

// Delete deletes a job by ID
func (r *JobsRepository) Delete(id uuid.UUID) error {
	if err := r.db.Where("id = ?", id).Delete(&model.Job{}).Error; err != nil {
		r.logger.Errorw("failed to delete job", "id", id, "error", err)
		return err
	}
	return nil
}

// GetByCompanyID retrieves jobs by company ID
func (r *JobsRepository) GetByCompanyID(companyID uuid.UUID) ([]*model.Job, error) {
	var jobs []*model.Job
	if err := r.db.Where("company_id = ?", companyID).Find(&jobs).Error; err != nil {
		r.logger.Errorw("failed to find jobs by company ID", "companyID", companyID, "error", err)
		return nil, err
	}
	return jobs, nil
}
