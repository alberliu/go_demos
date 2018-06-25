package json_test

import (
	"testing"
	"fmt"
	"github.com/json-iterator/go"
	"encoding/json"
)

type S struct {
	A int `json:"-"`
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

	var s = make([]S,5)
	s[0]=S{1,2}
	ss:="hello"
	str, _ := json.Marshal(ss)
	fmt.Println(string(str))

	var s1 string
	error:=json.Unmarshal(str, &s1)
	if error!=nil{
		fmt.Println(error)
	}
	fmt.Println(s1)
}
/**
反序列话数组和slice时，只需要将数组和slice的指针传进去
 */