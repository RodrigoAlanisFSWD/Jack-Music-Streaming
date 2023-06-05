package routers

import (
	"github.com/Jack-Music-Streaming/server/src/controllers"
	"github.com/Jack-Music-Streaming/server/src/middlewares"
	"github.com/labstack/echo/v4"
)

func AlbumRouter(api *echo.Group) {
	albumController := controllers.NewAlbumsController(albumService, authService)

	albums := api.Group("/albums")

	albums.GET("/author/:author", albumController.GetUserAlbums)
	albums.GET("/getOne/:id", albumController.GetAlbum)

	albums.Use(middlewares.JWTMiddleware())

	albums.PUT("/addSong/:album/:song", albumController.AddSong)
	albums.DELETE("/removeSong/:album/:song", albumController.RemoveSong)

}
