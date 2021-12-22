package routers

import (
	"ions/controllers"
	auth "ions/controllers/auth"
	"ions/controllers/cars"
	"ions/controllers/echarts"
	"ions/controllers/finance"
	login "ions/controllers/login"
	"ions/controllers/news"
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
	beego.Router("m/mecenter-edit",&user.MyCenterController{})
	beego.Router("m/salary-list",&user.SalaryController{})
	beego.Router("m/salary-detail",&user.SalaryController{},"get:Detail")
	//财务中心
	beego.Router("/m/finace/salary-list",&finance.FinaController{})
	beego.Router("/m/finace/salary-toimport",&finance.FinaController{},"get:ToImport")
	beego.Router("/m/finace/salary-doimport",&finance.FinaController{},"post:DoImport")
	//echarts
	beego.Router("/m/finace/echarts-list",&finance.CaiWuEchartDataController{})
	beego.Router("/m/finace/echarts-toimport",&finance.CaiWuEchartDataController{},"get:ToImportExcel")
	beego.Router("/m/finace/echarts-doimport",&finance.CaiWuEchartDataController{},"post:DoImportExcel")
	//内容模块
	beego.Router("/m/news/category-list",&news.CategoryController{})
	beego.Router("/m/news/category-toadd",&news.CategoryController{},"get:ToAdd")
	beego.Router("/m/news/category-doadd",&news.CategoryController{},"post:DoAdd")
	beego.Router("/m/news/news-list",&news.NewsController{})
	beego.Router("/m/news/news-toadd",&news.NewsController{},"get:ToAdd")
	beego.Router("/m/news/news-doadd",&news.NewsController{},"post:DoAdd")

	//车辆模块
	beego.Router("/m/cars/brand-list",&cars.CarBrandController{})
	beego.Router("/m/cars/brand-toadd",&cars.CarBrandController{},"get:ToAdd")
	beego.Router("/m/cars/brand-doadd",&cars.CarBrandController{},"post:DoAdd")

	beego.Router("/m/cars/car-list",&cars.CarsController{})
	beego.Router("/m/cars/car-toadd",&cars.CarsController{},"get:ToAdd")
	beego.Router("/m/cars/car-doadd",&cars.CarsController{},"post:DoAdd")

	beego.Router("/m/cars/apply-list",&cars.CarsApplyController{})
	beego.Router("/m/cars/apply-toapply",&cars.CarsApplyController{},"get:ToApply")
	beego.Router("/m/cars/apply-doapply",&cars.CarsApplyController{},"post:DoApply")
	beego.Router("/m/cars/myapply",&cars.CarsApplyController{},"get:MyApply")

	beego.Router("/m/cars/audit-list",&cars.CarsApplyController{},"get:AuditApply")
	beego.Router("/m/cars/audit-toaudit",&cars.CarsApplyController{},"get:ToAuditApply")
	beego.Router("/m/cars/audit-doaudit",&cars.CarsApplyController{},"post:DoAuditApply")
	beego.Router("/m/cars/audit-doreturn",&cars.CarsApplyController{},"get:DoReturn")

	beego.Router("/m/notify-list",&controllers.HomeController{},"get:NotifyList")
	beego.Router("/m/notify-read",&controllers.HomeController{},"get:ReadNotify")

	//echarts模块
	beego.Router("/m/echarts_business",&echarts.EchartsBusinessController{})
	beego.Router("/m/echarts_caiwu",&echarts.EchartsCaiwuController{})
	beego.Router("/m/echarts_course",&echarts.EchartsCourseController{})
	beego.Router("/m/echarts_business_getecharts",&echarts.EchartsBusinessController{},"get:GetBusinessChart")
	beego.Router("/m/echarts_caiwu_getecharts",&echarts.EchartsCaiwuController{},"get:GetCaiWuChart")
	beego.Router("/m/echarts_course_getecharts",&echarts.EchartsCourseController{},"get:GetCourseChart")
}
