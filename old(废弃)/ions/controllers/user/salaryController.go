package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"ions/models/auth"
	"ions/models/salary"
	"time"
)

type SalaryController struct {
	beego.Controller
}

func (s *SalaryController) Get() {
	month := s.GetString("month")
	if month == "" {
		month = time.Now().Format("2006-01")
	}
	usr := auth.User{}
	sal := salary.Salary{}
	uid := s.GetSession("id")
	o := orm.NewOrm()
	_ = o.QueryTable("sys_user").Filter("id",uid).One(&usr)
	_ = o.QueryTable("sys_salary").Filter("card_id",usr.CardId).Filter("pay_date",month).One(&sal)

	s.Data["sal"] = sal
	s.TplName = "user/salary-list.html"
}

func (s *SalaryController) Detail() {
	id,_ := s.GetInt("id")
	sal := salary.Salary{}
	o := orm.NewOrm()
	_ = o.QueryTable("sys_salary").Filter("id",id).One(&sal)
	s.Data["sal"] = sal
	s.TplName = "user/salary-detail.html"
}