package sync

import (
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	var m sync.Map
	for i := 0; i < 1000000; i++ {
		m.Store("hi", "hi")
	}

	m.Range(func(key, value interface{}) bool {
		return true
	})
}

func BenchmarkMap(b *testing.B) {
	var m sync.Map
	for i := 0; i < 2000000; i++ {
		m.Store(i, i)
	}

	for i := 0; i < b.N; i++ {
		m.Range(func(key, value interface{}) bool {
			return true
		})
	}
}
