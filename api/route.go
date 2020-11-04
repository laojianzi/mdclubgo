package api

import (
	"github.com/laojianzi/mdclubgo/internal/ping"
)

func (app *App) route() {
	app.server.GET("/ping", ping.Ping)
}
