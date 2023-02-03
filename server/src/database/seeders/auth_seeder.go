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
		{
			Name: "Seed Plans",
			Function: func(db *gorm.DB) error {
				plans := []models.Plan{
					{
						Name: "Basic",
					},
					{
						Name: "Pro",
					},
				}

				return db.Model(&plans).Create(&plans).Error
			},
		},
		{
			Name: "Seed Scopes",
			Function: func(db *gorm.DB) error {
				scopes := []models.Scope{
					{
						RoleID:      1,
						PlanID:      1,
						Permissions: "playlist=false;music=true;premium=false;others=false",
					},
					{
						RoleID:      2,
						PlanID:      1,
						Permissions: "playlist=false;music=true;premium=false;others=false;artist_premium=false",
					},
					{
						RoleID:      1,
						PlanID:      2,
						Permissions: "playlist=true;music=true;premium=true;others=true",
					},
					{
						RoleID:      2,
						PlanID:      2,
						Permissions: "playlist=true;music=true;premium=true;others=true;artist_premium=true",
					},
				}

				return db.Model(&scopes).Create(&scopes).Error
			},
		},
	},
}
