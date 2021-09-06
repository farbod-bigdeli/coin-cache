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

func newClient ()  *redis.Client{
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })
    
    return rdb
}
func storeRedis (key redisKeys, value *[]byte) {
	rdb := newClient() 
	err := rdb.Set(ctx, string(key), *value, 0).Err()
    if err != nil { 
        panic(err)
    }
    defer rdb.Close()

}

func getRedis (key redisKeys) (string, error) {
    rdb := newClient()
    defer rd
	return rdb.Get(ctx, string(key)).Result()

}
