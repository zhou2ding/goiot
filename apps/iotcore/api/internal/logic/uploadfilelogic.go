package logic

import (
	"context"
	"goiot/apps/iotcore/rpc/pb"

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
	rpcResp, err := l.svcCtx.RPC.UploadFile(l.ctx, &pb.UploadFileRequest{
		Bucket: req.Bucket,
	})
	resp = &types.UserUploadFileResponse{
		FileId: rpcResp.GetFileId(),
	}
	return
}
