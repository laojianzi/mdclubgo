package handle

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/laojianzi/mdclubgo/internal/captcha"
	"github.com/laojianzi/mdclubgo/log"
)

// NewCaptcha create a captcha
func NewCaptcha(ctx echo.Context) error {
	present, err := captcha.Generate(100, 36)
	if err != nil {
		log.Error(fmt.Errorf("new captcha: %w", err).Error())
		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusOK, present.Format())
}
