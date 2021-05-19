package models

import "github.com/astaxie/beego/orm"

//一对一模型定义
type ArticleMany struct {
	Id   int `orm:"pk;auto"`
	Name string
	Desc string
	//一对多，多是rel(fk)；一是reverse(many);
	//on_delete()是级联删除的配置，默认参数为级联删除cascade；set_null是删除主表数据后，子表的外键字段置位NULL；do_nothing是删除主表数据后，子表什么都不做
	// UserOne *UserOne `orm:"rel(fk);on_delete(set_null)"`

	//多对多，分别是rel(m2m)，和reverse(many)
	Tags []*Tags `orm:"rel(m2m);rel_table(relation_table)"` //还可以指定关系表的表名，rel_table
	// Tags []*Tags `orm:"rel(m2m);rel_through(beego_project/models.ArticleManyTags)"` //把第三张表通过ArticleManyTags模型映射出来，rel_through

}

// type ArticleManyTags struct {
// 	Id          int          `orm:"pk;auto"`
// 	ArticleMany *ArticleMany `orm:"rel(fk)"`
// 	Tags        *Tags        `orm:"rel(fk)"`
// }
// func (a *ArticleManyTags) TableName() string {
// 	return "article_many_tags"
// }

func (a *ArticleMany) TableName() string {
	return "article_many"
}

func init() {
	orm.RegisterModel(new(ArticleMany))

	// orm.RegisterModel(new(ArticleMany), new(ArticleMany))
}
