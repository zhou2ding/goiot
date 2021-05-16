package testorm

import "github.com/astaxie/beego"

type TestOrm6Controller struct {
	beego.Controller
}

func (t *TestOrm6Controller) Get() {
	t.TplName = "orm6.html"
}
