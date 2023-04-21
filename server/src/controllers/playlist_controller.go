package controllers

import (
	"fmt"

	"github.com/Jack-Music-Streaming/server/src/errors"
	"github.com/Jack-Music-Streaming/server/src/models"
	"github.com/labstack/echo/v4"
)

type playlistController struct {
	playlistService models.PlaylistService
	authService     models.AuthService
}

type PlaylistController interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
	UploadSongMedia(c echo.Context) error
	GetUserPlaylists(c echo.Context) error
	GetPlaylist(c echo.Context) error
	AddSong(c echo.Context) error
}

func NewPlaylistController(p models.PlaylistService, a models.AuthService) PlaylistController {
	return &playlistController{
		playlistService: p,
		authService:     a,
	}
}

func (p *playlistController) Create(c echo.Context) error {
	playlist := &models.Playlist{}

	if err := c.Bind(playlist); err != nil {
		fmt.Println(err)
		return err
	}

	createdPlaylist, err := p.playlistService.Create(playlist)

	if err != nil {
		return err
	}

	return c.JSON(200, createdPlaylist)
}

func (p *playlistController) Update(c echo.Context) error {
	playlistData := &models.Playlist{}

	if err := c.Bind(playlistData); err != nil {
		fmt.Println(err)
		return err
	}

	playlist, err := p.playlistService.GetPlaylist(c.Param("id"))

	if err != nil {
		return errors.NotFoundError()
	}

	user, err := p.authService.GetUserFromToken(c)

	if err != nil {
		return err
	}

	if playlist.AuthorID != user.ID {
		return errors.UnauthorizedError()
	}

	playlist.Name = playlistData.Name

	updatedPlaylist, err := p.playlistService.Update(playlist)

	if err != nil {
		return err
	}

	return c.JSON(200, updatedPlaylist)
}

func (p *playlistController) UploadSongMedia(c echo.Context) error {

}

func (p *playlistController) GetUserPlaylists(c echo.Context) error {

}

func (p *playlistController) GetPlaylist(c echo.Context) error {

}

func (p *playlistController) AddSong(c echo.Context) error {

}
