package api

import (
	"net/http"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/laojianzi/mdclubgo/conf"
)

// App api server
type App struct {
	server *fiber.App
	once   sync.Once
}

var fiberApp = new(App)

var defaultFiberConfig = fiber.Config{
	ServerHeader:         "MDClubGo",
	ReadTimeout:          5 * time.Second,
	WriteTimeout:         10 * time.Second,
	CompressedFileSuffix: ".mdclubgo.gz",
}

// Server return a api.App
func Server() *App {
	fiberApp.once.Do(func() {
		if conf.App.Name != "" {
			defaultFiberConfig.ServerHeader = conf.App.Name
		}

		fiberApp.server = fiber.New(defaultFiberConfig)
		fiberApp.route()
	})

	return fiberApp
}

// Add add handler to api server
func (app *App) Add(method string, path string, handlers ...fiber.Handler) fiber.Router {
	return app.server.Add(method, path, handlers...)
}

// Start listen api server
func (app *App) Start(addr string) error {
	return app.server.Listen(addr)
}

// Test send test request to api server
func (app *App) Test(req *http.Request, msTimeout ...int) (resp *http.Response, err error) {
	return app.server.Test(req, msTimeout...)
}
