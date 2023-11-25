package svc

import (
	"context"
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	rd "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-chat/common"
	"go-zero-chat/ent"
	"go-zero-chat/ent/migrate"
	"go-zero-chat/internal/config"
	"go-zero-chat/internal/middleware"
	"go-zero-chat/utils"
	"reflect"
)

type ServiceContext struct {
	Config         config.Config
	DB             *ent.Client
	Redis          *redis.Redis
	WSRedis        *rd.Client
	Validator      *validator.Validate
	Authority      rest.Middleware
	AuthorityQuery rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)
	// 自动迁移数据库表
	if err := db.Schema.Create(context.Background(),
		// 启用删除列
		migrate.WithDropColumn(true),
		// 启用删除索引
		migrate.WithDropIndex(true),
		// 关闭物理外键创建, 使用逻辑外键
		migrate.WithForeignKeys(false)); err != nil {
		logx.Errorw("创建表失败", logx.Field("detail", err.Error()))
	}
	// 初始化redis
	rds := redis.MustNewRedis(c.RedisConf)
	// 初始化wsredis
	wsrds := rd.NewClient(&rd.Options{
		Addr:         c.RedisConf.Host,
		PoolSize:     30,
		MinIdleConns: 30,
	})
	// 初始化定时任务
	utils.InitTimer()
	return &ServiceContext{
		Config:         c,
		DB:             db,
		Redis:          rds,
		WSRedis:        wsrds,
		Authority:      middleware.NewAuthorityMiddleware(rds, c).Handle,
		AuthorityQuery: middleware.NewAuthorityQueryMiddleware(rds, c).Handle,
	}
}

// Validate 错误翻译
func (svc ServiceContext) Validate(dataStruct interface{}) error {
	zhCh := zh.New()
	validate := validator.New()
	// 注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("json")
		return name
	})
	uni := ut.New(zhCh)
	trans, _ := uni.GetTranslator("zh")
	// 验证器注册翻译器
	zh_translations.RegisterDefaultTranslations(validate, trans)
	// 注册自定义校验函数
	common.RegisterValidation(validate, trans)
	err := validate.Struct(dataStruct)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(err.Translate(trans))
		}
	}
	return nil
}
