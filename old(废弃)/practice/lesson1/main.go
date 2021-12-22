package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"path"
)

func hello(c *gin.Context) {
	//第二个参数是模板名，没有define命名模板时就是文件名
	//第三个参数是传给前端的参数，gin.H就是一个快捷的map[string]interface{}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "ZhangSan",
	})
}

func json(c *gin.Context) {
	//方法1：使用map
	data := gin.H{
		"name":"张三",
		"age":18,
	}
	c.JSON(200,data)
	//方法2：结构体
	type msg struct{
		Name string
		Age int
	}
	data2 := msg{
		Name: "张四",
		Age: 19,
	}

	c.JSON(200,data2)
}

func queryString(c *gin.Context) {
	//方法1
	kw := c.Query("kw")
	//方法2
	kw = c.DefaultQuery("kw","默认值")
	//方法3
	kw, ok := c.GetQuery("kw")
	if !ok {
		kw = "默认值"
	}

	c.JSON(200,gin.H{
		"kewWord":kw,
	})
}

func queryStringPath(c *gin.Context) {
	kw1 := c.Param("kw1")
	kw2 := c.Param("kw2")
	c.JSON(200,gin.H{
		"keyword1":kw1,
		"keyword2":kw2,
	})
}

func toLogin(c *gin.Context) {
	c.HTML(http.StatusOK,"login.html",gin.H{})
}
func doLogin(c *gin.Context) {
	//方式1
	userName := c.PostForm("username")
	password := c.PostForm("password")
	//方式2
	userName = c.DefaultPostForm("username","默认值")
	password = c.DefaultPostForm("password","默认值")
	//方式3
	userName, ok1 := c.GetPostForm("username")
	password, ok2 := c.GetPostForm("password")
	if !ok1 {
		userName = "默认值"
	}
	if !ok2 {
		password = "默认值"
	}

	c.HTML(200,"welcome.html",gin.H{
		"username":userName,
		"pwd":password,
	})
}

type User struct {
	Username string	`form:"username"`
	Password string	`form:"password"`
}
func bindStruct(c *gin.Context) {
	u := User{}
	_ = c.ShouldBind(&u)
	c.JSON(200,gin.H{
		"msg":"ok",
	})
}

func toUpload(c *gin.Context) {
	c.HTML(200,"upload.html",nil)
}

func doUpload(c *gin.Context) {
	file,_ := c.FormFile("f1")
	filePath := path.Join("./",file.Filename)
	_ = c.SaveUploadedFile(file,filePath)
	c.JSON(200,gin.H{
		"msg":"ok",
	})
}

func toUploadMulti(c *gin.Context) {
	c.HTML(200,"uploadMulti.html",nil)
}
func doUploadMulti(c *gin.Context) {
	form,_ := c.MultipartForm()
	files := form.File["file"]
	for _, file := range files {
		filePath := path.Join("./",file.Filename)
		_ = c.SaveUploadedFile(file,filePath)
	}
	c.JSON(200,gin.H{
		"msg":"ok",
	})
}

func redirect (c *gin.Context) {
	//重定向某个域名，跳转后URL会改变
	c.Redirect(http.StatusMovedPermanently,"www.baidu.com")
	//重定向到/n的路由处理函数
	c.Request.URL.Path = "/n"	//修改URI
	gin.New().HandleContext(c)	//后续的处理
}

func main() {
	r := gin.Default()

	//自定义模板函数
	r.SetFuncMap(template.FuncMap{
		"str2html": func(s string) template.HTML {
			return template.HTML(s)
		},
	})

	//设置/static这个URL的路径为./static；beego的SetStaticPath
	r.Static("/static","./static")

	//模板解析
	//r.LoadHTMLFiles("template/index.tmpl")
	//解析一坨模板，**表示所有目录，*表示所有文件。template没有子目录时直接一个*
	r.LoadHTMLGlob("template/*")

	//模板渲染
	r.GET("/hello", hello)
	//给前端返回json
	r.GET("/json",json)

	//获取queryString
	r.GET("/queryString",queryString)
	//获取URL中的path参数，path前面最好不要直接是根路径，否则会路由冲突
	r.GET("/queryStringPath/:kw1/:kw2",queryStringPath)
	//获取form表单提交的数据
	r.GET("/login",toLogin)
	r.POST("/login",doLogin)
	//queryString的参数或者form表单提交的参数来绑定结构体。万能函数。还有其他的ShouldBindJSON,ShouldBindXML等
	r.GET("/bindStruct",bindStruct)
	r.POST("/bindStruct",bindStruct)
	//上传文件
	r.MaxMultipartMemory = 1 << 20	//上传文件的大小超过这个限制时，会先把已经接收到内存的落盘到临时文件中，再继续接收
	r.GET("/upload",toUpload)
	r.POST("/upload",doUpload)
	//上传多个文件
	r.GET("/uploadMulti",toUploadMulti)
	r.POST("/uploadMulti",doUploadMulti)
	//请求重定向
	r.GET("/m",redirect)

	//不想对一个URL写一堆get、post，就用any
	r.Any("/", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK,gin.H{"method":"GET"})
		case http.MethodPost:
			c.JSON(200,gin.H{"method":"POST"})
		}
	})
	//访问不存在的URL会路由到这里
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound,gin.H{"msg":"找不到"})
	})
	//路由组：某个路径下有多个不同的子路径（
	userGroup := r.Group("/user")
	{
		userGroup.GET("/cars", func(c *gin.Context) {
			c.JSON(200,gin.H{})
		})
		userGroup.GET("salary", func(c *gin.Context) {
			c.JSON(200,gin.H{})
		})
		//路由组的嵌套
		infoGroup := userGroup.Group("/info")
		infoGroup.GET("/username", func(c *gin.Context) {
			c.JSON(200,gin.H{})
		})
	}

	userGroup.GET("/index2")
	_ = r.Run("127.0.0.1:8080")
}
