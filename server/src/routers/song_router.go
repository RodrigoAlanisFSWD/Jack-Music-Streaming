package routers

import (
	"github.com/Jack-Music-Streaming/server/src/controllers"
	"github.com/Jack-Music-Streaming/server/src/middlewares"
	"github.com/Jack-Music-Streaming/server/src/services"
	"github.com/labstack/echo/v4"
)

func SongsRouter(api *echo.Group) {
	songService := services.NewSongService(songRepository, filesRepository)

	songsController := controllers.NewSongsController(songService)

	songs := api.Group("/songs")

	songs.Use(middlewares.JWTMiddleware())

	songs.POST("/", songsController.Create)
}
