package controllers

import "github.com/astaxie/beego"

type TemplateController struct {
	beego.Controller
}

type TestUsr struct {
	Name    string
	Age     int
	Phone   string
	Married bool
	Price   float64
}

func (t *TemplateController) Get() {
	t.Data["name"] = "zhangsan"
	t.Data["arr"] = []int{}
	testusr := TestUsr{Name: "ss", Age: 10, Phone: "132", Married: true, Price: 3.15}
	t.Data["testusr"] = testusr
	t.TplName = "template1.html"
}
