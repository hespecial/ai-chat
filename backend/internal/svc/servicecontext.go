package svc

import (
	"backend/internal/config"
	"backend/internal/model"
	"backend/pkg/llm"
	"backend/pkg/news"
	"backend/pkg/weather"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"log"
)

type ServiceContext struct {
	Config                   config.Config
	CharactersModel          model.CharactersModel
	ChatHistoryModel         model.ChatHistoryModel
	VoiceModel               model.VoiceModel
	AssocCharacterSkillModel model.AssocCharacterSkillModel
	SkillModel               model.SkillModel
	LLM                      *llm.LLM
	Weather                  *weather.ApiConfig
	News                     *news.ApiConfig
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
		Config:                   c,
		CharactersModel:          model.NewCharactersModel(conn),
		ChatHistoryModel:         model.NewChatHistoryModel(conn),
		VoiceModel:               model.NewVoiceModel(conn),
		AssocCharacterSkillModel: model.NewAssocCharacterSkillModel(conn),
		SkillModel:               model.NewSkillModel(conn),
		LLM:                      llm.NewLLM(c.LLM.BaseUrl, c.LLM.ApiKey),
		Weather:                  weather.New(c.WeatherApi.Url, c.WeatherApi.Key, c.WeatherApi.City),
		News:                     news.New(c.HotNewsApi.Url),
	}
}
