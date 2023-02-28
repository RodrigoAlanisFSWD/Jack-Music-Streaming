package models

type Profile struct {
	Model
	Songs  []Song `gorm:"foreignKey:AuthorID"`
	UserID uint   `json:"user_id"`
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
}
