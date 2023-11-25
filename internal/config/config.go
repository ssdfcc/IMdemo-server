package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-chat/common"
)

type Config struct {
	rest.RestConf
	RedisConf    redis.RedisConf
	DatabaseConf common.DatabaseConf
	Auth         struct {
		AccessSecret string
		AccessExpire int64
		RefreshTime  int64
	}
	MultiSignOnBlocking struct {
		Type bool
	}
	Local struct {
		Path string
	}
}
