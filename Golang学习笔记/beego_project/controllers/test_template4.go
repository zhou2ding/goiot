package controllers

import (
	"github.com/astaxie/beego"
)

type T4Controller struct {
	beego.Controller
}

func (t *T4Controller) Get() {
	t.Data["name"] = "王五"
	t.TplName = "t4.html"
}

func Hello(h string) string {
	newH := h + "，你好"
	return newH
}

func SubstrAndReplace(s string, i int) (res string) {
	res = string([]rune(s)[0:i]) + "..."
	return
}
