package ch

import (
	"testing"
	"fmt"
)

func TestFor(t *testing.T) {
	ch := make(chan int, 9)

	go in(ch)

	for v := range ch {
		fmt.Println(v)
	}


}

func in(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)

}
