package models

type Role struct {
	Model
	Name string `json:"name" gorm:"unique"`
}

type Plan struct {
	Model
	Name string `json:"name" gorm:"unique"`
}

type Scope struct {
	Model
	RoleID      int `json:"role_id"`
	Role        Role
	PlanID      int `json:"plan_id"`
	Plan        Plan
	Permissions string `json:"permissions"`
}

type ScopeRepository interface {
	FindScope(roleID interface{}, planID interface{}) (*Scope, error)
}

type ScopeService interface {
	GetScope(user *User) (*Scope, error)
}
