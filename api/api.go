package api

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	echolog "github.com/labstack/gommon/log"

	"github.com/laojianzi/mdclubgo/conf"
	"github.com/laojianzi/mdclubgo/internal/storage"
	"github.com/laojianzi/mdclubgo/internal/storage/local"
	"github.com/laojianzi/mdclubgo/log"
	"github.com/laojianzi/mdclubgo/middleware"
)

var instance = new(App)

// App api server
type App struct {
	server *echo.Echo
	once   sync.Once
}

// Server return a api.App
func Server() *App {
	instance.once.Do(func() {
		instance.server = echo.New()

		if conf.App.Name != "" {
			instance.server.Logger.SetPrefix(conf.App.Name)
		}

		instance.server.Logger.SetOutput(log.Output())
		level := echolog.INFO
		if conf.App.Debug {
			level = echolog.DEBUG
		}

		instance.server.Logger.SetLevel(level)
		instance.server.Debug = conf.App.Debug
		instance.server.HideBanner = false
		instance.server.Server.ReadTimeout = 5 * time.Second
		instance.server.Server.WriteTimeout = 10 * time.Second
		instance.server.HTTPErrorHandler = middleware.ErrorHandler
		instance.server.Pre(middleware.RemoveTrailingSlash())

		instance.server.Use(
			middleware.RequestID(),
			middleware.Recover(),
			middleware.Logger(),
			middleware.CORS(),
			middleware.Limiter(),
		)

		instance.route()

		// static url
		if conf.Storage.Type == storage.Local {
			root := conf.StorageLocal.URL
			if root == "" {
				root = local.DefaultPathPrefix
			}

			instance.server.Static(fmt.Sprintf("/%s", conf.Server.SiteStaticURL), root)
		}

		instance.server.Static("/", "public")
	})

	return instance
}

// Start listen api server
func (app *App) Start(addr string) error {
	if conf.Server.HTTPSEnable {
		if err := instance.server.StartTLS(addr, conf.Server.CertFile, conf.Server.KeyFile); err != nil {
			log.Fatal("can't read cert file and key file")
		}
	}

	return app.server.Start(addr)
}

// Shutdown close api server
func (app *App) Shutdown(ctx context.Context) error {
	return app.server.Shutdown(ctx)
}

// NewContext uses echo.Echo.NewContext
func (app *App) NewContext(r *http.Request, w http.ResponseWriter) echo.Context {
	return app.server.NewContext(r, w)
}
