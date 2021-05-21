package testcache

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
)

type RedisController struct {
	beego.Controller
}

func (r *RedisController) Get() {
	//设置缓存的引擎为redis；连接名称为zhou（所有key的名称都为zhou:<key>，默认为beeCacheRedis:<key>）；ip端口；选第0个数据库。还可以有password配置
	ra, _ := cache.NewCache("redis", `{"key":"zhou","conn":"47.114.134.146:6379","dbnum":"0"}`)
	ra.Put("name", "zhangsan", 60)
	r.TplName = "test_cache/cache.html"
}
