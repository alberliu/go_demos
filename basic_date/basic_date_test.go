package basic_date

import (
	"fmt"
	"strconv"
	"testing"
)

func TestInt64Max(t *testing.T) {
	var a uint64
	fmt.Println((a - 1) / 1000 / 365)
}

func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		strconv.FormatInt(1, 10)
	}
}

func Benchmark_Fmt(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		fmt.Sprint(1)
	}
}

func TestLock(t *testing.T) {
	var a = struct {
		A int
	}{}

	go func() {
		for {
			fmt.Println("a:", a.A)
		}
	}()

	go func() {
		for {
			a.A = a.A + 1
			//fmt.Println(a)
		}
	}()

	select {}
}

func TestAd(t *testing.T) {
	var a = 10000000000
	a = a >> 1
	fmt.Println(a)
}
