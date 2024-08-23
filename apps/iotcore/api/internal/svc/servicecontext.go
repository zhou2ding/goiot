package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"goiot/apps/iotcore/api/internal/config"
	"goiot/apps/iotcore/api/internal/middleware"
	"goiot/apps/iotcore/rpc/rpc"
)

type ServiceContext struct {
	Config                   config.Config
	ProcessReqRespMiddleware rest.Middleware
	RPC                      rpc.Rpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                   c,
		ProcessReqRespMiddleware: middleware.NewProcessReqRespMiddleware().Handle,
		RPC:                      rpc.NewRpc(zrpc.MustNewClient(c.RPC)),
	}
}
