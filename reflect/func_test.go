package reflect

import (
	"testing"
	"reflect"
	"fmt"
	"github.com/json-iterator/go"
	"time"
	"giftone/ico-audit/library"
)

type S1 struct {
	A int
	B int
}

func hello(s *S1) {
	fmt.Println(*s)
}

func TestFunc(t *testing.T) {
	//得到参数类型
	ty := reflect.TypeOf(hello)
	typ := ty.In(0)
	n := reflect.New(typ.Elem())
	s := n.Interface()

	//将json注入进去
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var s1 = S{1, 2}
	str, _ := json.Marshal(&s1)
	json.Unmarshal(str, s)

	//调用
	param := make([]reflect.Value, 1)
	param[0] = n

	v := reflect.ValueOf(hello)
	returnSlice := v.Call(param)

	return1 := returnSlice[1]
	fmt.Println(return1)

}

func TestFunc2(t *testing.T) {
	ty := reflect.TypeOf(hello)
	typ := ty.In(0)
	//reflect.Struct
	//v:=reflect.ValueOf(hello)

	fmt.Println(typ.NumIn())
}

type User struct {
	Id   int
	Name string
	Now  time.Time
}

func TestReflect(a *testing.T) {
	user := User{Id: 1, Name: "1", Now: time.Now()}
	fmt.Println(GetStrutsField(user))
}

func GetStrutsField(s interface{}) string {
	fieldStr := ""
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	for i := 0; i < t.NumField(); i++ {
		name:=t.Field(i).Name
		if  t.Field(i).Type.Name() == "Time" {
			fieldStr += name + ": " + lib.FormatTime(v.Field(i).Interface().(time.Time)) + ""
			break
		}
		fieldStr += name + ": " + fmt.Sprint(v.Field(i).Interface()) + ""
	}
	return fieldStr
}

func TestReflectCompare(a *testing.T) {
	time1:=time.Now()
	time2:=time.Now()
	user1 := User{Id: 1, Name: "1", Now: time1}
	user2 := User{Id: 6, Name: "1", Now: time2}
	fmt.Println(GetStrutsFieldChange(user1, user2))
}

// GetStrutsChange 获取结构体s1到结构体s2属性的变化
func GetStrutsFieldChange(s1, s2 interface{}) string {
	change := ""
	t1 := reflect.TypeOf(s1)
	v1 := reflect.ValueOf(s1)

	v2 := reflect.ValueOf(s2)

	for i := 0; i < t1.NumField(); i++ {
		name := t1.Field(i).Name
		if name == "Createtime" || name == "Updatetime" {
			continue
		}



		i1 := v1.Field(i).Interface()
		i2 := v2.FieldByName(name).Interface()
		if i1 != i2 {
			if t1.Field(i).Type.Name() == "Time" || lib.FormatTime(i1.(time.Time))!= lib.FormatTime(i1.(time.Time)){
				change += name + ": " + lib.FormatTime(i1.(time.Time)) + " -> " + lib.FormatTime(i1.(time.Time)) + ""
				continue
			}
			change += name + ": " + fmt.Sprint(i1) + " -> " + fmt.Sprint(i2) + ""
		}
	}
	return change
}

