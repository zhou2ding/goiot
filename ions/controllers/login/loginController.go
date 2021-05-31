package login

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"ions/models/auth"
	"ions/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	id, base64, err := utils.GetCaptcha()
	if err != nil {
		ret := fmt.Sprintf("get captcha failed, error:%v\n", err)
		logs.Error(ret)
		return
	}
	l.Data["captcha"] = utils.Captcha{Id: id, BS64: base64}
	l.TplName = "login/login.html"
}

func (l *LoginController) Post() {
	o := orm.NewOrm()
	qs := o.QueryTable("sys_user")
	usr := auth.User{}

	username := l.GetString("username")
	password := l.GetString("password")
	md5PWD := utils.GetMd5File(password)
	captcha := l.GetString("captcha")
	captcha_id := l.GetString("captcha_id")

	qs.Filter("user_name", username).Filter("password", password).One(&usr)
	uname_pwd_ok := qs.Filter("user_name", username).Filter("password", password).Exist()
	catcha_ok := utils.VerifyCaptcha(captcha_id, captcha)
	resp := make(map[string]interface{}, 1)
	if !uname_pwd_ok {
		ret := fmt.Sprintf("username:%v or password:%v wrong\n",username,md5PWD)
		logs.Info(ret)
		resp["code"] = 1001
		resp["msg"] = "用户名或密码错误"
	} else if !catcha_ok {
		ret := fmt.Sprintf("captcha is wrong:%v\n",catcha_ok)
		logs.Info(ret)
		resp["code"] = 1001
		resp["msg"] = "验证码错误"
	} else if usr.IsActive == 0 {
		ret := fmt.Sprintf("user:%v is deactived\n",username)
		logs.Info(ret)
		resp["code"] = 1001
		resp["msg"] = "该用户已停用，请联系管理员"
	} else {
		ret := fmt.Sprintf("username:%v or password:%v login success\n",username,password)
		logs.Info(ret)
		l.SetSession("id", usr.Id)
		resp["code"] = 200
		resp["msg"] = "登录成功"
	}
	l.Data["json"] = resp
	l.ServeJSON()

}

func (l *LoginController) ChangeCaptcha() {
	id, base64, err := utils.GetCaptcha()
	if err != nil {
		ret := fmt.Sprintf("get captcha:%v failed, error:%v\n", id,err)
		logs.Error(ret)
		l.Data["json"] = map[string]interface{}{"Code": 404, "msg": "生成验证码失败"}
		l.ServeJSON()
		return
	}
	l.Data["json"] = utils.Captcha{Id: id, BS64: base64, Code: 200}
	l.ServeJSON()
}

func (l *LoginController) LogOut() {
	l.DelSession("id")
	l.Redirect(beego.URLFor("LoginController.Get"),302)
}
