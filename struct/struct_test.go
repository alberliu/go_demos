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

type B struct {
	A
}



func TestS(t *testing.T){
	var a *A
	a=&B{}
}
