package challenge

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ChallengeModel = (*customChallengeModel)(nil)

type (
	// ChallengeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChallengeModel.
	ChallengeModel interface {
		challengeModel
		withSession(session sqlx.Session) ChallengeModel
	}

	customChallengeModel struct {
		*defaultChallengeModel
	}
)

// NewChallengeModel returns a model for the database table.
func NewChallengeModel(conn sqlx.SqlConn) ChallengeModel {
	return &customChallengeModel{
		defaultChallengeModel: newChallengeModel(conn),
	}
}

func (m *customChallengeModel) withSession(session sqlx.Session) ChallengeModel {
	return NewChallengeModel(sqlx.NewSqlConnFromSession(session))
}
