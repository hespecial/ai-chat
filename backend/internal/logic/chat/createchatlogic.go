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
	internalError := code.NewInternalError("failed to create chat")

	character, err := l.svcCtx.CharactersModel.FindOne(l.ctx, req.CharacterId)
	if errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(code.NewInvalidParamError(), "req: %+v", req)
	}
	if err != nil {
		return nil, errors.Wrapf(internalError, "get character by id error: %v", err)
	}

	response, err := l.svcCtx.LLM.CreateChat(prompt.Combine(character.Prompt, req.Content))
	if err != nil {
		return nil, errors.Wrapf(internalError, "call llm error: %v", err)
	}

	if len(response.Choices) == 0 {
		return nil, errors.Wrapf(internalError, "llm response message is empty")
	}
	content := response.Choices[0].Message.Content

	// 保存此轮对话记录
	historyId, err := l.saveRoundChat(character, req.Content, content)
	if err != nil {
		return nil, errors.Wrapf(internalError, "save round chat error: %v", err)
	}

	resp = &types.CreateChatResp{
		Id:      historyId,
		Content: content,
	}
	return
}

func (l *CreateChatLogic) saveRoundChat(character *model.Characters, userContent, assistantContent string) (int64, error) {
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
	return l.svcCtx.ChatHistoryModel.SaveRoundChat(context.Background(), &userChatHistory, &assistantChatHistory)
}
