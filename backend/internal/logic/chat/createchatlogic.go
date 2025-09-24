package chat

import (
	"backend/pkg/code"
	"backend/pkg/prompt"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateChatLogic {
	return &CreateChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateChatLogic) CreateChat(req *types.CreateChatReq) (resp *types.CreateChatResp, err error) {
	character, err := l.svcCtx.CharactersModel.FindOne(l.ctx, req.CharacterId)
	if errors.Is(err, sqlx.ErrNotFound) {
		return nil, errors.Wrapf(code.NewInvalidParamError(), "req: %+v", req)
	}
	if err != nil {
		return nil, errors.Wrapf(code.NewInternalError("failed to create chat"), "get character by id error: %v", err)
	}

	response, err := l.svcCtx.LLM.Call(prompt.Combine(character.Prompt, req.Content))
	if err != nil {
		return nil, errors.Wrapf(code.NewInternalError("failed to create chat"), "call llm error: %v", err)
	}

	resp = &types.CreateChatResp{
		Content: response.Choices[0].Message.Content,
	}
	return
}
