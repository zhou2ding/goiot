package config

import (
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
)

var RpcConf *Config

type Config struct {
	zrpc.RpcServerConf
}

func InitRpcConf(path string) {
	RpcConf = new(Config)
	conf.MustLoad(path, RpcConf)
}
