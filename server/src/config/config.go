package config

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

func GetConfig() {
	c := color.New(color.FgGreen)

	c.Println("Initializating Configuration...")

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	c.Println("Configuration Initialized!")
}
