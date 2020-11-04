package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/laojianzi/mdclubgo/conf"
)

// AllowHeaders CORS headers Access-Control-Allow-Methods
var AllowHeaders = []string{
	echo.HeaderContentType,
	echo.HeaderAccessControlAllowHeaders,
	echo.HeaderAccessControlExposeHeaders,
	echo.HeaderAuthorization,
	echo.HeaderXRequestedWith,
	echo.HeaderXRequestID,
}

// ExposeHeaders CORS headers Access-Control-Expose-Headers
var ExposeHeaders = []string{
	echo.HeaderContentLength,
	echo.HeaderLastModified,
	echo.HeaderContentType,
	echo.HeaderAccessControlAllowHeaders,
	echo.HeaderAccessControlExposeHeaders,
	echo.HeaderAuthorization,
	echo.HeaderXRequestedWith,
	echo.HeaderXRequestID,
}

// CORS cors settings for echo handler
func CORS() echo.MiddlewareFunc {
	cfg := middleware.CORSConfig{
		AllowHeaders:  AllowHeaders,
		ExposeHeaders: ExposeHeaders,
		MaxAge:        3600,
	}

	if conf.Server.AccessControlAllowOrigin != "" {
		cfg.AllowOrigins = strings.Split(strings.TrimSpace(conf.Server.AccessControlAllowOrigin), ",")
		cfg.AllowCredentials = true
	}

	return middleware.CORSWithConfig(cfg)
}
