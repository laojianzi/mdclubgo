package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// NotFound 404 not found handle
func NotFound(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"code":    100004,
		"message": "接口不存在",
	})
}
