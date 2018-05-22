package prof

import (
	"testing"
	"os"
	"runtime/pprof"
	"net/http"
	"time"
	_ "net/http/pprof"
	"fmt"
)

func TestPprof(t *testing.T) {

	// CPU
	cpu, _ := os.Create("cpu.out")
	defer cpu.Close()
	pprof.StartCPUProfile(cpu)
	defer pprof.StopCPUProfile()
	// Memory
	mem, _ := os.Create("mem.out")

	defer mem.Close()
	defer pprof.WriteHeapProfile(mem)

}

func TestPprofWeb(t *testing.T) {
	go http.ListenAndServe("localhost:6060", nil)
	var arr [1000]int
	fmt.Println(arr)
	for i := 0; i < 100; i++ {
		go f()
	}
	f()
}

func f() {
	for {
		time.Sleep(time.Second)
	}
}
