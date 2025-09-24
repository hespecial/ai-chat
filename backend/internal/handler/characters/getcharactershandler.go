package characters

import (
	"backend/pkg/response"
	"net/http"

	"backend/internal/logic/characters"
	"backend/internal/svc"
	"backend/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCharactersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCharactersReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(w, err)
			return
		}

		l := characters.NewGetCharactersLogic(r.Context(), svcCtx)
		resp, err := l.GetCharacters(&req)
		response.HttpResult(w, resp, err)
	}
}
