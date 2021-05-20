package models

import "github.com/astaxie/beego/orm"

//一对一模型定义
type UserO struct {
	Id       int `orm:"pk;auto"`
	Name     string
	Age      int
	ProfileO *ProfileO   `orm:"reverse(one)"`
	Articles []*Articles `orm:"reverse(many)"`
}

type ProfileO struct {
	Id     int `orm:"pk;auto"`
	Phone  int64
	IdCard string
	UserO  *UserO `orm:"rel(one);null;on_delete(set_null)"`
}

func (u *UserO) TableName() string {
	return "user_o"
}

func (p *ProfileO) TableName() string {
	return "profile_o"
}

func init() {
	orm.RegisterModel(new(UserO), new(ProfileO))
}
