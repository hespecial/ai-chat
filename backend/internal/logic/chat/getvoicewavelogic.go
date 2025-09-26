package chat

import (
	"backend/internal/model"
	"backend/pkg/code"
	"context"
	"github.com/pkg/errors"
	"net/http"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVoiceWaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	writer http.ResponseWriter
}

func NewGetVoiceWaveLogic(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter) *GetVoiceWaveLogic {
	return &GetVoiceWaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		writer: w,
	}
}

func (l *GetVoiceWaveLogic) GetVoiceWave(req *types.GetVoiceWaveReq) (err error) {
	internalError := code.NewInternalError("failed to get voice wave")

	chatHistory, err := l.svcCtx.ChatHistoryModel.FindOne(l.ctx, req.ChatHistoryId)
	if err != nil {
		return errors.Wrapf(internalError, "get chat history by id error: %v", err)
	}
	if chatHistory.Role != model.RoleAssistant {
		return errors.Wrapf(code.NewIgnoredError("invalid param"), "illegal request: %v", req)
	}

	character, err := l.svcCtx.CharactersModel.FindOne(l.ctx, chatHistory.CharacterId)
	if err != nil {
		return errors.Wrapf(internalError, "get character by id error: %v", err)
	}

	voice, err := l.svcCtx.VoiceModel.FindOne(l.ctx, character.VoiceId)
	if err != nil {
		return errors.Wrapf(internalError, "get voice by id error: %v", err)
	}

	binData, err := l.svcCtx.LLM.TransferTextToVoice(chatHistory.Content, voice.Type)
	if err != nil {
		return errors.Wrapf(internalError, "llm transfer text to voice error: %v", err)
	}

	l.writer.Header().Set("Content-Type", "audio/mpeg")
	l.writer.Header().Set("Content-Disposition", "inline")
	l.writer.Header().Set("Cache-Control", "public, max-age=3600")
	if _, err = l.writer.Write(binData); err != nil {
		return err
	}

	return
}
