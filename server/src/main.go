package main

import (
	"log"
	"net/http"

	"github.com/Jack-Music-Streaming/server/src/config"
	"github.com/Jack-Music-Streaming/server/src/database"
	"github.com/Jack-Music-Streaming/server/src/models"
	"github.com/Jack-Music-Streaming/server/src/routers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// HTTP Request Logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"set-cookie"},
	}))

	// Error Handler
	e.Use(middleware.Recover())

	// Get Env Vars
	config.GetConfig()

	// Initializing Database Connection And Migrations / Seeders / Rollback
	if err := database.InitDB(); err != nil {
		log.Panic(err)
		return
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	// Main Api Router
	api := e.Group("/api")

	api.GET("/file/:id", func(c echo.Context) error {
		var file models.File

		err := database.DB.Model(&file).Where("id = ?", c.Param("id")).First(&file).Error

		if err != nil {
			return err
		}

		return c.File(file.Src)
	})

	// Initialize Router Dependencies
	routers.InitializeRotuers()

	// Applying Routers
	routers.AuthRouter(api)
	routers.SongsRouter(api)
	routers.PlaylistRouter(api)
	routers.AlbumRouter(api)

	// Initializing The Server
	e.Logger.Fatal(e.Start(":8080"))
}
