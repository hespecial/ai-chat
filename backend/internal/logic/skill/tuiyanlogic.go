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

type TuiYanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTuiYanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TuiYanLogic {
	return &TuiYanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TuiYanLogic) TuiYan(req *types.SkillReq) (resp *types.SkillResp, err error) {
	internalError := code.NewInternalError("failed to use TuiYan skill")

	lastHistory, err := l.svcCtx.ChatHistoryModel.LastHistory(l.ctx, req.CharacterId)
	if err != nil {
		return nil, errors.Wrapf(internalError, fmt.Sprintf("find last chat history error: %v", err))
	}

	infoMap := make(map[string]string)
	if lastHistory == nil {
		infoMap["result"] = "现在是第一轮对话，我之前没有和其他角色对话过"
	} else {
		character, err := l.svcCtx.CharactersModel.FindOne(l.ctx, lastHistory.CharacterId)
		if err != nil {
			return nil, errors.Wrapf(internalError, "find character by id err: %v", req.CharacterId)
		}
		infoMap = map[string]string{
			"tip":                   "这是我上次交谈的人，你需要告诉我是你推演到了这个结果",
			"character_name":        "角色名称：" + character.Name,
			"character_description": "角色描述：" + character.Description,
		}
	}

	infoStr, err := json.Marshal(infoMap)
	if err != nil {
		return nil, errors.Wrapf(internalError, "marshal info map error: %v", infoMap)
	}

	historyId, content, err, errMsg := common.ExecCommonLogic(l.ctx, l.svcCtx, req.CharacterId, req.SkillId, string(infoStr))
	if err != nil {
		return nil, errors.Wrapf(internalError, fmt.Sprintf("%s: %v", errMsg, err))
	}

	resp = &types.SkillResp{
		Id:      historyId,
		Content: content,
	}

	return
}
