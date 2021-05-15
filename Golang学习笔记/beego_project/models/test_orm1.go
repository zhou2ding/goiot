package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id   int    `orm:"pk"` // - 是忽略字段的映射关系；pk 是主键
	Name string `orm:"column(username)"`
	Age  int
	Addr string
}

func (u *User) TableName() string {
	return "sys_user"
}

func init() {
	orm.RegisterModel(new(User))
}
