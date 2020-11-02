package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/laojianzi/mdclubgo/conf"
	"github.com/laojianzi/mdclubgo/log"
)

// AllowHeaders CORS headers Access-Control-Allow-Methods
var AllowHeaders = []string{
	fiber.HeaderContentType,
	fiber.HeaderAccessControlAllowHeaders,
	fiber.HeaderAccessControlExposeHeaders,
	fiber.HeaderAuthorization,
	fiber.HeaderXRequestedWith,
	fiber.HeaderXRequestID,
	fiber.HeaderUserAgent,
}

// ExposeHeaders CORS headers Access-Control-Expose-Headers
var ExposeHeaders = []string{
	fiber.HeaderCacheControl,
	fiber.HeaderContentLength,
	fiber.HeaderExpires,
	fiber.HeaderLastModified,
	fiber.HeaderPragma,
	fiber.HeaderContentType,
	fiber.HeaderAccessControlAllowHeaders,
	fiber.HeaderAccessControlExposeHeaders,
	fiber.HeaderAuthorization,
	fiber.HeaderXRequestedWith,
	fiber.HeaderXRequestID,
	fiber.HeaderUserAgent,
}

// CORS cors settings for fiber handler
func CORS() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		cors.ConfigDefault.MaxAge = 3600
		cors.ConfigDefault.AllowHeaders = strings.Join(AllowHeaders, ", ")
		cors.ConfigDefault.ExposeHeaders = strings.Join(ExposeHeaders, ", ")
		if conf.Server.AccessControlAllowOrigin != "" {
			cors.ConfigDefault.AllowOrigins = conf.Server.AccessControlAllowOrigin
			cors.ConfigDefault.AllowCredentials = true
		}

		err := cors.New()(ctx)
		if err != nil {
			return err
		}

		origin := ctx.Get(fiber.HeaderOrigin)
		accessControlAllowOrigin := ctx.Get(fiber.HeaderAccessControlAllowOrigin)
		if origin != accessControlAllowOrigin {
			log.Debug("cors %s is '%s' but %s is '%s'", fiber.HeaderOrigin, origin,
				fiber.HeaderAccessControlAllowOrigin, accessControlAllowOrigin)
		}

		return nil
	}
}
