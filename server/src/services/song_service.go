package services

import (
	"fmt"
	"mime/multipart"
	"sync"

	"github.com/Jack-Music-Streaming/server/src/database"
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

func (u *songService) UploadSongMedia(song *models.Song, songFormFile *multipart.FileHeader, logoFormFile *multipart.FileHeader) (*models.Song, error) {
	oldLogo := song.Logo
	oldMedia := song.Media

	songFile := models.FileDTO{
		Name: song.Name,
		Dst:  fmt.Sprintf("uploads/songs/%d/", song.AuthorID),
		Src:  songFormFile,
	}

	logoFile := models.FileDTO{
		Name: song.Name,
		Dst:  fmt.Sprintf("uploads/songs/%d/", song.AuthorID),
		Src:  logoFormFile,
	}

	u.filesRepository.CreateFileName(&songFile)
	u.filesRepository.CreateFileName(&logoFile)

	songMedia, err := u.filesRepository.Upload(&songFile)

	if err != nil {
		return song, err
	}

	logoMedia, err := u.filesRepository.Upload(&logoFile)

	if err != nil {
		return song, err
	}

	song.Media = *songMedia
	song.MediaID = songMedia.ID
	song.Logo = *logoMedia
	song.LogoID = logoMedia.ID

	updatedSong, err := u.songRepository.Update(song)

	if err != nil {
		return song, err
	}

	err = database.DB.Where("id != 1 AND id != 2").Delete(oldLogo).Error

	if err != nil {
		return song, err
	}

	err = database.DB.Where("id != 1 AND id != 2").Delete(oldMedia).Error

	return updatedSong, err
}

func (u *songService) Create(song *models.Song) (*models.Song, error) {
	return u.songRepository.Save(song)
}

func (u *songService) GetAll(page int) ([]models.Song, error) {
	return u.songRepository.GetPage(page)
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
	err := u.songRepository.Delete(song)

	u.filesRepository.Delete(&song.Media)
	u.filesRepository.Delete(&song.Logo)

	return err
}

func (u *songService) Update(song *models.Song) (*models.Song, error) {
	return u.songRepository.Update(song)
}
