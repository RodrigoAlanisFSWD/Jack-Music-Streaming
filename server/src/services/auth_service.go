package services

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/Jack-Music-Streaming/server/src/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	jwtRepository  models.JWTRepository
	userService    models.UserService
	profileService models.ProfileService
}

var authInstance *authService

var authOnce sync.Once

func NewAuthService(r models.JWTRepository, u models.UserService, p models.ProfileService) models.AuthService {

	authOnce.Do(func() {
		authInstance = &authService{
			jwtRepository:  r,
			userService:    u,
			profileService: p,
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

	if err != nil {
		fmt.Println(err)
		return &models.Tokens{}, err
	}

	return &models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *authService) RefreshTokens(c echo.Context) (*models.Tokens, error) {
	refresh := c.Get("user").(*jwt.Token)
	token := refresh.Raw

	user, err := a.GetUserFromToken(c)

	if err != nil {
		return &models.Tokens{}, err
	}

	old, err := a.jwtRepository.VerifyRefreshToken(token)

	if err != nil {
		return &models.Tokens{}, err
	}

	newTokens, err := a.GetTokens(user)

	if err != nil {
		return &models.Tokens{}, err
	}

	_, err = a.UpdateRefreshToken(newTokens.RefreshToken, old.Token)

	if err != nil {
		return &models.Tokens{}, err
	}

	return newTokens, nil
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

	fmt.Println(userID)

	return a.userService.GetUserByID(userID)
}

func (a *authService) UpdateUserRole(c echo.Context, u *models.User) (*models.User, error) {
	user, err := a.GetUserFromToken(c)

	if err != nil {
		return &models.User{}, err
	}

	user.RoleID = u.RoleID
	user.Role = u.Role

	return a.userService.Update(user)
}

func (a *authService) RegisterRefreshToken(token string) (*models.RefreshToken, error) {
	return a.jwtRepository.RegisterRefreshToken(&models.RefreshToken{
		Token: token,
	})
}

func (a *authService) UpdateRefreshToken(new string, old string) (*models.RefreshToken, error) {
	newToken := models.RefreshToken{
		Token: new,
	}

	return a.jwtRepository.UpdateRefreshToken(&newToken, old)
}

func (a *authService) CreateProfile(user *models.User) error {
	_, err := a.profileService.Create(&models.Profile{
		UserID: user.ID,
	})

	if err != nil {
		return err
	}

	return nil
}
