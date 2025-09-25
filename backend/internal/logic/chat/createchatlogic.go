package chat

import (
	"backend/internal/model"
	"backend/pkg/code"
	"backend/pkg/prompt"
	"context"
	"github.com/pkg/errors"
	"time"

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
	if errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(code.NewInvalidParamError(), "req: %+v", req)
	}
	if err != nil {
		return nil, errors.Wrapf(code.NewInternalError("failed to create chat"), "get character by id error: %v", err)
	}

	response, err := l.svcCtx.LLM.Call(prompt.Combine(character.Prompt, req.Content))
	if err != nil {
		return nil, errors.Wrapf(code.NewInternalError("failed to create chat"), "call llm error: %v", err)
	}

	if len(response.Choices) == 0 {
		return nil, errors.Wrapf(code.NewInternalError("failed to create chat"), "llm response message is empty")
	}
	content := response.Choices[0].Message.Content

	go func(character *model.Characters, userContent, assistantContent string) {
		assistantChatHistory := model.ChatHistory{
			CharacterId: character.Id,
			Role:        model.RoleAssistant,
			Content:     assistantContent,
			Created:     time.Now().Unix(),
		}
		userChatHistory := model.ChatHistory{
			CharacterId: character.Id,
			Role:        model.RoleUser,
			Content:     userContent,
			Created:     time.Now().Unix(),
		}
		if err = l.svcCtx.ChatHistoryModel.SaveRoundChat(context.Background(), &userChatHistory, &assistantChatHistory); err != nil {
			l.Logger.Errorf("failed to insert chat history, err: %v", err)
		}
	}(character, req.Content, content)

	resp = &types.CreateChatResp{
		Content: content,
	}
	return
}
