package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Tag struct {
	Id       int `orm:"pk;auto"`
	Name     string
	CreaTime time.Time `orm:"auto_now;type(datetime)"`
	Desc     string
	Post     []*Post `orm:"reverse(many)"`
}

type Post struct {
	Id       int `orm:"pk;auto"`
	Name     string
	CreaTime time.Time `orm:"auto_now;type(datetime)"`
	Content  string
	Tag      []*Tag `orm:"rel(m2m)"`
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
