package errors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func BadRequest(c echo.Context) error {
	return echo.NewHTTPError(http.StatusBadRequest, "bad request")
}

func ServerError(c echo.Context) error {
	return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
}
