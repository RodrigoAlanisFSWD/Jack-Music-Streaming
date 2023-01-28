package services

import (
	"net/http"
	"time"

	"github.com/Jack-Music-Streaming/server/src/models"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	jwtRepository models.JWTRepository
}

var auth *authService

func NewAuthService(r models.JWTRepository) models.AuthService {

	once.Do(func() {
		auth = &authService{
			jwtRepository: r,
		}
	})

	return auth
}

func (a *authService) SetTokens(c echo.Context, tokens *models.Tokens) {
	c.SetCookie(&http.Cookie{
		Name:    "@jack/access_token",
		Value:   tokens.AccessToken,
		Expires: time.Now().Add(30 * time.Minute),
	})

	c.SetCookie(&http.Cookie{
		Name:    "@jack/refresh_token",
		Value:   tokens.RefreshToken,
		Expires: time.Now().Add(48 * time.Hour),
	})
}

func (a *authService) GetTokens(user *models.User) (*models.Tokens, error) {
	accessToken, err := a.jwtRepository.CreateAccessToken(user)

	if err != nil {
		return &models.Tokens{}, err
	}

	refreshToken, err := a.jwtRepository.CreateRefreshToken(user)

	if err != nil {
		return &models.Tokens{}, err
	}

	return &models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *authService) RefreshTokens(user *models.User) (*models.Tokens, error) {
	return &models.Tokens{}, nil
}

func (a *authService) LogoutUser(user *models.User) error {
	return nil
}

func (a *authService) EncryptPassword(user *models.User) error {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)

	if err != nil {
		return err
	}

	user.Password = string(passwordBytes)

	return nil
}
