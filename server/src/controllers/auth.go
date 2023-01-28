package controllers

import (
	"net/http"

	"github.com/Jack-Music-Streaming/server/src/models"
	"github.com/labstack/echo/v4"
)

type authController struct {
	userService models.UserService
	authService models.AuthService
}

type AuthController interface {
	SignIn(c echo.Context) error
	SignUp(c echo.Context) error
}

func NewAuthController(s models.UserService, auth models.AuthService) AuthController {
	return &authController{
		userService: s,
		authService: auth,
	}
}

func (a *authController) SignIn(c echo.Context) error {
	return c.String(http.StatusOK, "SignIn")
}

func (a *authController) SignUp(c echo.Context) error {
	return c.String(http.StatusOK, "SignUp")
}
