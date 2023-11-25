package chat

import (
	"go-zero-chat/utils"
	"net/http"

	"go-zero-chat/internal/svc"
)

func ChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.Chat(w, r, svcCtx.DB, svcCtx.WSRedis)
	}
}
