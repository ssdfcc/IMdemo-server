package user

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"

	"go-zero-chat/internal/svc"
	"go-zero-chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRedisMsgListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRedisMsgListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRedisMsgListLogic {
	return &GetRedisMsgListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRedisMsgListLogic) GetRedisMsgList(req *types.GetRedisMsgListReq) (resp *types.GetRedisMsgListResp, err error) {
	var msgList []types.MsgInfo
	offset := -1 * req.Offset
	end := -1 * (req.Offset + 1 - 10)
	if req.Type == 1 {
		userId := l.ctx.Value("userId").(int)
		ids := []int{userId, req.ID}
		sort.Ints(ids)
		val := l.svcCtx.WSRedis.LRange(l.ctx, fmt.Sprintf("msg_%v_%v", ids[0], ids[1]), offset, end).Val()
		for i := range val {
			var msg types.MsgInfo
			if err := json.Unmarshal([]byte(val[i]), &msg); err != nil {
				return nil, err
			}
			msgList = append(msgList, msg)
		}
	} else {
		val := l.svcCtx.WSRedis.LRange(l.ctx, fmt.Sprintf("group_msg_%v", req.ID), offset, end).Val()
		for i := range val {
			var msg types.MsgInfo
			if err := json.Unmarshal([]byte(val[i]), &msg); err != nil {
				return nil, err
			}
			msgList = append(msgList, msg)
		}
	}

	return &types.GetRedisMsgListResp{
		List: msgList,
	}, nil
}
