package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// RequestID request id settings for fiber handler
func RequestID() fiber.Handler {
	return requestid.New()
}

// RequestIDFromCtx get request id from current fiber.Ctx
func RequestIDFromCtx(ctx *fiber.Ctx) string {
	value := ctx.Locals(requestid.ConfigDefault.ContextKey)
	if value == nil {
		return ""
	}

	requestID, _ := value.(string)
	return requestID
}
