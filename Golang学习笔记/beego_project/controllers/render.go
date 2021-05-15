package controllers

import "github.com/astaxie/beego"

type RenderController struct {
	beego.Controller
}
type UserRender struct {
	Id    int    `form:"id,text,学号："` //三个参数都带全的话会自动带默认的零值
	Name  string `form:"username,text,姓名："`
	Sex   string `form:"-"`
	Age   int    `form:"age,text,年龄："`
	Intro string `form:",textarea,"`
}

func (r *RenderController) Get() {
	usr := &UserRender{}
	r.Data["Form"] = usr
	r.TplName = "render.html"
}
