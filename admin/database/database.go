package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"semay.com/config"
)

var (
	DBConn *gorm.DB
)

func ReturnSession() *gorm.DB {

	var DBSession *gorm.DB
	db, _ := gorm.Open(sqlite.Open(config.Config("SQLITE_URI")), &gorm.Config{})

	DBSession = db

	return DBSession
}
