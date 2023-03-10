package repository

import (
	"errors"
	"time"

	"github.com/Jack-Music-Streaming/server/src/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type jwtRepository struct {
	DB *gorm.DB
}

func NewJWTRepository(db *gorm.DB) models.JWTRepository {
	return &jwtRepository{
		DB: db,
	}
}

func (j *jwtRepository) CreateAccessToken(user *models.User) (string, error) {
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

func (j *jwtRepository) CreateRefreshToken(user *models.User) (string, error) {
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

func (j *jwtRepository) RegisterRefreshToken(token *models.RefreshToken) (*models.RefreshToken, error) {
	err := j.DB.Create(token).Error

	if err != nil {
		return token, err
	}

	return token, err
}

func (j *jwtRepository) UpdateRefreshToken(new *models.RefreshToken, old string) (*models.RefreshToken, error) {
	var exists models.RefreshToken

	err := j.DB.Where("token = ?", old).First(&exists).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return j.RegisterRefreshToken(new)
	} else if err != nil {
		return &exists, err
	}

	j.DB.Where("token = ?", old).Delete(&exists)

	return j.RegisterRefreshToken(new)
}

func (j *jwtRepository) VerifyRefreshToken(token string) (*models.RefreshToken, error) {
	var refresh *models.RefreshToken

	err := j.DB.Where("token = ?", token).First(&refresh).Error

	return refresh, err
}
