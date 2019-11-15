package main

import (
	"fmt"
	"github.com/alberliu/go_demos/main/a"
	"github.com/alberliu/go_demos/main/b"
)

func main() {
	a := a.A{A: 1, B: 1}
	var b2 b.A = b.A(a)
	fmt.Println(b2)
}
