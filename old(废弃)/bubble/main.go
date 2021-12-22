package main

import (
	"bubble/dao"
	"bubble/router"
)

func main() {
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	dao.InitModel()
	defer dao.Close()
	r := router.SetupRouter()
	_ = r.Run()
}
