package main


import (
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()


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
	rdb.Set(ctx, string(key), *value, 0)

}

func getRedis (key redisKeys) *redis.StringCmd {
    rdb := newClient()
	return rdb.Get(ctx, string(key))

}
