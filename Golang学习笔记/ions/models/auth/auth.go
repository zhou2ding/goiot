package auth

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id       int    `orm:"pk;auto"`
	CardId   string `orm:"size(64)"`
	UserName string `orm:"unique;column(user_name);size(64)" json:"username"`
	Password string `orm:"size(32)" json:"password"`
	Age      int    `orm:"null"`
	Gender   string `orm:"null;description(1:男,2:女,3:未知)"`
	Phone    int64  `orm:"null"`
	Addr     string `orm:"null"`
	IsActive int    `orm:"default(1);description(1:启用,0:停用)"`
	IsDel    int    `orm:"default(0);description(1:已删,0:未删)"`

	Creatime time.Time `orm:"type(datetime);auto_now;null"`
	Role     []*Role   `orm:"reverse(many)"`
}

type Auth struct {
	Id         int    `orm:"pk;auto"`
	AuthName   string `orm:"unique;size(64)"`
	UrlFor     string
	ParentId   int
	Desc       string    `orm:"null"`
	CreateTime time.Time `orm:"type(datetime);auto_now"`
	IsActive   int       `orm:"default(1)"`
	IsDel      int       `orm:"default(0)"`
	Weight     int
	Role       []*Role `orm:"reverse(many)"`
}

type Role struct {
	Id         int    `orm:"pk;auto"`
	RoleName   string `orm:"size(64)"`
	Desc       string
	IsActive   int       `orm:"default(1)"`
	IsDel      int       `orm:"default(0)"`
	CreateTime time.Time `orm:"type(datetime);auto_now"`
	Auth       []*Auth   `orm:"rel(m2m)"`
	User       []*User   `orm:"rel(m2m)"`
}

func (u *User) TableName() string {
	return "sys_user"
}

func (a *Auth) TableName() string {
	return "sys_auth"
}

func (r *Role) TableName() string {
	return "sys_role"
}

func init() {
	orm.RegisterModel(new(Auth), new(Role), new(User))
}
