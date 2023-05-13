package repository

import (
	"fmt"

	"github.com/Jack-Music-Streaming/server/src/models"
	"gorm.io/gorm"
)

type songRepository struct {
	DB *gorm.DB
}

func NewSongRepository(db *gorm.DB) models.SongRepository {
	return &songRepository{
		DB: db,
	}
}

func (u *songRepository) GetPage(page int) ([]models.Song, error) {
	var songs []models.Song

	err := u.DB.Model(&songs).Limit(15).Offset((page - 1) * 10).Preload("Author").Find(&songs).Error

	if err != nil {
		fmt.Println(err)
		return songs, err
	}

	return songs, nil
}

func (u *songRepository) Save(song *models.Song) (*models.Song, error) {
	err := u.DB.Create(song).Error

	if err != nil {
		fmt.Println(err)
		return &models.Song{}, err
	}

	return song, nil
}

func (u *songRepository) FindAll(query string, args ...interface{}) ([]models.Song, error) {
	var songs []models.Song

	err := u.DB.Model(&songs).Where(query, args).Preload("Author").Find(&songs).Error

	if err != nil {
		return songs, err
	}

	return songs, nil
}

func (u *songRepository) FindOne(query string, args ...interface{}) (*models.Song, error) {
	var song models.Song

	err := u.DB.Model(&song).Where(query, args).Preload("Media").Preload("Logo").First(&song).Error

	if err != nil {
		return &song, err
	}

	return &song, nil
}

func (u *songRepository) Delete(song *models.Song) error {
	u.DB.Model(song).Session(&gorm.Session{FullSaveAssociations: true}).Association("Playlists").Clear()
	u.DB.Where("id = ?", song.ID).Select("Media").Select("Logo").Delete(song)
	return nil
}

func (u *songRepository) Update(song *models.Song) (*models.Song, error) {
	err := u.DB.Save(song).Error

	if err != nil {
		fmt.Println(err)
		return &models.Song{}, err
	}

	return song, nil
}

func (u *songRepository) SearchSongs(query string) ([]models.Song, error) {
	var songs []models.Song

	err := u.DB.Model(&songs).Preload("Author").Where("songs.name LIKE ?", query).Find(&songs).Error

	return songs, err
}
