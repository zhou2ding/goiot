package testorm

import "github.com/astaxie/beego"

type M2MController struct {
	beego.Controller
}

func (m *M2MController) Get() {
	m.TplName = "test_orm/pM2M.html"
}
