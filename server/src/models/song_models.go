package models

type Song struct {
	Model
	Name     string `json:"name" gorm:"unique"`
	AuthorID uint   `json:"author_id"`
	Author   User   `json:"author"`
}
