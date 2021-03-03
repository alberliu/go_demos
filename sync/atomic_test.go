package sync

import (
	"fmt"
	"math"
	"os"
	"sync"
	"sync/atomic"
	"testing"
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

// 1,2
// 4,5
func TestRWMutex(t *testing.T) {
	r := (1*4 + 2*5) / (math.Sqrt(1*1+2*2) * math.Sqrt(4*4+5*5))
	fmt.Println(r)
	fmt.Println(math.Sqrt(4))
}

func TestEnv(t *testing.T) {
	fmt.Println(os.Getenv("PATH"))
}
