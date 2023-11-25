package ws

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-chat/internal/svc"
)

type WSLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWSLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WSLogic {
	return &WSLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WSLogic) WS() error {
	// todo: add your logic here and delete this line

	return nil
}
