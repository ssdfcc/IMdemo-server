package user

import (
	"context"
	"errors"
	"go-zero-chat/ent/user"

	"go-zero-chat/internal/svc"
	"go-zero-chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLIstLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLIstLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLIstLogic {
	return &GetUserLIstLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLIstLogic) GetUserLIst(req *types.GetUserListReq) (resp *types.GetUserListResp, err error) {
	db := l.svcCtx.DB.User.Query()
	if len(req.Name) > 0 {
		db.Where(user.NameContainsFold(req.Name), user.DeletedAtIsNil())
	}
	all, err := db.All(l.ctx)
	if err != nil {
		logx.Errorf("[DB ERROR]: %v", err.Error())
		return nil, errors.New("获取失败")
	}
	var list []types.UserInfo
	for i, _ := range all {
		list = append(list, types.UserInfo{
			ID:     all[i].ID,
			Name:   all[i].Name,
			Phone:  all[i].Phone,
			Email:  all[i].Email,
			Avatar: all[i].Avatar,
		})
	}
	return &types.GetUserListResp{
		List: list,
	}, nil
}
