package chat

import (
	"backend/pkg/code"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetChatHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatHistoryLogic {
	return &GetChatHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChatHistoryLogic) GetChatHistory(req *types.GetChatHistoryReq) (resp *types.GetChatHistoryResp, err error) {
	records, err := l.svcCtx.ChatHistoryModel.List(l.ctx, req.Id)
	if err != nil {
		return nil, errors.Wrapf(code.NewInternalError("failed to get chat history"), "err: %v, req: %+v", err, req)
	}

	resp = &types.GetChatHistoryResp{
		Histories: make([]*types.ChatHistory, 0),
	}
	for _, record := range records {
		history := types.ChatHistory{}
		_ = copier.Copy(&history, record)
		resp.Histories = append(resp.Histories, &history)
	}
	return
}
