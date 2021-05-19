package testorm

import (
	"beego_project/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type TestOrmO2MController struct {
	beego.Controller
}

func (t *TestOrmO2MController) Get() {
	o := orm.NewOrm()
	// uid := 2 //一般由前端获取，下面的更新和删除要用到
	aid := 4

	//orm插入
	// usr := models.UserOne{UserName: "王五", Age: 37} //也可直接用已有的usr数据
	// art := models.ArticleMany{Name: "海贼王", Desc: "日本漫画", UserOne: &usr}
	// o.Insert(&usr)
	// o.Insert(&art)

	//orm修改
	// qs := o.QueryTable("article_many")
	// qs.Filter("id__exact", aid).Update(orm.Params{"user_one_id": uid})

	//orm删除
	// qs := o.QueryTable("article_many")
	// qs.Filter("user_one_id__exact", uid).Delete() //没有级联删除的功能

	//orm查询：RelatedSel()（正向查询）

	//1.已知子表id，查询主表和子表信息
	// art := models.ArticleMany{}
	// qs := o.QueryTable("article_many")
	// qs.Filter("id__exact", aid).RelatedSel().One(&art)
	// fmt.Println(art, art.UserOne)

	//2.已知子表外键的值，查询主表和子表信息（正向查询）
	// arts := []models.ArticleMany{}
	// qs := o.QueryTable("article_many")
	// qs.Filter("user_one_id__exact", aid).RelatedSel().All(&arts)
	// fmt.Println(arts, arts[0].UserOne)

	//3.已知子表id，通过外键查询（反向查询，只能查出主表信息）
	usr := models.UserOne{}
	qs := o.QueryTable("user_one")
	qs.Filter("ArticleMany__Id", aid).One(&usr)
	fmt.Println(usr)

	t.TplName = "test_orm/o2m.html"
}
