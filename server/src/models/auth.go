package models

import "github.com/golang-jwt/jwt/v4"

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
	GetTokens(user *User) (*Tokens, error)
	RefreshTokens(user *User) (*Tokens, error)
	LogoutUser(user *User) error
}
