package de

import (
	"fmt"
	"testing"
)

type Hello interface {
	SayHello()
}

func TestDefer(t *testing.T) {
	if true {
		defer fmt.Println("hello")
	}

	fmt.Print("world")
}
