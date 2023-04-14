package controllers

import (
	"net/http"

	"github.com/Jack-Music-Streaming/server/src/errors"
	"github.com/Jack-Music-Streaming/server/src/models"
	"github.com/labstack/echo/v4"
)

type authController struct {
	userService    models.UserService
	authService    models.AuthService
	profileService models.ProfileService
}

type AuthController interface {
	SignIn(c echo.Context) error
	SignUp(c echo.Context) error
	GetProfile(c echo.Context) error
	UpdateUser(c echo.Context) error
	RefreshTokens(c echo.Context) error
	UpdateUserLogo(c echo.Context) error
}

func NewAuthController(s models.UserService, auth models.AuthService, p models.ProfileService) AuthController {
	return &authController{
		userService:    s,
		authService:    auth,
		profileService: p,
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

	_, err = a.authService.RegisterRefreshToken(tokens.RefreshToken, validUser)

	if err != nil {
		return errors.ServerError()
	}

	a.authService.SetTokens(c, tokens)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"name":  validUser.Name,
		"email": validUser.Email,
		"role":  validUser.Role,
	})
}

func (a *authController) SignUp(c echo.Context) error {
	u := &models.User{
		RoleID: 1,
	}

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

	_, err = a.authService.RegisterRefreshToken(tokens.RefreshToken, user)

	if err != nil {
		return errors.ServerError()
	}

	a.authService.SetTokens(c, tokens)

	err = a.authService.CreateProfile(user)

	if err != nil {
		return errors.ServerError()
	}

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

	return c.JSON(http.StatusOK, user)
}

func (a *authController) UpdateUser(c echo.Context) error {
	u := &models.User{}

	if err := c.Bind(u); err != nil {
		return errors.BadRequest()
	}

	user, err := a.authService.UpdateUserRole(c, u)

	if err != nil {
		return errors.ServerError()
	}

	tokens, err := a.authService.GetTokens(user)

	if err != nil {
		return errors.ServerError()
	}

	_, err = a.authService.RegisterRefreshToken(tokens.RefreshToken, user)

	if err != nil {
		return errors.ServerError()
	}

	a.authService.SetTokens(c, tokens)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"created_at": user.CreatedAt,
	})
}

func (a *authController) RefreshTokens(c echo.Context) error {
	tokens, err := a.authService.RefreshTokens(c)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, tokens)
}

func (a *authController) UpdateUserLogo(c echo.Context) error {
	user, err := a.authService.GetUserFromToken(c)

	if err != nil {
		return errors.UnauthorizedError()
	}

	logoFormFile, err := c.FormFile("logo")

	if err != nil {
		return err
	}

	_, err = a.profileService.UpdateUserLogo(user, logoFormFile)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
