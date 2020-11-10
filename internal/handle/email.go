package handle

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/laojianzi/mdclubgo/conf"
	"github.com/laojianzi/mdclubgo/internal/email"
	"github.com/laojianzi/mdclubgo/internal/email/validator"
	"github.com/laojianzi/mdclubgo/log"
)

// RegisterEmail send register email
func RegisterEmail(ctx echo.Context) error {
	var form email.Form
	if err := ctx.Bind(&form); err != nil {
		log.Error(fmt.Errorf("register email form bind: %w", err).Error())
		return echo.ErrBadRequest
	}

	if validErr := form.Validate(); validErr != nil {
		return validErr
	}

	err := validator.Send(form.Email, conf.App.Name, validator.GenerateCode(form.Email))
	if err != nil {
		log.Error(fmt.Errorf("register email send: %w", err).Error())
		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusOK, new(email.Present).Format())
}
