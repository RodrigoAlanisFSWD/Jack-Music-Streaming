package database

import (
	"github.com/fatih/color"
	"gorm.io/gorm"
)

var p = color.New(color.FgCyan)

func Rollback(db *gorm.DB) error {
	var err error

	p.Println("Executing Rollback...")

	for _, table := range entities {
		err = db.Migrator().DropTable(table)

		if err != nil {
			return err
		}
	}

	p.Println("Rollback Finished!")

	return err
}

func Migrate(db *gorm.DB) error {
	var err error

	p.Println("Executing Migrations...")

	for _, table := range entities {
		err = db.AutoMigrate(table)

		if err != nil {
			return err
		}
	}

	p.Println("Migrations Finished!")

	return err
}
