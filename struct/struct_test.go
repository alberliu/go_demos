package _struct

import (
	"testing"
	"fmt"
)

type A struct {
	I int
	J int
}

func (a *A)hello(){
	fmt.Println("hello world")
}

func TestA(t *testing.T){
	var a *A
	a.hello()
}
