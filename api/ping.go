package api

import "github.com/gofiber/fiber/v2"

// Ping for server status check
func Ping(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]string{"version": "no version"})
}
