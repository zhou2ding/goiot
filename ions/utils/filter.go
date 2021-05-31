package utils

import "github.com/astaxie/beego/context"

func LoginFilter(ctx *context.Context) {
	id := ctx.Input.Session("id")
	if id == nil {
		ctx.Redirect(302, "/")
	}
}
