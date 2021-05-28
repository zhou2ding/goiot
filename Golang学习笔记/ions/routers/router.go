package routers

import (
	"ions/controllers"
	auth "ions/controllers/auth"
	login "ions/controllers/login"
	user "ions/controllers/user"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &login.LoginController{})
	beego.Router("/m/user-logout", &login.LoginController{},"get:LogOut")
	beego.Router("/change_captcha", &login.LoginController{}, "get:ChangeCaptcha")

	//后台首页
	beego.Router("/m/home", &controllers.HomeController{})
	beego.Router("/m/welcome", &controllers.HomeController{}, "get:Welcome")

	// 未登录需要拦截
	//user模块
	beego.Router("/m/user-list", &user.UserController{}, "get:List")
	beego.Router("/m/user-add", &user.UserController{}, "get:Add")
	beego.Router("/m/user-submit", &user.UserController{}, "post:Submit")
	beego.Router("/m/user-active", &user.UserController{}, "post:ActiveOrDeActive")
	beego.Router("/m/user-delete", &user.UserController{}, "post:Delete")
	beego.Router("/m/user-resetpwd", &user.UserController{}, "post:ResetPWD")
	beego.Router("/m/user-edit", &user.UserController{}, "get:Edit")
	beego.Router("/m/user-editdone", &user.UserController{}, "post:EditDone")
	beego.Router("/m/user-multidel", &user.UserController{}, "post:MultiDel")

	//权限模块
	beego.Router("/m/auth-list", &auth.AuthController{},"get:List")
	beego.Router("/m/auth-toadd", &auth.AuthController{},"get:ToAdd")
	beego.Router("/m/auth-doadd", &auth.AuthController{},"post:DoAdd")
	beego.Router("/m/auth-delete", &auth.AuthController{},"post:Delete")
	//角色模块
	beego.Router("m/role-list",&auth.RoleController{},"get:List")
	beego.Router("m/role-toadd",&auth.RoleController{},"get:ToAdd")
	beego.Router("m/role-doadd",&auth.RoleController{},"post:DoAdd")
	//角色模块的用户配置和权限配置
	beego.Router("m/role-to-user-add",&auth.RoleController{},"get:ToRoleUser")
	beego.Router("m/role-do-user-add",&auth.RoleController{},"post:DoRoleUser")
	beego.Router("m/role-to-auth-add",&auth.RoleController{},"get:ToRoleAuth")
	beego.Router("m/role-do-auth-add",&auth.RoleController{},"post:DoRoleAuth")
	beego.Router("m/role-get-auth-json",&auth.RoleController{},"get:GetAuthJson")

	//个人中心
	beego.Router("m/mycenter-edit",&user.MyCenterController{})
}
