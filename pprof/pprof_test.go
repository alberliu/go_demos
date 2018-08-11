package pprof

import (
	"testing"
	"os"
	"runtime/pprof"
	"net/http"
	"time"
	_ "net/http/pprof"
	"fmt"
	"log"
)

// CPU
func TestPprofCPU(t *testing.T) {
	f, err := os.Create("cpu.out")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
}

// 内存 算命先生 慎用
func TestPprofMem(t *testing.T) {
	file, err := os.Create("mem.out")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	defer pprof.WriteHeapProfile(file)
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
