package middleware

import (
	"context"
	"goiot/pkg/defs"
	"goiot/pkg/utils"
	"net/http"
)

type ProcessReqRespMiddleware struct {
	localIp string
}

func NewProcessReqRespMiddleware() *ProcessReqRespMiddleware {
	return &ProcessReqRespMiddleware{utils.GetLocalIP()}
}

func (m *ProcessReqRespMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), defs.RequestIdCtx, utils.GetUUIDFull())
		ctx = context.WithValue(ctx, defs.RemoteIpCtx, m.localIp)
		next(w, r.WithContext(ctx))
	}
}
