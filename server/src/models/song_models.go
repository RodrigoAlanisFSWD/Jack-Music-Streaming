package models

import "mime/multipart"

type Song struct {
	Model
	Name     string `json:"name" gorm:"unique"`
	AuthorID uint   `json:"author_id"`
	Author   User   `json:"author"`
}

type SongRepository interface {
	Save(song *Song) (*Song, error)
	FindAll(query string, args ...interface{}) ([]Song, error)
	FindOne(query string, args ...interface{}) (*Song, error)
	Delete(song *Song) error
	Update(song *Song) (*Song, error)
}

type SongService interface {
	Create(song *Song, formFile *multipart.FileHeader) (*Song, error)
	GetAll() ([]Song, error)
	GetSongByID(id interface{}) (*Song, error)
	GetSongsByAuthor(authorID interface{}) ([]Song, error)
	Delete(song *Song) error
	Update(song *Song) (*Song, error)
}
