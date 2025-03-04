// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	game "ruiMiBack2/internal/handler/game"
	rank "ruiMiBack2/internal/handler/rank"
	user "ruiMiBack2/internal/handler/user"
	"ruiMiBack2/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/game/index",
					Handler: game.GetIndexDataHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/game/info",
					Handler: game.SendGameInfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/rank/list",
					Handler: rank.GetRankListHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/info",
				Handler: user.SendUserInfoHandler(serverCtx),
			},
		},
	)
}
