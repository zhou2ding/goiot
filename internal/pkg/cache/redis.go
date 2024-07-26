package cache

import (
	"github.com/go-redis/redis"
	"goiot/internal/pkg/conf"
	"goiot/internal/pkg/logger"
)

var gRedis *redis.Client

func init() {
	sentinelAddrs := conf.Conf.GetStringSlice("redis.sentinel.addrs")
	masterName := conf.Conf.GetString("redis.sentinel.master")
	pwd := conf.Conf.GetString("redis.password")
	num := conf.Conf.GetInt("redis.num")

	cli := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    masterName,
		SentinelAddrs: sentinelAddrs,
		Password:      pwd,
		DB:            num, // 添加数据库选择
	})
	if _, err := cli.Ping().Result(); err != nil {
		logger.Logger.Errorf("connect to redis error: %v", err)
		return
	}

	gRedis = cli
}

func GetRedis() *redis.Client {
	return gRedis
}
