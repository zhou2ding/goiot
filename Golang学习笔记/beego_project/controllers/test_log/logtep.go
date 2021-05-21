package testlog

import (
	logtemplate "beego_project/utils/log_template"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type LogTepController struct {
	beego.Controller
}

func (l *LogTepController) Get() {
	cnt := 6
	for i := 1; i <= 6; i++ {
		start_time := time.Now()
		s := start_time.Format("2006-01-02 15:04:05")
		time.Sleep(time.Second * 3)
		end_time := time.Now()
		e := end_time.Format("2006-01-02 15:04:05")
		cur := i
		use_time := end_time.Sub(start_time)
		u := use_time.String()
		ret := logtemplate.LogProcess(s, cur, cnt, u, e)
		logs.Info(ret)
	}
	l.TplName = "success.html"
}
