package contact

import (
	"context"
	"errors"
	"go-zero-chat/ent/group"
	"go-zero-chat/ent/grouprelation"
	"go-zero-chat/internal/svc"
	"go-zero-chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGroupLogic {
	return &AddGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddGroupLogic) AddGroup(req *types.IDReq) error {
	userID := l.ctx.Value("userId").(int)
	if _, err := l.svcCtx.DB.Group.Query().Where(group.ID(req.ID)).Only(l.ctx); err != nil {
		return errors.New("群组不存在")
	}
	count, err := l.svcCtx.DB.GroupRelation.Query().Where(grouprelation.GroupID(req.ID), grouprelation.UserID(userID), grouprelation.Type(1)).Count(l.ctx)
	if err != nil {
		logx.Errorf("[DB ERROR]: %v", err.Error())
		return err
	}
	if count > 0 {
		return errors.New("已经加入该群了")
	}
	_, err = l.svcCtx.DB.GroupRelation.Create().SetUserID(userID).SetGroupID(req.ID).Save(l.ctx)
	if err != nil {
		logx.Errorf("[DB ERROR]: %v", err.Error())
		return errors.New("添加好友失败")
	}
	return nil
}
