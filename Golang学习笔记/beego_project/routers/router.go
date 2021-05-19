package routers

import (
	"beego_project/controllers"
	testorm "beego_project/controllers/test_orm"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/static_test", &controllers.StaticController{})

	//固定路由，访问方式：/param，/param?name=zhangsan
	beego.Router("/param", &controllers.ParamsController{})

	//正则路由，访问方式：/param，/param/zhangsan；但去掉？的话就必须带参数了，否则访问不到
	// "/param/?:id:int"可以限制类型，限制类型是string的话等于没限制；用正则的方式限制类型,int是id:([\d]+)或id:([0-9]+)，string是id:([\w]+)
	// beego.Router("/param/?:id", &controllers.ParamsController{})

	//自动路由，访问方式：/param，/param/get，/param/get/zhangsan；这里的param是截取的ParamsController前面的字符串，controller变的话前面的字符串也要变。/param/post的话必须提交表单才行，不能直接用域名
	// beego.AutoRouter(&controllers.ParamsController{})

	//自定义路由，要传第三个参数，把请求映射到controller的方法；controller必须有Login和Register方法，*:My代表所有请求都会映射到My方法
	// beego.Router("/param/?:id:int", &controllers.ParamsController{}, "get:Login; post:Register; get,post:Util; *:My")

	beego.Router("/other", &controllers.OtherTypeController{})
	beego.Router("/flash", &controllers.FlashController{})

	beego.Router("/t-router/?:id:int", &controllers.RouterController{}, "get,post:Util")
	beego.Router("/upload", &controllers.UploadController{})
	beego.Router("/sess", &controllers.SessionController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/validate", &controllers.ValidateController{})
	beego.Router("/abort", &controllers.AbortController{})

	beego.Router("/template1", &controllers.TemplateController{})
	beego.Router("/template2", &controllers.T2Controller{})
	beego.Router("/template3", &controllers.T3Controller{})
	beego.Router("/template4", &controllers.T4Controller{})
	beego.Router("/layout", &controllers.TestLayoutController{})
	beego.Router("/render", &controllers.RenderController{})
	beego.Router("/orm1", &testorm.TestOrm1Controller{})
	beego.Router("/orm2-list", &testorm.TestOrm2ListController{})
	beego.Router("/orm2-add", &testorm.TestOrm2AddController{})
	beego.Router("/orm2-update", &testorm.TestOrm2UpdateController{})
	beego.Router("/orm2-delete", &testorm.TestOrm2DeleteController{})
	beego.Router("/orm3", &testorm.TestOrm3Controller{})  // exeper表达式
	beego.Router("/orm4", &testorm.TestOrm4Controller{})  // querytable接口函数
	beego.Router("/orm5", &testorm.TestOrm5Controller{})  // 原生sql
	beego.Router("/o2o", &testorm.TestOrmO2OController{}) // 一对一模型
	beego.Router("/o2m", &testorm.TestOrmO2MController{}) // 一对多模型
	beego.Router("/m2m", &testorm.TestOrmM2MController{}) // 多对多模型
	//实战
	beego.Router("/po2o", &testorm.O2OController{}) // 一对一实战
	beego.Router("/po2m", &testorm.O2MController{}) // 一对一实战
	beego.Router("/pm2m", &testorm.M2MController{}) // 一对一实战
}
