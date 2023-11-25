package user

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"go-zero-chat/internal/svc"
	"go-zero-chat/utils"
	"net/http"
	"time"
)

var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ws, err := upGrade.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer func(ws *websocket.Conn) {
			if err := ws.Close(); err != nil {
				fmt.Println(err.Error())
			}
		}(ws)
		MsgHandler(svcCtx.WSRedis, ws, r.Context())
	}
}

func MsgHandler(red *redis.Client, ws *websocket.Conn, ctx context.Context) {
	msg, err := utils.Subscribe(red, ctx, utils.PublishKey)
	if err != nil {
		fmt.Println(err.Error())
	}
	tm := time.Now().Format("2006-01-02 15:04:05")
	m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
	err = ws.WriteMessage(1, []byte(m))
	if err != nil {
		fmt.Println(err.Error())
	}
}
