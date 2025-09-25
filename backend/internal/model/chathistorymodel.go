package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ChatHistoryModel = (*customChatHistoryModel)(nil)

type (
	// ChatHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChatHistoryModel.
	ChatHistoryModel interface {
		chatHistoryModel
		withSession(session sqlx.Session) ChatHistoryModel
		SaveRoundChat(ctx context.Context, userRecord, assistantRecord *ChatHistory) error
		List(ctx context.Context, characterId int64) ([]*ChatHistory, error)
		TruncateChat(ctx context.Context, characterId int64) error
	}

	customChatHistoryModel struct {
		*defaultChatHistoryModel
	}
)

// NewChatHistoryModel returns a model for the database table.
func NewChatHistoryModel(conn sqlx.SqlConn) ChatHistoryModel {
	return &customChatHistoryModel{
		defaultChatHistoryModel: newChatHistoryModel(conn),
	}
}

func (m *customChatHistoryModel) withSession(session sqlx.Session) ChatHistoryModel {
	return NewChatHistoryModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customChatHistoryModel) SaveRoundChat(ctx context.Context, userRecord, assistantRecord *ChatHistory) error {
	return m.conn.Transact(func(session sqlx.Session) error {
		tx := m.withSession(session)
		if _, err := tx.Insert(ctx, userRecord); err != nil {
			return err
		}
		if _, err := tx.Insert(ctx, assistantRecord); err != nil {
			return err
		}
		return nil
	})
}

func (m *customChatHistoryModel) List(ctx context.Context, characterId int64) ([]*ChatHistory, error) {
	query := fmt.Sprintf("select %s from %s where `character_id` = ? order by `created` asc, id asc", chatHistoryRows, m.table)
	var resp []*ChatHistory
	return resp, m.conn.QueryRowsCtx(ctx, &resp, query, characterId)
}

func (m *customChatHistoryModel) TruncateChat(ctx context.Context, characterId int64) error {
	query := fmt.Sprintf("delete from %s where `character_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, characterId)
	return err
}
