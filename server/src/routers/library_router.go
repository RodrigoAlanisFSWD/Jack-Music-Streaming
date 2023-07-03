package routers

import (
	"github.com/Jack-Music-Streaming/server/src/controllers"
	"github.com/Jack-Music-Streaming/server/src/middlewares"
	"github.com/labstack/echo/v4"
)

func LibraryRouter(api *echo.Group) {
	libraryController := controllers.NewLibraryController(libraryService, authService)

	library := api.Group("/library")

	library.Use(middlewares.JWTMiddleware())

	library.GET("/", libraryController.GetUserLibrary)
	library.POST("/addSong/:song", libraryController.AddSong)
	library.POST("/addPlaylist/:playlist", libraryController.AddPlaylist)
	library.POST("/addAlbum/:album", libraryController.AddAlbum)
	library.POST("/removeSong/:song", libraryController.RemoveSong)
	library.POST("/removePlaylist/:playlist", libraryController.RemovePlaylist)
	library.POST("/removeAlbum/:album", libraryController.RemoveAlbum)
}
