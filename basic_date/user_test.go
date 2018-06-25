package basic_date

import (
	"testing"
	"fmt"
)

type User struct {
	A int
}

func TestInt2(t *testing.T){
	var a int =1022000000
	b:=float64(a)
	fmt.Println(b/100000000)
}