package _struct

import (
	"testing"
	"fmt"
)

type S struct {
	A int
	B int
}

func TestS(t *testing.T){
	s:=&S{
		A:1,
		B:2,
	}
	str:=fmt.Sprintf("%+v\n",*s)
	fmt.Println(str)
}
