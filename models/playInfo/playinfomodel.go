package playInfo

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"ruiMiBack2/models/user"
)

var _ PlayInfoModel = (*customPlayInfoModel)(nil)

type (
	// PlayInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPlayInfoModel.
	PlayInfoModel interface {
		playInfoModel
		withSession(session sqlx.Session) PlayInfoModel
		FindIndexData(ctx context.Context, challengeId int64) (*[]*FindIndexDataRes, error)
		FindRankList(ctx context.Context, challengeId, page, size int64) (*[]*FindRankListRes, int64, error)
	}

	customPlayInfoModel struct {
		*defaultPlayInfoModel
	}
)

// NewPlayInfoModel returns a model for the database table.
func NewPlayInfoModel(conn sqlx.SqlConn) PlayInfoModel {
	return &customPlayInfoModel{
		defaultPlayInfoModel: newPlayInfoModel(conn),
	}
}

func (m *customPlayInfoModel) withSession(session sqlx.Session) PlayInfoModel {
	return NewPlayInfoModel(sqlx.NewSqlConnFromSession(session))
}

type FindIndexDataRes struct {
	UserId    int64 `json:"userId"`
	BestScore int64 `json:"bestScore"`
}

// 查询首页数据信息 包括已游戏玩家数 剩余挑战次数 当前最好成绩 当前最好排名
func (m *customPlayInfoModel) FindIndexData(ctx context.Context, challengeId int64) (*[]*FindIndexDataRes, error) {
	query := fmt.Sprintf("SELECT `userId`,MIN(score) as bestScore FROM %s WHERE challengeId = ? GROUP BY userId ORDER BY bestScore", m.tableName())
	resp := make([]*FindIndexDataRes, 0)
	err := m.conn.QueryRowsPartialCtx(ctx, &resp, query, challengeId)
	return &resp, err
}

type FindRankListRes struct {
	BestScore int64  `json:"bestScore"`
	Name      string `json:"name"`
	Age       int64  `json:"age"`
	AvatorUrl string `json:"avatorUrl"`
}

func (m *customPlayInfoModel) FindRankList(ctx context.Context, challengeId, page, size int64) (*[]*FindRankListRes, int64, error) {
	// 查询总条数
	res := struct {
		Count int64 `json:"count"`
	}{}
	query := fmt.Sprintf("SELECT count(DISTINCT(userId)) as count FROM %s where challengeId = ?", m.tableName())
	err := m.conn.QueryRowPartialCtx(ctx, &res, query, challengeId)
	if err != nil {
		return nil, 0, err
	}
	if res.Count > 100 {
		res.Count = 100
	}

	userModel := user.NewUserModel(m.conn)
	query = fmt.Sprintf("SELECT MIN(t1.score) as bestScore,t2.name as name,t2.age as age,t2.avatorUrl as avatorUrl FROM %s as t1 JOIN %s as t2 on t1.userId = t2.id WHERE t1.challengeId = ? GROUP BY t1.userId ORDER BY bestScore limit ? offset ? ", m.tableName(), userModel.TableName())
	resp := make([]*FindRankListRes, 0)
	err = m.conn.QueryRowsPartialCtx(ctx, &resp, query, challengeId, size, (page-1)*size)
	return &resp, res.Count, err
}
