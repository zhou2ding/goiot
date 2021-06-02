# net.http

```go
//HandleFunc把域名路由到对应的函数，函数必须接收响应和请求两个参数
func login(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "pls login")
}
//ListenAndServe监听端口，接收两个参数：ip端口；一个handler(一般置为nil)
func main() {
    _,_ = http.HandleFunc("/login",login)
    http.ListenAndServe("127.0.0.1:8080",nil)
}
```

# Gin初识

> 当前后端是只返回json格式的数据给代理服务器（nginx等），然后前台拿到数据后渲染预期的效果
>
> RESTful API风格：用4种http的请求来代表业务中不同的动作（get、put、post、delete

```go
func login(c *gin.Context){
    //把gin.H结构体对象封装成json返回给前端
    r.JSON(gin.H{
        "code":200,
        "msg":"hello gin!",
    })
}

func main() {
    r := gin.Deatult()		//初始化一个默认路由引擎
    r.Get("/login",login)	//把get请求的url路由到login函数
    r.Run("127.0.0.1:8080")	//监听端口
}
```

# template模板

> 模板：使用相应的数据去替换html文档中事先准备好的标记
>
> html/template包有防恶意注入功能；text/template没有

## 使用

1. 定义模板

   - `tmpl`或`tpl`文件

2. 解析模板

   - 解析字符串：`t,_ := template.Parse`

   - 解析文件：

     - `t,_ := template.ParseFiles(filenames ...string)`

     - `t,_ := template.New(name string).ParseFiles(filenames ...string)`

       > new的作用是创建模板对象，接收的参数要和parsefile的参数一致。（一般先new，然后可以调用模板对象的其他方法）

   - `t,_ := template.ParseGlog(pattern string)`：用正则的方式解析模板文件名

3. 模板渲染

   - `t.Execute(wr io.Writer, data interface{})`：前端通过`{{.}}`获取数据

     > 相当于beego的`t.Data["data"]`
     >
     > 结构体的字段名需要大写；map的key或普通变量不用大写
     >
     > 有多个数据要传是时，execute的第二个参数是`map[string]interface{}`，key随便叫，value是要传的变量，前端通过`{{.value}}`可获取数据

   - `t.ExecuteTemplate(wr io.Writer, name string, data interface{})`：有继承模板时使用

## 语法

- 自定义模板函数

  ```go
  // 1. 定义的函数要么只有一个返回值，要么两个且第二个必须是error
  // 2. 注册这个函数
  t := template.New("f.tmpl")
  t.Func(
      template.FuncMap{
          "前端用的函数名":定义好的函数名
  	}
  )
  ```

- 模板嵌套和自定义模板

  -  template.ParseFiles的时候，传两个参数，第一个是主模板文件名，第二个是被嵌套的模板
  - 自定义模板是，要带文件后缀`{{define "my.tmpl"}}`

- 模板继承（类似beego的layout）

  - 根模板：在`base.tmpl`中放上`{{block "content" .}} {{end}}`，作为各个页面中显示内容不同的地方（content可以改成其他，但继承此模板的文件define的名字和此处必须一致）

  - 在其他需要用到base的模板文件中：先继承根模板，再重新定义区块

    ```html
    {{template "base.html" .}}
    {{define "content"}}
    	这里的内容会被替换到base中{{block "content" .}} {{end}}的部分
    {{end}}
    ```

  - 解析文件时：`t.ParseFIles()`先传根模板文件名，再传其他模板文件名

  - 渲染模板时：要使用`t.ExecuteTemplate`

- 修改标识符：`template.New("h.tmpl").Delims("[[","]]").ParseFiles`

  > beego中是beego.BConfig.WebConfig.TemplateLeft("[[")

- `html/template`包有防恶意注入功能，不会转义html格式的参数，如果确实需要转义，则自定义一个模板函数：接收string，把参数通过`template.HTML()`强转后返回