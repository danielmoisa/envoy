package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	UsersRepository        *UsersRepository
	CompaniesRepository    *CompaniesRepository
	JobsRepository         *JobsRepository
	ApplicationsRepository *ApplicationsRepository
	CandidatesRepository   *CandidatesRepository
}

func NewRepository(postgresDriver *gorm.DB, logger *zap.SugaredLogger) *Repository {
	return &Repository{
		UsersRepository:        NewUsersRepository(logger, postgresDriver),
		CompaniesRepository:    NewCompaniesRepository(logger, postgresDriver),
		JobsRepository:         NewJobsRepository(logger, postgresDriver),
		ApplicationsRepository: NewApplicationsRepository(logger, postgresDriver),
		CandidatesRepository:   NewCandidatesRepository(logger, postgresDriver),
	}
}
