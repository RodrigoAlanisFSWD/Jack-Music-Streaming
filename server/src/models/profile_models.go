package models

import "mime/multipart"

type Profile struct {
	Model
	Songs     []Song     `gorm:"foreignKey:AuthorID;references:UserID"`
	Playlists []Playlist `gorm:"foreignKey:AuthorID;references:UserID"`
	Albums    []Album    `gorm:"foreignKey:AuthorID;references:UserID"`
	UserID    uint       `json:"user_id"`
	LogoID    uint       `json:"logo_id"`
	Logo      File       `json:"logo" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Library   Library    `gorm:"foreignKey:UserID;references:UserID"`
}

type ProfileRepository interface {
	Save(profile *Profile) (*Profile, error)
	FindOne(query string, args ...interface{}) (*Profile, error)
	Update(profile *Profile) (*Profile, error)
}

type ProfileService interface {
	Create(profile *Profile) (*Profile, error)
	GetUserProfile(id interface{}) (*Profile, error)
	Update(profile *Profile) (*Profile, error)
	UpdateUserLogo(user *User, logoFormFile *multipart.FileHeader) (*User, error)
}
