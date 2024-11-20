package dao

import (
	"github.com/redis/go-redis/v9"
	"server-zys/internal/core"
	"sync"
	"time"
)

var (
	rdb     *redis.Client
	rdbOnce sync.Once
)

func GetRdb() *redis.Client {
	if rdb != nil {
		return rdb
	}
	rdbOnce.Do(func() {
		initRedis()
	})
	return rdb
}

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:         core.GlobalConfig.Redis.Addr,
		DialTimeout:  time.Millisecond * time.Duration(core.GlobalConfig.Redis.DialTimeOut),
		ReadTimeout:  time.Millisecond * time.Duration(core.GlobalConfig.Redis.ReadTimeOut),
		WriteTimeout: time.Millisecond * time.Duration(core.GlobalConfig.Redis.WriteTimeOut),
	})
}
