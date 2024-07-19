package database

import (
	"goiot/internal/pkg/conf"
	"goiot/internal/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var gDB *gorm.DB

func init() {
	logger.Logger.Infof("linke %s", conf.Conf.GetString("database.link"))
	db, err := gorm.Open(mysql.Open(conf.Conf.GetString("database.link")), &gorm.Config{})
	if err != nil {
		logger.Logger.Fatalf("init database err %s", err)
	}
	if conf.Conf.GetBool("log.stdout") {
		db.Logger = gormLogger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), gormLogger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  gormLogger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		})
	}

	gDB = db
}

func GetDB() *gorm.DB {
	return gDB
}
