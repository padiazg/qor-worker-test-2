package application

import (
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/jinzhu/gorm"
	"github.com/padiazg/qor-worker-test/config/settings"
	"github.com/qor/admin"
)

// Config application config
type Config struct {
	HostPort  uint
	Router    *chi.Mux
	StripeKey string
	Domain    string
	Settings  *settings.Settings
	Workers   map[string]interface{}
	Mutex     *sync.RWMutex

	Admin *admin.Admin
	DB    *gorm.DB

	// Auth           *auth.Auth
	// Logger *zap.SugaredLogger
	// Storage        *storage.Storage
	// SessionManager session.ManagerInterface
	// Licenser       *lm.Licenser
}
