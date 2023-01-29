package services

import (
	"net/http"
	"sync"
	"time"

	"github.com/Jack-Music-Streaming/server/src/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	jwtRepository   models.JWTRepository
	scopeRepository models.ScopeRepository
	userService     models.UserService
}

var authInstance *authService

var authOnce sync.Once

func NewAuthService(r models.JWTRepository, u models.UserService, s models.ScopeRepository) models.AuthService {

	authOnce.Do(func() {
		authInstance = &authService{
			jwtRepository:   r,
			userService:     u,
			scopeRepository: s,
		}
	})

	return authInstance
}

func (a *authService) SetTokens(c echo.Context, tokens *models.Tokens) {
	cookie := new(http.Cookie)

	cookie.Name = "jack_access_token"
	cookie.Value = tokens.AccessToken
	cookie.Expires = time.Now().Add(30 * time.Minute)
	cookie.Path = "/"
	cookie.Domain = "localhost"
	c.SetCookie(cookie)

	cookie.Name = "jack_refresh_token"
	cookie.Value = tokens.RefreshToken
	cookie.Expires = time.Now().Add(48 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = "localhost"
	c.SetCookie(cookie)
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

func (a *authService) ComparePasswords(user *models.User, validUser *models.User) error {
	return bcrypt.CompareHashAndPassword([]byte(validUser.Password), []byte(user.Password))
}

func (a *authService) ValidateUser(user *models.User) (*models.User, error) {
	validUser, err := a.userService.GetUserByEmail(user.Email)

	if err != nil {
		return validUser, err
	}

	err = a.ComparePasswords(user, validUser)

	if err != nil {
		return validUser, err
	}

	return validUser, err
}

func (a *authService) GetUserFromToken(c echo.Context) (*models.User, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JWTClaims)
	userID := claims.User

	return a.userService.GetUserByID(userID)
}

func (a *authService) UpdateUserScope(c echo.Context, u *models.User) (*models.User, error) {
	user, err := a.GetUserFromToken(c)

	if err != nil {
		return &models.User{}, err
	}

	user.RoleID = u.RoleID
	user.PlanID = u.PlanID

	return a.userService.Update(user)
}

func (a *authService) GetUserScope(c echo.Context) (*models.Scope, error) {
	user, err := a.GetUserFromToken(c)

	if err != nil {
		return &models.Scope{}, err
	}

	return a.scopeRepository.FindScope(user.RoleID, user.PlanID)
}
