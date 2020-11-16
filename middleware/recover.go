package middleware

import (
	"fmt"
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/laojianzi/mdclubgo/log"
)

// Recover recover settings for echo handler
func Recover() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			defer RecoverHandle(ctx)

			return next(ctx)
		}
	}
}

// RecoverHandle recover all panic and print error log
func RecoverHandle(ctx echo.Context) {
	if r := recover(); r != nil {
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
		}

		stack := make([]byte, middleware.DefaultRecoverConfig.StackSize)
		length := runtime.
			Stack(stack, !middleware.DefaultRecoverConfig.DisableStackAll)
		if !middleware.DefaultRecoverConfig.DisablePrintStack {
			msg := fmt.
				Sprintf("[PANIC RECOVER] %v [STACK] %s", err, stack[:length])
			if requestID := RequestIDFromCtx(ctx); requestID != "" {
				msg = fmt.Sprintf("[REQUEST-ID] %s\t%s", requestID, msg)
			}

			log.Error(msg)
		}

		if ctx != nil {
			ctx.Error(err)
		}
	}
}
