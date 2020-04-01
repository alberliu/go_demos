package sync

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	c := make(chan int, 100)
	go func() {
		for {
			select {
			case <-c:
				fmt.Println("case1")
				return
			default:
				time.Sleep(time.Second)
				fmt.Println("default2")
			}
		}
	}()

	go func() {
		for {
			select {
			case <-c:
				fmt.Println("case2")
				return
			default:
				time.Sleep(time.Second)
				fmt.Println("default2")
			}
		}
	}()

	time.Sleep(5 * time.Second)
	close(c)
	select {}
}
