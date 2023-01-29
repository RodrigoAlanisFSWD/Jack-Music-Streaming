package models

type User struct {
	Model
	Name     string `json:"name" gorm:"unique"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique"`
	RoleID   int    `json:"role_id"`
	Role     Role
	PlanID   int `json:"plan_id"`
	Plan     Plan
}

type UserService interface {
	Create(user *User) (*User, error)
	GetAll() ([]User, error)
	GetUserByID(id interface{}) (*User, error)
	GetUserByEmail(email string) (*User, error)
	Delete(user *User) error
	Update(user *User) (*User, error)
}

type UserRepository interface {
	Save(user *User) (*User, error)
	FindAll(query string, args ...interface{}) ([]User, error)
	FindOne(query string, args ...interface{}) (*User, error)
	Delete(user *User) error
}
