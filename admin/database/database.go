package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"semay.com/config"
)

var (
	DBConn *gorm.DB
)

func ReturnSession() *gorm.DB {
	//  This is file to output gorm logger on to
	gormfile, err := os.OpenFile("gormblue.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	gormFileLogger := log.Logger{}
	gormFileLogger.SetOutput(gormfile)
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

	// io.Copy(gormfile, gorm_logs)
	// fmt.Println(log_buffer.String())
	var DBSession *gorm.DB
	db, _ := gorm.Open(sqlite.Open(config.Config("SQLITE_URI")), &gorm.Config{
		Logger:                 newLogger,
		SkipDefaultTransaction: true,
	})

	DBSession = db
	return DBSession

}
