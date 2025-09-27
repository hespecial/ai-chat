package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AssocCharacterSkillModel = (*customAssocCharacterSkillModel)(nil)

type (
	// AssocCharacterSkillModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAssocCharacterSkillModel.
	AssocCharacterSkillModel interface {
		assocCharacterSkillModel
		withSession(session sqlx.Session) AssocCharacterSkillModel
		GetSkillIds(ctx context.Context, characterId int64) ([]int64, error)
	}

	customAssocCharacterSkillModel struct {
		*defaultAssocCharacterSkillModel
	}
)

// NewAssocCharacterSkillModel returns a model for the database table.
func NewAssocCharacterSkillModel(conn sqlx.SqlConn) AssocCharacterSkillModel {
	return &customAssocCharacterSkillModel{
		defaultAssocCharacterSkillModel: newAssocCharacterSkillModel(conn),
	}
}

func (m *customAssocCharacterSkillModel) withSession(session sqlx.Session) AssocCharacterSkillModel {
	return NewAssocCharacterSkillModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customAssocCharacterSkillModel) GetSkillIds(ctx context.Context, characterId int64) ([]int64, error) {
	query := fmt.Sprintf("select `skill_id` from %s where `character_id` = ?", m.table)
	var skillIds []int64
	return skillIds, m.conn.QueryRowsCtx(ctx, &skillIds, query, characterId)
}
