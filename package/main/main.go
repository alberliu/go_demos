package main

import (
	"github.com/alberliu/go_demos/packa/a"
	"github.com/alberliu/go_demos/packa/b"
)

func main() {
	A:=a.A{}
	a.Print(A)
	b.CallA()
}
