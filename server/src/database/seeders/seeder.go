package seeders

import (
	"github.com/fatih/color"
	"gorm.io/gorm"
)

var seeders = []Seeder{
	AuthSeeder,
}

var c = color.New(color.FgWhite)

func SeedDatabase(db *gorm.DB) error {
	c.Println("Seed Database:")

	for _, seeder := range seeders {
		c.Println(c.Sprintf("- Running %s Seeder", seeder.Name))
		for _, seed := range seeder.Seeds {
			c.Println(c.Sprintf("-- Running %s Seed", seed.Name))
			err := seed.Function(db)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
