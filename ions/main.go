package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	_ "ions/models/auth"
	_ "ions/models/news"
	_ "ions/models/profit"
	_ "ions/models/salary"
	_ "ions/routers"
	"ions/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	username := beego.AppConfig.String("username")
	pwd := beego.AppConfig.String("pwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	db := beego.AppConfig.String("db")
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	dataSource := username + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + db + "?charset=utf8"
	_ = orm.RegisterDataBase("default", "mysql", dataSource)
	ret := fmt.Sprintf("connect to databse success, host:%v, post:%v, db:%v",host,port,db)
	logs.Info(ret)
}

func main() {
	orm.RunCommand()
	orm.Debug = true
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.InsertFilter("/m/*", beego.BeforeRouter, utils.LoginFilter)

	//日志
	logs.SetLogFuncCallDepth(3)
	_ = logs.SetLogger(logs.AdapterMultiFile,`{"filename":"logs/ions.log","separate":["error","info"]}`)

	beego.Run()
}
