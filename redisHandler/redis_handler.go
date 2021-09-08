package redisHandler


import (
	r "github.com/go-redis/redis/v8"
    "context"
)

const NotFound = r.Nil



var ctx = context.Background()


func newClient ()  *r.Client{
    rdb := r.NewClient(&r.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })
    
    return rdb
}
func Store (key string, value *[]byte) {
	rdb := newClient() 
	err := rdb.Set(ctx, string(key), *value, 0).Err()
    defer rdb.Close()
    if err != nil { 
        panic(err)
    }


}

func Get (key string) (string, error) {
    rdb := newClient()
    defer rdb.Close()
	res, err := rdb.Get(ctx, string(key)).Result()
    return res, err

}
