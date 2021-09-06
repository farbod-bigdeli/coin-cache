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

var redisHandler struct {
    redisClient *redis.Client
}

// var redisHandler interface{
//     set()
//     get()
// }

func newRedis() redisHandler{
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })
    
    return red
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
    defer rdb.Close()
	return rdb.Get(ctx, string(key)).Result()

}
