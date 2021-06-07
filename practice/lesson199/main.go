package main

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var rdb *redis.Client

func initRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: "",
		Password: "",
		DB: 0,
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
	return
}

var logger *zap.Logger
func initLogger() {
	logger,_ = zap.NewProduction()
}

func main() {
	//redis相关操作
	if err := initRedis();err != nil {
		return
	}
	defer rdb.Close()
	pipe := rdb.Pipeline()
	pipe.Incr("")
	pipe.Set("","",0)
	pipe.Get("").Val()
	_,_ = pipe.Exec()
	err := rdb.Watch(func(tx *redis.Tx) error {
		_, err := rdb.Pipelined(func(pipeliner redis.Pipeliner) error {
			pipeliner.Set("cn1",1,0)
			pipeliner.Incr("cn1")
			return nil
		})
		return err
	},"key")
	if err != nil {
		return
	}

	//zap相关操作
	initLogger()
}
