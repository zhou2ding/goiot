package models

import "github.com/astaxie/beego/orm"

//一对一模型定义
type Articles struct {
	Id      int `orm:"pk;auto"`
	Name    string
	Content string
	Desc    string
	UserO   *UserO `orm:"rel(fk)"`
}

func (a *Articles) TableName() string {
	return "articles"
}

func init() {
	orm.RegisterModel(new(Articles))
}
