package user

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-chat/ent/user"
	"go-zero-chat/internal/svc"
	"go-zero-chat/internal/types"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) (resp *types.UserInfo, err error) {
	userId := l.ctx.Value("userId").(int)
	res, err := l.svcCtx.DB.User.Query().Where(user.IDEQ(userId)).Only(l.ctx)
	if err != nil {
		logx.Errorf("[DB ERROR]: %v", err.Error())
		return nil, errors.New("修改失败")
	}
	if res.Name != req.Name {
		count, err := l.svcCtx.DB.User.Query().Where(user.NameEQ(req.Name)).Count(l.ctx)
		if err != nil {
			logx.Errorf("[DB ERROR]: %v", err.Error())
			return nil, errors.New("修改失败")
		}
		if count > 0 {
			return nil, errors.New("用户名重复")
		}
	}
	only, err := l.svcCtx.DB.User.UpdateOneID(res.ID).SetName(req.Name).SetAvatar(req.Avatar).Save(l.ctx)
	if err != nil {
		logx.Errorf("[DB ERROR]: %v", err.Error())
		return nil, errors.New("修改失败")
	}
	return &types.UserInfo{
		ID:     only.ID,
		Name:   only.Name,
		Avatar: only.Avatar,
		Phone:  only.Phone,
		Email:  only.Email,
	}, nil
}
