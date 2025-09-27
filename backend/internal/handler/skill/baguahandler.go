package skill

import (
	"backend/pkg/response"
	"net/http"

	"backend/internal/logic/skill"
	"backend/internal/svc"
	"backend/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BaGuaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SkillReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(w, err)
			return
		}

		l := skill.NewBaGuaLogic(r.Context(), svcCtx)
		resp, err := l.BaGua(&req)
		response.HttpResult(w, resp, err)
	}
}
