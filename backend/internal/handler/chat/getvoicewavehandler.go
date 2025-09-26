package chat

import (
	"backend/pkg/response"
	"net/http"

	"backend/internal/logic/chat"
	"backend/internal/svc"
	"backend/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetVoiceWaveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetVoiceWaveReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(w, err)
			return
		}

		l := chat.NewGetVoiceWaveLogic(r.Context(), svcCtx, w)
		if err := l.GetVoiceWave(&req); err != nil {
			response.HttpResult(w, nil, err)
		}
	}
}
