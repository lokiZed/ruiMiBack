package game

import (
	"context"
	"encoding/json"
	"errors"
	"ruiMiBack2/database"
	"ruiMiBack2/internal/define"
	"ruiMiBack2/models/challengeUser"
	"ruiMiBack2/models/playInfo"

	"ruiMiBack2/internal/svc"
	"ruiMiBack2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendGameInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendGameInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendGameInfoLogic {
	return &SendGameInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendGameInfoLogic) SendGameInfo(req *types.SendGameInfoReq) (resp *types.SendGameInfoRes, err error) {
	userIdStr := l.ctx.Value(define.JwtUserId)
	userId, _ := userIdStr.(json.Number).Int64()
	challengeIdStr := l.ctx.Value(define.JwtChallengeId)
	challengeId, _ := challengeIdStr.(json.Number).Int64()

	challengeUserModel := challengeUser.NewChallengeUserModel(database.GetMysqlConn())
	challengeUser, err := challengeUserModel.FindIndexData(l.ctx, challengeId, userId)
	if err != nil {
		return nil, err
	}
	if challengeUser.LeaveNum.Int64 == 0 {
		return nil, errors.New("没有挑战机会了")
	}

	newPlayInfoModel := playInfo.NewPlayInfoModel(database.GetMysqlConn())
	newPlayInfo := &playInfo.PlayInfo{
		ChallengeId: challengeId,
		UserId:      userId,
		Score:       req.Score,
	}
	_, err = newPlayInfoModel.Insert(l.ctx, newPlayInfo)
	if err != nil {
		return nil, err
	}

	err = challengeUserModel.ReduceLeaveNum(l.ctx, challengeId, userId)
	if err != nil {
		return nil, err
	}

	res := &types.SendGameInfoRes{
		Status: define.ResponseStatusOk,
	}
	return res, nil
}
