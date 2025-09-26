package svc

import (
	"backend/internal/config"
	"backend/internal/model"
	"backend/pkg/llm"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"log"
)

type ServiceContext struct {
	Config           config.Config
	CharactersModel  model.CharactersModel
	ChatHistoryModel model.ChatHistoryModel
	VoiceModel       model.VoiceModel
	LLM              *llm.LLM
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn, err := sqlx.NewConn(sqlx.SqlConf{
		DataSource: c.DB.DataSource,
		DriverName: "mysql",
		Replicas:   nil,
	})
	if err != nil {
		log.Fatalf("connect database err: %v", err)
	}
	return &ServiceContext{
		Config:           c,
		CharactersModel:  model.NewCharactersModel(conn),
		ChatHistoryModel: model.NewChatHistoryModel(conn),
		VoiceModel:       model.NewVoiceModel(conn),
		LLM:              llm.NewLLM(c.LLM.BaseUrl, c.LLM.ApiKey),
	}
}
