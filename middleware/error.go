package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/laojianzi/mdclubgo/internal/exception"
	"github.com/laojianzi/mdclubgo/log"
)

// ErrorHandler for echo error handler
func ErrorHandler(err error, ctx echo.Context) {
	mdclubgoErr := errorHandle(err)

	// Send response
	if ctx.Response().Committed {
		return
	}

	if ctx.Request().Method == http.MethodHead {
		code := http.StatusInternalServerError
		if e, ok := err.(*echo.HTTPError); ok {
			code = e.Code
		}

		err = ctx.NoContent(code)
	} else {
		err = ctx.JSON(http.StatusOK, mdclubgoErr)
	}

	if err != nil {
		log.Error("[REQUEST-ID] %s\t[METHOD] %s\t[URI] %s", RequestIDFromCtx(ctx), ctx.Request().Method, ctx.Request().RequestURI)
	}
}

func errorHandle(err error) *exception.MDClubGoError {
	if mdclubgoErr, ok := err.(*exception.MDClubGoError); ok {
		return mdclubgoErr
	}

	e, ok := err.(*echo.HTTPError)
	if !ok {
		e = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	if e.Internal != nil {
		if mdclubgoErr, ok := e.Internal.(*exception.MDClubGoError); ok {
			return mdclubgoErr
		}

		if herr, ok := e.Internal.(*echo.HTTPError); ok {
			return exception.HTTPCodeToMDClubGoError(herr.Code)
		}
	}

	return exception.HTTPCodeToMDClubGoError(e.Code)
}
