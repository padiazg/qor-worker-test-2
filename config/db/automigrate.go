package db

import (
	um "github.com/padiazg/qor-worker-test/model/user"

	"github.com/jinzhu/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&um.User{})
}
