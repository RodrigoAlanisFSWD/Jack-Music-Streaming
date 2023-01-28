package main

import (
	"log"
	"net/http"

	"github.com/Jack-Music-Streaming/server/src/config"
	"github.com/Jack-Music-Streaming/server/src/database"
	"github.com/Jack-Music-Streaming/server/src/routers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	config.GetConfig()

	if err := database.InitDB(); err != nil {
		log.Panic(err)
		return
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	api := e.Group("/api")

	routers.AuthRouter(api)

	e.Logger.Fatal(e.Start(":8080"))
}
