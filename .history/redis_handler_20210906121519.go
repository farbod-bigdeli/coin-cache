package main


import (
	"github.com/go-redis/redis/v8"
)

func newClient () {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // us default DB
    })
    return rd
}
func storeRedis (key redisKeys, value *[]byte) {
	rdb := 
	rdb.Set(ctx, string(key), *value, 0)

}

func getRedis (key redisKeys) *redis.StringCmd {
	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
	return rdb.Get(ctx, string(key))

}
