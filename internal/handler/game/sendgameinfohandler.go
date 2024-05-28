package game

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ruiMiBack2/internal/logic/game"
	"ruiMiBack2/internal/svc"
	"ruiMiBack2/internal/types"
)

func SendGameInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendGameInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := game.NewSendGameInfoLogic(r.Context(), svcCtx)
		resp, err := l.SendGameInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
