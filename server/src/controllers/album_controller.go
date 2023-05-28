package controllers

import (
	"fmt"

	"github.com/Jack-Music-Streaming/server/src/errors"
	"github.com/Jack-Music-Streaming/server/src/models"
	"github.com/labstack/echo/v4"
)

type albumsController struct {
	albumService models.AlbumService
	authService  models.AuthService
}

type AlbumsController interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
	GetUserAlbums(c echo.Context) error
	GetAlbum(c echo.Context) error
	UpdateAlbumLogo(c echo.Context) error
}

func NewAlbumsController(p models.AlbumService, a models.AuthService) AlbumsController {
	return &albumsController{
		albumService: p,
		authService:  a,
	}
}

func (p *albumsController) Create(c echo.Context) error {
	album := &models.Album{}

	if err := c.Bind(album); err != nil {
		fmt.Println(err)
		return err
	}

	user, err := p.authService.GetUserFromToken(c)

	if err != nil {
		return err
	}

	album.AuthorID = user.ID
	album.LogoID = 2

	createdAlbum, err := p.albumService.Create(album)

	if err != nil {
		return err
	}

	return c.JSON(200, createdAlbum)
}

func (p *albumsController) Update(c echo.Context) error {
	albumData := &models.Album{}

	if err := c.Bind(albumData); err != nil {
		fmt.Println(err)
		return err
	}

	album, err := p.albumService.GetAlbum(c.Param("id"))

	if err != nil {
		return errors.NotFoundError()
	}

	user, err := p.authService.GetUserFromToken(c)

	if err != nil {
		return err
	}

	if album.AuthorID != user.ID {
		return errors.UnauthorizedError()
	}

	album.Name = albumData.Name

	updatedAlbum, err := p.albumService.Update(album)

	if err != nil {
		return err
	}

	return c.JSON(200, updatedAlbum)
}

func (p *albumsController) GetUserAlbums(c echo.Context) error {
	authorID := c.Param("author")

	albums, err := p.albumService.GetUserAlbums(authorID)

	if err != nil {
		return err
	}

	return c.JSON(200, albums)
}

func (p *albumsController) GetAlbum(c echo.Context) error {
	albumID := c.Param("id")

	album, err := p.albumService.GetAlbum(albumID)

	if err != nil {
		return errors.NotFoundError()
	}

	return c.JSON(200, album)
}

func (p *albumsController) UpdateAlbumLogo(c echo.Context) error {
	albumID := c.Param("album")
	user, err := p.authService.GetUserFromToken(c)

	if err != nil {
		return errors.UnauthorizedError()
	}

	album, err := p.albumService.GetAlbum(albumID)

	if err != nil {
		return errors.NotFoundError()
	}

	if album.AuthorID != user.ID {
		return errors.UnauthorizedError()
	}

	logoFormFile, err := c.FormFile("logo")

	if err != nil {
		return err
	}

	updatedAlbum, err := p.albumService.UploadAlbumLogo(album, logoFormFile)

	if err != nil {
		return err
	}

	return c.JSON(200, updatedAlbum)
}
