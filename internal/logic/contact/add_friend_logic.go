package contact

import (
	"context"
	"errors"
	"go-zero-chat/ent"
	"go-zero-chat/ent/contact"
	"go-zero-chat/utils"

	"go-zero-chat/internal/svc"
	"go-zero-chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendLogic {
	return &AddFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFriendLogic) AddFriend(req *types.IDReq) error {
	userID := l.ctx.Value("userId").(int)
	if userID == req.ID {
		return errors.New("不能添加自己为好友")
	}
	count, err := l.svcCtx.DB.Contact.Query().Where(contact.TargetIDEQ(req.ID), contact.OwnerIDEQ(userID), contact.Type(1)).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("已经是好友了")
	}
	err = utils.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		_, err = tx.Contact.Create().SetOwnerID(userID).SetTargetID(req.ID).SetType(1).Save(l.ctx)

		if err != nil {
			return err
		}
		_, err = tx.Contact.Create().SetOwnerID(req.ID).SetTargetID(userID).SetType(1).Save(l.ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logx.Errorf("[DB ERROR]: %v", err.Error())
		return errors.New("添加好友失败")
	}
	return nil
}
