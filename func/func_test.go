package func_test

import (
	"fmt"
	"testing"
	"reflect"
	"encoding/json"
)

func hello(a int) {
	fmt.Println("hello world")
}

//比较函数和函数类型的变量是否一致
func TestFunc(t *testing.T) {
	var f func(int)
	fty := reflect.TypeOf(f)
	ty := reflect.TypeOf(hello)
	fmt.Println(fty)
	fmt.Println(ty)
	fmt.Println(fty == ty)
}

type S struct {
	A int
	B int
}

func (s S)f1(){
	s.f2()
}

func (s S)f2(){

}

//json-iterator测试
func TestEnJson(t *testing.T) {
	var s = S{1, 2}
	str, _ := json.Marshal(&s)
	fmt.Println(string(str))
}
