package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func index(c *gin.Context) {
	name,ok := c.Get("name")
	name = c.MustGet("name")
	name = c.GetString("name")
	fmt.Println(name,ok)

	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

//定义一个m1中间件
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	//中间件常用做法：计算耗时
	start := time.Now()
	c.Next() //马上调用后续的处理函数，即m2。不调用的话则是m1先走完，再根据注册的顺序走下一个函数
	//c.Abort()	//阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Println(cost)
	fmt.Println("m1 out ...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in ...")
	c.Set("name","zhangsan")
	c.Next()	//调用后续的处理函数，即index。
	fmt.Println("m2 out ...")
}

//实际业务的中间件，用闭包
func authMiddleware(doCheck bool)gin.HandlerFunc {
	//先进行一些其他工作，如连接数据库等
	return func(c *gin.Context) {
		if doCheck {
			//存放具体的逻辑，如判断用户是否是vip等
			//if 满足auth，则c.Next()
			//else，则c.Abort()
		}else {
			c.Next()
		}
	}
}
func main() {
	r := gin.Default()
	//全局注册m1中间件函数，这样的话就不用在r.GET中注册m1了。所有的请求都会先走中间件
	r.Use(m1, m2,authMiddleware(false))
	r.GET("/index", index)

	//为此路由单独注册一个中间件
	r.GET("/index", m1, index)

	//为此路由组注册一个中间件，方法1：
	xxGroup := r.Group("/xx",m1)
	//为此路由组注册一个中间件，方法2：
	xxGroup = r.Group("/xx")
	xxGroup.Use(m1)
	{
		xxGroup.GET("/abc", func(c *gin.Context) {
			c.JSON(200,"")
		})
	}
	_ = r.Run()
}
