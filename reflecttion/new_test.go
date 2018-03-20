package reflecttion

import (
	"testing"
	"reflect"
	"fmt"
)

type S struct {
	A int
	B int
}

//根据结构体类型构造结构体
func TestNew(t *testing.T) {
	s:=S{}
	s1:=reflect.TypeOf(s)
	new:=reflect.New(s1)
	fmt.Println(new)
}

//根据结构体指针类型构造实体
func TestNewByP(t *testing.T) {
	s:=new(S)
	s1:=reflect.TypeOf(s)
	new:=reflect.New(s1.Elem())
	fmt.Println(new)
}


func TestSlice(t *testing.T){
	ss:=make([]S,4)
	ty:=reflect.TypeOf(ss).Elem()
	fmt.Println(ty)
}