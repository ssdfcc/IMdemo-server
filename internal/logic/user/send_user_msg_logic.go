package user

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-chat/internal/svc"
)

type SendUserMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendUserMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendUserMsgLogic {
	return &SendUserMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendUserMsgLogic) SendUserMsg() error {
	// todo: add your logic here and delete this line

	return nil
}
