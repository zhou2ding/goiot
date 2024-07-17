package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./dao",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	gormdb, _ := gorm.Open(mysql.Open(""))
	g.UseDB(gormdb)

	g.ApplyBasic(
		g.GenerateAllTable()...,
	)
	g.Execute()
}
