package migration

import (
	"github.com/congmanh18/XNM-Express/User_Profile/entity"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.Exec("CREATE SCHEMA IF NOT EXISTS user_profile_service;")
	err := db.AutoMigrate(
		&entity.User{},
	)
	if err != nil {
		panic(err)
	}
}
