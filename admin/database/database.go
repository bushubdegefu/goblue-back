package database

import (
	"log"
	"os"
	"time"

	// "gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"semay.com/config"
)

var (
	DBConn *gorm.DB
)

func GormLoggerFile() *os.File {

	gormLogFile, gerr := os.OpenFile("gormblue.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if gerr != nil {
		log.Fatalf("error opening file: %v", gerr)
	}
	return gormLogFile
}

func ReturnSession() *gorm.DB {

	//  This is file to output gorm logger on to
	gormlogger := GormLoggerFile()
	gormFileLogger := log.Logger{}
	gormFileLogger.SetOutput(gormlogger)
	gormFileLogger.Writer()
	gormLogger := log.New(gormFileLogger.Writer(), "\r\n", log.LstdFlags|log.Ldate|log.Ltime|log.Lshortfile)
	newLogger := logger.New(
		gormLogger, // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			Colorful:                  true,        // Enable color
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			// ParameterizedQueries:      true,        // Don't include params in the SQL log

		},
	)

	var DBSession *gorm.DB
	//  this is for postgresql connection
	// conn := "host=192.168.49.2 user=blueuser password=default dbname=bluev5 port=30432 sslmode=disable"
	conn := config.Config("PSQL_URI")
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  conn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage,

	}), &gorm.Config{
		Logger:                 newLogger,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	//  this is sqlite connection
	// db, _ := gorm.Open(sqlite.Open(config.Config("SQLITE_URI")), &gorm.Config{
	// 	Logger:                 newLogger,
	// 	SkipDefaultTransaction: true,
	// })
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(4)
	sqlDB.SetConnMaxLifetime(2 * time.Second)

	DBSession = db

	return DBSession

}
