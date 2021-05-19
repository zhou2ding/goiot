package models

import "github.com/astaxie/beego/orm"

type Tags struct {
	Id   int `orm:"pk;auto"`
	Name string
	//orm会自动生成第三个表，来存放多对多的关系，命名规则是正向表名+下划线+反向表名+s
	ArticleMany []*ArticleMany `orm:"reverse(many)"` //多对多，分别是rel(m2m)，和reverse(many)
}

func (t *Tags) TableName() string {
	return "tags"
}

func init() {
	orm.RegisterModel(new(Tags))
}
