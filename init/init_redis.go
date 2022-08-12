package init

import (
	"github.com/go-redis/redis/v9"
	"github.com/shanlongpan/catgin/config"
)

//Redis 7, install go-redis/v9
//Redis 6, install go-redis/v8
func init() {
	//
	config.RedisGO = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
