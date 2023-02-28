package services

import (
	"sync"

	"github.com/Jack-Music-Streaming/server/src/models"
)

type profileService struct {
	profileRepository models.ProfileRepository
}

var profileInstance *profileService

var profileOnce sync.Once

func NewProfileService(r models.ProfileRepository) models.ProfileService {

	profileOnce.Do(func() {
		profileInstance = &profileService{
			profileRepository: r,
		}
	})

	return profileInstance
}

func (u *profileService) Create(profile *models.Profile) (*models.Profile, error) {
	return u.profileRepository.Save(profile)
}

func (u *profileService) GetUserProfile(id interface{}) (*models.Profile, error) {
	profile, err := u.profileRepository.FindOne("id = ?", id)

	if err != nil {
		return profile, err
	}

	return profile, nil
}

func (u *profileService) Update(profile *models.Profile) (*models.Profile, error) {
	return u.profileRepository.Update(profile)
}
