package basic_date

import (
	"testing"
	"fmt"
	"strings"
	"reflect"
	"log"
)

var (
	s S
)

func TestSlice(t *testing.T) {
	//var m map[int]int

	//var a []ints:=new(S)
	//fmt.Println(a==nil)

	fmt.Println(s)
}

func TestString1(t *testing.T) {
	var s1 = "11"
	var s2 = "11"
	fmt.Println(s1 == s2)
}

func TestString2(t *testing.T) {
	s := `
		select
			id,
			name,
		from
			table
		where
			id=:{}
	`
	fmt.Println(s)

}

type S struct {
	A int
	B int
}

func TestStrut(t *testing.T) {
	fmt.Println(reflect.TypeOf(S{}) == reflect.TypeOf(S{}))
}

//将url的第一个字符串和后面字符串分割
func ufirst(url string) (u1 string, u2 string) {
	index := strings.Index(url[1:], "/")
	if index == -1 {
		u1 = url[0:]
		u2 = ""
		return
	}

	u1 = url[0:index+1]
	u2 = url[index+1:]
	return
}

func TestInt(t *testing.T) {
	var a int64 =8
	//var b int
	ty:=reflect.ValueOf(a);
	fmt.Println(ty)
}

func TestData(t *testing.T) {
	a,b:=1,2
	fmt.Println(a,&b)
	c,b:=4,3
	fmt.Println(a,&b,c)
}

func TestLog(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("hello","world")
	fmt.Println("fmt")
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("hello","world")
}
