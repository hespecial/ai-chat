package skill

import (
	"backend/pkg/response"
	"net/http"

	"backend/internal/logic/skill"
	"backend/internal/svc"
	"backend/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TuiYanHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SkillReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(w, err)
			return
		}

		l := skill.NewTuiYanLogic(r.Context(), svcCtx)
		resp, err := l.TuiYan(&req)
		response.HttpResult(w, resp, err)
	}
}
