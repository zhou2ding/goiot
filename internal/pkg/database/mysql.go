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

func InitIotDB() *gorm.DB {
	logger.Logger.Infof("linke %s", conf.Conf.GetString("database.link"))
	if db, err := gorm.Open(mysql.Open(conf.Conf.GetString("database.link")), &gorm.Config{}); err != nil {
		logger.Logger.Fatalf("init database err %s", err)
		return nil
	} else {
		if conf.Conf.GetBool("log.stdout") {
			db.Logger = gormLogger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), gormLogger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  gormLogger.Info,
				IgnoreRecordNotFoundError: false,
				Colorful:                  true,
			})
		}
		return db
	}
}
