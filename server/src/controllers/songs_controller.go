package controllers

import (
	"net/http"
	"strconv"

	"github.com/Jack-Music-Streaming/server/src/errors"
	"github.com/Jack-Music-Streaming/server/src/models"
	"github.com/labstack/echo/v4"
)

type songsController struct {
	songService models.SongService
	authService models.AuthService
}

type SongsController interface {
	Create(c echo.Context) error
	UploadSongMedia(c echo.Context) error
	GetAll(c echo.Context) error
	GetSongMedia(c echo.Context) error
}

func NewSongsController(s models.SongService, a models.AuthService) SongsController {
	return &songsController{
		songService: s,
		authService: a,
	}
}

func (s songsController) Create(c echo.Context) error {
	song := &models.Song{}

	if err := c.Bind(song); err != nil {
		return err
	}

	createdSong, err := s.songService.Create(song)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, createdSong)
}

func (s songsController) UploadSongMedia(c echo.Context) error {
	songID := c.Param("songID")
	user, err := s.authService.GetUserFromToken(c)

	if err != nil {
		return errors.UnauthorizedError()
	}

	song, err := s.songService.GetSongByID(songID)

	if err != nil {
		return errors.NotFoundError()
	}

	if song.AuthorID != user.ID {
		return errors.UnauthorizedError()
	}

	formFile, err := c.FormFile("songMedia")

	if err != nil {
		return err
	}

	_, err = s.songService.UploadSongMedia(song, formFile)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, song)
}

func (s songsController) GetAll(c echo.Context) error {
	page := c.Param("page")

	pageInt, err := strconv.Atoi(page)

	if err != nil {
		return err
	}

	songs, err := s.songService.GetAll(pageInt)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, songs)
}

func (s songsController) GetSongMedia(c echo.Context) error {
	song, err := s.songService.GetSongByID(c.Param("id"))

	if err != nil {
		return err
	}

	return c.File(song.Media.Src)
}
