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

func (a *albumRepository) Save(album *models.Album) (*models.Album, error) {
	err := a.DB.Create(album).Error

	if err != nil {
		fmt.Println(err)
		return &models.Album{}, err
	}

	return album, nil
}

func (a *albumRepository) FindOne(query string, args ...interface{}) (*models.Album, error) {
	var album models.Album

	err := a.DB.Model(&album).Where(query, args).Preload("Author").Preload("Songs").Preload("Songs.Album").Preload("Songs.Author").Preload("Logo").First(&album).Error

	if err != nil {
		return &album, err
	}

	return &album, nil
}

func (a *albumRepository) FindMany(query string, args ...interface{}) ([]models.Album, error) {
	var albums []models.Album

	err := a.DB.Model(&albums).Preload("Author").Preload("Songs").Where(query, args).Find(&albums).Error

	if err != nil {
		return albums, err
	}

	return albums, nil
}

func (a *albumRepository) Update(album *models.Album) (*models.Album, error) {
	err := a.DB.Save(album).Error

	if err != nil {
		fmt.Println(err)
		return &models.Album{}, err
	}

	return album, nil
}

func (a *albumRepository) AddSong(album *models.Album, song *models.Song) (*models.Album, error) {
	album.Songs = append(album.Songs, *song)

	return a.Update(album)
}

func (a *albumRepository) RemoveSong(album *models.Album, song *models.Song) (*models.Album, error) {
	err := a.DB.Session(&gorm.Session{FullSaveAssociations: true}).Model(album).Association("Songs").Delete(song)

	if err != nil {
		fmt.Println(err)
		return album, err
	}

	return a.Update(album)
}
