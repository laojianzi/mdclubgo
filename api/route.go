package api

import (
	"github.com/laojianzi/mdclubgo/internal/handle"
)

func (app *App) route() {
	app.server.GET("/ping", handle.Ping)
	app.server.POST("/api/captchas", handle.NewCaptcha)
}
