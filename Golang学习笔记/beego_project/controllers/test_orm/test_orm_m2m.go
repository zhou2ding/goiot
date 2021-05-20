package testorm

import (
	"github.com/astaxie/beego"
)

type TestOrmM2MController struct {
	beego.Controller
}

//多对多关系中，所有的操作都能反着来
func (t *TestOrmM2MController) Get() {
	// o := orm.NewOrm()
	//插入数据
	// tags := models.Tags{Name: "详解多对多"}
	// arts := models.ArticleMany{Name: "火影忍者", Desc: "Japan漫画"}
	// o.Insert(&tags)
	// o.Insert(&arts)

	//表A和表B多对多，方式1，一边插入新数据，一边插入关联关系
	// m2m := o.QueryM2M(&arts, "Tags") //用A查B，第二个参数必须是模型名
	// m2m.Add(&tags)                   //查到后把B插入关系表

	//表A和表B多对多，方式2，可以用已有的数据来插入关联关系
	// arts := models.ArticleMany{Id: 1}
	// tags := models.Tags{Id: 2}
	// m2m := o.QueryM2M(&tags, "ArticleMany") //用B查A，第二个参数必须是模型名
	// m2m.Add(&arts)                          //查到后把A插入关系表

	//修改数据：只更新关系表的数据，根据两个外键定位一条数据后进行修改
	// aid := 1
	// tid := 3
	// newAid := 3
	// newTid := 2
	// qs := o.QueryTable("relation_table")
	// qs.Filter("article_many_id__exact", aid).Filter("tags_id__exact", tid).Update(orm.Params{
	// 	"article_many_id": newAid,
	// 	"tags_id":         newTid,
	// })

	//删除数据，只删除中间表的数据

	//Remove()，表A和表B多对多，根据表A的id查完后，只能手动指定表B的关联关系
	// arts := models.ArticleMany{Id: 2}
	// tags := []models.Tags{{Id: 2}, {Id: 4}}
	// m2m := o.QueryM2M(&arts, "Tags")
	// m2m.Remove(&tags)

	//Clear()，表A和表B多对多，根据表A的id查完后，能删除所有表B的关联关系，常用
	// arts := models.ArticleMany{Id: 2}
	// m2m := o.QueryM2M(&arts, "Tags")
	// m2m.Clear()

	//查询
	//1.已知A表id，从A表中查询A表id关联的所有B：正向查询
	// aid := 3
	// arts := models.ArticleMany{}
	// qs := o.QueryTable("article_many")
	// qs.Filter("id__exact", aid).One(&arts) //先把查出的其他字段扔进arts
	// o.LoadRelated(&arts, "Tags")           //再关联查询Tags（切片），再扔进arts
	// fmt.Println(arts, arts.Tags)

	//2.已知B表id，从A表中查询B表id关联的所有A：反向查询
	// tid := 2
	// arts := []models.ArticleMany{}
	// qs := o.QueryTable("article_many")
	// qs.Filter("Tags__tags__id", tid).All(&arts) //参数是：模型中的关系名__表名__表中字段名；三个名称均不能包含_
	// fmt.Println(arts)

	//m2m的其他方法
	// arts := models.ArticleMany{Id: 4}
	//Count()查询数量
	// m2m := o.QueryM2M(&arts, "Tags")
	// cnt, _ := m2m.Count()
	// fmt.Println(cnt)
	//Exists()判断是否存在，需要关系表的两个外键共同确认一条数据
	// m2m := o.QueryM2M(&arts, "Tags")
	// tags := models.Tags{Id: 5} //不能是切片，因为只能判断指定的一行数据是否存在
	// exists := m2m.Exist(&tags)
	// fmt.Println(exists)

	t.TplName = "test_orm/m2m.html"
}
