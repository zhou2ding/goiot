package testorm

import (
	"beego_project/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type O2OController struct {
	beego.Controller
}

func (c *O2OController) Prepare() {
	c.EnableXSRF = false
}

func (c *O2OController) Get() {
	c.SetSession("id", 1)
	user_id := c.GetSession("id")
	pf := models.ProfileO{}

	o := orm.NewOrm()
	o.QueryTable(new(models.ProfileO)).Filter("user_o_id", user_id).One(&pf)
	c.Data["pf"] = pf

	c.TplName = "test_orm/pO2O.html"
}

func (c *O2OController) Post() {
	user_id := c.GetSession("id")
	id_card := c.GetString("id_card")
	phone, _ := strconv.ParseInt(c.GetString("phone"), 10, 64)

	o := orm.NewOrm()
	usr := models.UserO{}
	o.QueryTable(new(models.UserO)).Filter("id", user_id).One(&usr)
	profile := models.ProfileO{IdCard: id_card, Phone: phone, UserO: &usr}
	o.InsertOrUpdate(&profile, "user_one_id")

	c.TplName = "test_orm/success.html"
}
