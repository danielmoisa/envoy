package repository

import (
	"github.com/danielmoisa/envoy/src/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UsersRepository struct {
	logger *zap.SugaredLogger
	db     *gorm.DB
}

func NewUsersRepository(logger *zap.SugaredLogger, db *gorm.DB) *UsersRepository {
	return &UsersRepository{
		logger: logger,
		db:     db,
	}
}

func (repository *UsersRepository) RetrieveUsers() ([]*model.User, error) {
	var resources []*model.User
	if err := repository.db.Find(&resources).Error; err != nil {
		return nil, err
	}
	return resources, nil
}

func (repository *UsersRepository) RetrieveByUserID(userID int) (*model.User, error) {
	var user *model.User

	if err := repository.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return &model.User{}, err
	}
	return user, nil
}

func (repository *UsersRepository) Create(name, email, password, avatar string) (*model.User, error) {
	user := &model.User{
		Nickname:       name,
		Email:          email,
		PasswordDigest: password,
		Avatar:         avatar,
	}
	if err := repository.db.Create(user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (repository *UsersRepository) DeleteByID(userID int) error {
	if err := repository.db.Where("id = ?", userID).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}
