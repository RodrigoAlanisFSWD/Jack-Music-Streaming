package routers

import (
	"github.com/Jack-Music-Streaming/server/src/controllers"
	"github.com/Jack-Music-Streaming/server/src/middlewares"
	"github.com/labstack/echo/v4"
)

func SongsRouter(api *echo.Group) {
	songsController := controllers.NewSongsController(songService, authService)

	songs := api.Group("/songs")

	songs.GET("/:page", songsController.GetAll)
	songs.GET("/getOne/:id", songsController.GetSongByID)
	songs.GET("/media/:id", songsController.GetSongMedia)
	songs.GET("/user/:user", songsController.GetUserSongs)

	songs.Use(middlewares.JWTMiddleware())

	songs.PUT("/update", songsController.UpdateSong)
	songs.POST("/", songsController.Create)
	songs.POST("/uploadMedia/:songID", songsController.UploadSongMedia)
	songs.DELETE("/:id", songsController.DeleteSong)
}
