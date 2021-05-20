package controllers

import "github.com/astaxie/beego"

type T2Controller struct {
	beego.Controller
}

func (t *T2Controller) Get() {
	t.Data["name"] = "张三"
	t.Data["age"] = "五十"
	t.Data["f"] = T
	t.Data["mapp"] = map[string]string{"name": "李四"}
	t.Data["slicee"] = []int{5, 9, 30}
	t.Data["stringg"] = "abcdef"
	t.Data["is_bool"] = false

	t.TplName = "t2.html"
}

func T(a int) string {
	return "f调用"
}
