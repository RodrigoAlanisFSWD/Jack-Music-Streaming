package errors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func BadRequest() error {
	return echo.NewHTTPError(http.StatusBadRequest, "bad request")
}

func ServerError() error {
	return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
}

func UnauthorizedError() error {
	return echo.NewHTTPError(http.StatusUnauthorized, "user unauthorized")
}

func NotFoundError() error {
	return echo.NewHTTPError(http.StatusNotFound, "not found")
}
