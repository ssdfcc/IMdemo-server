package user

import (
	"context"
	"errors"
	"time"

	"go-zero-chat/internal/svc"
	"go-zero-chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.IDReq) error {
	if _, err := l.svcCtx.DB.User.UpdateOneID(req.ID).SetDeletedAt(time.Now()).Save(l.ctx); err != nil {
		logx.Errorf("[DB ERROR]: %v", err.Error())
		return errors.New("删除失败")
	}
	return nil
}
