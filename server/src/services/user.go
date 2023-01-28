package services

import (
	"sync"

	"github.com/Jack-Music-Streaming/server/src/models"
)

type userService struct {
	userRepository models.UserRepository
}

var instance *userService

var once sync.Once

func NewUserService(r models.UserRepository) models.UserService {

	once.Do(func() {
		instance = &userService{
			userRepository: r,
		}
	})

	return instance
}

func (u *userService) Create(user models.User) (*models.User, error) {
	return u.userRepository.Save(&user)
}

func (u *userService) GetAll() ([]models.User, error) {
	return nil, nil
}

func (u *userService) GetOne(id int) (models.User, error) {
	return models.User{}, nil
}

func (u *userService) Delete(user *models.User) error {
	return nil
}
