package middlewares

import (
	"fmt"

	"github.com/Jack-Music-Streaming/server/src/models"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := echojwt.Config{
		SigningKey: []byte(viper.Get("ACCESS_KEY").(string)),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			fmt.Println(c.Get("user"))
			return new(models.JWTClaims)
		},
		ErrorHandler: func(c echo.Context, err error) error {
			fmt.Println(err)
			return c.String(403, "TOKEN_ERROR")
		},
	}

	return echojwt.WithConfig(config)
}

func RefreshMiddleware() echo.MiddlewareFunc {
	config := echojwt.Config{
		SigningKey: []byte(viper.Get("REFRESH_KEY").(string)),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.JWTClaims)
		},
	}

	return echojwt.WithConfig(config)
}
