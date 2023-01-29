package models

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type JWTClaims struct {
	User uint
	Role uint
	Plan uint
	jwt.RegisteredClaims
}

type JWTRepository interface {
	CreateAccessToken(user *User) (string, error)
	CreateRefreshToken(user *User) (string, error)
}

type AuthService interface {
	SetTokens(c echo.Context, tokens *Tokens)
	GetTokens(user *User) (*Tokens, error)
	RefreshTokens(user *User) (*Tokens, error)
	LogoutUser(user *User) error
	EncryptPassword(user *User) error
	ComparePasswords(user *User, validUser *User) error
	ValidateUser(user *User) (*User, error)
	GetUserFromToken(c echo.Context) (*User, error)
	UpdateUserScope(c echo.Context, u *User) (*User, error)
	GetUserScope(c echo.Context) (*Scope, error)
}
