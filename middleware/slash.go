package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// RemoveTrailingSlash uses echo middleware.RemoveTrailingSlash
func RemoveTrailingSlash() echo.MiddlewareFunc {
	return middleware.RemoveTrailingSlash()
}
