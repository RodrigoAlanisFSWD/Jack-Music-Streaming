package seeders

import "gorm.io/gorm"

type Seed struct {
	Name     string
	Function func(db *gorm.DB) error
}

type Seeder struct {
	Name  string
	Seeds []Seed
}
