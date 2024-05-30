package middleware

import (
	"encoding/json"
	"net/http"
	"ruiMiBack2/internal/define"
	"time"
)

type JwtMiddleware struct {
}

func NewJwtMiddleware() *JwtMiddleware {
	return &JwtMiddleware{}
}

func (m *JwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expireAt := r.Context().Value(define.JwtExpireAt).(json.Number)
		expireNum, _ := expireAt.Int64()
		if expireNum < time.Now().Unix() {
			// 过期了
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		//challengeIdStr := r.Context().Value(define.JwtChallengeId).(json.Number)
		//challengeId, _ := challengeIdStr.Int64()
		//if challengeId != 1 {
		//	w.WriteHeader(http.StatusUnauthorized)
		//	return
		//}

		// 允许来自任何源的请求（仅用于测试，生产环境请限制来源）
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// 如果需要，还可以设置其他 CORS 头部，如：
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// 如果是 OPTIONS 请求，则直接返回
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}
