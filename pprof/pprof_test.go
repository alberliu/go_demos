package pprof

import (
	"testing"
	"os"
	"runtime/pprof"
	"net/http"
	"time"
	_ "net/http/pprof"
	"fmt"
)

// CPU
func TestPprofCPU(t *testing.T) {
	mem, _ := os.Create("mem.out")
	defer mem.Close()
	defer pprof.WriteHeapProfile(mem)
}

// 内存
func TestPprofMem(t *testing.T) {
	mem, _ := os.Create("mem.out")
	defer mem.Close()
	defer pprof.WriteHeapProfile(mem)
}

func TestPprofWeb(t *testing.T) {
	go http.ListenAndServe("localhost:6060", nil)

	for i := 0; i < 100; i++ {
		go f()
	}
	select {}

}

func f() {
	for {
		time.Sleep(time.Second)
		a := make([]int, 100)
		fmt.Println(a)
	}
}
