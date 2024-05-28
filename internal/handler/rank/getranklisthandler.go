package rank

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ruiMiBack2/internal/logic/rank"
	"ruiMiBack2/internal/svc"
	"ruiMiBack2/internal/types"
)

func GetRankListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRankListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := rank.NewGetRankListLogic(r.Context(), svcCtx)
		resp, err := l.GetRankList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
