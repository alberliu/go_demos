package reflect

import (
	"fmt"
	"testing"
)

type A struct {
	a int
}

//
func TestInterface(t *testing.T) {
	var i interface{}
	a := &A{a: 1}
	fmt.Printf("%p\n", a)

	i = a

	a.a = 2

	a2 := i.(*A)
	fmt.Printf("%p\n", a2)
	fmt.Println(a2)
}
