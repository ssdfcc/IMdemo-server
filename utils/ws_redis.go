package utils

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var PublishKey = "websocker"
var OnlineTime = 4

func Publish(red *redis.Client, ctx context.Context, channel, msg string) error {
	return red.Publish(ctx, channel, msg).Err()
}

func Subscribe(red *redis.Client, ctx context.Context, channel string) (string, error) {
	sub := red.Subscribe(ctx, channel)
	msg, err := sub.ReceiveMessage(ctx)
	fmt.Println("Subscribe...", msg.Payload)
	return msg.Payload, err
}

// 设置在线用户到缓存
func SetUserOnlineInfo(red *redis.Client, key string, val []byte) {
	ctx := context.Background()
	red.Set(ctx, key, val, time.Duration(OnlineTime)*time.Hour)
}
