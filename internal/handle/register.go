package handle

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/laojianzi/mdclubgo/cache"
	"github.com/laojianzi/mdclubgo/conf"
	"github.com/laojianzi/mdclubgo/db"
	"github.com/laojianzi/mdclubgo/internal/avatar"
	"github.com/laojianzi/mdclubgo/internal/database"
	"github.com/laojianzi/mdclubgo/internal/exception"
	"github.com/laojianzi/mdclubgo/internal/register"
	"github.com/laojianzi/mdclubgo/internal/web"
	"github.com/laojianzi/mdclubgo/log"
)

// Register create a user
func Register(ctx echo.Context) error {
	var form register.Form
	if err := ctx.Bind(&form); err != nil {
		log.Error(fmt.Errorf("register bind form: %w", err).Error())
		return exception.ErrBadRequest
	}

	user, err := database.Register(db.Instance(), ctx, form.Username, form.Email, form.Password)
	if err != nil {
		log.Error(fmt.Errorf("register user and save to db: %w", err).Error())
		return exception.ErrInternalServerError
	}

	user, err = avatar.DeleteUserAvatar(user.ID)
	if err != nil {
		log.Error(fmt.Errorf("delete user avatar: %w", err).Error())
		return exception.ErrInternalServerError
	}

	user.Password = ""
	err = register.Send(user.Email, conf.App.Name, user.Username)
	if err != nil {
		log.Error(fmt.Errorf("send register email: %w", err).Error())
		return exception.ErrInternalServerError
	}

	// delete user throttle cache
	key := fmt.Sprintf("throttle_create_token_%s", web.IPSign(ctx))
	err = cache.Delete(key)
	if err != nil {
		log.Error(fmt.Errorf("delete cache %s: %w", key, err).Error())
		return exception.ErrInternalServerError
	}

	return ctx.JSON(http.StatusOK, register.NewPresent(ctx, user).Format())
}
