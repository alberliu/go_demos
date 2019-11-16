package sync

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestAtomic(t *testing.T) {
	var a int64 = 0

	w := sync.WaitGroup{}
	w.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 5; j++ {
				atomic.AddInt64(&a, 1)
			}
			w.Done()
		}()
	}
	w.Wait()
	fmt.Println(a)
}

func TestRWMutex(t *testing.T) {
	i := 0
	go func() {
		for {
			time.Sleep(time.Second)
			i++
			fmt.Println("i++", i)
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println("i=", i)
		}
	}()

	select {}
}
