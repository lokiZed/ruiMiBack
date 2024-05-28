package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ruiMiBack2/internal/logic/user"
	"ruiMiBack2/internal/svc"
	"ruiMiBack2/internal/types"
)

func SendUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendUserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewSendUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.SendUserInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
