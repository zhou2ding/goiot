package models

import "github.com/astaxie/beego/orm"

type Article struct {
	Id     int `orm:"pk"`
	Title  string
	Author string
}

func (a *Article) TableName() string {
	return "article"
}

func init() {
	orm.RegisterModel(new(Article))
}
