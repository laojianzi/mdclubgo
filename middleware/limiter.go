package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/laojianzi/mdclubgo/internal/exception"
	"github.com/laojianzi/mdclubgo/limiter"
)

// AllowSkipLimiter set ip/domain for skip limiter handle
var AllowSkipLimiter = []string{"127.0.0.1", "localhost"}

// Limiter limiter settings for echo handler
func Limiter() echo.MiddlewareFunc {
	// 每 1 秒重置，最大 20 个令牌
	ipLimiter := limiter.NewIPRateLimiter(1, 20)
	userLimiter := limiter.NewUserRateLimiter(1, 20)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ip := ctx.RealIP()
			if skipLimiter(ip) {
				return next(ctx)
			}

			if !ipLimiter.GetLimiter(ip).Allow() {
				return echo.NewHTTPError(http.StatusTooManyRequests)
			}

			if !userLimiter.GetLimiter("user key").Allow() {
				return echo.NewHTTPError(http.StatusTooManyRequests).SetInternal(exception.ErrUserTooManyRequests)
			}

			return next(ctx)
		}
	}
}

var skipKey = make(map[string]int)

func skipLimiter(ip string) bool {
	if len(skipKey) != len(AllowSkipLimiter) {
		for k, v := range AllowSkipLimiter {
			skipKey[v] = k
		}
	}

	_, ok := skipKey[ip]
	return ok
}
