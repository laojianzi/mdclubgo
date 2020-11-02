package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/laojianzi/mdclubgo/conf"
	"github.com/laojianzi/mdclubgo/log"
)

// Recover recover settings for fiber handler
func Recover() fiber.Handler {
	// Return new handler
	return func(c *fiber.Ctx) (err error) {
		// Catch panics
		defer func() {
			if r := recover(); r != nil {
				if conf.App.Debug {
					log.Error("panic recovered: \n%s\n%s\n%s", c.Request().String(), r, zap.Stack("").String)
				} else {
					log.Error("panic recovered: \n%s\n%s", r, zap.Stack("").String)
				}

				err = c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"code":    100000,
					"message": "服务器错误",
				})
			}
		}()

		// Return err if exist, else move to next handler
		return c.Next()
	}
}
