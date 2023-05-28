package models

import "mime/multipart"

type Song struct {
	Model
	Name      string      `json:"name" gorm:"unique"`
	Duration  string      `json:"duration"`
	AuthorID  uint        `json:"author_id"`
	Author    User        `json:"author"`
	MediaID   uint        `json:"media_id"`
	Media     File        `json:"media"`
	LogoID    uint        `json:"logo_id"`
	Logo      File        `json:"logo"`
	Playlists []*Playlist `json:"playlists" gorm:"many2many:playlist_songs;constraint:OnDelete:CASCADE;"`
	AlbumID   uint        `json:"album_id"`
	Album     Album       `json:"album"`
}

type SongRepository interface {
	Save(song *Song) (*Song, error)
	FindAll(query string, args ...interface{}) ([]Song, error)
	FindOne(query string, args ...interface{}) (*Song, error)
	GetPage(page int) ([]Song, error)
	Delete(song *Song) error
	Update(song *Song) (*Song, error)
	SearchSongs(query string) ([]Song, error)
}

type SongService interface {
	UploadSongMedia(song *Song, songFormFile *multipart.FileHeader, logoFormFile *multipart.FileHeader) (*Song, error)
	Create(song *Song) (*Song, error)
	GetAll(page int) ([]Song, error)
	GetSongByID(id interface{}) (*Song, error)
	GetSongsByAuthor(authorID interface{}) ([]Song, error)
	Delete(song *Song) error
	Update(song *Song) (*Song, error)
	SearchSongs(query string) ([]Song, error)
}
