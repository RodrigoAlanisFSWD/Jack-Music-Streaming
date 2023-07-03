package services

import (
	"sync"

	"github.com/Jack-Music-Streaming/server/src/database"
	"github.com/Jack-Music-Streaming/server/src/models"
	"gorm.io/gorm"
)

type libraryService struct {
	libraryRepository models.LibraryRepository
	songService       models.SongService
	playlistService   models.PlaylistService
	albumService      models.AlbumService
}

var libraryInstance *libraryService

var libraryOnce sync.Once

func NewLibraryRepository(l models.LibraryRepository, s models.SongService, a models.AlbumService, p models.PlaylistService) models.LibraryService {
	libraryOnce.Do(func() {
		libraryInstance = &libraryService{
			libraryRepository: l,
			songService:       s,
			albumService:      a,
			playlistService:   p,
		}
	})

	return libraryInstance
}

func (l *libraryService) Create(library *models.Library) (*models.Library, error) {
	return l.libraryRepository.Save(library)
}

func (l *libraryService) GetUserLibrary(id interface{}) (*models.Library, error) {
	library, err := l.libraryRepository.FindOne("id = ?", id)

	if err != nil {
		return library, err
	}

	return library, nil
}

func (l *libraryService) Update(library *models.Library) (*models.Library, error) {
	return l.libraryRepository.Update(library)
}

func (l *libraryService) AddSong(songID interface{}, library *models.Library) (*models.Library, error) {
	song, err := l.songService.GetSongByID(songID)

	if err != nil {
		return library, err
	}

	for _, s := range library.Songs {
		if song.ID == s.ID {
			return library, err
		}
	}

	library.Songs = append(library.Songs, *song)

	return l.libraryRepository.Save(library)
}

func (l *libraryService) AddPlaylist(playlistID interface{}, library *models.Library) (*models.Library, error) {
	playlist, err := l.playlistService.GetPlaylist(playlistID)

	if err != nil {
		return library, err
	}

	for _, p := range library.Playlists {
		if playlist.ID == p.ID {
			return library, err
		}
	}

	library.Playlists = append(library.Playlists, *playlist)

	return l.libraryRepository.Save(library)
}

func (l *libraryService) AddAlbum(albumID interface{}, library *models.Library) (*models.Library, error) {
	album, err := l.albumService.GetAlbum(albumID)

	if err != nil {
		return library, err
	}

	for _, a := range library.Albums {
		if album.ID == a.ID {
			return library, err
		}
	}

	library.Albums = append(library.Albums, *album)

	return l.libraryRepository.Save(library)
}

func (l *libraryService) RemoveSong(songID interface{}, library *models.Library) (*models.Library, error) {
	song, err := l.songService.GetSongByID(songID)

	if err != nil {
		return library, err
	}

	err = database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Model(library).Association("Songs").Delete(song)

	if err != nil {
		return library, err
	}

	return library, err
}

func (l *libraryService) RemovePlaylist(playlistID interface{}, library *models.Library) (*models.Library, error) {
	playlist, err := l.playlistService.GetPlaylist(playlistID)

	if err != nil {
		return library, err
	}

	err = database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Model(library).Association("Playlists").Delete(playlist)

	if err != nil {
		return library, err
	}

	return library, err
}

func (l *libraryService) RemoveAlbum(albumID interface{}, library *models.Library) (*models.Library, error) {
	album, err := l.albumService.GetAlbum(albumID)

	if err != nil {
		return library, err
	}

	err = database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Model(library).Association("Albums").Delete(album)

	if err != nil {
		return library, err
	}

	return library, err
}
