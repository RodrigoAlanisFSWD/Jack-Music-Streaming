package models

import "mime/multipart"

type Album struct {
	Model
	Name     string `json:"name"`
	AuthorID uint   `json:"author_id"`
	Author   User   `json:"author"`
	Songs    []Song `gorm:"constraint:OnDelete:SET NULL;" json:"songs"`
	LogoID   uint   `json:"logo_id"`
	Logo     File   `json:"logo"`
}

type AlbumRepository interface {
	Save(album *Album) (*Album, error)
	FindOne(query string, args ...interface{}) (*Album, error)
	FindMany(query string, args ...interface{}) ([]Album, error)
	Update(album *Album) (*Album, error)
}

type AlbumService interface {
	Create(album *Album) (*Album, error)
	GetAlbum(id interface{}) (*Album, error)
	GetUserAlbums(id interface{}) ([]Album, error)
	Update(album *Album) (*Album, error)
	UploadAlbumLogo(album *Album, logoFormFile *multipart.FileHeader) (*Album, error)
}
