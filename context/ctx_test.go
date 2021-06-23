package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func f(ctx context.Context, str string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("down", str)
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("start run")
		}
	}
}

func TestCtx(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	ctxChild, _ := context.WithTimeout(ctx, 10*time.Second)
	go f(ctx, "ctx")
	go f(ctxChild, "child")
	//time.Sleep(5 * time.Second)
	//cancal()
	select {}
}

func TestCtxValue(t *testing.T) {
	var ctx context.Context = context.TODO()
	ctx = context.WithValue(ctx, "data_key", "data")
	fmt.Println(ctx.Value("data_key1"))
}

type S struct {
	A int
	B int
}

func TestInterFace(t *testing.T) {
	s1 := &S{}
	var i interface{} = s1
	s2 := i.(*S)
	s2.A = 1
	fmt.Println(i)

	s1.A = 2
	fmt.Println(i)
}

func TestCase(t *testing.T) {
	var userId = 290902
	key := (2 << 48) | userId
	fmt.Println(key)
}
