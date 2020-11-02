package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/laojianzi/mdclubgo/conf"
	"github.com/laojianzi/mdclubgo/log"
)

var LoggerTimeFormat = "2006-01-02T15:04:05.000Z0700"

// Logger logger settings for fiber handler
func Logger() fiber.Handler {
	if conf.App.Debug {
		logger.ConfigDefault.Format = "${status} - ${latency} ${method} ${path}\n"
	} else {
		logger.ConfigDefault.Format = `{"STATUS": ${status}, "LATENCY": "${latency}", "METHOD": "${method}", "PATH": "${path}}"}` + "\n"
	}
	logger.ConfigDefault.TimeFormat = LoggerTimeFormat
	logger.ConfigDefault.Output = new(fiberLogger)

	return logger.New(logger.ConfigDefault)
}

type fiberLogger struct{}

// Write fiberLogger implement io.Writer
func (fiberLogger) Write(p []byte) (n int, err error) {
	log.Info(string(p))
	return len(p), nil
}
