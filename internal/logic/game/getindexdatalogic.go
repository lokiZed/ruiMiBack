package game

import (
	"context"
	"encoding/json"
	"ruiMiBack2/database"
	"ruiMiBack2/internal/define"
	"ruiMiBack2/models/challengeUser"
	"ruiMiBack2/models/playInfo"

	"ruiMiBack2/internal/svc"
	"ruiMiBack2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIndexDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetIndexDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIndexDataLogic {
	return &GetIndexDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetIndexDataLogic) GetIndexData(req *types.GetIndexDataReq) (resp *types.GetIndexDataRes, err error) {
	userIdStr := l.ctx.Value(define.JwtUserId)
	userId, _ := userIdStr.(json.Number).Int64()
	challengeIdStr := l.ctx.Value(define.JwtChallengeId)
	challengeId, _ := challengeIdStr.(json.Number).Int64()

	db := database.GetMysqlConn()
	// 查询GamerCount
	playInfoModel := playInfo.NewPlayInfoModel(db)
	dataList, err := playInfoModel.FindIndexData(l.ctx, challengeId)
	if err != nil {
		return nil, err
	}

	challengeUserModel := challengeUser.NewChallengeUserModel(db)
	challengeUser, err := challengeUserModel.FindIndexData(l.ctx, challengeId, userId)
	if err != nil {
		return nil, err
	}

	var gamerCount, bestRank, bestScore int64
	for i, data := range *dataList {
		gamerCount++
		if data.UserId == userId {
			bestRank = int64(i + 1)
			bestScore = data.BestScore
		}
	}
	resData := types.GetIndexDataResData{
		GamerCount:   gamerCount,
		ChallengeNum: challengeUser.LeaveNum.Int64,
		BestScore:    bestScore,
		BestRank:     bestRank,
	}

	resp = &types.GetIndexDataRes{
		Status: define.ResponseStatusOk,
		Data:   resData,
	}
	return resp, nil
}
