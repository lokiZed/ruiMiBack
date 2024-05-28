package challengeUser

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ChallengeUserModel = (*customChallengeUserModel)(nil)

type (
	// ChallengeUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChallengeUserModel.
	ChallengeUserModel interface {
		challengeUserModel
		withSession(session sqlx.Session) ChallengeUserModel
		FindIndexData(ctx context.Context, challengeId, userId int64) (*ChallengeUser, error)
		ReduceLeaveNum(ctx context.Context, challengeId, userId int64) error
	}

	customChallengeUserModel struct {
		*defaultChallengeUserModel
	}
)

// NewChallengeUserModel returns a model for the database table.
func NewChallengeUserModel(conn sqlx.SqlConn) ChallengeUserModel {
	return &customChallengeUserModel{
		defaultChallengeUserModel: newChallengeUserModel(conn),
	}
}

func (m *customChallengeUserModel) withSession(session sqlx.Session) ChallengeUserModel {
	return NewChallengeUserModel(sqlx.NewSqlConnFromSession(session))
}

// 查询首页数据信息 包括已游戏玩家数 剩余挑战次数 当前最好成绩 当前最好排名
func (m *customChallengeUserModel) FindIndexData(ctx context.Context, challengeId, userId int64) (*ChallengeUser, error) {
	query := fmt.Sprintf("select `leaveNum` from %s where `challengeId` = ? and `userId` = ? ", m.table)
	var resp ChallengeUser
	err := m.conn.QueryRowPartialCtx(ctx, &resp, query, challengeId, userId)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// 根据challengeId和userId减少指定用户的剩余游戏次数
func (m *customChallengeUserModel) ReduceLeaveNum(ctx context.Context, challengeId, userId int64) error {
	query := fmt.Sprintf("UPDATE %s set leaveNum = leaveNum-1 where `challengeId` = ? and `userId` = ? ", m.table)
	_, err := m.conn.ExecCtx(ctx, query, challengeId, userId)
	return err
}
