package chat

import (
	"backend/internal/model"
	"backend/pkg/code"
	"backend/pkg/prompt"
	"context"
	"encoding/json"
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

	// 查询该角色的历史对话消息
	histories, err := l.svcCtx.ChatHistoryModel.List(l.ctx, character.Id)
	if err != nil {
		return nil, errors.Wrapf(internalError, "list history by characterId error: %v", err)
	}

	// 对历史消息进行分组、拼接
	historiesStr, err := l.historyGroupSplice(histories)
	if err != nil {
		return nil, errors.Wrapf(internalError, "history group error: %v", err)
	}

	// 进行完全上下文召回
	response, err := l.svcCtx.LLM.CreateChat(prompt.Combine(character.Prompt, historiesStr, req.Content))
	if err != nil {
		return nil, errors.Wrapf(internalError, "call llm error: %v", err)
	}

	if len(response.Choices) == 0 {
		return nil, errors.Wrapf(internalError, "llm response message is empty")
	}
	content := response.Choices[0].Message.Content

	// 保存此轮对话记录
	historyId, err := l.SaveRoundChat(character, req.Content, content)
	if err != nil {
		return nil, errors.Wrapf(internalError, "save round chat error: %v", err)
	}

	resp = &types.CreateChatResp{
		Id:      historyId,
		Content: content,
	}
	return
}

func (l *CreateChatLogic) SaveRoundChat(character *model.Characters, userContent, assistantContent string) (int64, error) {
	now := time.Now().Unix()
	assistantChatHistory := model.ChatHistory{
		CharacterId: character.Id,
		Role:        model.RoleAssistant,
		Content:     assistantContent,
		Created:     now,
	}
	userChatHistory := model.ChatHistory{
		CharacterId: character.Id,
		Role:        model.RoleUser,
		Content:     userContent,
		Created:     now,
	}
	return l.svcCtx.ChatHistoryModel.SaveRoundChat(context.Background(), &userChatHistory, &assistantChatHistory)
}

type message struct {
	Round   int    `json:"round"`
	Role    string `json:"role"`
	Content string `json:"content"`
}

func (l *CreateChatLogic) historyGroupSplice(histories []*model.ChatHistory) (string, error) {
	if len(histories) == 0 {
		return "", nil
	}
	groupMap := make(map[int][]*message)
	for i, history := range histories {
		round := i/2 + 1
		if (history.Role == model.RoleUser && len(groupMap[round]) > 0) || (history.Role == model.RoleAssistant && len(groupMap[round]) == 0) {
			continue
		}
		groupMap[round] = append(groupMap[round], &message{
			Round:   round,
			Role:    history.Role,
			Content: history.Content,
		})
	}
	historiesStr, err := json.Marshal(groupMap)
	if err != nil {
		return "", err
	}
	return string(historiesStr), nil
}
