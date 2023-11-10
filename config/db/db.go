package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/padiazg/qor-worker-test/config/settings"
)

func NewDB(settings *settings.DatabaseSetting) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)

	if db, err = gorm.Open(settings.Dialect, settings.ConnectionString); err != nil {
		return nil, err
	}

	AutoMigrate(db)

	return db, nil
}
