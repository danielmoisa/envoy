package repository

import (
	"time"

	"github.com/danielmoisa/envoy/src/model"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CompaniesRepository struct {
	logger *zap.SugaredLogger
	db     *gorm.DB
}

func NewCompaniesRepository(logger *zap.SugaredLogger, db *gorm.DB) *CompaniesRepository {
	return &CompaniesRepository{
		logger: logger,
		db:     db,
	}
}

// GetAll retrieves all companies
func (r *CompaniesRepository) GetAll() ([]*model.Company, error) {
	var resources []*model.Company
	if err := r.db.Find(&resources).Error; err != nil {
		r.logger.Errorw("failed to retrieve companies", "error", err)
		return nil, err
	}
	return resources, nil
}

// GetByID retrieves a single company by ID
func (r *CompaniesRepository) GetByID(id uuid.UUID) (*model.Company, error) {
	var company *model.Company
	if err := r.db.Where("id = ?", id).First(&company).Error; err != nil {
		r.logger.Errorw("failed to retrieve company", "id", id, "error", err)
		return &model.Company{}, err
	}
	return company, nil
}

// Create creates a new company
func (r *CompaniesRepository) Create(newCompany *model.Company) (*model.Company, error) {
	company := &model.Company{
		CompanyName:    newCompany.CompanyName,
		Industry:       newCompany.Industry,
		CompanySize:    newCompany.CompanySize,
		CompanyWebsite: newCompany.CompanyWebsite,
		CompanyLogo:    newCompany.CompanyLogo,
		UserID:         newCompany.UserID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	if err := r.db.Create(company).Error; err != nil {
		r.logger.Errorw("failed to create company", "error", err)
		return company, err
	}
	return company, nil
}

// Update updates an existing company
func (r *CompaniesRepository) Update(updatedCompany *model.Company) (*model.Company, error) {
	company := &model.Company{
		CompanyName:    updatedCompany.CompanyName,
		Industry:       updatedCompany.Industry,
		CompanySize:    updatedCompany.CompanySize,
		CompanyWebsite: updatedCompany.CompanyWebsite,
		CompanyLogo:    updatedCompany.CompanyLogo,
		UpdatedAt:      time.Now(),
	}
	if err := r.db.Where("id = ?", updatedCompany.ID).Updates(company).Error; err != nil {
		r.logger.Errorw("failed to update company", "id", updatedCompany.ID, "error", err)
		return company, err
	}
	return company, nil
}

// Delete deletes a company by ID
func (r *CompaniesRepository) Delete(id uuid.UUID) error {
	if err := r.db.Where("id = ?", id).Delete(&model.Company{}).Error; err != nil {
		r.logger.Errorw("failed to delete company", "id", id, "error", err)
		return err
	}
	return nil
}

// GetByUserID retrieves a company by user ID
func (r *CompaniesRepository) GetByUserID(userID uuid.UUID) (*model.Company, error) {
	var company model.Company
	if err := r.db.Where("user_id = ?", userID).First(&company).Error; err != nil {
		r.logger.Errorw("failed to find company by user ID", "userID", userID, "error", err)
		return nil, err
	}
	return &company, nil
}
