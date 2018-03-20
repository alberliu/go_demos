package _struct

import "fmt"

type S struct {
	A int
}

type B struct {
	*S
	A int
}

func (b B)f(a int){
	fmt.Println(b.S.A)
}

func do(f func(int)){

}
