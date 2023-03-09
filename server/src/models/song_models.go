package models

import "mime/multipart"

type Song struct {
	Model
	Name     string `json:"name" gorm:"unique"`
	AuthorID uint   `json:"author_id"`
	Author   User   `json:"author"`
	MediaID  uint   `json:"media_id"`
	Media    File   `json:"media"`
}

type SongRepository interface {
	Save(song *Song) (*Song, error)
	FindAll(query string, args ...interface{}) ([]Song, error)
	FindOne(query string, args ...interface{}) (*Song, error)
	GetPage(page int) ([]Song, error)
	Delete(song *Song) error
	Update(song *Song) (*Song, error)
}

type SongService interface {
	UploadSongMedia(song *Song, formFile *multipart.FileHeader) (*Song, error)
	Create(song *Song) (*Song, error)
	GetAll(page int) ([]Song, error)
	GetSongByID(id interface{}) (*Song, error)
	GetSongsByAuthor(authorID interface{}) ([]Song, error)
	Delete(song *Song) error
	Update(song *Song) (*Song, error)
}
