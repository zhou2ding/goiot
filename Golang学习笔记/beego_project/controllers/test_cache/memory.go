package testcache

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
)

type MemoryCacheController struct {
	beego.Controller
}

func (m *MemoryCacheController) Get() {
	ca, _ := cache.NewCache("memory", `{"interval":60}`) //缓存引擎设为:内存，60s后缓存过期
	//增加，所有引擎的方法都是统一的
	ca.Put("name", "张三", 60)
	ca.Put("age", 18, 60)
	//获取单个key
	name := ca.Get("name")
	//获取多个key，接收一个string切片，返回一个切片
	keys := []string{"name", "age"}
	vals := ca.GetMulti(keys)
	//判断是否存在
	ex := ca.IsExist("name")
	//删除
	ca.Delete("name")
	ca.ClearAll()
	//自增1、自减1
	ca.Incr("age")
	ca.Decr("age")

	fmt.Println(name, vals, ex)
	m.TplName = "test_cache/cache.html"
}
