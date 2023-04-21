package services

import (
	"sync"

	"github.com/Jack-Music-Streaming/server/src/models"
)

type playlistService struct {
	playlistRepository models.PlaylistRepository
	songService        models.SongService
}

var playlistInstance *playlistService

var playlistOnce sync.Once

func NewPlaylistService(p models.PlaylistRepository, s models.SongService) models.PlaylistService {

	playlistOnce.Do(func() {
		playlistInstance = &playlistService{
			playlistRepository: p,
			songService:        s,
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
