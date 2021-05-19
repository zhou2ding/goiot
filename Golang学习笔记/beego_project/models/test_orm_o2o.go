package models

import "github.com/astaxie/beego/orm"

//一对一模型定义
type UserOne struct {
	Id         int `orm:"pk;auto"`
	UserName   string
	Age        int
	ProfileOne *ProfileOne `orm:"reverse(one)"` //只是用来反向查询的，建表时会自动省略
}

type ProfileOne struct {
	Id    int `orm:"pk;auto"`
	Phone string
	// Phone int64 //int64才对应bigint，其他都对应int

	//字段必须是结构体指针类型，orm会自动把字段名转换成user_one_id（先进行大小写和加下划线的转换，再自动在尾部跟_id）
	UserOne *UserOne `orm:"rel(one)"` //rel(one)是正向查询，会自动添加unique约束；但需要建表后手动加外键，不过orm会自动把此字段的值和主表对应值保持一致
}

func (u *UserOne) TableName() string {
	return "user_one"
}

func (p *ProfileOne) TableName() string {
	return "profile_one"
}

func init() {
	orm.RegisterModel(new(UserOne), new(ProfileOne))
}
