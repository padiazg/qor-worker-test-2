package static

import (
	"net/http"
	"strings"

	"github.com/padiazg/qor-worker-test/config/application"
)

// Config home config struct
type Config struct {
	Prefixs []string
	Handler http.Handler
}

// App home app
type App struct {
	Config *Config
}

// New new home app
func New(config *Config) *App {
	return &App{Config: config}
}

// ConfigureApplication configure application
func (app App) ConfigureApplication(application *application.Application) {
	for _, prefix := range app.Config.Prefixs {
		application.Router.Mount("/"+strings.TrimPrefix(prefix, "/"), app.Config.Handler)
	}
}
