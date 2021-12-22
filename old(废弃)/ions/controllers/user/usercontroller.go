package controllers

import (
	"fmt"
	auth "ions/models/auth"
	"ions/utils"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) List() {
	var usr []auth.User
	o := orm.NewOrm()
	qs := o.QueryTable("sys_user")
	//当前页
	curPage, err := u.GetInt("page")
	if err != nil {
		curPage = 1 //没有获取到页码数时，给定为1
	}
	//数据总数
	var all int64
	//每页显示量
	pageSize := 6
	//搜索显示
	kw := u.GetString("kw")
	ret := fmt.Sprintf("current page:%v, search keyword:%v\n", curPage, kw)
	logs.Info(ret)
	all, _ = qs.Filter("is_del", 0).Filter("user_name__contains", kw).Count()
	_,_ = qs.Filter("is_del", 0).Filter("user_name__contains", kw).Limit(pageSize, (curPage-1)*pageSize).All(&usr)
	//总页数
	allPage := int(math.Ceil(float64(all) / float64(pageSize)))
	//上一页
	var prePage int
	if curPage == 1 {
		prePage = 1
	} else {
		prePage = curPage - 1
	}
	//下一页
	var nextPage int
	if curPage < allPage {
		nextPage = curPage + 1
	} else {
		nextPage = curPage
	}
	u.Data["usr"] = usr
	u.Data["prePage"] = prePage
	u.Data["curPage"] = curPage
	u.Data["nextPage"] = nextPage
	u.Data["allPage"] = allPage
	u.Data["all"] = all
	pageMap := utils.Paginator(curPage, pageSize, all)
	u.Data["pageMap"] = pageMap
	u.Data["kw"] = kw //传到前端后，前端在分页的href中加上kw，这样换页时就不会清楚搜索结果

	u.TplName = "user/user-list.html"
}

func (u *UserController) Add() {
	u.TplName = "user/user-add.html"
}

func (u *UserController) Submit() {
	o := orm.NewOrm()
	username := u.GetString("username")
	password := u.GetString("password")
	md5PWD := utils.GetMd5File(password)
	age, _ := u.GetInt("age")
	gender := u.GetString("gender")
	phone, _ := u.GetInt("phone")
	addr := u.GetString("addr")
	isactive, _ := u.GetInt("isactive")
	usr := auth.User{
		UserName: username,
		Password: password,
		Age:      age,
		Gender:   gender,
		Phone:    int64(phone),
		Addr:     addr,
		IsActive: isactive,
		Creatime: time.Now(),
	}
	_, err := o.Insert(&usr)
	resp := make(map[string]interface{}, 2)
	if err != nil {
		info := fmt.Sprintf("insert info failed: username:%v, password:%v, age:%v, gender:%v, phone:%v, addr:%v, isactive:%v\n", username, md5PWD, age, gender, phone, addr, isactive)
		ret := fmt.Sprintf("insert into sys_user failed, error:%v\n", err)
		logs.Error(info)
		logs.Error(ret)
		resp["code"] = 1001
		resp["msg"] = fmt.Sprintf("insert into sys_user failed, error:%v\n", err)
	} else {
		info := fmt.Sprintf("insert info success: username:%v, password:%v, age:%v, gender:%v, phone:%v, addr:%v, isactive:%v\n", username, md5PWD, age, gender, phone, addr, isactive)
		logs.Info(info)
		resp["code"] = 200
		resp["msg"] = "添加成功"
	}
	u.Data["json"] = resp
	u.ServeJSON()
}

func (u *UserController) ActiveOrDeActive() {
	isactive, _ := u.GetInt("isactive_val")
	id, _ := u.GetInt("id")
	o := orm.NewOrm()
	qs := o.QueryTable("sys_user").Filter("id", id)
	resp := make(map[string]interface{}, 1)
	if isactive == 1 {
		_,_ = qs.Update(orm.Params{
			"is_active": 0,
		})
		ret := fmt.Sprintf("userId:%v deactice success\n",id)
		logs.Info(ret)
		resp["msg"] = "停用成功"
	} else {
		_,_ = qs.Update(orm.Params{
			"is_active": 1,
		})
		ret := fmt.Sprintf("userId:%v actice success\n",id)
		logs.Info(ret)
		resp["msg"] = "启用成功"
	}
	u.Data["json"] = resp
	u.ServeJSON()
}

func (u *UserController) Delete() {
	id, _ := u.GetInt("id")
	o := orm.NewOrm()
	_,_ = o.QueryTable("sys_user").Filter("id", id).Update(orm.Params{"is_del": 1})
	ret := fmt.Sprintf("userId:%v delete success\n",id)
	logs.Info(ret)
	resp := map[string]interface{}{"code": 200, "msg": "删除用户成功！"}
	u.Data["json"] = resp
	u.ServeJSON()
}

func (u *UserController) ResetPWD() {
	id, _ := u.GetInt("id")
	o := orm.NewOrm()
	_,_ = o.QueryTable("sys_user").Filter("id", id).Update(orm.Params{"password": "12345"})
	ret := fmt.Sprintf("userId:%v reset password success\n",id)
	logs.Info(ret)
	resp := map[string]interface{}{"code": 200, "msg": "重置密码成功！"}
	u.Data["json"] = resp
	u.ServeJSON()
}

func (u *UserController) Edit() {
	id, _ := u.GetInt("id")
	o := orm.NewOrm()
	usr := auth.User{}
	_ = o.QueryTable("sys_user").Filter("id", id).One(&usr)
	u.Data["usr"] = usr
	ret := fmt.Sprintf("userId:%v try to update\n",id)
	logs.Info(ret)
	u.TplName = "user/user-edit.html"
}

func (u *UserController) EditDone() {
	o := orm.NewOrm()
	username := u.GetString("username")
	password := u.GetString("password")
	age, _ := u.GetInt("age")
	gender := u.GetString("gender")
	phone, _ := u.GetInt("phone")
	addr := u.GetString("addr")
	isactive, _ := u.GetInt("isactive")
	uid, _ := u.GetInt("uid")
	qs := o.QueryTable("sys_user").Filter("id", uid)
	var err error
	_, err = qs.Update(orm.Params{
		"user_name": username,
		"password":  password,
		"age":       age,
		"gender":    gender,
		"phone":     phone,
		"addr":      addr,
		"is_active": isactive,
	})
	resp := make(map[string]interface{}, 2)
	if err == nil {
		ret := fmt.Sprintf("userId:%v update failed, error%v\n",uid,err)
		logs.Error(ret)
		resp["code"] = 200
		resp["msg"] = "更新成功"
	} else {
		ret := fmt.Sprintf("userId:%v update success\n",uid)
		logs.Info(ret)
		resp["code"] = 1001
		resp["msg"] = err
	}
	u.Data["json"] = resp
	u.ServeJSON()
}

func (u *UserController) MultiDel() {
	ids := strings.Split(u.GetString("ids"), ",")
	o := orm.NewOrm()
	qs := o.QueryTable("sys_user")
	var idInt int
	for i := 1; i < len(ids); i++ {
		idInt, _ = strconv.Atoi(ids[i])
		_,_ = qs.Filter("id", idInt).Update(orm.Params{
			"is_del": 1,
		})
	}
	resp := make(map[string]interface{}, 2)
	resp["code"] = 200
	resp["msg"] = "批量删除成功"
	u.Data["json"] = resp
	ret := fmt.Sprintf("userIds:%v multi delete success\n",ids)
	logs.Info(ret)
	u.ServeJSON()
}
