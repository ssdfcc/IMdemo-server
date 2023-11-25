package middleware

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/x/errors"
	xhttp "github.com/zeromicro/x/http"
	"go-zero-chat/internal/config"
	"go-zero-chat/utils"
	"net/http"
	"strings"
	"time"
)

type AuthorityMiddleware struct {
	Redis  *redis.Redis
	Config config.Config
}

func NewAuthorityMiddleware(rds *redis.Redis, c config.Config) *AuthorityMiddleware {
	return &AuthorityMiddleware{
		Redis:  rds,
		Config: c,
	}
}

func (m *AuthorityMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			logx.Errorf("[AUTH ERROR]: %v", "token不能为空")
			w.WriteHeader(http.StatusForbidden)
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New(403, "无权访问"))
			return
		}
		parts := strings.SplitN(token, " ", 2)
		if len(parts) == 2 && parts[0] == "Bearer" {
			token = parts[1]
		}
		// token校验
		claims, err := utils.ParseToken(token, m.Config.Auth.AccessSecret)
		if err != nil {
			logx.Errorf("[AUTH ERROR]: %v", err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New(401, "token已失效"))
			return
		}
		// userId写入上下文
		ctx := context.WithValue(r.Context(), "userId", claims.UserId)
		r = r.WithContext(ctx)
		// 判断是否启用多点登录拦截
		if m.Config.MultiSignOnBlocking.Type {
			// 判断token是否在resdis中
			get, err := m.Redis.Get(fmt.Sprintf("%v_accessToken_%v", m.Config.RestConf.Name, claims.UserId))
			if err != nil {
				logx.Errorf("[AUTH ERROR]: %v", err.Error())
				w.WriteHeader(http.StatusUnauthorized)
				xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New(401, "token已失效"))
				return
			}
			// 判断token是否在白名单中
			if token != get && token != "Bearer "+get {
				logx.Errorf("[AUTH ERROR]: %v", err.Error())
				w.WriteHeader(http.StatusUnauthorized)
				xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New(401, "token已失效"))
				return
			}
		}
		// 判断token时间是否需要重新生成新token
		if claims.ExpAt-time.Now().Unix() < m.Config.Auth.RefreshTime {
			jwtToken, _ := utils.GetJwtToken(m.Config.Auth.AccessSecret, m.Config.Auth.AccessExpire, claims.UserId)
			w.Header().Set("new-token", jwtToken)
			w.Header().Set("new-token-expires", fmt.Sprintf("%v", time.Now().Add(time.Second*time.Duration(m.Config.Auth.AccessExpire)).Unix()))
		}
		next(w, r)
	}
}
