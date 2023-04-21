package routers

import (
	"github.com/Jack-Music-Streaming/server/src/controllers"
	"github.com/Jack-Music-Streaming/server/src/middlewares"
	"github.com/labstack/echo/v4"
)

func PlaylistRouter(api *echo.Group) {
	playlistController := controllers.NewPlaylistController(playlistService, authService)

	playlist := api.Group("/playlist")

	playlist.Use(middlewares.JWTMiddleware())

	playlist.POST("/", playlistController.Create)
	playlist.PUT("/:id", playlistController.Update)
	playlist.GET("/:id", playlistController.GetPlaylist)
	playlist.GET("/author/:author", playlistController.GetUserPlaylists)
	playlist.PUT("/addSong/:playlist/:song", playlistController.AddSong)
}
