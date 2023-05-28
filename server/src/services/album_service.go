package services

import (
	"fmt"
	"mime/multipart"
	"sync"

	"github.com/Jack-Music-Streaming/server/src/database"
	"github.com/Jack-Music-Streaming/server/src/models"
)

type albumService struct {
	albumRepository models.AlbumRepository
	songService     models.SongService
	filesRepository models.FilesRepository
}

var albumInstance *albumService

var albumOnce sync.Once

func NewAlbumService(a models.AlbumRepository, s models.SongService, f models.FilesRepository) models.AlbumService {

	albumOnce.Do(func() {
		albumInstance = &albumService{
			albumRepository: a,
			songService:     s,
			filesRepository: f,
		}
	})

	return albumInstance
}

func (a *albumService) Create(album *models.Album) (*models.Album, error) {
	return a.albumRepository.Save(album)
}

func (a *albumService) GetAlbum(id interface{}) (*models.Album, error) {
	return a.albumRepository.FindOne("id = ?", id)
}

func (a *albumService) GetUserAlbums(id interface{}) ([]models.Album, error) {
	return a.albumRepository.FindMany("author_id = ?", id)
}

func (a *albumService) Update(album *models.Album) (*models.Album, error) {
	return a.albumRepository.Update(album)
}

func (a *albumService) UploadAlbumLogo(album *models.Album, logoFormFile *multipart.FileHeader) (*models.Album, error) {
	oldLogo := album.Logo

	logoFile := models.FileDTO{
		Name: album.Name + "-Album",
		Dst:  fmt.Sprintf("uploads/albums/%d/", album.AuthorID),
		Src:  logoFormFile,
	}

	a.filesRepository.CreateFileName(&logoFile)

	albumLogo, err := a.filesRepository.Upload(&logoFile)

	if err != nil {
		return album, err
	}

	album.Logo = *albumLogo
	album.LogoID = albumLogo.ID

	updatedAlbum, err := a.albumRepository.Update(album)

	if err != nil {
		return album, err
	}

	err = database.DB.Where("id != 1 AND id != 2").Delete(oldLogo).Error

	if err != nil {
		fmt.Println(err)
		return album, err
	}

	return updatedAlbum, err
}
