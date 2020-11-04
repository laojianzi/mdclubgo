package middleware

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/color"

	"github.com/laojianzi/mdclubgo/log"
)

// Logger logger settings for echo handler
func Logger() echo.MiddlewareFunc {
	middleware.Logger()
	colorer := color.New()
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			res := c.Response()
			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}

			stop := time.Now()
			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}

			printFunc := log.Info
			status := colorer.Green(res.Status)
			switch {
			case res.Status >= 500:
				status = colorer.Red(res.Status)
				printFunc = log.Error
			case res.Status >= 400:
				status = colorer.Yellow(res.Status)
				printFunc = log.Warn
			case res.Status >= 300:
				status = colorer.Cyan(res.Status)
			}

			msg := fmt.Sprintf("[REQUEST-ID] %s\t[REMOTE-IP] %s\t[HOST] %s\t[METHOD] %s\t[URI] %s\t[STATUS] %s\t[LATENCY] %s",
				id, c.RealIP(), req.Host, req.Method, req.RequestURI, status, stop.Sub(start).String())

			if err != nil {
				// Error may contain invalid JSON e.g. `"`
				b, _ := json.Marshal(err.Error())
				b = b[1 : len(b)-1]
				msg = fmt.Sprintf("%s\t[ERROR] %s", msg, string(b))
			}

			printFunc(msg)
			return
		}
	}
}
