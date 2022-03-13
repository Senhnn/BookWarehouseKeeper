package DB

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"sync"
)

const (
	BookNumHash = "BookNum_hash"
	BookHash    = "Book_hash"
)

var rdb *redis.Client

var redisInitOnce sync.Once

func RedisClient() *redis.Client {
	return rdb
}

func RedisInit() {
	redisInitOnce.Do(func() {
		res := viper.GetStringMapString("redis")
		var redisIp string = res["ip"]
		var redisPort string = res["port"]
		rdb = redis.NewClient(&redis.Options{
			Addr:     redisIp + ":" + redisPort,
			PoolSize: 100,
		})
		fmt.Println("Redis Pool init successful")
	})
}
