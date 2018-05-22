package ctx

import (
	"testing"
	"context"
	"fmt"
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

//对方法进行超时控制
func CtxUse() {
	timeout := 5 * time.Second
	ctx, caccel := context.WithTimeout(context.Background(), timeout)
	defer caccel()

	done := make(chan int, 1)
	go func() {
		do(ctx)
		done <- 1
	}()

	select {
	case <-done:
		fmt.Println("ok")
	case <-ctx.Done():
		fmt.Println("timeout")
		return
	}
}

func do(ctx context.Context) {
	for i := 1; i <= 7; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i, "work")
	}
}

func TestCtx2(t *testing.T) {
	go CtxUse()
	select {}
}