package reflecttion

import (
	"testing"
	"reflect"
	"fmt"
	"github.com/json-iterator/go"
)

type S1 struct {
	A int
	B int
}

func hello(s *S1){
	fmt.Println(*s)
}

func TestFunc(t *testing.T){
	//得到参数类型
	ty:=reflect.TypeOf(hello)
	typ:=ty.In(0)
	n:=reflect.New(typ.Elem())
	s:=n.Interface()

	//将json注入进去
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var s1=S{1,2}
	str,_:=json.Marshal(&s1)
	json.Unmarshal(str, s)

	//调用
	param:=make([]reflect.Value,1)
	param[0]=n

	v:=reflect.ValueOf(hello)
	returnSlice:=v.Call(param)

	return1:=returnSlice[1]
	fmt.Println(return1)


}

func TestFunc2(t *testing.T){
	ty:=reflect.TypeOf(hello)
	typ:=ty.In(0)
	//reflect.Struct
	//v:=reflect.ValueOf(hello)


	fmt.Println(typ.NumIn())
}