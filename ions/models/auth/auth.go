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
	CarsApply []*CarsApply `orm:"reverse(many)"`
	MessageNotify []*MessageNotify `orm:"reverse(many)"`
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
// 车辆品牌
type CarBrand struct {
	Id int `orm:"pk;auto"`
	Name string `orm:"description(品牌名称);size(64)"`
	Desc string `orm:"description(品牌描述);size(255)"`
	Cars []*Cars `orm:"reverse(many)"`
	IsActive int `orm:"default(1);description(启用:1,停用:0)"`
	IsDelete int `orm:"default(0);description(删除:1,未删除:0)"`
	CreateTime time.Time `orm:"description(创建时间);type(datetime);auto_now"`
}
// 车辆列表
type Cars struct {
	Id int `orm:"pk;auto"`
	Name string `orm:"description(车辆名称);size(64)"`
	CarBrand *CarBrand `orm:"rel(fk);description(车辆品牌外键)"`
	CarsApply []*CarsApply `orm:"reverse(many)"`
	Status int `orm:"description(1:可借,2:不可借);default(1)"`
	IsActive int `orm:"default(1);description(启用:1,停用:0)"`
	IsDelete int `orm:"default(0);description(删除:1,未删除:0)"`
	CreateTime time.Time `orm:"description(创建时间);type(datetime);auto_now"`
}


// 车辆申请、审核
type CarsApply struct {
	Id int `orm:"pk;auto"`
	User *User `orm:"rel(fk)"`
	Cars *Cars `orm:"rel(fk)"`
	Reason string `orm:"description(申请理由);size(255)"`
	Destination string `orm:"description(目的地);size(64)"`
	ReturnDate time.Time `orm:"type(date);auto_now;description(归还日期)"`
	ReturnStatus int `orm:"descrption(1:已归还，0：未归还);default(0)"`
	AuditStatus int `orm:"description(1:同意，2:未同意，3:未审批);default(3)"`
	AuditOption string `orm:"description(审批意见);size(255)"`
	IsActive int `orm:"default(1);description(启用:1,停用:0)"`
	IsDelete int `orm:"default(0);description(删除:1,未删除:0)"`
	CreateTime time.Time `orm:"description(创建时间);type(datetime);auto_now"`
	NotifyTag int `orm:"description(1:已发送通知，0：未发送通知);default(0)"`

}

// 消息通知
type MessageNotify struct {
	Id int `orm:"pk;auto"`
	Flag int `orm:"description(1:车辆逾期，2:所有通知);default(1)"`
	Title string `orm:"size(64);description(消息标题)"`
	Content string `orm:"type(text);description(消息内容)"`
	User *User `orm:"rel(fk);desription(用户外键)"`
	ReadTag int `orm:"description(1:已读，0:未读)"`

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
func (u *CarBrand) TableName() string {
	return "sys_cars_brand"
}

func (u *Cars) TableName() string {
	return "sys_cars"
}

func (u *CarsApply) TableName() string {
	return "sys_cars_apply"
}

func (u *MessageNotify) TableName() string {
	return "sys_message_notify"
}

func init() {
	orm.RegisterModel(new(Auth), new(Role), new(User), new(Cars), new(CarBrand), new(CarsApply),new(MessageNotify))
}
