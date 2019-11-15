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

func Benchmark_Pool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := pool.Get()
		defer pool.Put(a)
	}
}

func Benchmark_Make(b *testing.B) {
	var s [][]byte
	for i := 0; i < b.N; i++ {
		s = append(s, make([]byte, 1024))
	}
}
