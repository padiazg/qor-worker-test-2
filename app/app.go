package app

import (
	"context"
	"log"
	"net/http"

	"github.com/padiazg/qor-worker-test/app/admin"
	"github.com/padiazg/qor-worker-test/app/static"
	"github.com/padiazg/qor-worker-test/app/user"
	snl "github.com/padiazg/qor-worker-test/app/workers/send-newsletter"
	"github.com/padiazg/qor-worker-test/config/application"
	qa "github.com/qor/admin"
)

func Mount(Application *application.Application, ctx context.Context) {
	log.Println("Setting up regular server")

	// connect the DB
	if err := Application.ConnectDB(ctx); err != nil {
		log.Fatalf("connecting db: %+v", err)
	}

	// static
	Application.Use(static.New(&static.Config{
		Prefixs: []string{Application.Config.Settings.Static.Route},
		Handler: http.StripPrefix(
			Application.Config.Settings.Static.Route,
			http.FileServer(http.Dir(Application.Config.Settings.Static.Path)),
		),
	}))

	// Admin
	Application.Use(admin.New(&admin.Config{}))

	// Workers
	Application.Use(snl.New(&snl.Config{
		AdminConfig: &qa.Config{
			Menu: []string{"Workers"},
			Name: "Send Newsletter",
		},
	}))

	Application.Use(user.New(&user.Config{}))

	// Application.PrintRoutes()
}
