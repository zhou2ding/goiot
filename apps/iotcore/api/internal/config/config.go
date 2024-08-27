package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		SigningKey   string
		AccessSecret string
		AccessExpire int64
		Issuer       string
	}
	RPC zrpc.RpcClientConf
}
