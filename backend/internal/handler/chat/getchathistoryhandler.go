package chat

import (
	"backend/pkg/response"
	"net/http"

	"backend/internal/logic/chat"
	"backend/internal/svc"
	"backend/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetChatHistoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetChatHistoryReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(w, err)
			return
		}

		l := chat.NewGetChatHistoryLogic(r.Context(), svcCtx)
		resp, err := l.GetChatHistory(&req)
		response.HttpResult(w, resp, err)
	}
}
