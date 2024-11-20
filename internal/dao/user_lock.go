package dao

import (
	"context"
	"fmt"
	"server-zys/internal/constant"
	"server-zys/logs"
	"time"
)

func Lock(ctx context.Context, id string, expire time.Duration) bool {
	key := fmt.Sprintf(constant.RedisUserLockKey, id)
	result, err := GetRdb().SetNX(ctx, key, 1, expire).Result()
	if err != nil {
		logs.Error(ctx, fmt.Sprintf("[DAO][Lock] Redis Err, error: %s", err.Error()))
	}
	return result
}

func Unlock(ctx context.Context, id string) {
	key := fmt.Sprintf(constant.RedisUserLockKey, id)
	_ = GetRdb().Del(ctx, key)
}
