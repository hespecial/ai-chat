package common

import (
	"backend/internal/model"
	"backend/internal/svc"
	"backend/pkg/prompt"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"time"
)

func ExecCommonLogic(ctx context.Context, svcCtx *svc.ServiceContext, characterId, skillId int64, extraInfo string) (int64, string, error, string) {
	character, err := svcCtx.CharactersModel.FindOne(ctx, characterId)
	if err != nil {
		return 0, "", err, "find character error"
	}

	skill, err := svcCtx.SkillModel.FindOne(ctx, skillId)
	if err != nil {
		return 0, "", err, "find skill error"
	}

	skillMap := map[string]string{
		"name":        skill.Name,
		"description": skill.Description,
	}
	skillStr, err := json.Marshal(skillMap)
	if err != nil {
		return 0, "", err, "json marshal skill map error"
	}

	response, err := svcCtx.LLM.CreateChat(prompt.Combine(character.Prompt, "", skill.Name+extraInfo, string(skillStr)))
	if err != nil {
		return 0, "", err, "call llm error"
	}

	if len(response.Choices) == 0 {
		return 0, "", errors.New("llm response message is empty"), "call llm error"
	}
	content := response.Choices[0].Message.Content

	historyId, err := saveRoundChat(svcCtx, character, skill.Name, content)
	if err != nil {
		return 0, "", err, "save round chat error"
	}

	return historyId, content, nil, ""
}

func saveRoundChat(svcCtx *svc.ServiceContext, character *model.Characters, userContent, assistantContent string) (int64, error) {
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
	return svcCtx.ChatHistoryModel.SaveRoundChat(context.Background(), &userChatHistory, &assistantChatHistory)
}
