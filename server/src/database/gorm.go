package database

import (
	"fmt"

	"github.com/Jack-Music-Streaming/server/src/models"
	"github.com/fatih/color"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	c := color.New(color.FgBlue)

	c.Println("Connecting To Database...")

	var (
		host     = viper.Get("DATABASE_HOST")
		user     = viper.Get("DATABASE_USER")
		port     = viper.Get("DATABASE_PORT")
		password = viper.Get("DATABASE_PASSWORD")
		name     = viper.Get("DATABASE_DB")
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	db.AutoMigrate(&models.Role{}, &models.Plan{}, &models.Scope{}, &models.User{})

	if err != nil {
		return err
	}

	DB = db

	c.Println("Database Connected!")

	return nil
}
