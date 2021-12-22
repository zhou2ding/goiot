package auth

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"ions/models/auth"
	"ions/utils"
	"math"
	"time"
)

type AuthController struct {
	beego.Controller
}

func (a *AuthController) List() {
	var au []auth.Auth
	o := orm.NewOrm()
	qs := o.QueryTable("sys_auth")
	pageSize := 8
	curPage, err := a.GetInt("page")
	if err != nil {
		curPage = 1
	}
	kw := a.GetString("kw")
	var all int64
	all, _ = qs.Filter("is_del", 0).Filter("auth_name__contains", kw).Count()
	_, _ = qs.Filter("is_del", 0).Filter("auth_name__contains", kw).Limit(pageSize, (curPage-1)*pageSize).All(&au)
	ret := fmt.Sprintf("current page:%v, key word:%v\n", curPage, kw)
	logs.Info(ret)
	allPage := int(math.Ceil(float64(all) / float64(pageSize)))
	pageMap := utils.Paginator(curPage, pageSize, all)

	var prevPage int
	if curPage == 1 {
		prevPage = 1
	} else {
		prevPage = curPage - 1
	}
	var nexPage int
	if curPage >= allPage {
		nexPage = curPage
	} else {
		nexPage = curPage + 1
	}

	a.Data["pageMap"] = pageMap
	a.Data["allPage"] = allPage
	a.Data["prevPage"] = prevPage
	a.Data["nextPage"] = nexPage
	a.Data["all"] = all
	a.Data["au"] = au
	a.Data["kw"] = kw
	a.TplName = "auth/auth-list.html"
}

func (a *AuthController) ToAdd() {
	var au []auth.Auth
	o := orm.NewOrm()
	qs := o.QueryTable("sys_auth")
	_, _ = qs.Filter("is_del", 0).All(&au)
	a.Data["au"] = au
	a.TplName = "auth/auth-add.html"
}

func (a *AuthController) DoAdd() {
	parentId,_ := a.GetInt("auth_parent_id")
	authName := a.GetString("auth_name")
	authURL := a.GetString("auth_url")
	desc := a.GetString("auth_desc")
	active,_ := a.GetInt("is_active")
	weight,_ := a.GetInt("auth_weight")
	au := auth.Auth{
		ParentId:   parentId,
		AuthName:   authName,
		UrlFor:     authURL,
		Desc:       desc,
		IsActive:   active,
		Weight:     weight,
		CreateTime: time.Now(),
	}

	o := orm.NewOrm()
	_,err := o.Insert(&au)
	ret := fmt.Sprintf("insert into sys_auth info:%v\n",au)
	logs.Info(ret)
	resp := make(map[string]interface{},2)
	if err != nil {
		resp["code"] = 1001
		resp["msg"] = "插入数据有误！"
		ret := fmt.Sprintf("insert into sys_auth failed, error:%v\n",err)
		logs.Error(ret)
	} else {
		resp["code"] = 200
		resp["msg"] = "插入数据成功！"
		ret := fmt.Sprintf("insert into sys_auth success\n")
		logs.Info(ret)
	}
	a.Data["json"] = resp
	a.ServeJSON()
}

func (a *AuthController) Delete() {
	id,_ := a.GetInt("id")
	o := orm.NewOrm()
	qs := o.QueryTable("sys_auth")
	_,_ = qs.Filter("id",id).Update(orm.Params{"is_del":1})
	resp := map[string]interface{}{"code":200,"msg":"删除成功！"}
	a.Data["json"] = resp
	a.ServeJSON()
}
