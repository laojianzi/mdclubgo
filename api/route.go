package api

func (app *App) route() {
	app.server.Get("/ping", Ping)
}
