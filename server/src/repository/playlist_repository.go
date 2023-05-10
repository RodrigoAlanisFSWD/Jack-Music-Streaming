package repository

import (
	"fmt"

	"github.com/Jack-Music-Streaming/server/src/models"
	"gorm.io/gorm"
)

type playlistRepository struct {
	DB *gorm.DB
}

func NewPlaylistRepository(db *gorm.DB) models.PlaylistRepository {
	return &playlistRepository{
		DB: db,
	}
}

func (p *playlistRepository) Save(playlist *models.Playlist) (*models.Playlist, error) {
	err := p.DB.Create(playlist).Error

	if err != nil {
		fmt.Println(err)
		return &models.Playlist{}, err
	}

	return playlist, nil
}

func (p *playlistRepository) FindOne(query string, args ...interface{}) (*models.Playlist, error) {
	var playlist models.Playlist

	err := p.DB.Model(&playlist).Where(query, args).Preload("Author").Preload("Songs").Preload("Songs.Author").Preload("Logo").First(&playlist).Error

	if err != nil {
		return &playlist, err
	}

	return &playlist, nil
}

func (p *playlistRepository) FindMany(query string, args ...interface{}) (*[]models.Playlist, error) {
	var playlists []models.Playlist

	err := p.DB.Model(&playlists).Preload("Author").Preload("Songs").Where(query, args).Find(&playlists).Error

	if err != nil {
		return &playlists, err
	}

	return &playlists, nil
}

func (p *playlistRepository) Update(playlist *models.Playlist) (*models.Playlist, error) {
	err := p.DB.Save(playlist).Error

	if err != nil {
		fmt.Println(err)
		return &models.Playlist{}, err
	}

	return playlist, nil
}

func (p *playlistRepository) AddSong(playlist *models.Playlist, song *models.Song) (*models.Playlist, error) {
	playlist.Songs = append(playlist.Songs, *song)

	return p.Update(playlist)
}

func (p *playlistRepository) RemoveSong(playlist *models.Playlist, song *models.Song) (*models.Playlist, error) {
	err := p.DB.Session(&gorm.Session{FullSaveAssociations: true}).Model(playlist).Association("Songs").Delete(song)

	if err != nil {
		fmt.Println(err)
		return playlist, err
	}

	// updated, err := p.Update(playlist)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return playlist, err
	// }

	// var filteredArray []models.Song

	// for _, s := range playlist.Songs {
	// 	if s.ID != song.ID {
	// 		filteredArray = append(filteredArray, s)
	// 	}
	// }

	// fmt.Println(len(filteredArray))

	// updated.Songs = filteredArray

	return p.Update(playlist)
}

func (p *playlistRepository) GetPage(page int) ([]models.Playlist, error) {
	var playlists []models.Playlist

	err := p.DB.Model(&playlists).Limit(15).Offset((page - 1) * 10).Preload("Author").Find(&playlists).Error

	if err != nil {
		fmt.Println(err)
		return playlists, err
	}

	return playlists, nil
}
