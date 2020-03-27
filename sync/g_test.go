package sync

import (
	"fmt"
	"testing"
	"time"
)

const lenChan = 10

func TestG(t *testing.T) {
	ch := make(chan int, lenChan)

	for i := 0; i < 5; i++ {
		go print(ch)
	}

	for i := 0; i < 10000; i++ {
		time.Sleep(time.Millisecond * 100)
		if len(ch) == lenChan {
			go print(ch)
			fmt.Println("create g")
		}
		ch <- 1
		//fmt.Println("in:", i)
	}
	select {}
}

func print(ch chan int) {
	for {
		time.Sleep(time.Second)
		<-ch
		// a := <-ch
		// fmt.Println("out:", a)
	}
}
