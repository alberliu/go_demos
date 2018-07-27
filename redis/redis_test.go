package redis

import (
	"testing"
	"github.com/go-redis/redis"
	"fmt"
)

func TestRedis(t *testing.T){
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := client.Set("key", "value", 0).Err()

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

}