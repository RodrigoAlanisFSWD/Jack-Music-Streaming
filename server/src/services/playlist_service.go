package services

import (
	"fmt"
	"mime/multipart"
	"sync"

	"github.com/Jack-Music-Streaming/server/src/database"
	"github.com/Jack-Music-Streaming/server/src/models"
)

type playlistService struct {
	playlistRepository models.PlaylistRepository
	songService        models.SongService
	filesRepository    models.FilesRepository
}

var playlistInstance *playlistService

var playlistOnce sync.Once

func NewPlaylistService(p models.PlaylistRepository, s models.SongService, f models.FilesRepository) models.PlaylistService {

	playlistOnce.Do(func() {
		playlistInstance = &playlistService{
			playlistRepository: p,
			songService:        s,
			filesRepository:    f,
		}
	})

	return playlistInstance
}

func (p *playlistService) Create(playlist *models.Playlist) (*models.Playlist, error) {
	return p.playlistRepository.Save(playlist)
}

func (p *playlistService) GetPlaylist(id interface{}) (*models.Playlist, error) {
	return p.playlistRepository.FindOne("id = ?", id)
}

func (p *playlistService) GetUserPlaylists(id interface{}) (*[]models.Playlist, error) {
	return p.playlistRepository.FindMany("author_id = ?", id)
}

func (p *playlistService) GetAll(page int) ([]models.Playlist, error) {
	return p.playlistRepository.GetPage(page)
}

func (p *playlistService) Update(playlist *models.Playlist) (*models.Playlist, error) {
	return p.playlistRepository.Update(playlist)
}

func (p *playlistService) AddSong(playlist *models.Playlist, songID interface{}) (*models.Playlist, error) {
	song, err := p.songService.GetSongByID(songID)

	if err != nil {
		return playlist, err
	}

	return p.playlistRepository.AddSong(playlist, song)
}

func (p *playlistService) UploadPlaylistLogo(playlist *models.Playlist, logoFormFile *multipart.FileHeader) (*models.Playlist, error) {
	oldLogo := playlist.Logo

	logoFile := models.FileDTO{
		Name: playlist.Name + "-Playlist",
		Dst:  fmt.Sprintf("uploads/playlists/%d/", playlist.AuthorID),
		Src:  logoFormFile,
	}

	p.filesRepository.CreateFileName(&logoFile)

	playlistLogo, err := p.filesRepository.Upload(&logoFile)

	if err != nil {
		return playlist, err
	}

	playlist.Logo = *playlistLogo
	playlist.LogoID = playlistLogo.ID

	updatedPlaylist, err := p.playlistRepository.Update(playlist)

	if err != nil {
		return playlist, err
	}

	err = database.DB.Where("id != 1 AND id != 2").Delete(oldLogo).Error

	if err != nil {
		fmt.Println(err)
		return playlist, err
	}

	return updatedPlaylist, err
}
