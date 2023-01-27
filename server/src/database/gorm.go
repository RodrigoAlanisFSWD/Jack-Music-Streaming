package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var (
		host     = "localhost"
		user     = "postgres"
		port     = 5432
		password = "password"
		name     = "jack"
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name),
		PreferSimpleProtocol: true,
	}))

	if err != nil {
		return err
	}

	DB = db

	return nil
}
