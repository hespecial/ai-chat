package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SkillModel = (*customSkillModel)(nil)

type (
	// SkillModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSkillModel.
	SkillModel interface {
		skillModel
		withSession(session sqlx.Session) SkillModel
	}

	customSkillModel struct {
		*defaultSkillModel
	}
)

// NewSkillModel returns a model for the database table.
func NewSkillModel(conn sqlx.SqlConn) SkillModel {
	return &customSkillModel{
		defaultSkillModel: newSkillModel(conn),
	}
}

func (m *customSkillModel) withSession(session sqlx.Session) SkillModel {
	return NewSkillModel(sqlx.NewSqlConnFromSession(session))
}
