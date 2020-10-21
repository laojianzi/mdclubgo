package api

import (
	"net/http"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

// App api server
type App struct {
	server *fiber.App
	once   sync.Once
}

var app = new(App)

// Server return a api.App
func Server() *App {
	app.once.Do(func() {
		app.server = fiber.New(fiber.Config{
			ServerHeader:         "MDClubGo",
			ReadTimeout:          5 * time.Second,
			WriteTimeout:         10 * time.Second,
			CompressedFileSuffix: ".mdclubgo.gz",
		})

		app.route()
	})

	return app
}

// Add add handler to api server
func (a *App) Add(method string, path string, handlers ...fiber.Handler) fiber.Router {
	return a.server.Add(method, path, handlers...)
}

// Start listen api server
func (a *App) Start(addr string) error {
	return a.server.Listen(addr)
}

// Test send test request to api server
func (a *App) Test(req *http.Request, msTimeout ...int) (resp *http.Response, err error) {
	return a.server.Test(req, msTimeout...)
}
