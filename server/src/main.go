package main

import (
	"log"
	"net/http"

	"github.com/Jack-Music-Streaming/server/src/database"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	if err := database.InitDB(); err != nil {
		log.Panic(err)
		return
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
