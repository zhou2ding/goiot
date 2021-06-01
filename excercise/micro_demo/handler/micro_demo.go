package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	micro_demo "micro_demo/proto"
)

type Micro_demo struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Micro_demo) Call(ctx context.Context, req *micro_demo.Request, rsp *micro_demo.Response) error {
	log.Info("Received Micro_demo.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Micro_demo) Stream(ctx context.Context, req *micro_demo.StreamingRequest, stream micro_demo.Micro_demo_StreamStream) error {
	log.Infof("Received Micro_demo.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&micro_demo.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Micro_demo) PingPong(ctx context.Context, stream micro_demo.Micro_demo_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&micro_demo.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
