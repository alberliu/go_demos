package gor_test

import (
	"sync"
	"testing"
)

func TestQuece(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(10)
	quece := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go func() {
			for a := range quece {
				// 处理任务
			}
			group.Done()
		}()
	}

	go func() {
		for i := 0; i < 100; i++ {
			// 进任务
		}
		close(quece)
	}()
	group.Wait()
}
