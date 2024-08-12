package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goiot/apps/iotcore/api/internal/logic"
	"goiot/apps/iotcore/api/internal/svc"
	"goiot/apps/iotcore/api/internal/types"
)

func uploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadFileRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUploadFileLogic(r.Context(), svcCtx)
		resp, err := l.UploadFile(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
