package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"ions/models/auth"
	"strings"
)

type MyCenterController struct {
	beego.Controller
}

func (m *MyCenterController) Get() {
	user := auth.User{}
	o := orm.NewOrm()
	qs := o.QueryTable("sys_user")
	uid := m.GetSession("id")
	_ = qs.Filter("id", uid).One(&user)

	m.Data["user"] = user
	m.TplName = "user/my_center_edit.html"
}

func (m *MyCenterController) Post() {
	uid, _ := m.GetInt("uid")
	userName := m.GetString("username")
	oldPwd := m.GetString("old_pwd")
	newPwd := m.GetString("new_pwd")
	age, _ := m.GetInt("age")
	gender := m.GetString("gender")
	phone, _ := m.GetInt64("phone")
	addr := strings.TrimSpace(m.GetString("addr"))
	isActive, _ := m.GetInt("is_active")
	usr := auth.User{}
	o := orm.NewOrm()
	qs := o.QueryTable("sys_user")
	_ = qs.Filter("id", uid).One(&usr)
	resp := make(map[string]interface{}, 2)
	if oldPwd != usr.Password {
		resp["code"] = 1001
		resp["msg"] = "原始密码错误！"
	} else {
		_, _ = qs.Filter("id", uid).Update(orm.Params{
			"id":        uid,
			"user_name": userName,
			"password":  newPwd,
			"age":       age,
			"gender":    gender,
			"phone":     phone,
			"addr":      addr,
			"is_active": isActive,
		})
		resp["code"] = 200
		resp["msg"] = "个人信息更新成功！"
	}
	m.Data["json"] = resp
	m.ServeJSON()
}
