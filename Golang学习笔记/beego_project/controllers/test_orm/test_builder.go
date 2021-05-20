package testorm

import (
	"beego_project/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type TestBuilderController struct {
	beego.Controller
}

func (t *TestBuilderController) Get() {
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").From("user")
	sql := qb.String()
	o := orm.NewOrm()
	usr := models.User{}
	o.Raw(sql).QueryRow(&usr)
	fmt.Println(usr)
	t.TplName = "test_orm/builder.html"
}
