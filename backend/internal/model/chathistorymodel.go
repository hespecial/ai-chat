package model

import (
	"context"
	"database/sql"
	"errors"
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
		SaveRoundChat(ctx context.Context, userRecord, assistantRecord *ChatHistory) (int64, error)
		List(ctx context.Context, characterId int64) ([]*ChatHistory, error)
		TruncateChat(ctx context.Context, characterId int64) error
		LastHistory(ctx context.Context, characterId int64) (*ChatHistory, error)
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

func (m *customChatHistoryModel) SaveRoundChat(ctx context.Context, userRecord, assistantRecord *ChatHistory) (int64, error) {
	var historyId int64
	err := m.conn.Transact(func(session sqlx.Session) error {
		tx := m.withSession(session)
		if _, err := tx.Insert(ctx, userRecord); err != nil {
			return err
		}
		result, err := tx.Insert(ctx, assistantRecord)
		if err != nil {
			return err
		}

		if historyId, err = result.LastInsertId(); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, nil
	}
	return historyId, nil
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

func (m *customChatHistoryModel) LastHistory(ctx context.Context, characterId int64) (*ChatHistory, error) {
	query := fmt.Sprintf("select %s from %s where `character_id` != ? and role = ? order by `created` desc limit 1", chatHistoryRows, m.table)
	var chatHistory ChatHistory
	err := m.conn.QueryRowCtx(ctx, &chatHistory, query, characterId, RoleAssistant)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &chatHistory, nil
}
