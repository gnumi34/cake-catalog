package database

import (
	"log"
	"os"
	"time"

	"github.com/astaxie/beego"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Conn *gorm.DB

func Connect() (db *gorm.DB, err error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,
		},
	)
	db, err = gorm.Open(postgres.Open(beego.AppConfig.String("sqlconn")), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Panic("Cannot connect to the database!")
	}
	Conn = db
	return
}
