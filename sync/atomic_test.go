package sync

import (
	"testing"
	"sync/atomic"
	"fmt"
	"sync"
)

func TestAtomic(t *testing.T){
	var a int64=0

	w:=sync.WaitGroup{}
	w.Add(100)

	for i:=0;i<100;i++{
		go func(){
			for j:=0;j<5;j++{
				atomic.AddInt64(&a, 1)
			}
			w.Done()
		}()
	}
	w.Wait()
	fmt.Println(a)
}