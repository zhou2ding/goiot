package controllers

import (
	"time"

	"github.com/astaxie/beego"
)

type T3Controller struct {
	beego.Controller
}

func (t *T3Controller) Get() {
	t.Data["nowtime"] = time.Now()
	t.TplName = "t3.html"
}
