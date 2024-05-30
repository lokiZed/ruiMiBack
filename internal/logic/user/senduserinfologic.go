package user

import (
	"context"
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"ruiMiBack2/database"
	"ruiMiBack2/internal/config"
	"ruiMiBack2/internal/define"
	"ruiMiBack2/internal/svc"
	"ruiMiBack2/internal/types"
	"ruiMiBack2/models/challengeUser"
	"ruiMiBack2/models/user"
	"time"
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
	if req.AccountName != "147258" || req.AccountPass != "88888888" {
		return nil, errors.New("账号密码错误")
	}

	newUser := &user.User{
		Name:        req.UserName,
		AvatorUrl:   req.UserHeadUrl,
		Age:         req.UserAge,
		AccountName: req.AccountName,
		AccountPass: req.AccountPass,
	}
	err = database.GetMysqlConn().TransactCtx(l.ctx, doSendUserInfo(newUser))
	if err != nil {
		return nil, err
	}

	token, err := getJwtToken(config.GetConfig(), newUser, req.ChallengeId)
	res := &types.SendUserInfoRes{
		Status: define.ResponseStatusOk,
		Data: types.SendUserInfoResData{
			Token: token,
		},
	}
	return res, err
}

func doSendUserInfo(newUser *user.User) func(context.Context, sqlx.Session) error {
	return func(ctx context.Context, session sqlx.Session) error {
		newDb := sqlx.NewSqlConnFromSession(session)
		newUserModel := user.NewUserModel(newDb)
		result, err := newUserModel.Insert(ctx, newUser)
		if err != nil {
			return err
		}
		newUser.Id, _ = result.LastInsertId()

		challengeUserModel := challengeUser.NewChallengeUserModel(newDb)
		newChallengeUser := &challengeUser.ChallengeUser{
			UserId:      sql.NullInt64{Int64: newUser.Id, Valid: true},
			ChallengeId: sql.NullInt64{Int64: 1, Valid: true},
			LeaveNum:    sql.NullInt64{Int64: 3, Valid: true},
		}
		_, err = challengeUserModel.Insert(ctx, newChallengeUser)

		return err
	}
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
