package testorm

import (
	"beego_project/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type O2MController struct {
	beego.Controller
}

func (c *O2MController) Prepare() {
	c.EnableXSRF = false
}
func (c *ArticleAddController) Prepare() {
	c.EnableXSRF = false
}

func (c *O2MController) Get() {
	c.SetSession("uid", 1)
	user_id := c.GetSession("uid")

	o := orm.NewOrm()
	arts := []models.Articles{}
	o.QueryTable(new(models.Articles)).Filter("user_o_id", user_id).All(&arts)

	c.Data["arts"] = arts

	c.TplName = "test_orm/pO2M.html"
}

type ArticleAddController struct {
	beego.Controller
}

func (a *ArticleAddController) Get() {
	a.TplName = "test_orm/articleadd.html"
}

func (a *ArticleAddController) Post() {
	uid := a.GetSession("uid")
	name := a.GetString("name")
	content := a.GetString("content")
	desc := a.GetString("desc")

	o := orm.NewOrm()
	usr := models.UserO{}
	o.QueryTable("user_o").Filter("id", uid).One(&usr)
	arts := models.Articles{Name: name, Content: content, Desc: desc, UserO: &usr}
	o.Insert(&arts)

	a.TplName = "test_orm/success.html"
}

type ArticleUpdateController struct {
	beego.Controller
}

func (a *ArticleUpdateController) Get() {
	id, _ := strconv.Atoi(a.GetString("id"))
	o := orm.NewOrm()
	art := models.Articles{}
	o.QueryTable("articles").Filter("id", id).One(&art)
	a.Data["art"] = art

	a.TplName = "test_orm/articleupdate.html"
}

func (a *ArticleUpdateController) Post() {
	//从前端获取到id、name、content、desc，user_o_id通过getsession获取，然后根据获取到的变量去o.QueryTable().Filter().Update(orm.Params{})
	a.TplName = "test_orm/success.html"
}
