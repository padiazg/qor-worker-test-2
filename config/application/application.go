package application

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	db "github.com/padiazg/qor-worker-test/config/db"
)

// Application main application
type Application struct {
	*Config
	Server *http.Server
}

func (application *Application) Use(app MicroAppInterface) {
	app.ConfigureApplication(application)
}

func (application *Application) UseWorker(wrk MicroWorkerInterface) {
	wrk.ConfigureWorker(application)
}

func (application *Application) SetWorker(name string, wrk interface{}) {
	application.Mutex.Lock()
	defer application.Mutex.Unlock()

	application.Workers[name] = wrk
}

func (a *Application) Run() {
	log.Printf("Listening on port %s\n", a.Server.Addr)
	if err := a.Server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("HTTP server error: %v", err)
	}
}

func (a *Application) Shutdown(ctx context.Context) error {
	if err := a.DB.Close(); err != nil {
		return fmt.Errorf("closing db: %+v", err)
	}

	if err := a.Server.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutting down server: %+v", err)
	}

	return nil
}

func (a *Application) PrintRoutes() {
	// 👇 the walking function 🚶‍♂️
	chi.Walk(a.Router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("[%s]: '%s' has %d middlewares\n", method, route, len(middlewares))
		return nil

	})
}

func (application *Application) ConnectDB(ctx context.Context) error {
	var err error

	application.DB, err = db.NewDB(&application.Settings.Database)
	if err != nil {
		return fmt.Errorf("connecting db: %+v", err)
	}

	return nil
}
