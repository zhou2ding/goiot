package testcache

import (
	"errors"
	"time"

	"github.com/astaxie/beego/cache"
)

type MyCache struct {
}

//必须要注册一下
func init() {
	cache.Register("mycache", NewOwnCache)
}

//构造函数
func NewOwnCache() cache.Cache {
	return &MyCache{}
}

//实现cache接口的所有方法
func (m *MyCache) Put(key string, val interface{}, timeout time.Duration) error {
	err := errors.New("")
	return err
}
func (m *MyCache) Get(key string) interface{} {
	return ""
}
func (m *MyCache) GetMulti(keys []string) []interface{} {
	v := make([]interface{}, 0, 10)
	return v

}
func (m *MyCache) Delete(key string) error {
	err := errors.New("")
	return err
}

func (m *MyCache) IsExist(key string) bool {
	return true
}

func (m *MyCache) ClearAll() error {
	err := errors.New("")
	return err
}
func (m *MyCache) StartAndGC(config string) error {
	err := errors.New("")
	return err
}
func (m *MyCache) Incr(key string) error {
	err := errors.New("")
	return err
}
func (m *MyCache) Decr(key string) error {
	err := errors.New("")
	return err
}
