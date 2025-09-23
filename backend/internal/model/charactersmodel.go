package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CharactersModel = (*customCharactersModel)(nil)

type (
	// CharactersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCharactersModel.
	CharactersModel interface {
		charactersModel
		withSession(session sqlx.Session) CharactersModel
		List(ctx context.Context) ([]*Characters, error)
	}

	customCharactersModel struct {
		*defaultCharactersModel
	}
)

// NewCharactersModel returns a model for the database table.
func NewCharactersModel(conn sqlx.SqlConn) CharactersModel {
	return &customCharactersModel{
		defaultCharactersModel: newCharactersModel(conn),
	}
}

func (m *customCharactersModel) withSession(session sqlx.Session) CharactersModel {
	return NewCharactersModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customCharactersModel) List(ctx context.Context) ([]*Characters, error) {
	query := fmt.Sprintf("select %s from %s", charactersRows, m.table)
	var list []*Characters
	return list, m.conn.QueryRowsCtx(ctx, &list, query)
}
