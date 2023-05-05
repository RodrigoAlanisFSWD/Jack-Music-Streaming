package controllers

import (
	"fmt"
	"strconv"

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
	GetUserPlaylists(c echo.Context) error
	GetPlaylist(c echo.Context) error
	AddSong(c echo.Context) error
	UpdatePlaylistLogo(c echo.Context) error
	GetAll(c echo.Context) error
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

	user, err := p.authService.GetUserFromToken(c)

	if err != nil {
		return err
	}

	playlist.AuthorID = user.ID
	playlist.LogoID = 2

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

func (p *playlistController) GetUserPlaylists(c echo.Context) error {
	authorID := c.Param("author")

	playlists, err := p.playlistService.GetUserPlaylists(authorID)

	if err != nil {
		return err
	}

	return c.JSON(200, playlists)
}

func (p *playlistController) GetPlaylist(c echo.Context) error {
	playlistID := c.Param("id")

	playlist, err := p.playlistService.GetPlaylist(playlistID)

	if err != nil {
		return errors.NotFoundError()
	}

	return c.JSON(200, playlist)
}

func (p *playlistController) AddSong(c echo.Context) error {
	playlistID := c.Param("playlist")
	songID := c.Param("song")

	playlist, err := p.playlistService.GetPlaylist(playlistID)

	if err != nil {
		fmt.Println(err)
		return errors.NotFoundError()
	}

	user, err := p.authService.GetUserFromToken(c)

	if err != nil {
		return err
	}

	if playlist.AuthorID != user.ID {
		return errors.UnauthorizedError()
	}

	updatedPlaylist, err := p.playlistService.AddSong(playlist, songID)

	if err != nil {
		return err
	}

	return c.JSON(200, updatedPlaylist)
}

func (p *playlistController) UpdatePlaylistLogo(c echo.Context) error {
	playlistID := c.Param("playlist")
	user, err := p.authService.GetUserFromToken(c)

	if err != nil {
		return errors.UnauthorizedError()
	}

	playlist, err := p.playlistService.GetPlaylist(playlistID)

	fmt.Println(playlist)

	if err != nil {
		return errors.NotFoundError()
	}

	if playlist.AuthorID != user.ID {
		return errors.UnauthorizedError()
	}

	logoFormFile, err := c.FormFile("logo")

	if err != nil {
		return err
	}

	updatedPlaylist, err := p.playlistService.UploadPlaylistLogo(playlist, logoFormFile)

	if err != nil {
		return err
	}

	return c.JSON(200, updatedPlaylist)
}

func (p *playlistController) GetAll(c echo.Context) error {
	page := c.Param("page")

	pageInt, err := strconv.Atoi(page)

	if err != nil {
		return err
	}

	playlists, err := p.playlistService.GetAll(pageInt)

	if err != nil {
		return err
	}

	return c.JSON(200, playlists)
}
