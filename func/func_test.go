package func_test

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func hello(a int) {
	fmt.Println("hello world")
}

// 比较函数和函数类型的变量是否一致
func TestFunc(t *testing.T) {
	var f func(int)
	fty := reflect.TypeOf(f)
	ty := reflect.TypeOf(hello)
	fmt.Println(fty)
	fmt.Println(ty)
	fmt.Println(fty == ty)
}

// 比较函数和函数类型的变量是否一致
func TestFunc2(t *testing.T) {
	f()

}

var i = 0

func f() {
	// 记录下首次开启事务的函数
	pc, _, _, _ := runtime.Caller(0)
	fmt.Println(pc)
	fmt.Println(runtime.FuncForPC(pc))
	if i == 1 {
		panic("stop")
	}
	i++
	f2()
}

func f2() { // 记录下首次开启事务的函数
	pc, _, _, _ := runtime.Caller(0)
	fmt.Println(pc)
	fmt.Println(runtime.FuncForPC(pc))

	f()
}

func deferf() int {
	return 2
}

func TestDefer(t *testing.T) {
	fmt.Println(deferf())
}
