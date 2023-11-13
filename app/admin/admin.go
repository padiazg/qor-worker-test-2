package admin

import (
	"encoding/json"
	"fmt"
	"html/template"
	"reflect"

	"github.com/go-chi/chi/v5"
	"github.com/padiazg/qor-worker-test/config/application"
	"github.com/qor/admin"
	"github.com/qor/assetfs"
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
	app.ConfigureAdmin(application)
	app.ConfigureFuncMaps(application)

	application.Router.Route(app.Config.Prefix, func(r chi.Router) {
		// r.Use(func(next http.Handler) http.Handler { return mw.AdminAuthedMiddleware(next, application) })
		r.Mount("/", application.Admin.NewServeMux(app.Config.Prefix))
	})

}

// ConfigureAdmin configure admin interface
func (app App) ConfigureAdmin(application *application.Application) {
	fs := assetfs.AssetFS()
	fs.RegisterPath("app/views")

	application.Admin = admin.New(&admin.AdminConfig{
		DB:      application.DB,
		AssetFS: fs,
	})

}

func (app App) ConfigureFuncMaps(application *application.Application) {
	application.Admin.RegisterFuncMap("marshal", func(v interface{}) template.JS {
		var res string

		if a, err := json.Marshal(v); err != nil {
			res = fmt.Sprintf("error marshalling: %+v", err)
		} else {
			res = string(a)
		}

		return template.JS(res)
	})

	application.Admin.RegisterFuncMap("avail", func(name string, data interface{}) bool {
		v := reflect.ValueOf(data)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if v.Kind() != reflect.Struct {
			return false
		}
		return v.FieldByName(name).IsValid()
	})

}
