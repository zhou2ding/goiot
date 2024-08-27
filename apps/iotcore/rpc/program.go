package main

import (
	microSvc "github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"goiot/apps/iotcore/rpc/internal/config"
	"goiot/apps/iotcore/rpc/internal/server"
	"goiot/apps/iotcore/rpc/internal/svc"
	"goiot/apps/iotcore/rpc/pb"
	"goiot/interceptor"
	"goiot/pkg/cache"
	_ "goiot/pkg/cache"
	_ "goiot/pkg/database"
	"goiot/pkg/logger"
	"goiot/pkg/mq"
	"goiot/pkg/oss"
	"goiot/pkg/push"
	"goiot/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Program struct {
}

func (p *Program) Start(s service.Service) error {
	logger.Logger.Warnf("starting programme...")

	logger.Logger.Warnf("start programme...")

	c := *config.RpcConf
	ctx := svc.NewServiceContext(c)

	rpcSvc := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterRpcServer(grpcServer, server.NewRpcServer(ctx))
		if c.Mode == microSvc.DevMode || c.Mode == microSvc.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer rpcSvc.Stop()

	rpcSvc.AddUnaryInterceptors(interceptor.PreProcess)

	if err := cache.InitRedis(); err != nil {
		logger.Logger.Errorf("init redis error %s", err)
		_ = p.Stop(s)
		return err
	}

	if err := oss.InitS3Client(); err != nil {
		logger.Logger.Warnf("init s3 client error: %v", err)
		_ = p.Stop(s)
		return err
	}

	if err := push.InitSnsClient(); err != nil {
		logger.Logger.Warnf("init sns client error: %v", err)
		_ = p.Stop(s)
		return err
	}

	if err := push.InitSESClient(); err != nil {
		logger.Logger.Warnf("init ses client error: %v", err)
		_ = p.Stop(s)
		return err
	}

	if err := mq.InitSqsClient(); err != nil {
		logger.Logger.Warnf("init sqs client error: %v", err)
		_ = p.Stop(s)
		return err
	}

	return nil
}

func (p *Program) Stop(s service.Service) (e error) {
	logger.Logger.Warnf("stoping programme...")
	return nil
}
