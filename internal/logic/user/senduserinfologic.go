package user

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"ruiMiBack2/database"
	"ruiMiBack2/internal/config"
	"ruiMiBack2/internal/define"
	"ruiMiBack2/internal/svc"
	"ruiMiBack2/internal/types"
	"ruiMiBack2/models/user"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendUserInfoLogic {
	return &SendUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendUserInfoLogic) SendUserInfo(req *types.SendUserInfoReq) (resp *types.SendUserInfoRes, err error) {
	db := database.GetMysqlConn()
	newUserModel := user.NewUserModel(db)

	oldUser, err := newUserModel.FindByAccount(l.ctx, req.AccountName)
	if err != nil {
		return nil, err
	}
	if oldUser.AccountPass != req.AccountPass {
		return nil, errors.New("账号密码错误")
	}
	oldUser.Age = req.UserAge
	oldUser.Name = req.UserName
	oldUser.AvatorUrl = req.UserHeadUrl
	err = newUserModel.Update(l.ctx, oldUser)
	if err != nil {
		return nil, err
	}
	token, err := getJwtToken(config.GetConfig(), oldUser, req.ChallengeId)
	res := &types.SendUserInfoRes{
		Status: define.ResponseStatusOk,
		Data: types.SendUserInfoResData{
			Token: token,
		},
	}
	return res, err
}

func getJwtToken(c *config.Config, userInfo *user.User, challengeId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims[define.JwtUserId] = userInfo.Id
	claims[define.JwtChallengeId] = challengeId
	claims[define.JwtExpireAt] = time.Now().Add(time.Duration(c.JwtAuth.AccessExpire) * time.Second).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(c.JwtAuth.AccessSecret))
}
