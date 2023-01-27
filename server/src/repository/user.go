package repository

import (
	"github.com/Jack-Music-Streaming/server/src/models"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) models.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (u *userRepository) Save(user *models.User) (*models.User, error) {
	return user, nil
}

func (u *userRepository) FindAll() ([]models.User, error) {
	return nil, nil
}

func (u *userRepository) FindOne(id int) (models.User, error) {
	return models.User{}, nil
}

func (u *userRepository) Delete(user *models.User) error {
	return nil
}
