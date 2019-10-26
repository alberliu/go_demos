package sync

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

var pool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 1024)
		return b
	},
}

func TestPool(t *testing.T) {
	b := []byte{1, 2}
	fmt.Println(b)
	pool.Put(b)
	runtime.GC()
	c := pool.Get().([]byte)
	fmt.Println(c)
	d := pool.Get().([]byte)
	fmt.Println(d)
}
