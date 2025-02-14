package storage

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Storage struct {
	UsersStorage *UsersStorage
}

func NewStorage(postgresDriver *gorm.DB, logger *zap.SugaredLogger) *Storage {
	return &Storage{
		UsersStorage: NewUsersStorage(logger, postgresDriver),
	}
}
