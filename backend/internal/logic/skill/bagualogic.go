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

type BaGuaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBaGuaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaGuaLogic {
	return &BaGuaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BaGuaLogic) BaGua(req *types.SkillReq) (resp *types.SkillResp, err error) {
	internalError := code.NewInternalError("failed to use BaGua skill")

	newsResp, err := l.svcCtx.News.Query()
	if err != nil {
		return nil, errors.Wrapf(internalError, "query news error: %v", err)
	}
	hotNews := newsResp.Data

	if len(hotNews) > 10 {
		hotNews = hotNews[:10]
	}
	hotNewsStr, err := json.Marshal(hotNews)
	if err != nil {
		return nil, errors.Wrapf(internalError, "marshal news error: %v", err)
	}

	historyId, content, err, errMsg := common.ExecCommonLogic(l.ctx, l.svcCtx, req.CharacterId, req.SkillId, string(hotNewsStr))
	if err != nil {
		return nil, errors.Wrapf(internalError, fmt.Sprintf("%s: %v", errMsg, err))
	}

	resp = &types.SkillResp{
		Id:      historyId,
		Content: content,
	}
	return
}
