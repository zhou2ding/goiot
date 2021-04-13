# 常用命令

`msconfig`

`shell:startup`

# Go bilibili day01学习笔记

## 课程物料

### 预习URL

[讲师博客](https://www.liwenzhou.com/)

### 视频

[哔哩哔哩](https://www.bilibili.com/video/BV16E411H7og)

### 课后作业

在[路飞学城](www.luffycity.com)注册账号

## Go语言特点

1. 语法简洁，自动格式化（`go fmt main.go`）

2. 开发效率高

3. 执行性能高（最接近Java的）

4. 天生支持并发（07年立项，05年AMD发布第一个双核处理器）

5. 编译型语言（C等，没有虚拟机，解释型语言有：PHP、Python等）

## Go语言开发环境搭建

### 安装Go开发包

1. Go镜像站：https://golang.google.cn
2. `go version`查看版本

## 配置GOPATH

1. 新建一个目录存放代码，里面新建三个文件夹，`bin`、`src`、`pkg`
2. 环境变量PATH中添加.\bin
3. 环境变量里新建一项：`GOPATH:D\go`
4. Go 1.8版本后有默认值`%USERPROFILE%/go`，删掉即可

`go env`

![1609861832477](C:\Users\ZhouDing\Desktop\Golang学习笔记\Go bilibili学习笔记.assets\1609861832477.png)

## Go项目结构

### 目前流行

![1609861976911](C:\Users\ZhouDing\Desktop\Golang学习笔记\Go bilibili学习笔记.assets\1609861976911.png)

### 企业常用

![1609862069823](C:\Users\ZhouDing\Desktop\Golang学习笔记\Go bilibili学习笔记.assets\1609862069823.png)



## VScode

### 常用设置

![1610032701901](C:\Users\ZhouDing\Desktop\Golang学习笔记\Go bilibili学习笔记.assets\1610032701901.png)

![1610032725303](C:\Users\ZhouDing\Desktop\Golang学习笔记\Go bilibili学习笔记.assets\1610032725303.png)

![1610275372290](C:\Users\ZhouDing\Desktop\Golang学习笔记\Go bilibili学习笔记.assets\1610275372290.png)

![1610275421943](C:\Users\ZhouDing\Desktop\Golang学习笔记\Go bilibili学习笔记.assets\1610275421943.png)

### 安装Go扩展

### VScode界面

1. `Ctrl + Shift + E` 文件资源管理器
2. `Ctrl + Shift + F` 跨文件搜素
3. `Ctrl + Shift + G` 源代码管理器
4. `Ctrl + Shift + D` 启动和调试
5. `Ctrl + Shift +X` 管理扩展
6. `Ctrl + Shift + M` 查看错误和警告
7. `Ctrl + Shift + P`查找并运行所有命令
8. `Ctrl + ·` 切换集成终端

### 配置代码快捷键

`ctrl + shift + p`，输入`snippets`，选择`Go`，在go.json中添加如下代码块；`$0`表示光标最终停留的位置。

``` json
{
	"println":{
		"prefix": "pln",
		"body":"fmt.Println($0)",
		"description": "println"
	},
	"printf":{
		"prefix": "plf",
		"body": "fmt.Printf(\"$0\")",
		"description": "printf"
	}
}
```

### 编译

#### go build

1. 首先在根目录下`go mod init go.mod`
2. 在项目目录下执行
3. 在其他路径下执行，需要在后面加上项目的路径（从GOPATH/src后开始写起，编译之后的exe就在当前目录下）
4. `go build -o name.exe`

#### go run

向执行脚本一样执行Go代码

`go run main.go`

#### go install

1. 先编译得到一个可执行文件
2. 将可执行文件拷贝到`GOPATH/bin`

#### 交叉编译

Go支撑跨平台编译

```bash
SET CGO_ENABLE=0 // 禁用CGO
SET GOOS=linux // 目标平台是linux
SET GOARCH=arm64 // 目标处理器架构是arm64
```

# Go bilibili day02学习笔记

## 变量和常量

Go语言中的变量必须先声明再使用

Go语言中非全局变量声明必须使用，不使用就编译不过去

### 声明变量

`var s1 string`：声明一个保存字符串类型数据的变量

`var s1 string = "ok"`：声明的时候同时赋值（不是很推荐，因为有类型推导）

`var s1 = "ok"`：类型推导

`s1 := "ok"`：简短变量声明，只能在函数里用

`_`：匿名变量（多用于占位，表示忽略

``` go
// 批量声明
var (
	name1 string
	age1  int
	isOk1 bool
)
```

推荐使用小驼峰

同一作用域不能重复声明同一个变量

常用：批量（全局） + 简短（非全局）

### 常量

###  **Tips：**观察下面的代码，体味此处`_`的妙用

```go
// 摘自gin框架routergroup.go
type IRouter interface{ ... }

type RouterGroup struct { ... }

var _ IRouter = &RouterGroup{}  // 确保RouterGroup实现了接口IRouter
```















   



