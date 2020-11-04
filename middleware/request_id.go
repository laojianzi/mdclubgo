package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// RequestID request id settings for echo handler
func RequestID() echo.MiddlewareFunc {
	return middleware.RequestID()
}

// RequestIDFromCtx get request id from current echo.Context
func RequestIDFromCtx(ctx echo.Context) string {
	req := ctx.Request()
	res := ctx.Response()

	id := req.Header.Get(echo.HeaderXRequestID)
	if id == "" {
		id = res.Header().Get(echo.HeaderXRequestID)
	}

	return id
}
