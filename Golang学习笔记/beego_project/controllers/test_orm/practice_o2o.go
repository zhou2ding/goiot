package testorm

import (
	"beego_project/models"
	"fmt"
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
	fmt.Println(user_id)
	c.TplName = "test_orm/pO2O.html"
}

func (c *O2OController) Post() {
	user_id := c.GetSession("id")
	fmt.Println(user_id)
	id_card := c.GetString("id_card")
	phone, _ := strconv.ParseInt("phone", 10, 64)
	fmt.Println(id_card, phone)

	o := orm.NewOrm()
	profile := models.ProfileO{IdCard: id_card, Phone: phone}
	o.InsertOrUpdate(&profile, "user_one_id")

	c.TplName = "test_orm/success.html"
}
