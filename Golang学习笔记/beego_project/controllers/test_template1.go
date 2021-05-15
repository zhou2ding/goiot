package controllers

import "github.com/astaxie/beego"

type TemplateController struct {
	beego.Controller
}

func (t *TemplateController) Get() {
	t.Data["name"] = "zhangsan"
	t.Data["arr"] = []int{}
	t.TplName = "template1.html"
}
