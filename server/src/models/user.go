package models

type User struct {
	Model
	Name     string `json:"name" gorm:"unique"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique"`
}

type UserService interface {
	Create(user User) (*User, error)
	GetAll() ([]User, error)
	GetOne(id int) (User, error)
	Delete(user *User) error
}

type UserRepository interface {
	Save(user *User) (*User, error)
	FindAll() ([]User, error)
	FindOne(id int) (User, error)
	Delete(user *User) error
}
