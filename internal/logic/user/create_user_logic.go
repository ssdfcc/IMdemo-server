package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-chat/ent/user"
	"go-zero-chat/internal/svc"
	"go-zero-chat/internal/types"
	"go-zero-chat/utils"
	"math/rand"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserReq) error {
	count, err := l.svcCtx.DB.User.Query().Where(user.NameEQ(req.Name)).Count(l.ctx)
	if err != nil {
		logx.Errorf("[DB ERROR]: %v", err.Error())
		return errors.New("创建失败")
	}
	if count > 0 {
		return errors.New("用户名重复")
	}
	salt := fmt.Sprintf("%06d", rand.Int31())
	if _, err := l.svcCtx.DB.User.Create().SetName(req.Name).SetPassword(utils.Md5Salt(req.Password, salt)).SetSalt(salt).Save(l.ctx); err != nil {
		logx.Errorf("[DB ERROR]: %v", err.Error())
		return errors.New("创建失败")
	}
	return nil
}
