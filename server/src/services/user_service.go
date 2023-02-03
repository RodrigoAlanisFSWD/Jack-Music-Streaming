package services

import (
	"sync"

	"github.com/Jack-Music-Streaming/server/src/models"
)

type userService struct {
	userRepository models.UserRepository
}

var instance *userService

var userOnce sync.Once

func NewUserService(r models.UserRepository) models.UserService {

	userOnce.Do(func() {
		instance = &userService{
			userRepository: r,
		}
	})

	return instance
}

func (u *userService) Create(user *models.User) (*models.User, error) {
	return u.userRepository.Save(user)
}

func (u *userService) GetAll() ([]models.User, error) {
	return nil, nil
}

func (u *userService) GetUserByID(id interface{}) (*models.User, error) {
	user, err := u.userRepository.FindOne("id = ?", id)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userService) GetUserByEmail(email string) (*models.User, error) {
	user, err := u.userRepository.FindOne("email = ?", email)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userService) Delete(user *models.User) error {
	return nil
}

func (u *userService) Update(user *models.User) (*models.User, error) {
	return u.userRepository.Update(user)
}
