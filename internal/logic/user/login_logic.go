package user

import (
	"context"
	"errors"
	"fmt"
	"go-zero-chat/ent/user"
	"go-zero-chat/utils"
	"time"

	"go-zero-chat/internal/svc"
	"go-zero-chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	res, err := l.svcCtx.DB.User.Query().Where(user.NameEQ(req.Name)).Only(l.ctx)
	if err != nil {
		logx.Errorf("[DB ERROR]: %v", err.Error())
		return nil, errors.New("用户不存在")
	}
	password := utils.Md5Salt(req.Password, res.Salt)
	if password != res.Password {
		return nil, errors.New("密码错误")
	}
	token, _ := utils.GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire, res.ID)
	//redis存入Token
	err = l.svcCtx.Redis.Setex(fmt.Sprintf("%v_accessToken_%v", l.svcCtx.Config.Name, res.ID), token, int(l.svcCtx.Config.Auth.AccessExpire))
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		Token:     token,
		ExpiresAt: time.Now().Add(time.Second * time.Duration(l.svcCtx.Config.Auth.AccessExpire)).Unix(),
	}, nil
}
