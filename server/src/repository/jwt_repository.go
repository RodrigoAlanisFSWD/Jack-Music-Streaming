package repository

import (
	"time"

	"github.com/Jack-Music-Streaming/server/src/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type jwtRepository struct {
}

func NewJWTRepository() models.JWTRepository {
	return &jwtRepository{}
}

func (s *jwtRepository) CreateAccessToken(user *models.User) (string, error) {
	claims := &models.JWTClaims{
		User: user.ID,
		Role: user.RoleID,
		Plan: user.PlanID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(viper.Get("ACCESS_KEY").(string)))

	if err != nil {
		return "", err
	}

	return signed, nil
}

func (s *jwtRepository) CreateRefreshToken(user *models.User) (string, error) {
	claims := &models.JWTClaims{
		User: user.ID,
		Role: user.RoleID,
		Plan: user.PlanID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 48)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(viper.Get("REFRESH_KEY").(string)))

	if err != nil {
		return "", err
	}

	return signed, nil
}
