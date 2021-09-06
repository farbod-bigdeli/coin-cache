package main


import (
	"github.com/go-redis/redis/v8"
    "context"
)

var ctx = context.Background()

type redisKeys string

const (
	TopCryptos redisKeys = "top_five_cryptos"
	AllCryptos redisKeys= "all_cryptos"	
)

type redisHandler struct {
    redisClient *redis.Client
    a int
}


func newRedis() *redisHandler{
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })
    
    return &redisHandler{redisClient: rdb}
}


func (*redisHandler) storeRedis (key redisKeys, value *[]byte) {
    a := (*redisHandler.redisClient 
	err := redisHandler.redisClient.Set(ctx, string(key), *value, 0).Err()
    if err != nil { 
        panic(err)
    }
    defer rdb.Close()

}

func getRedis (key redisKeys) (string, error) {
    rdb := newClient()
    defer rdb.Close()
	return rdb.Get(ctx, string(key)).Result()

}
