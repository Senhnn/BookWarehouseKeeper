package DB

import (
	"BWKV1/MiddleWare"
	"BWKV1/Model"
	BWKErr "BWKV1/error"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"sync"
	"time"
)

const (
	BookNumHash = "BookNumHash" // 书籍数量缓存，hash
	BookHash    = "BookHash"    // 书籍缓存，hash
	BookZset    = "BookZset"    // 书籍排行榜，zset
)

var rIns *RedisIns

var redisInitOnce sync.Once

func RedisClient() *RedisIns {
	return rIns
}

type RedisIns struct {
	rdb *redis.Client
}

func RedisInit() {
	redisInitOnce.Do(func() {
		rIns := new(RedisIns)
		res := viper.GetStringMapString("redis")
		var redisIp string = res["ip"]
		var redisPort string = res["port"]
		rIns.rdb = redis.NewClient(&redis.Options{
			Addr:     redisIp + ":" + redisPort,
			PoolSize: 100,
		})
		fmt.Println("Redis Pool init successful")
	})
}

func (r *RedisIns) SetBookCache(book *Model.Book) int32 {
	// 字符串key
	strKey := book.Isbn
	// 字符串value
	strValue, err := json.Marshal(book)
	if err != nil {
		MiddleWare.GLog.Errorf("json marshal err:%s", err.Error())
		return BWKErr.JSON_MARSHAL_ERROR
	}
	status := r.rdb.Set(context.Background(), strKey, strValue, time.Hour)
	if status.Err() != nil {
		MiddleWare.GLog.Errorf("Redis set key:%s err:%s", strKey, status.Err().Error())
		return BWKErr.REDIS_SET_BOOK_CACHE_FAIL
	}
	MiddleWare.GLog.Debugf("Redis set key:%s success", strKey)
	return BWKErr.SUCCESS
}

func (r *RedisIns) GetBookCache(isbn string) (int32, string) {
	strKey := isbn
	val, err := r.rdb.Get(context.Background(), isbn).Result()
	if err != nil {
		MiddleWare.GLog.Errorf("Redis get book:%s err", strKey)
		return BWKErr.REDIS_GET_BOOK_CACHE_FAIL, val
	}
	return BWKErr.SUCCESS, val
}

func (r *RedisIns) DelBookCache(isbn string) (int32, int64) {
	strKey := isbn
	val, err := r.rdb.Del(context.Background(), isbn).Result()
	if err != nil {
		MiddleWare.GLog.Errorf("Redis get book:%s err", strKey)
		return BWKErr.REDIS_DEL_BOOK_CACHE_FALI, val
	}
	return BWKErr.SUCCESS, val
}
