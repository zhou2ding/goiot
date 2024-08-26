package logic

import (
	"bytes"
	"context"
	"goiot/apps/iotcore/rpc/pb"
	"goiot/pkg/defs"
	"goiot/pkg/logger"
	"io"

	"goiot/apps/iotcore/api/internal/svc"
	"goiot/apps/iotcore/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadFileLogic) UploadFile(req *types.UploadFileRequest) (resp *types.UserUploadFileResponse, err error) {
	reqId := l.ctx.Value(defs.RequestIdCtx)

	stream, err := l.svcCtx.RPC.UploadFile(l.ctx)
	if err != nil {
		logger.Logger.Errorf("requestId: %v get stream error: %v", reqId, err)
		return
	}

	var (
		n      int
		buff   = make([]byte, 1024)
		reader = bytes.NewReader(req.File.Body)
	)
	for {
		n, err = reader.Read(buff)
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Logger.Errorf("requestId: %v read file body error: %v", reqId, err)
			return
		}

		if err = stream.Send(&pb.UploadFileRequest{Chunk: buff[:n]}); err != nil {
			logger.Logger.Errorf("requestId: %v send chunk to stream error: %v", reqId, err)
			return
		}
	}

	rpcResp, err := stream.CloseAndRecv()
	resp = &types.UserUploadFileResponse{
		FileId: rpcResp.GetFileId(),
	}
	return
}
