package models

type Playlist struct {
	Model
	Name     string `json:"name"`
	AuthorID uint   `json:"author_id"`
	Author   User   `json:"author"`
	Type     string `json:"type"`
	Songs    []Song `json:"songs" gorm:"many2many:playlist_songs;"`
}

type PlaylistRepository interface {
	Save(playlist *Playlist) (*Playlist, error)
	FindOne(query string, args ...interface{}) (*Playlist, error)
	FindMany(query string, args ...interface{}) (*[]Playlist, error)
	Update(playlist *Playlist) (*Playlist, error)
	AddSong(playlist *Playlist, song *Song) (*Playlist, error)
}

type PlaylistService interface {
	Create(playlist *Playlist) (*Playlist, error)
	GetPlaylist(id interface{}) (*Playlist, error)
	GetUserPlaylists(id interface{}) (*[]Playlist, error)
	Update(playlist *Playlist) (*Playlist, error)
	AddSong(playlist *Playlist, songID interface{}) (*Playlist, error)
}
