package skill

import (
	"backend/internal/logic/skill/common"
	"backend/pkg/code"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GuanXingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGuanXingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GuanXingLogic {
	return &GuanXingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GuanXingLogic) GuanXing(req *types.SkillReq) (resp *types.SkillResp, err error) {
	internalError := code.NewInternalError("failed to use GuanXing skill")

	weatherResp, err := l.svcCtx.Weather.Query()
	if err != nil {
		return nil, errors.Wrapf(internalError, "query weather error: %v", err)
	}

	weatherInfo, err := json.Marshal(weatherResp.Lives[0])
	if err != nil {
		return nil, errors.Wrapf(internalError, "marshal weather error: %v", err)
	}

	historyId, content, err, errMsg := common.ExecCommonLogic(l.ctx, l.svcCtx, req.CharacterId, req.SkillId, string(weatherInfo))
	if err != nil {
		return nil, errors.Wrapf(internalError, fmt.Sprintf("%s: %v", errMsg, err))
	}

	resp = &types.SkillResp{
		Id:      historyId,
		Content: content,
	}
	return
}
