package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
	"time"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func TestRedis(t *testing.T) {
	b, err := client.SetNX("key", "value", 6*time.Second).Result()
	fmt.Println(b, err)

	time.Sleep(3 * time.Second)
	b, err = client.SetNX("key", "value", 6*time.Second).Result()
	fmt.Println(b, err)

	time.Sleep(3 * time.Second)
	b, err = client.SetNX("key", "value", 6*time.Second).Result()
	fmt.Println(b, err)

}

/**
POLLIN   = 0x1
	POLLPRI  = 0x2
	POLLOUT  = 0x4
	POLLERR  = 0x8
	POLLHUP  = 0x10
	POLLNVAL = 0x20
*/
func TestPanic(t *testing.T) {
	fmt.Printf("%b\n", 0x1|0x2)
	fmt.Printf("%b\n", 0x2)
	fmt.Printf("%b\n", 0x8)
	fmt.Printf("%b\n", 0x10)
	fmt.Printf("%b\n", 0x20)
}
