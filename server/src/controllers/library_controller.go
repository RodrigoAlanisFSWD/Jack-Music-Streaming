package controllers

import (
	"github.com/Jack-Music-Streaming/server/src/models"
	"github.com/labstack/echo/v4"
)

type libraryController struct {
	libraryService models.LibraryService
	authService    models.AuthService
}

type LibraryController interface {
	GetUserLibrary(c echo.Context) error
	AddSong(c echo.Context) error
	AddPlaylist(c echo.Context) error
	AddAlbum(c echo.Context) error
	RemoveSong(c echo.Context) error
	RemovePlaylist(c echo.Context) error
	RemoveAlbum(c echo.Context) error
}

func NewLibraryController(l models.LibraryService, a models.AuthService) LibraryController {
	return &libraryController{
		libraryService: l,
		authService:    a,
	}
}

func (l *libraryController) GetUserLibrary(c echo.Context) error {
	user, err := l.authService.GetUserFromToken(c)

	if err != nil {
		return err
	}

	library, err := l.libraryService.GetUserLibrary(user.ID)

	if err != nil {
		return err
	}

	return c.JSON(200, library)
}

func (l *libraryController) AddSong(c echo.Context) error {
	songID := c.Param("song")

	user, err := l.authService.GetUserFromToken(c)

	if err != nil {
		return err
	}

	library, err := l.libraryService.GetUserLibrary(user.ID)

	if err != nil {
		return err
	}

	updatedLibrary, err := l.libraryService.AddSong(songID, library)

	if err != nil {
		return err
	}

	return c.JSON(200, updatedLibrary)
}

func (l *libraryController) AddPlaylist(c echo.Context) error {
	playlistID := c.Param("playlist")

	user, err := l.authService.GetUserFromToken(c)

	if err != nil {
		return err
	}

	library, err := l.libraryService.GetUserLibrary(user.ID)

	if err != nil {
		return err
	}

	updatedLibrary, err := l.libraryService.AddPlaylist(playlistID, library)

	if err != nil {
		return err
	}

	return c.JSON(200, updatedLibrary)
}

func (l *libraryController) AddAlbum(c echo.Context) error {
	albumID := c.Param("album")

	user, err := l.authService.GetUserFromToken(c)

	if err != nil {
		return err
	}

	library, err := l.libraryService.GetUserLibrary(user.ID)

	if err != nil {
		return err
	}

	updatedLibrary, err := l.libraryService.AddAlbum(albumID, library)

	if err != nil {
		return err
	}

	return c.JSON(200, updatedLibrary)
}

func (l *libraryController) RemoveSong(c echo.Context) error {
	songID := c.Param("song")

	user, err := l.authService.GetUserFromToken(c)

	if err != nil {
		return err
	}

	library, err := l.libraryService.GetUserLibrary(user.ID)

	if err != nil {
		return err
	}

	updatedLibrary, err := l.libraryService.RemoveSong(songID, library)

	if err != nil {
		return err
	}

	return c.JSON(200, updatedLibrary)
}

func (l *libraryController) RemovePlaylist(c echo.Context) error {
	playlistID := c.Param("playlist")

	user, err := l.authService.GetUserFromToken(c)

	if err != nil {
		return err
	}

	library, err := l.libraryService.GetUserLibrary(user.ID)

	if err != nil {
		return err
	}

	updatedLibrary, err := l.libraryService.RemovePlaylist(playlistID, library)

	if err != nil {
		return err
	}

	return c.JSON(200, updatedLibrary)
}

func (l *libraryController) RemoveAlbum(c echo.Context) error {
	albumID := c.Param("album")

	user, err := l.authService.GetUserFromToken(c)

	if err != nil {
		return err
	}

	library, err := l.libraryService.GetUserLibrary(user.ID)

	if err != nil {
		return err
	}

	updatedLibrary, err := l.libraryService.RemoveAlbum(albumID, library)

	if err != nil {
		return err
	}

	return c.JSON(200, updatedLibrary)
}
