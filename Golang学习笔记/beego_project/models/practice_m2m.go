package models

import "github.com/astaxie/beego/orm"

type Tag struct {
	Id   int `orm:"pk;auto"`
	Name string
	Post []*Post `orm:"reverse(many)"`
}

type Post struct {
	Id   int `orm:"pk;auto"`
	Name string
	Tag  []*Tag `orm:"rel(m2m)"`
}

func (t *Tag) TableName() string {
	return "tag"
}

func (p *Post) TableName() string {
	return "post"
}

func init() {
	orm.RegisterModel(new(Tag), new(Post))
}
