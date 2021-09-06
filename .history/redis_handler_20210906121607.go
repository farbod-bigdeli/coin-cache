package main


import (
	"github.com/go-redis/redis/v8"
)

func newClient ()  *redis.Client{
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // us default DB
    })
    return rdb
}
func storeRedis (key redisKeys, value *[]byte) {
	rdb := newClient() 
	rdb.Set(ctx, string(key), *value, 0)

}

func getRedis (key redisKeys) *redis.StringCmd {
    newClient()
	return rdb.Get(ctx, string(key))

}