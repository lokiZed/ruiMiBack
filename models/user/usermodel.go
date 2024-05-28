package user

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		withSession(session sqlx.Session) UserModel
		TableName() string
		FindByAccount(ctx context.Context, account string) (*User, error)
		FindMany(ctx context.Context) ([]*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}

func (m *customUserModel) withSession(session sqlx.Session) UserModel {
	return NewUserModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customUserModel) TableName() string {
	return m.tableName()
}

func (m *customUserModel) FindByAccount(ctx context.Context, account string) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `accountName` = ? limit 1", userRows, m.table)
	var resp User
	err := m.conn.QueryRowCtx(ctx, &resp, query, account)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUserModel) FindMany(ctx context.Context) ([]*User, error) {
	query := fmt.Sprintf("select %s from %s", userRows, m.table)
	var resp = make([]*User, 0)
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	return resp, err
}
