package chat

import (
	"context"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TruncateChatHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTruncateChatHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TruncateChatHistoryLogic {
	return &TruncateChatHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TruncateChatHistoryLogic) TruncateChatHistory(req *types.TruncateChatHistoryReq) (resp *types.TruncateChatHistoryResp, err error) {
	return &types.TruncateChatHistoryResp{}, l.svcCtx.ChatHistoryModel.TruncateChat(l.ctx, req.Id)
}
