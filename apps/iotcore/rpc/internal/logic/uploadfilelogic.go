package logic

import (
	"context"

	"goiot/apps/iotcore/rpc/internal/svc"
	"goiot/apps/iotcore/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadFileLogic) UploadFile(in *pb.UploadFileRequest) (*pb.UploadFileResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.UploadFileResponse{}, nil
}
