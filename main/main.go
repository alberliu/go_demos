package main

import (
	"os"
	"log"
	"runtime/pprof"
)

var sliceList [][]byte

func main() {
	f, err := os.Create("cpu.out")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i < 10000; i++ {
		f1()
		f2()
		k := 0
		for k < 10000 {
			k++
		}

	}
}

func f1() {
	k := 0
	for k < 10000 {
		k++
	}
}

func f2() {
	k := 0
	for k < 10000 {
		k++
	}
}
