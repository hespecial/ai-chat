package chat

import (
	"backend/pkg/response"
	"net/http"

	"backend/internal/logic/chat"
	"backend/internal/svc"
	"backend/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateChatReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(w, err)
			return
		}

		l := chat.NewCreateChatLogic(r.Context(), svcCtx)
		resp, err := l.CreateChat(&req)
		response.HttpResult(w, resp, err)
	}
}
