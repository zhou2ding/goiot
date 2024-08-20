package cache

import (
	"github.com/go-redis/redis"
	"goiot/pkg/conf"
	"goiot/pkg/logger"
	"sync"
)

var (
	defaultClient *redis.Client
	gClients      = make(map[string]*redis.Client)
	mu            sync.RWMutex
)

const (
	BlackListKey = "TokenBlackList:"
	APIKeyKey    = "APIKey:"
	AppIdKey     = "AppId:"
)

const (
	PermissionDB = "permission"
	IotDB        = "iot"
)

func InitRedis() error {
	sentinelAddrs := conf.Conf.GetStringSlice("redis.sentinel.addrs")
	masterName := conf.Conf.GetString("redis.sentinel.master")
	password := conf.Conf.GetString("redis.password")
	num := conf.Conf.GetInt("redis.num")

	defaultClient = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    masterName,
		SentinelAddrs: sentinelAddrs,
		Password:      password,
		DB:            num,
	})

	err := initInstance(PermissionDB, masterName, sentinelAddrs, num, password)
	if err != nil {
		return err
	}
	err = initInstance(IotDB, masterName, sentinelAddrs, num+1, password)
	if err != nil {
		return err
	}

	return nil
}

func initInstance(clientName, masterName string, sentinelAddrs []string, num int, password string) error {
	client := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    masterName,
		SentinelAddrs: sentinelAddrs,
		Password:      password,
		DB:            num,
	})
	if _, err := client.Ping().Result(); err != nil {
		logger.Logger.Errorf("connect to redis error: %v", err)
		return err
	}

	mu.Lock()
	gClients[clientName] = client
	mu.Unlock()

	return nil
}

func GetRedis(name string) *redis.Client {
	mu.RLock()
	defer mu.RUnlock()

	client, ok := gClients[name]
	if !ok {
		return defaultClient
	}

	return client
}
