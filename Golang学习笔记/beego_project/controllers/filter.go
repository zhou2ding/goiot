package controllers

import "github.com/astaxie/beego/context"

func FilterUser(ctx *context.Context) {
	url := ctx.Request.RequestURI
	if url != "/login" {
		ctx.Redirect(302, "/login/get")
	}
}
