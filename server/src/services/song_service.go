package services

import (
	"fmt"
	"mime/multipart"
	"sync"

	"github.com/Jack-Music-Streaming/server/src/models"
)

type songService struct {
	songRepository  models.SongRepository
	filesRepository models.FilesRepository
}

var songInstance *songService

var songOnce sync.Once

func NewSongService(r models.SongRepository, f models.FilesRepository) models.SongService {

	songOnce.Do(func() {
		songInstance = &songService{
			songRepository:  r,
			filesRepository: f,
		}
	})

	return songInstance
}

func (u *songService) Create(song *models.Song, formFile *multipart.FileHeader) (*models.Song, error) {
	file := models.File{
		Name: song.Name,
		Dst:  fmt.Sprintf("/uploads/songs/%d/", song.AuthorID),
		Src:  formFile,
	}

	u.filesRepository.CreateFileName(&file)

	u.filesRepository.Upload(&file)

	return u.songRepository.Save(song)
}

func (u *songService) GetAll() ([]models.Song, error) {
	return nil, nil
}

func (u *songService) GetSongByID(id interface{}) (*models.Song, error) {
	song, err := u.songRepository.FindOne("id = ?", id)

	if err != nil {
		return song, err
	}

	return song, nil
}

func (u *songService) GetSongsByAuthor(authorID interface{}) ([]models.Song, error) {
	songs, err := u.songRepository.FindAll("author_id = ?", authorID)

	if err != nil {
		return songs, err
	}

	return songs, nil
}

func (u *songService) Delete(song *models.Song) error {
	return nil
}

func (u *songService) Update(song *models.Song) (*models.Song, error) {
	return u.songRepository.Update(song)
}
