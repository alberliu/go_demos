package json_test

import (
	"testing"
	"fmt"
	"github.com/json-iterator/go"
	"encoding/json"
)

type S struct {
	A interface{} `json:"a"`
	B int `json:"b"`
}

//json-iterator测试
func TestEnJson(t *testing.T) {

	var s = S{1, 2}
	str, _ := json.Marshal(&s)
	fmt.Println(string(str))
}

func TestDeJson(t *testing.T) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	//var s = S{1, 2}
	//str, _ := json.Marshal(&s)
	str:=[]byte("{\"C\":1,\"D\":2}")
	sd := new(S)
	error:=json.Unmarshal(str, sd)
	fmt.Println(error)
	fmt.Println(sd)
}

//
func TestSlice(t *testing.T){
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	var s = []int{0,1,2,3,4,5}
	str, _ := json.Marshal(s)
	fmt.Println(string(str))

	var s1 []int
	error:=json.Unmarshal(str, s1)
	if error!=nil{
		fmt.Println(error)
	}
	fmt.Println(s1)
}
/**
反序列话数组和slice时，只需要将数组和slice的指针传进去
 */


var jsonStr =`{"a":"1","b":2}`


func TestSliceJson(t *testing.T){
	var s S
	fmt.Println(json.Unmarshal([]byte(jsonStr),&s))
	if i,ok:=s.A.(float64);ok{
		fmt.Println("float")
		fmt.Println(i)
	}
	if str,ok:=s.A.(string);ok{
		fmt.Println("str")
		fmt.Println(str)
	}
	fmt.Println(s)
}