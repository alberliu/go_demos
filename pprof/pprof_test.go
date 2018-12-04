package pprof

import (
	"log"
	"math"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"testing"
)

// CPU
func TestPprofCPU(t *testing.T) {
	file, err := os.Create("cpu.out")
	if err != nil {
		log.Fatal(err)
	}
	err = pprof.StartCPUProfile(file)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	defer pprof.StopCPUProfile()
	useCPU()
}

func useCPU() {
	for i := 0; i < 100000000; i++ {
		math.Pow(1024*1024*1024, 1024*1024*1024)
	}
}

// 内存
func TestPprofMem(t *testing.T) {
	file, err := os.Create("mem.out")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	defer pprof.WriteHeapProfile(file)
	useMem()
}

func useMem() {
	var s [][]byte
	for i := 0; i < 1000; i++ {
		a := make([]byte, 1024*1024)
		s = append(s, a)
	}
}

func TestPprofWeb(t *testing.T) {
	http.ListenAndServe("localhost:6060", nil)
}
