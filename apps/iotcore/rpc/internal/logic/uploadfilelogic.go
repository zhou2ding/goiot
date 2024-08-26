package logic

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"goiot/interceptor"
	"goiot/pkg/defs"
	"goiot/pkg/errcode"
	"goiot/pkg/logger"
	"goiot/pkg/oss"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
	"strings"

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

func (l *UploadFileLogic) UploadFile(stream pb.Rpc_UploadFileServer) error {
	reqId := interceptor.GetMetadata(l.ctx, strings.ToLower(defs.RequestIdCtx))

	md, ok := metadata.FromIncomingContext(l.ctx)
	if !ok {
		logger.Logger.Warnf("requestId: %s missing metadata", reqId)
		return status.Error(codes.Code(errcode.ParamError), errcode.ParamError.String())
	}

	var data []byte
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		data = append(data, req.Chunk...)
	}

	_, err := oss.GetS3Client().PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(md["bucket"][0]),
		Key:    aws.String(md["uploadkey"][0]),
		Body:   bytes.NewReader(data),
	})
	if err != nil {
		logger.Logger.Errorf("requestId %s upload file to oss error: %v", reqId, err)
		return err
	}

	return nil
}
