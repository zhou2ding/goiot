package testorm

import "github.com/astaxie/beego"

type O2MController struct {
	beego.Controller
}

func (c *O2MController) Get() {
	c.TplName = "test_orm/pO2M.html"
}
