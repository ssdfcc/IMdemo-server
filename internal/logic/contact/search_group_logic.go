package contact

import (
	"context"
	"go-zero-chat/ent/grouprelation"

	"go-zero-chat/internal/svc"
	"go-zero-chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchGroupLogic {
	return &SearchGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchGroupLogic) SearchGroup() (resp *types.SearchGroupResp, err error) {
	userID := l.ctx.Value("userId").(int)
	res, err := l.svcCtx.DB.GroupRelation.Query().Where(grouprelation.UserID(userID)).QueryGroup().WithGroupRelationGroup().All(l.ctx)
	if err != nil {
		logx.Errorf("[DB ERROR]: %v", err.Error())
		return nil, err
	}
	var list []types.GroupInfo
	for i, _ := range res {
		var num int
		groupRelations, err := res[i].Edges.GroupRelationGroupOrErr()
		if err == nil {
			num = len(groupRelations)
		}
		list = append(list, types.GroupInfo{
			ID:          res[i].ID,
			Name:        res[i].Name,
			Img:         res[i].Img,
			MemberCount: num,
		})
	}
	return &types.SearchGroupResp{
		Total: len(list),
		List:  list,
	}, nil
}
