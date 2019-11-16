package sync

import (
	"fmt"
	"sync/atomic"
	"testing"
)

type TokenBucket struct {
	tokenNum int64     // 初始令牌数
	c        chan int8 // 令牌缓冲区
}

func (b *TokenBucket) Get() {
	num := atomic.LoadInt64(&b.tokenNum)
	if num > 0 {
		atomic.AddInt64(&b.tokenNum, -1)
		b.c <- 1
	}

}

func TestAtomic(t *testing.T) {
	c := make(chan int, 1)
	c <- 1
	fmt.Println(len(c))
}
