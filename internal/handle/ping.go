package handle

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/laojianzi/mdclubgo/conf"
)

// Ping for server status check
func Ping(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{"version": conf.App.Version})
}
