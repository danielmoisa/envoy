package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	UsersRepository *UsersRepository
}

func NewRepository(postgresDriver *gorm.DB, logger *zap.SugaredLogger) *Repository {
	return &Repository{
		UsersRepository: NewUsersRepository(logger, postgresDriver),
	}
}
