package repository

import (
	"fmt"

	"github.com/Jack-Music-Streaming/server/src/models"
	"gorm.io/gorm"
)

type albumRepository struct {
	DB *gorm.DB
}

func NewAlbumRepository(db *gorm.DB) models.AlbumRepository {
	return &albumRepository{
		DB: db,
	}
}

func (p *albumRepository) Save(album *models.Album) (*models.Album, error) {
	err := p.DB.Create(album).Error

	if err != nil {
		fmt.Println(err)
		return &models.Album{}, err
	}

	return album, nil
}

func (p *albumRepository) FindOne(query string, args ...interface{}) (*models.Album, error) {
	var album models.Album

	err := p.DB.Model(&album).Where(query, args).Preload("Author").Preload("Songs").Preload("Songs.Author").Preload("Logo").First(&album).Error

	if err != nil {
		return &album, err
	}

	return &album, nil
}

func (p *albumRepository) FindMany(query string, args ...interface{}) ([]models.Album, error) {
	var albums []models.Album

	err := p.DB.Model(&albums).Preload("Author").Preload("Songs").Where(query, args).Find(&albums).Error

	if err != nil {
		return albums, err
	}

	return albums, nil
}

func (p *albumRepository) Update(album *models.Album) (*models.Album, error) {
	err := p.DB.Save(album).Error

	if err != nil {
		fmt.Println(err)
		return &models.Album{}, err
	}

	return album, nil
}
