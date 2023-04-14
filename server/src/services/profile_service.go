package services

import (
	"mime/multipart"
	"sync"

	"github.com/Jack-Music-Streaming/server/src/database"
	"github.com/Jack-Music-Streaming/server/src/models"
)

type profileService struct {
	profileRepository models.ProfileRepository
	filesRepository   models.FilesRepository
}

var profileInstance *profileService

var profileOnce sync.Once

func NewProfileService(r models.ProfileRepository, f models.FilesRepository) models.ProfileService {

	profileOnce.Do(func() {
		profileInstance = &profileService{
			profileRepository: r,
			filesRepository:   f,
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

func (u *profileService) UpdateUserLogo(user *models.User, logoFormFile *multipart.FileHeader) (*models.User, error) {
	logoFile := models.FileDTO{
		Name: user.Name + "-(Logo)",
		Dst:  "uploads/logos/",
		Src:  logoFormFile,
	}

	u.filesRepository.CreateFileName(&logoFile)

	userLogo, err := u.filesRepository.Upload(&logoFile)

	if err != nil {
		return user, err
	}

	profile, err := u.profileRepository.FindOne("user_id = ?", user.ID)

	if err != nil {
		return user, err
	}

	err = database.DB.Where("id = 7").Delete(profile.Logo).Error

	if err != nil {
		return user, err
	}

	profile.Logo = *userLogo
	profile.LogoID = userLogo.ID

	_, err = u.profileRepository.Update(profile)

	user.Profile = *profile

	return user, err
}
