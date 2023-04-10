package seeders

import (
	"github.com/Jack-Music-Streaming/server/src/models"
	"gorm.io/gorm"
)

var AuthSeeder = Seeder{
	Name: "Auth",
	Seeds: []Seed{
		{
			Name: "Seed Roles",
			Function: func(db *gorm.DB) error {
				roles := []models.Role{
					{
						Name: "LISTENER",
					},
					{
						Name: "ARTIST",
					},
				}

				return db.Model(&roles).Create(&roles).Error
			},
		},
	},
}
