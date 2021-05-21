package testcache

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
)

type FileCacheController struct {
	beego.Controller
}

func (f *FileCacheController) Get() {
	//设置缓存引擎为file，缓存目录是./cache；文件名后缀是.cache；目录层级是2；过期时间是60s
	fa, _ := cache.NewCache("file", `{"CachePath":"./cache",""FileSuffix:".cache","DirectoryLevel":"2","EmbedExpire":"60"}`)
	fa.Put("name", "zhangsan", 60)

	f.TplName = "test_cache/cache.html"
}
