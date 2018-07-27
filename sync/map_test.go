package sync

import (
	"testing"
	"sync"
	"fmt"
)

func TestMap(t *testing.T){
	var m sync.Map
	m.Store("hi","hi")
	fmt.Println(m.Load("hi"))

	m.Delete("hi")
	fmt.Println(m.Load("hi"))
}
