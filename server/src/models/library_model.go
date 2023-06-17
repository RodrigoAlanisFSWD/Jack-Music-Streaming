package models

type Library struct {
	Model
	Songs     []Song     `gorm:"many2many:library_songs;" json:"songs"`
	Playlists []Playlist `gorm:"many2many:library_playlists;" json:"playlists"`
	Albums    []Album    `gorm:"many2many:library_albums;" json:"albums"`
	UserID    uint       `json:"user_id"`
}

type LibraryRepository interface {
	Save(library *Library) (*Library, error)
	FindOne(query string, args ...interface{}) (*Library, error)
	Update(library *Library) (*Library, error)
}

type LibraryService interface {
	Create(library *Library) (*Library, error)
	GetUserLibrary(id interface{}) (*Library, error)
	Update(library *Library) (*Library, error)
	AddSong(songID interface{}, library *Library) (*Library, error)
	AddPlaylist(playlistID interface{}, library *Library) (*Library, error)
	AddAlbum(albumID interface{}, library *Library) (*Library, error)
}
