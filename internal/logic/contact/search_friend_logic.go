package contact

import (
	"context"
	"errors"
	"go-zero-chat/ent/contact"
	"go-zero-chat/internal/svc"
	"go-zero-chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchFriendLogic {
	return &SearchFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchFriendLogic) SearchFriend() (resp *types.SearchFriendResp, err error) {
	userId := l.ctx.Value("userId").(int)
	all, err := l.svcCtx.DB.Contact.Query().Where(contact.OwnerID(userId), contact.Type(1)).QueryTargetUser().All(l.ctx)
	if err != nil {
		logx.Errorf("[DB ERROR]: %v", err.Error())
		return nil, errors.New("获取失败")
	}
	var list []types.FrientInfo
	for i, _ := range all {
		list = append(list, types.FrientInfo{
			ID:     all[i].ID,
			Name:   all[i].Name,
			Avatar: all[i].Avatar,
		})
	}
	return &types.SearchFriendResp{
		Total: len(list),
		List:  list,
	}, nil
}
