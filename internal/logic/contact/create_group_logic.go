package contact

import (
	"context"
	"errors"
	"go-zero-chat/ent"
	"go-zero-chat/internal/svc"
	"go-zero-chat/internal/types"
	"go-zero-chat/utils"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateGroupLogic) CreateGroup(req *types.CreateGroupReq) error {
	userID := l.ctx.Value("userId").(int)
	err := utils.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		// 创建群组
		db := tx.Group.Create().SetName(req.Name).SetUserID(userID)
		if len(strings.TrimSpace(req.Img)) > 0 {
			db.SetImg(req.Img)
		}
		res, err := db.Save(l.ctx)
		if err != nil {
			return err
		}
		// 添加群组关系
		_, err = tx.GroupRelation.Create().SetGroupID(res.ID).SetUserID(userID).SetType(2).Save(l.ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logx.Errorf("[DB ERROR]: %v", err.Error())
		return errors.New("创建群组失败")
	}
	return nil
}
