package repository

import (
	"github.com/Jack-Music-Streaming/server/src/models"
	"gorm.io/gorm"
)

type libraryRepository struct {
	DB *gorm.DB
}

func NewLibraryRepository(db *gorm.DB) models.LibraryRepository {
	return &libraryRepository{
		DB: db,
	}
}

func (l *libraryRepository) Save(library *models.Library) (*models.Library, error) {
	err := l.DB.Save(library).Error

	return library, err
}

func (l *libraryRepository) FindOne(query string, args ...interface{}) (*models.Library, error) {
	var library models.Library

	err := l.DB.Model(&library).Where(query, args).Preload("Songs").Preload("Songs.Author").Preload("Playlists").Preload("Albums").First(&library).Error

	return &library, err
}

func (l *libraryRepository) Update(library *models.Library) (*models.Library, error) {
	err := l.DB.Save(library).Error

	return library, err
}
