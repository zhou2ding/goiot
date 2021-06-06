package dao

import (
	"bubble/models"
	"github.com/go-orm/gorm"
	_ "github.com/go-orm/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitMysql() (err error) {
	dsn := "root:564710@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func InitModel() {
	DB.AutoMigrate(&models.Todo{})
}

func Close() {
	DB.Close()
}
