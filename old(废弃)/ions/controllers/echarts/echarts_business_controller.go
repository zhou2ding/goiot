package echarts

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type EchartsBusinessController struct {

	beego.Controller

}

func (e *EchartsBusinessController) Get()  {
	e.TplName = "echarts/echarts_business.html"

}

func (e *EchartsBusinessController) GetBusinessChart()  {

	var caiwu_date orm.ParamsList
	var student_incress orm.ParamsList
	o := orm.NewOrm()
	o.Raw("select mon from sys_caiwu_data").ValuesFlat(&caiwu_date)
	o.Raw("select stu_incr from sys_caiwu_data").ValuesFlat(&student_incress)
	fmt.Println(caiwu_date)
	fmt.Println(student_incress)

	map_data := map[string]interface{}{}

	map_data["caiwu_date"]  = caiwu_date
	map_data["student_incress"] = student_incress

	e.Data["json"] = map_data
	e.ServeJSON()


}