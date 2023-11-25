package user

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-chat/internal/svc"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout() error {
	userId := l.ctx.Value("userId")
	_, err := l.svcCtx.Redis.Del(fmt.Sprintf("%v_accessToken_%v", l.svcCtx.Config.Name, userId))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
