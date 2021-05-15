package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type UploadController struct {
	beego.Controller
}

func (u *UploadController) Get() {
	u.Layout = "base.html"
	fmt.Println(beego.URLFor("LoginController.Get", "name", "zhangsan", "age", "18"))
	u.SetSession("username", "zhangsan") //设置session，然后在default.go的SessionController的Get方法中获取session
	u.TplName = "upload.html"
}

func (u *UploadController) Prepare() {
	u.EnableXSRF = false
}

func (u *UploadController) Post() {
	//获取上传的文件
	f, h, err := u.GetFile("upload_file")
	if err != nil {
		fmt.Println("获取上传的问失败", err)
		return
	}
	defer f.Close()
	fileName := h.Filename
	fmt.Println(fileName)
	//保存获取到的文件
	timeStr := strconv.FormatInt(int64(time.Now().Unix()), 10)

	u.SaveToFile("upload_file", "static/upload/"+timeStr+"-"+fileName)
	u.Data["json"] = map[string]interface{}{"code": 200, "message": "上传成功"}
	u.ServeJSON()
	// u.TplName = "success.html" // 使用了ajax就不需要了
}
