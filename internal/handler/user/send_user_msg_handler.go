package user

import (
	"go-zero-chat/utils"
	"net/http"

	"go-zero-chat/internal/svc"
)

func SendUserMsgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.Chat(w, r, svcCtx.DB, svcCtx.WSRedis)
	}
}
