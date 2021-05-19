package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id   int    `orm:"pk;auto"`                //pk是主键，auto是自增。字段是Id的话默认用作主键且自增；不是的话要手动添加pk和auto
	Name string `orm:"column(username);index"` //在此字段上添加索引
	Age  int    `orm:"null"`                   //允许为null；digits是int类型的长度
	Addr string `orm:"unique;size(32)"`        //unique约束；size是string类型的长度，即varchar

	// Price      float64   `orm:"digits(12);decimals(4)"`  //digits是总长度，decimals是小数部分的长度；整数部分超长会报错，小数部分超长会四舍五入
	// XXX        string    `orm:"-;description(姓名)"`       //- 是忽略字段的映射关系；description是注释
	// CreateTime time.Time `orm:"auto_now;type(datetime)"` //auto_now是每次model保存时都会自动更新时间；type(datetime)是datetime类型
	// CreateDate time.Time `orm:"auto_now_add;type(date)"` //auto_now_add是第一次保存时才设置时间；type(date)是date类型
	// Gender     int       `orm:"default(1)"`              //指定默认值
}

func (u *User) TableName() string {
	return "sys_user"
}

func init() {
	orm.RegisterModel(new(User))
}
