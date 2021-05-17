package main

import (
	"beego_project/controllers"
	_ "beego_project/routers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
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
	// beego.SetViewsPath("front")	//配置文件中也可配置
	// beego.SetStaticPath("/static", "front")
	// beego.BConfig.WebConfig.AutoRender = false //配置文件中也可配置
	// beego.AddTemplateExt("jsp")
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
	// beego.InsertFilter("*/", beego.BeforeRouter, controllers.FilterUser)

	beego.ErrorController(&controllers.ErrorController{}) //注册自定义错误的controller

	beego.AddFuncMap("hi", controllers.Hello) //把T4Controller的自定义view模板函数Hello注册一下，第一个参数是前端中使用的函数名；第二个参数不能传引用
	beego.AddFuncMap("sub", controllers.SubstrAndReplace)

	orm.Debug = true //开启sql语句的打印
	orm.RunCommand() //开启通过orm的命令行操作数据库
	//不通过命令行，直接在代码中自动建表
	name := "default"
	force := false
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
		return
	}

	beego.Run()
}
