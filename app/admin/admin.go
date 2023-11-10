package admin

import (
	"github.com/go-chi/chi/v5"
	"github.com/padiazg/qor-worker-test/config/application"
	"github.com/qor/admin"
)

type Config struct {
	Prefix string
}

type App struct {
	*Config
}

func New(config *Config) *App {
	if config.Prefix == "" {
		config.Prefix = "/admin"
	}

	return &App{Config: config}
}

func (app App) ConfigureApplication(application *application.Application) {
	// assign new Admin insantance
	application.Admin = admin.New(&admin.AdminConfig{
		DB: application.DB,
	})

	application.Router.Route(app.Config.Prefix, func(r chi.Router) {
		// r.Use(func(next http.Handler) http.Handler { return mw.AdminAuthedMiddleware(next, application) })
		r.Mount("/", application.Admin.NewServeMux(app.Config.Prefix))
	})

}

// ConfigureAdmin configure admin interface
func (app App) ConfigureAdmin(application *application.Application) {
	// Add Dashboard
	// application.Admin.AddResource(&asset_manager.AssetManager{}, &admin.Config{Invisible: true})
}

func (app App) ConfigureFuncMaps(application *application.Application) {
	// // Frontend FuncMaps
	// fm.GlobalFuncs.Add(map[string]interface{}{
	// 	"AuthURL": application.Auth.AuthURL,
	// 	"Flashes": application.Auth.Flashes,
	// 	"IsFullstoryEnabled": func() bool {
	// 		// TODO: get this flag from the license
	// 		return application.AllSettings.Settings.FullstoryEnabled
	// 	},
	// })
}
