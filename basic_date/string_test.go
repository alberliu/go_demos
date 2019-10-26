package basic_date

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestStringBuild(t *testing.T) {
	var builder strings.Builder
	builder.WriteString("hello")
	builder.WriteString("123456789123456789123456789")
	fmt.Println(builder.String())
}

func TestSplit(t *testing.T) {
	fmt.Println(strings.Split("1|2", "|"))
}

func TestInt(t *testing.T) {
	var slice *[]int
	var i interface{} = slice
	t1 := reflect.TypeOf(i)
	fmt.Println(t1)
	v1 := reflect.ValueOf(i)
	fmt.Println(v1)
}

func TestBytes(t *testing.T) {
	var str = "1号asdjfljaldjsf"
	fmt.Println(len([]rune(str)))
}
