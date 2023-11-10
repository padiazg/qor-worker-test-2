package user

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/auth/auth_identity"
)

type User struct {
	gorm.Model
	auth_identity.Basic
	First string `gorm:"column:first"`
	Last  string `gorm:"column:last"`
	Email string `gorm:"column:email;index"`
}
