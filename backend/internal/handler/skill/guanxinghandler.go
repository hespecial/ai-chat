package skill

import (
	"backend/pkg/response"
	"net/http"

	"backend/internal/logic/skill"
	"backend/internal/svc"
	"backend/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GuanXingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SkillReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(w, err)
			return
		}

		l := skill.NewGuanXingLogic(r.Context(), svcCtx)
		resp, err := l.GuanXing(&req)
		response.HttpResult(w, resp, err)
	}
}
