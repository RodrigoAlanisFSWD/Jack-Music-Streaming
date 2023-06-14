package repository

import (
	"fmt"

	"github.com/Jack-Music-Streaming/server/src/models"
	"gorm.io/gorm"
)

type profileRepository struct {
	DB *gorm.DB
}

func NewProfileRepository(db *gorm.DB) models.ProfileRepository {
	return &profileRepository{
		DB: db,
	}
}

func (u *profileRepository) Save(profile *models.Profile) (*models.Profile, error) {
	err := u.DB.Create(profile).Error

	if err != nil {
		fmt.Println(err)
		return &models.Profile{}, err
	}

	return profile, nil
}

func (u *profileRepository) FindOne(query string, args ...interface{}) (*models.Profile, error) {
	var profile models.Profile

	err := u.DB.Model(&profile).Where(query, args).Preload("Songs").Preload("Playlists").Preload("Albums").First(&profile).Error

	if err != nil {
		return &profile, err
	}

	return &profile, nil
}

func (u *profileRepository) Update(profile *models.Profile) (*models.Profile, error) {
	err := u.DB.Save(profile).Error

	if err != nil {
		fmt.Println(err)
		return &models.Profile{}, err
	}

	return profile, nil
}
