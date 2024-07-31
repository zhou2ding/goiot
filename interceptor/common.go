package interceptor

import (
	"context"
	"goiot/pkg/defs"
	"goiot/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

func PreProcess(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	logger.Logger.Infof("request %s from %s", GetMetadata(ctx, strings.ToLower(defs.RequestIdCtx)), GetMetadata(ctx, strings.ToLower(defs.RemoteIpCtx)))
	return handler(ctx, req)
}

func GetMetadata(ctx context.Context, key string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if val, ok := md[key]; ok && len(val) > 0 {
			return val[0]
		}
	}
	return ""
}

func GetTokenFromMetadata(ctx context.Context) string {
	var token string
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		tokenSlice, ok := md[strings.ToLower(defs.TokenCtx)]
		if ok && len(tokenSlice) != 0 {
			tokenStr := tokenSlice[0]
			parts := strings.SplitN(tokenStr, " ", 2)
			if len(parts) == 2 && parts[0] == "Bearer" {
				token = parts[1]
			}
		}
	}
	return token
}
