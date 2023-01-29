package controllers

import (
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
	GetProfile(c echo.Context) error
}

func NewAuthController(s models.UserService, auth models.AuthService) AuthController {
	return &authController{
		userService: s,
		authService: auth,
	}
}

func (a *authController) SignIn(c echo.Context) error {
	u := &models.User{}

	if err := c.Bind(&u); err != nil {
		return errors.BadRequest()
	}

	validUser, err := a.authService.ValidateUser(u)

	if err != nil {
		return errors.UnauthorizedError()
	}

	tokens, err := a.authService.GetTokens(validUser)

	if err != nil {
		return errors.ServerError()
	}

	a.authService.SetTokens(c, tokens)

	return c.JSON(http.StatusOK, map[string]string{
		"name":  validUser.Name,
		"email": validUser.Email,
	})
}

func (a *authController) SignUp(c echo.Context) error {
	u := &models.User{}

	if err := c.Bind(&u); err != nil {
		return errors.BadRequest()
	}

	a.authService.EncryptPassword(u)

	user, err := a.userService.Create(u)

	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "User Already Exists")
	}

	tokens, err := a.authService.GetTokens(user)

	if err != nil {
		return errors.ServerError()
	}

	a.authService.SetTokens(c, tokens)

	return c.JSON(http.StatusOK, map[string]string{
		"name":  user.Name,
		"email": user.Email,
	})
}

func (a *authController) GetProfile(c echo.Context) error {
	user, err := a.authService.GetUserFromToken(c)

	if err != nil {
		return errors.UnauthorizedError()
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"created_at": user.CreatedAt,
	})
}
