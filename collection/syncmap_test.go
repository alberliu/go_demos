package collection

import (
	"testing"
	"sync"
	"fmt"
)

func TestMap(t *testing.T) {
	var m sync.Map
	m.Store("1", "1")
	m.Store("2", "2")
	m.Store("3", "3")
	m.Range(func(key, value interface{}) bool {
		a := value.(*string)
		fmt.Println(key, a)
		return false
	})
}
