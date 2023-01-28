package controllers

import (
	"fmt"
	"net/http"

	"github.com/Jack-Music-Streaming/server/src/errors"
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

	u := models.User{}

	if err := c.Bind(&u); err != nil {
		fmt.Println(err)
		return errors.BadRequest(c)
	}

	a.authService.EncryptPassword(&u)

	user, err := a.userService.Create(u)

	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "User Already Exists")
	}

	tokens, err := a.authService.GetTokens(user)

	if err != nil {
		return errors.ServerError(c)
	}

	a.authService.SetTokens(c, tokens)

	return c.JSON(http.StatusCreated, map[string]string{
		"name":  user.Name,
		"email": user.Email,
	})
}
