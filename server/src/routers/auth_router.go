package routers

import (
	"github.com/Jack-Music-Streaming/server/src/controllers"
	"github.com/Jack-Music-Streaming/server/src/middlewares"
	"github.com/labstack/echo/v4"
)

func AuthRouter(api *echo.Group) {
	authController := controllers.NewAuthController(userService, authService, profileService)

	auth := api.Group("/auth")

	auth.POST("/signUp", authController.SignUp)
	auth.POST("/signIn", authController.SignIn)

	auth.Use(middlewares.JWTMiddleware())

	auth.GET("/profile", authController.GetProfile)
	auth.PUT("/update", authController.UpdateUser)
	auth.PUT("/logo", authController.UpdateUserLogo)
	auth.GET("/profile/full", authController.GetFullProfile)

	auth.Use(middlewares.RefreshMiddleware())

	auth.POST("/refresh", authController.RefreshTokens)
}
