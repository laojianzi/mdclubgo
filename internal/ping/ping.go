package ping

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Ping for server status check
func Ping(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{"version": "no version"})
}
