package result

import (
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"goiot/pkg/defs"
	"goiot/pkg/errcode"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
)

func HttpResult(ctx context.Context, w http.ResponseWriter, data any, err error) {
	if err == nil {
		httpx.WriteJson(w, http.StatusOK, successResp(ctx, data))
	} else {
		st, ok := status.FromError(err)
		if ok {
			httpx.WriteJson(w, http.StatusInternalServerError, errorResp(ctx, int32(st.Code()), st.Message()))
		} else {
			ec := errcode.ServerError
			httpx.WriteJson(w, http.StatusInternalServerError, errorResp(ctx, int32(ec), ec.String()))
		}
	}
}

type ResponseSuccess struct {
	Code      int32  `json:"code"`
	Data      any    `json:"data,omitempty"`
	RequestId string `json:"requestId"`
	Duration  int64  `json:"duration"`
}

func successResp(ctx context.Context, data any) *ResponseSuccess {
	return &ResponseSuccess{200, data, getStrCtxVal(ctx, defs.RequestIdCtx), getCosts(ctx)}
}

type ResponseError struct {
	Code      int32  `json:"code"`
	Message   any    `json:"message"`
	RequestId string `json:"requestId"`
	Duration  int64  `json:"duration"`
}

func errorResp(ctx context.Context, code int32, msg string) *ResponseError {
	return &ResponseError{code, msg, getStrCtxVal(ctx, defs.RequestIdCtx), getCosts(ctx)}
}

func ParamErrorResult(ctx context.Context, w http.ResponseWriter, err error) {
	httpx.WriteJson(w, http.StatusBadRequest, errorResp(ctx, int32(errcode.ParamError), err.Error()))
}

func ErrorResultWithCode(ctx context.Context, w http.ResponseWriter, status int, code errcode.ErrCode) {
	httpx.WriteJson(w, status, errorResp(ctx, int32(code), code.String()))
}

func getStrCtxVal(ctx context.Context, ctxKey string) string {
	val, ok := ctx.Value(ctxKey).(string)
	if !ok {
		return ""
	}
	return val
}

func getCosts(ctx context.Context) int64 {
	startTime, ok := ctx.Value(defs.StartTimeCtx).(time.Time)
	if !ok {
		return 0
	}
	return time.Since(startTime).Milliseconds()
}
