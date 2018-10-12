package redis

import (
	"testing"
	"github.com/go-redis/redis"
	"fmt"
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

func TestRedis(t *testing.T){
	b,err := client.SetNX("key", "value", 6*time.Second).Result()
	fmt.Println(b,err)

	time.Sleep(3*time.Second)
	b,err = client.SetNX("key", "value", 6*time.Second).Result()
	fmt.Println(b,err)

	time.Sleep(3*time.Second)
	b,err = client.SetNX("key", "value", 6*time.Second).Result()
	fmt.Println(b,err)

}

func TestPanic(t *testing.T){
	defer fmt.Println("defer")
	A()
}

func A(){
	panic("panic")
}