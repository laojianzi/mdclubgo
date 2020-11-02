package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	"github.com/laojianzi/mdclubgo/log"
)

// AllowSkipLimiter set ip/domain for skip limiter handle
var AllowSkipLimiter = []string{"127.0.0.1", "localhost"}

// LimitReached is called when a request hits the limit
var LimitReached = func(ctx *fiber.Ctx) error {
	log.Debug("ip '%s' too many requests", ctx.IP())

	return ctx.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
		"code":    100002,
		"message": "IP 请求超过上限",
	})
}

// Limiter limiter settings for fiber handler
func Limiter() fiber.Handler {
	limiter.ConfigDefault.Next = func(c *fiber.Ctx) bool {
		return skipLimiter(c.IP())
	}
	limiter.ConfigDefault.Max = 20
	limiter.ConfigDefault.LimitReached = LimitReached

	return limiter.New()
}

var skipKey = make(map[string]int)

func skipLimiter(ip string) bool {
	if skipKey == nil {
		for k, v := range AllowSkipLimiter {
			skipKey[v] = k
		}
	}

	_, ok := skipKey[ip]
	return ok
}
