package testorm

import (
	"beego_project/models"
	"beego_project/utils"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//小实战
type TestOrm2ListController struct {
	beego.Controller
}

type TestOrm2AddController struct {
	beego.Controller
}

type TestOrm2UpdateController struct {
	beego.Controller
}

type TestOrm2DeleteController struct {
	beego.Controller
}

func (t *TestOrm2AddController) Prepare() {
	t.EnableXSRF = false
}
func (t *TestOrm2UpdateController) Prepare() {
	t.EnableXSRF = false
}

func (t *TestOrm2ListController) Get() {
	o := orm.NewOrm()
	article := models.Article{Id: 4}
	o.Read(&article)
	// article := []models.Article{}
	// qs := o.QueryTable(new(models.Article))
	// qs.Filter("id__gt", "0").All(&article)

	t.Data["article"] = article
	t.TplName = "test_orm/orm2_list.html"
}

func (t *TestOrm2AddController) Get() {
	t.TplName = "test_orm/orm2_add.html"
}

func (t *TestOrm2AddController) Post() {
	title := t.GetString("title")
	author := t.GetString("author")
	fmt.Println(title, author)

	article := models.Article{Title: title, Author: author}
	o := orm.NewOrm()
	id, err := o.Insert(&article)
	fmt.Println(id, err)
	t.TplName = "test_orm/success.html"
}

func (t *TestOrm2UpdateController) Get() {
	//此处获得的id是从URL中来的
	id := t.GetString("id")
	// id_int, _ := strconv.ParseInt(id, 10, 0)
	id_int := utils.Str2Int(id)
	article := models.Article{Id: id_int}

	o := orm.NewOrm()
	o.Read(&article)
	t.Data["article"] = article

	t.TplName = "test_orm/orm2_update.html"
}

func (t *TestOrm2UpdateController) Post() {
	//此处的id，是点了get跳转到update页面后，get传了一个article结构体进来，结构体中有id字段，再把这个id字段从前端传到后端，后端先根据这个id查询，再修改指定字段
	id := t.GetString("id")
	title := t.GetString("title")
	author := t.GetString("author")
	// id_int, _ := strconv.ParseInt(id, 10, 0)
	id_int := utils.Str2Int(id)
	article := models.Article{Id: id_int}

	o := orm.NewOrm()
	err := o.Read(&article)
	if err == nil {
		article.Title = title
		article.Author = author
		o.Update(&article, "Title", "Author")
	}
	t.TplName = "test_orm/success.html"
}

func (t *TestOrm2DeleteController) Get() {
	id := t.GetString("id")
	// id_int, _ := strconv.ParseInt(id, 10, 0)
	id_int := utils.Str2Int(id)
	article := models.Article{Id: id_int}

	o := orm.NewOrm()
	o.Delete(&article)

	t.TplName = "test_orm/success.html"
}
