package json_test

import (
	"fmt"
	"testing"
)

type S struct {
	A int
	B string
}

//json-iterator测试
func TestJsoniter(t *testing.T) {
	m := map[string]string{"1": "1"}
	fmt.Println(m["1=2"])
}
