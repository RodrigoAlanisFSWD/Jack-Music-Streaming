package controllers

import (
	"net/http"

	"github.com/Jack-Music-Streaming/server/src/models"
	"github.com/labstack/echo/v4"
)

type songsController struct {
	songService models.SongService
}

type SongsController interface {
	Create(c echo.Context) error
}

func NewSongsController(s models.SongService) SongsController {
	return &songsController{
		songService: s,
	}
}

func (s songsController) Create(c echo.Context) error {
	song := &models.Song{}

	if err := c.Bind(song); err != nil {
		return err
	}

	formFile, err := c.FormFile("songMedia")

	if err != nil {
		return err
	}

	createdSong, err := s.songService.Create(song, formFile)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, createdSong)
}
