package svc

import (
	"backend/internal/config"
	"backend/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"log"
)

type ServiceContext struct {
	Config          config.Config
	CharactersModel model.CharactersModel
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
		Config:          c,
		CharactersModel: model.NewCharactersModel(conn),
	}
}
