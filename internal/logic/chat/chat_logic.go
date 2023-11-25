package chat

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-chat/internal/svc"
)

type ChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatLogic {
	return &ChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatLogic) Chat() error {
	// todo: add your logic here and delete this line

	return nil
}
