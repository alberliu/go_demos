package collection

import (
	"fmt"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	var m sync.Map
	m.Store("1", "1")
	m.Store("2", "2")
	m.Store("3", "3")

	m.Load("f")

	m.Range(func(key, value interface{}) bool {
		a := value.(string)
		fmt.Println(key, a)
		return true
	})
}
