package _type

import (
	"testing"
	"fmt"
)

type A string

func TestA(T *testing.T){
	var a A="aaa"
	fmt.Println(a)
}
