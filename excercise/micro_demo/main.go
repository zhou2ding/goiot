package main

import (
	"micro_demo/handler"
	pb "micro_demo/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("micro_demo"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterMicro_demoHandler(srv.Server(), new(handler.Micro_demo))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
