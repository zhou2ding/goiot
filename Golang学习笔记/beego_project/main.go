package main

import (
	"beego_project/controllers"
	_ "beego_project/routers"

	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//默认引入了models的init，如果models下面有嵌套文件夹的话，要手动引入其init
)

func init() {
	username := beego.AppConfig.String("username")
	pwd := beego.AppConfig.String("pwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	db := beego.AppConfig.String("db")
	dataSource := username + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + db + "?charset=utf8"
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		fmt.Println("连接mysql驱动出错：", err)
		return
	}
	//不管注册几个数据库，必须要有一个default。参数分别是：数据库别名。驱动。连接信息（用户名、密码、tcp、ip、域名、数据库名、字符集）。
	err = orm.RegisterDataBase("default", "mysql", dataSource)
	if err != nil {
		fmt.Println("连接数据库出错：", err)
		return
	}
}

func main() {
	// beego.SetViewsPath("front") //配置文件中也可配置
	// beego.SetStaticPath("/static", "front")
	// beego.BConfig.WebConfig.AutoRender = false //关闭自动渲染。配置文件中也可配置
	// beego.AddTemplateExt("jsp")                //把模板后缀从.tpl修改为.jsp
	// port, _ := beego.AppConfig.Int("httpport") //就是固定读取的app.conf，改名的话就读取不到

	// xsrf配置，也可在app.conf中设置
	// beego.BConfig.WebConfig.EnableXSRF = true
	// 设为true后，POST请求时必须携带xsrf参数
	// 接上，携带方法：先在html增加一个post表单和submit，然后在Get方法中增加：
	// （1）t.Data["xsfrdata"] = template.HTML(t.XSRFFormHTML())
	// （2）html页面的<form method="POST">中获取这个key：{{.xsfrdata}}
	//  (3)需要关闭时，在controller中写一个Prepare方法把此controller的EnableXSRF设为false
	// beego.BConfig.WebConfig.XSRFKey = "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o" //默认值"beegoxsrf"
	// beego.BConfig.WebConfig.XSRFExpire = 4800                                    //默认值3600秒

	beego.BConfig.MaxMemory = 1 << 22 //文件上传后是放在系统的内存里，设置缓存中的内存大小

	beego.BConfig.WebConfig.Session.SessionOn = true //配置文件中也可配置
	// beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 4800 // 其他session配置
	// beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 4800
	beego.InsertFilter("unlogin/", beego.BeforeRouter, controllers.FilterUser)

	beego.ErrorController(&controllers.ErrorController{}) //注册自定义错误的controller

	beego.AddFuncMap("hi", controllers.Hello) //把T4Controller的自定义view模板函数Hello注册一下，第一个参数是前端中使用的函数名；第二个参数不能传引用
	beego.AddFuncMap("sub", controllers.SubstrAndReplace)

	orm.Debug = true //开启sql语句的打印
	orm.RunCommand() //开启通过orm的命令行操作数据库
	//不通过命令行，直接在代码中自动建表
	// name := "default"
	// force := false
	// verbose := true
	// err := orm.RunSyncdb(name, force, verbose)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	logs.SetLogFuncCallDepth(3) //默认是4；直接调用logs包则改为3；对logs包封装后则是2，以此类推
	// logs.SetLevel(logs.LevelCritical) //设置日志级别，推荐用此方式

	//console引擎；开启颜色；级别是2（或者写LevelCritical）
	// logs.SetLogger(logs.AdapterConsole, `{"level":"2","color":true}`)

	//file引擎
	//filename：日志文件保存的路径和文件名
	//maxlines:保存的最大行数，默认100万行，超过后覆盖前面的行
	//maxsize：最大尺寸，默认1<<28（256MB）
	//daily：是否每天logrotate，默认true
	//maxdays：最多保存几天，默认7天
	//rotate：是否开启logrotate（循环覆盖），默认true
	//level：级别，默认Trace
	//perm：权限，默认775
	// logs.SetLogger(logs.AdapterFile, `{"filename":"logs/my.log"}`)

	//multifile引擎，最常用；separate：把数组中给的级别分别写入不同的文件（调用log打印的函数也按照给的级别来，比如error级别往上的都用error来输出）
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/beeproject.log","separate":["error","info"]}`)

	beego.Run()
}
