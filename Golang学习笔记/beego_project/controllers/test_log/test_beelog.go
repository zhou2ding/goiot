package testlog

import "github.com/astaxie/beego"

//Beego自带日志处理功能，会被日志模块替代
type TestBeeLogController struct {
	beego.Controller
}

func (t *TestBeeLogController) Get() {
	beego.SetLogger("file", `{"filename":"logs/test1.log"}`) //file是引擎，可换成其他的；filenam是固定的；路径可以改
	beego.BeeLogger.DelLogger("console")                     //日志不在console打印，console是引擎，可换成其他的
	beego.SetLevel(beego.LevelCritical)                      //设置日志级别，值输出此级别和往上级别的
	// beego.SetLevel(2)                                     //级别从0（Emergency）开始增加

	beego.Emergency("Emergency")
	beego.Alert("Alert")
	beego.Critical("Critical")
	beego.Error("Error")
	beego.Warn("Warn")
	beego.Notice("Notice")
	beego.Informational("Informational")
	beego.Debug("Debug")

	t.TplName = "test_log/beelog.html"
}
