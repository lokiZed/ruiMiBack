package rank

import (
	"context"
	"encoding/json"
	"ruiMiBack2/database"
	"ruiMiBack2/internal/define"
	"ruiMiBack2/models/playInfo"

	"ruiMiBack2/internal/svc"
	"ruiMiBack2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRankListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRankListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRankListLogic {
	return &GetRankListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRankListLogic) GetRankList(req *types.GetRankListReq) (resp *types.GetRankListRes, err error) {
	challengeIdStr := l.ctx.Value(define.JwtChallengeId)
	challengeId, _ := challengeIdStr.(json.Number).Int64()

	playInfoModel := playInfo.NewPlayInfoModel(database.GetMysqlConn())
	res, count, err := playInfoModel.FindRankList(l.ctx, challengeId, req.Page, req.Size)
	if err != nil {
		return nil, err
	}
	resData := make([]types.GetRankListResData, len(*res))
	for i, data := range *res {
		resData[i] = types.GetRankListResData{
			HeadUrl:  data.AvatorUrl,
			UserName: data.Name,
			Score:    data.BestScore,
			UserAge:  data.Age,
		}
	}
	resp = &types.GetRankListRes{
		Status: define.ResponseStatusOk,
		Data:   resData,
		Count:  count,
	}
	return
}
