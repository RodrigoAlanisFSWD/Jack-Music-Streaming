package services

import (
	"github.com/Jack-Music-Streaming/server/src/models"
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
