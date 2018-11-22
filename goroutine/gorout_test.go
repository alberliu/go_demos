package gor_test

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/astaxie/beego/logs"
)

func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], true)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
func Test_Id(t *testing.T) {
	var m sync.Map
	wg := sync.WaitGroup{}
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func(i int) {
			id := GoID()
			fmt.Println(i)
			m.Store(id, 0)
			wg.Done()
		}(i)
	}
	wg.Wait()
	count := 0
	m.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	fmt.Println("count:", count)
}

func TestGoroutine(t *testing.T) {
	go new(Goroutine).do()
	select {}
}

type Goroutine struct {
}

func (*Goroutine) do() {
	fmt.Println("hello")
}

func BenchmarkGetID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logs.Info("")
	}
}
