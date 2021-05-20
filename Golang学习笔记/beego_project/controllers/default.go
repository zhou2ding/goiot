package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type UserStruct struct {
	Age      int    `form:"age"`
	UserName string `form:"username"`
	Price    string `form:"price"`
	Married  string `form:"married"`
}

type UserController struct {
	beego.Controller
}

func (u *UserController) Get() {
	//字符串渲染
	// u.Data["Username"] = "zhangsan"
	// u.Data["Age"] = 18

	//数组渲染
	arr := [...]int{9, 5, 6, 456, 0, 63}
	u.Data["arr"] = arr

	//结构体渲染
	user := UserStruct{
		Age:      23,
		UserName: "lisi",
	}
	u.Data["user"] = user
	//map渲染
	tmap := make(map[string]string, 1)
	tmap["name"] = "wangwu"
	tmap["age"] = "22"
	u.Data["tmap"] = tmap

	u.TplName = "user.html"
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type StaticController struct {
	beego.Controller
}

func (s *StaticController) Get() {
	s.TplName = "static_test.html"
}

type ParamsController struct {
	beego.Controller
}

func (p *ParamsController) Get() {
	//获取前端发送的数据，http://127.0.0.1:8080/param?name=xxx
	//方式1
	name := p.GetString("name")
	fmt.Println("====", name)
	//方式2
	age := p.Input().Get("age")
	fmt.Println(">>>>>>>>>>>>>", age)

	//获取前端发送的数据，http://127.0.0.1:8080/param/1111111，此外router.go中的url也要改
	//方式1
	id := p.Ctx.Input.Param(":id")
	fmt.Println(id)
	//方式2
	id2 := p.GetString(":id")
	fmt.Println(id2)
	p.TplName = "param_test.html"
}

func (p *ParamsController) Post() {
	//获取前端提交的表单，方式1
	//GetString里传的参数是param_test.html中的<form action="/param" method="POST">，里面的<input type="text" name="username">
	// username := p.GetString("username")   //获取单个
	// usernames := p.GetStrings("username") //获取多个，如果前端的input，多个name都是"username"，则都传进GetStrings返回的数组中
	// age := p.GetString("age")
	// fmt.Printf("姓名：%v，批量姓名：%v，年龄：%v\n", username, usernames, age)

	//方式2
	// username1 := p.Input().Get("username")
	// age1 := p.Input().Get("age")
	// fmt.Printf("姓名：%v，年龄：%v\n", username1, age1)

	//方式3，提交的表单绑定结构体
	// usr := UserStruct{}
	// err := p.ParseForm(&usr)
	// if err != nil {
	// 	fmt.Println("解析出错：", err)
	// 	return
	// }
	// fmt.Println(usr)
	// p.TplName = "success.html"

	//方式4，获取ajax的数据，app.conf要加个copyrequestbody=true，且不需要p.TplName了
	var usr UserStruct
	body := p.Ctx.Input.RequestBody //二进制的json数据
	err := json.Unmarshal(body, &usr)
	if err != nil {
		fmt.Println("json解析失败", err)
		return
	}
	fmt.Println(usr)
	result := map[string]interface{}{"code": "200", "message": "处理成功"}
	p.Data["json"] = result //key必须是"json"
	p.ServeJSON()
}

//其他数据类型的获取方式，也可以都用GetString
// func (p *ParamsController) Post() {
// 	name := p.GetString("username")
// 	age, _ := p.GetInt("age")
// 	price, _ := p.GetFloat("price")
// 	married, _ := p.GetBool("married")
// 	fmt.Printf("姓名：%v，年龄：%v，价格：%v，婚姻状况：%v\n", name, age, price, married)

// 	p.TplName = "success.html"
// }

//非json的其他数据类型的传输方式
type OtherTypeController struct {
	beego.Controller
}

func (o *OtherTypeController) Get() {
	usr := UserStruct{Age: 18, UserName: "张三", Price: "3.14", Married: "true"}
	//json格式
	// o.Data["json"] = &usr
	// o.ServeJSON()

	//xml格式
	o.Data["xml"] = &usr
	o.ServeXML()

	//jsonp格式
	// o.Data["jsonp"] = &usr
	// o.ServeJSONP()

	//yaml格式，运行时要下载一个文件，下载后打开才是yaml
	// o.Data["yml"] = &usr
	// o.ServeYAML()
}

type FlashController struct {
	beego.Controller
}
type FlashStruc struct {
	UserName string `json:"username"`
	Password string `json:"pwd"`
}

func (f *FlashController) Get() {
	flash := beego.ReadFromRequest(&f.Controller)
	err := flash.Data["error"]
	notice := flash.Data["notice"]
	if len(notice) != 0 {
		f.TplName = "success.html"
	} else if len(err) != 0 {
		f.TplName = "error.html"
	} else {
		f.TplName = "flash.html"
	}
}

func (f *FlashController) Post() {
	//旧的获取数据方式
	username := f.GetString("username")
	pwd := f.GetString("pwd")
	fmt.Println(username, pwd)
	//初始化flash
	flash := beego.NewFlash()
	if len(username) == 0 {
		flash.Error("用户名不能为空")
		flash.Store(&f.Controller)
		f.Redirect("/flash", 302)
	} else if pwd != "123456" {
		flash.Error("密码错误")
		flash.Store(&f.Controller)
		f.Redirect("/flash", 302)
	} else {
		fmt.Println("成功")
		flash.Notice("成功")
		flash.Store(&f.Controller)
		f.Redirect("/flash", 302)
	}

	//新的通过ajax获取数据方式
	// var flash_struct FlashStruc
	// body := f.Ctx.Input.RequestBody
	// err := json.Unmarshal(body, &flash_struct)
	// if err != nil {
	// 	fmt.Println("json解析失败：", err)
	// 	return
	// }
	// fmt.Printf("Name:%v, Pwd:%v\n", flash_struct.UserName, flash_struct.Password)
	// // flash := beego.NewFlash()
	// if len(flash_struct.UserName) == 0 {
	// 	result := map[string]interface{}{"code": "401", "message": "用户名不能为空"}
	// 	f.Data["json"] = result
	// 	f.ServeJSON()
	// 	// flash.Error("用户名不能为空")
	// 	// flash.Store(&f.Controller)
	// 	// f.Redirect("/flash", 302)
	// } else if flash_struct.Password != "123456" {
	// 	result := map[string]interface{}{"code": "401", "message": "密码错误"}
	// 	f.Data["json"] = result
	// 	f.ServeJSON()
	// 	// flash.Error("密码错误")
	// 	// flash.Store(&f.Controller)
	// 	// f.Redirect("/flash", 302)
	// } else {
	// 	result := map[string]interface{}{"code": "200", "message": "处理成功"}
	// 	f.Data["json"] = result
	// 	f.ServeJSON()
	// 	// flash.Notice("成功")
	// 	// flash.Store(&f.Controller)
	// 	// f.Redirect("/flash", 302)
	// }
}

type SessionController struct {
	beego.Controller
}

func (s *SessionController) Get() {
	usrName := s.GetSession("username") //获取session
	fmt.Println(usrName)

	s.DelSession("username") //删除session
	usrName2 := s.GetSession("username")
	fmt.Println(usrName2)
	s.TplName = "success.html"
}
