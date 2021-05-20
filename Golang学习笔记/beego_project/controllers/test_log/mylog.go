package testlog

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MyLogController struct {
	beego.Controller
}

func (m *MyLogController) Get() {
	logs.Emergency("this is a Emergency")
	logs.Alert("this is a Alert")
	logs.Critical("this is a Critical")
	logs.Error("this is a Error")
	logs.Warn("this is a Warning") //或者Warning()
	logs.Notice("this is a Notice")
	logs.Info("this is a Info")   //或者Informational()
	logs.Trace("this is a Debug") //或者Debug()

	m.TplName = "test_log/mylog.html"
}
