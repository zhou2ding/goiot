package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	fmt.Println(beego.URLFor("LoginController.Get", "name", "zhangsan", "age", "18"))
	l.TplName = "login.html"
}
