package routers

import (
	"github.com/Jack-Music-Streaming/server/src/controllers"
	"github.com/Jack-Music-Streaming/server/src/middlewares"
	"github.com/Jack-Music-Streaming/server/src/services"
	"github.com/labstack/echo/v4"
)

func SongsRouter(api *echo.Group) {
	songService := services.NewSongService(songRepository, filesRepository)

	songsController := controllers.NewSongsController(songService, authService)

	songs := api.Group("/songs")

	songs.GET("/:page", songsController.GetAll)

	songs.Use(middlewares.JWTMiddleware())

	songs.POST("/", songsController.Create)
	songs.POST("/uploadMedia", songsController.UploadSongMedia)
}
