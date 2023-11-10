package application

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
)

func New(config *Config) *Application {

	if config == nil {
		config = &Config{}
	}

	if config.Router == nil {
		config.Router = chi.NewRouter()
	}

	if config.HostPort == 0 {
		config.HostPort = 8080
	}

	if config.Domain == "" {
		config.Domain = "http://localhost:8080"
	}

	if config.Mutex == nil {
		config.Mutex = &sync.RWMutex{}
	}

	a := &Application{
		Config: config,
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%d", config.HostPort),
			Handler: config.Router,
		},
	}

	a.Workers = make(map[string]interface{})

	// a.initializeRoutes()

	return a
}
