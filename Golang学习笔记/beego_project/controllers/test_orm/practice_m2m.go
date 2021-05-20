package testorm

import (
	"beego_project/models"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type PostListController struct {
	beego.Controller
}

func (m *PostListController) Get() {
	o := orm.NewOrm()
	posts := []models.Post{}
	o.QueryTable("post").All(&posts)
	for k, _ := range posts {
		o.LoadRelated(&posts[k], "Tag")
	}
	m.Data["posts"] = posts
	m.TplName = "test_orm/pM2M.html"
}

type PostAddController struct {
	beego.Controller
}

func (p *PostAddController) Get() {
	o := orm.NewOrm()
	tag := []models.Tag{}
	o.QueryTable("tag").All(&tag)
	p.Data["tag"] = tag
	p.TplName = "test_orm/postadd.html"
}

func (p *PostAddController) Post() {
	name := p.GetString("name")
	context := p.GetString("content")
	o := orm.NewOrm()
	post := models.Post{Name: name, Content: context, CreaTime: time.Now()}
	o.Insert(&post)

	tag := p.GetStrings("tag")
	for i := 0; i < len(tag); i++ {
		tagId, _ := strconv.Atoi(tag[i])
		tagObj := models.Tag{Id: tagId}
		m2m := o.QueryM2M(&post, "Tag")
		m2m.Add(&tagObj)
	}
	p.TplName = "test_orm/success.html"
}

type TagListController struct {
	beego.Controller
}

func (t *TagListController) Get() {
	o := orm.NewOrm()
	tag := []models.Tag{}
	o.QueryTable("tag").All(&tag)
	t.Data["tag"] = tag
	t.TplName = "test_orm/taglist.html"
}

type TagAddController struct {
	beego.Controller
}

func (t *TagAddController) Get() {
	t.TplName = "test_orm/tagadd.html"
}

func (t *TagAddController) Post() {
	name := t.GetString("name")
	desc := t.GetString("desc")
	o := orm.NewOrm()
	tag := models.Tag{Name: name, Desc: desc, CreaTime: time.Now()}
	o.Insert(&tag)

	t.TplName = "test_orm/success.html"
}

type PostUpdateController struct {
	beego.Controller
}

func (p *PostUpdateController) Get() {
	o := orm.NewOrm()
	post_id, _ := strconv.Atoi(p.GetString("id"))
	post := models.Post{}
	o.QueryTable("post").Filter("id", post_id).One(&post)
	o.LoadRelated(&post, "Tag")
	p.Data["post"] = post

	uncheckTag := models.Tag{}
	o.QueryTable("tag").Exclude("id__in", post.Tag).All(&uncheckTag)
	p.Data["uncheckTag"] = uncheckTag

	p.TplName = "test_orm/postupdate.html"
}

func (p *PostUpdateController) Post() {
	id, _ := strconv.Atoi(p.GetString("id"))
	tag := p.GetStrings("tag")
	o := orm.NewOrm()
	post := models.Post{Id: id}
	o.QueryM2M(&post, "Tag").Clear()
	for i := 0; i < len(tag); i++ {
		tagId, _ := strconv.Atoi(tag[i])
		tagObj := models.Tag{Id: tagId}
		m2m := o.QueryM2M(&post, "Tag")
		m2m.Add(&tagObj)
	}

	p.TplName = "test_orm/success.html"
}

type PostDelController struct {
	beego.Controller
}

func (p *PostDelController) Get() {
	id, _ := strconv.Atoi(p.GetString("id"))
	o := orm.NewOrm()
	o.QueryTable("post").Filter("id", id).Delete()

	post := models.Post{Id: id}
	o.QueryM2M(&post, "Tag").Clear()

	p.Redirect(beego.URLFor("PostListController.Get"), 302)
}

func (t *TagAddController) Prepare() {
	t.EnableXSRF = false
}

func (t *PostAddController) Prepare() {
	t.EnableXSRF = false
}

func (t *PostUpdateController) Prepare() {
	t.EnableXSRF = false
}
