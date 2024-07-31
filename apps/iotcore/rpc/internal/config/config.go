package config

import (
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

var RpcConf *Config

type Config struct {
	zrpc.RpcServerConf
	Gateway rest.RestConf
}

func InitRpcConf(path string) {
	RpcConf = new(Config)
	conf.MustLoad(path, RpcConf)
}
