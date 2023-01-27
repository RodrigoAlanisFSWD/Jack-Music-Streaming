package controllers

import (
	"net/http"

	"github.com/Jack-Music-Streaming/server/src/models"
	"github.com/labstack/echo/v4"
)

type authController struct {
	userService models.UserService
}

type AuthController interface {
	SignIn(c echo.Context) error
}

func NewAuthController(s models.UserService) AuthController {
	return &authController{
		userService: s,
	}
}

func (a *authController) SignIn(c echo.Context) error {
	return c.String(http.StatusOK, "SignIn")
}
