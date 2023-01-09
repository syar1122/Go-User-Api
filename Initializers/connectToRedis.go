package initializers

import (
	"fmt"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func ConnectToRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := Client.Ping().Result()
	if err == nil {
		fmt.Println("Redis Connecdted !!!", pong)

	} else {
		fmt.Println("Redis Error ", err)
	}
}
