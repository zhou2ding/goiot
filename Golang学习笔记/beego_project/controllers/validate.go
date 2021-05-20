package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type ValidateController struct {
	beego.Controller
}

type UserInfo struct {
	UserName string `form:"username" valid:"Required"`
	Age      int    `form:"age" valid:"Required"`
	Email    string `form:"email" valid:"Required;Email"`
	Phone    string `form:"phone" valid:"Required;Phone"`
}

func (v *ValidateController) Get() {
	v.TplName = "validate.html"
}

func (v *ValidateController) Prepare() {
	v.EnableXSRF = false
}

func (v *ValidateController) Post() {
	// //数据校验
	// username := v.GetString("username")
	// age := v.GetString("age")
	// email := v.GetString("email")
	// phone := v.GetString("phone")
	// //初始化validation
	// valid := validation.Validation{}
	// valid.Required(username, "用户名").Message("不能为空") //判断前端提交过来的username这个变量是否为空，第二个参数是下面err的key
	// valid.Required(age, "年龄").Message("不能为空")
	// valid.Required(email, "邮箱").Message("不能为空")
	// valid.Required(phone, "电话").Message("不能为空")
	// valid.Email(email, "邮箱").Message("邮箱格式不合法")
	// valid.Phone(phone, "手机号").Message("手机号格式不合法")
	// if valid.HasErrors() {
	// 	for _, err := range valid.Errors {
	// 		fmt.Printf("%v:%v\n", err.Key, err.Message) //key是Required中的第二个参数
	// 	}
	// }

	//解析到结构体并校验数据
	// usr := UserInfo{}
	// v.ParseForm(&usr)
	// fmt.Println(usr)
	// valid := validation.Validation{}
	// valid.Required(usr.UserName, "用户名").Message("不能为空")
	// valid.Required(usr.Age, "年龄").Message("不能为空")
	// valid.Required(usr.Email, "邮箱").Message("不能为空")
	// valid.Required(usr.Phone, "电话").Message("不能为空")
	// if valid.HasErrors() {
	// 	for _, err := range valid.Errors {
	// 		fmt.Printf("%v: %v\n", err.Key, err.Message)
	// 	}
	// }

	//使用struct Tag校验
	usr := UserInfo{}
	v.ParseForm(&usr)
	valid := validation.Validation{}
	msg := map[string]string{
		"Required": "不能为空",
		"Phone":    "格式错误",
		"Email":    "格式错误",
	}
	validation.SetDefaultMessage(msg)
	b, err := valid.Valid(&usr) //一个是校验结果，如果校验失败，b=false；一个是前面的struct tag写的是否有问题
	if err != nil {
		fmt.Println(err)
	}
	if !b {
		for _, err := range valid.Errors {
			fmt.Printf("%v:%v\n", err.Key, err.Message)
		}
	}

	v.TplName = "success.html"
}
