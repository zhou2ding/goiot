package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"goiot/apps/iotcore/api/internal/config"
	"goiot/apps/iotcore/api/internal/middleware"
)

type ServiceContext struct {
	Config                   config.Config
	ProcessReqRespMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                   c,
		ProcessReqRespMiddleware: middleware.NewProcessReqRespMiddleware().Handle,
	}
}
