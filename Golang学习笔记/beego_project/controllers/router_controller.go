package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type RouterController struct {
	beego.Controller
}

func (r *RouterController) Util() {
	data := r.Ctx.Input.Param(":id")
	fmt.Println("===========================", data)
	if r.Ctx.Request.Method == "GET" {
		// r.StopRun() //终止逻辑，比如用户认证不通过
		r.TplName = "test_router_get.html"
	} else if r.Ctx.Request.Method == "POST" {
		r.TplName = "test_router_post.html"
	}
	r.TplName = "test_router.html"
}
