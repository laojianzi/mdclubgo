package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/laojianzi/mdclubgo/internal/exception"
	"github.com/laojianzi/mdclubgo/log"
)

// ErrorHandler for echo error handler
func ErrorHandler(err error, ctx echo.Context) {
	e, ok := err.(*echo.HTTPError)
	if !ok {
		e = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	httpCode := e.Code
	content := exception.HTTPCodeToMDClubGoError(e.Code)
	if e.Internal != nil {
		if mdclubgoErr, ok := e.Internal.(*exception.MDClubGoError); ok {
			httpCode = e.Code
			content = mdclubgoErr
		}

		if herr, ok := e.Internal.(*echo.HTTPError); ok {
			httpCode = herr.Code
			content = exception.HTTPCodeToMDClubGoError(herr.Code)
		}
	}

	// Send response
	if ctx.Response().Committed {
		return
	}

	if ctx.Request().Method == http.MethodHead {
		err = ctx.NoContent(e.Code)
	} else {
		err = ctx.JSON(httpCode, content)
	}

	if err != nil {
		log.Error("[REQUEST-ID] %s\t[METHOD] %s\t[URI] %s", RequestIDFromCtx(ctx), ctx.Request().Method, ctx.Request().RequestURI)
	}
}
