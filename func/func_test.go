package _func

import (
	"fmt"
	"testing"
	"reflect"
	"github.com/json-iterator/go"
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

//json-iterator测试
func TestEnJson(t *testing.T) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	var s = S{1, 2}
	str, _ := json.Marshal(&s)
	fmt.Println(string(str))
}
