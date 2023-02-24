package repository

import (
	"fmt"

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
	err := u.DB.Create(user).Error

	if err != nil {
		fmt.Println(err)
		return &models.User{}, err
	}

	return user, nil
}

func (u *userRepository) FindAll(query string, args ...interface{}) ([]models.User, error) {
	var users []models.User

	err := u.DB.Model(&users).Where(query, args).Find(&users).Error

	if err != nil {
		return users, err
	}

	return users, nil
}

func (u *userRepository) FindOne(query string, args ...interface{}) (*models.User, error) {
	var user models.User

	err := u.DB.Model(&user).Where(query, args).Preload("Role").First(&user).Error

	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (u *userRepository) Delete(user *models.User) error {
	return nil
}

func (u *userRepository) Update(user *models.User) (*models.User, error) {
	err := u.DB.Save(user).Error

	if err != nil {
		fmt.Println(err)
		return &models.User{}, err
	}

	return user, nil
}
