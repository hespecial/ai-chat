package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ VoiceModel = (*customVoiceModel)(nil)

type (
	// VoiceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVoiceModel.
	VoiceModel interface {
		voiceModel
		withSession(session sqlx.Session) VoiceModel
	}

	customVoiceModel struct {
		*defaultVoiceModel
	}
)

// NewVoiceModel returns a model for the database table.
func NewVoiceModel(conn sqlx.SqlConn) VoiceModel {
	return &customVoiceModel{
		defaultVoiceModel: newVoiceModel(conn),
	}
}

func (m *customVoiceModel) withSession(session sqlx.Session) VoiceModel {
	return NewVoiceModel(sqlx.NewSqlConnFromSession(session))
}
