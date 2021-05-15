package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	l.TplName = "login.html"
}
