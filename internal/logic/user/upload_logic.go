package user

import (
	"context"

	"go-zero-chat/internal/svc"
	"go-zero-chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(req *types.UploadReq) (resp *types.UploadResp, err error) {
	userId := l.ctx.Value("userId").(int)
	res, err := l.svcCtx.DB.File.Create().SetHash(req.Hash).SetName(req.Name).SetExt(req.Ext).SetSize(req.Size).SetPath(req.Path).SetUserID(userId).Save(l.ctx)
	if err != nil {
		return nil, err
	}
	return &types.UploadResp{
		ID:   res.ID,
		Name: res.Name,
		Ext:  res.Ext,
		Size: res.Size,
		Path: res.Path,
	}, nil
}
