package routers

import (
	"github.com/Jack-Music-Streaming/server/src/controllers"
	"github.com/Jack-Music-Streaming/server/src/database"
	"github.com/Jack-Music-Streaming/server/src/middlewares"
	"github.com/Jack-Music-Streaming/server/src/repository"
	"github.com/Jack-Music-Streaming/server/src/services"
	"github.com/labstack/echo/v4"
)

func AuthRouter(api *echo.Group) {
	jwtRepository := repository.NewJWTRepository(database.DB)
	authService := services.NewAuthService(jwtRepository, userService, profileService)

	authController := controllers.NewAuthController(userService, authService)

	auth := api.Group("/auth")

	auth.POST("/signUp", authController.SignUp)
	auth.POST("/signIn", authController.SignIn)

	auth.Use(middlewares.JWTMiddleware())

	auth.GET("/profile", authController.GetProfile)
	auth.PUT("/update", authController.UpdateUser)

	auth.Use(middlewares.RefreshMiddleware())

	auth.POST("/refresh", authController.RefreshTokens)
}
