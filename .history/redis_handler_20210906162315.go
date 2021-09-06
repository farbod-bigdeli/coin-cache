package main


import (
	"github.com/go-redis/redis/v8"
    "context"
)

var ctx = context.Background()

type redisKeys string

type redisClient redis.Client

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
func (*redisClient) storeRedis (key redisKeys, value *[]byte) {
 	err := redisClient.Set(redisClient{},ctx,string(key),value, 0 ).Err()
    if err != nil { 
        panic(err)
    }
    defer redisClient.Close(redisClient{})

}

func ()) getRedis (key redisKeys) (string, error) {
    rdb := newClient()
    defer rdb.Close()
	return rdb.Get(ctx, string(key)).Result()

}
