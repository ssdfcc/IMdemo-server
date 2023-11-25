package user

import (
	"context"
	"errors"
	"go-zero-chat/ent/user"

	"go-zero-chat/internal/svc"
	"go-zero-chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.UserInfo, err error) {
	userId := l.ctx.Value("userId").(int)
	res, err := l.svcCtx.DB.User.Query().Where(user.ID(userId)).Only(l.ctx)
	if err != nil {
		logx.Errorf("[DB ERROR]: %v", err.Error())
		return nil, errors.New("用户信息获取失败")
	}
	return &types.UserInfo{
		ID:     res.ID,
		Name:   res.Name,
		Phone:  res.Phone,
		Email:  res.Email,
		Avatar: res.Avatar,
	}, nil
}
