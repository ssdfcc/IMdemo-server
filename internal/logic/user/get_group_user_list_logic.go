package user

import (
	"context"
	"errors"
	"go-zero-chat/ent/grouprelation"

	"go-zero-chat/internal/svc"
	"go-zero-chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGroupUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupUserListLogic {
	return &GetGroupUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupUserListLogic) GetGroupUserList(req *types.IDReq) (resp *types.GetGroupUserListResp, err error) {
	res, err := l.svcCtx.DB.GroupRelation.Query().Where(grouprelation.IDEQ(req.ID)).QueryUser().All(l.ctx)
	if err != nil {
		logx.Errorf("[DB ERROR]: %v", err.Error())
		return nil, errors.New("获取失败")
	}
	var list []types.UserInfo
	for i := range res {
		list = append(list, types.UserInfo{
			ID:     res[i].ID,
			Name:   res[i].Name,
			Phone:  res[i].Phone,
			Email:  res[i].Email,
			Avatar: res[i].Avatar,
		})
	}
	return &types.GetGroupUserListResp{
		List: list,
	}, nil
}
