package routers

import (
	"github.com/Jack-Music-Streaming/server/src/controllers"
	"github.com/labstack/echo/v4"
)

func AuthRouter(api echo.Group) {
	authController := controllers.NewAuthController(userService)

	auth := api.Group("/auth")

	auth.POST("/signIn", authController.SignIn)
}
