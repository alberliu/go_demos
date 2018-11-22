package json_test

import (
	"fmt"
	"testing"

	"github.com/json-iterator/go"
)

type S struct {
	A int
	B string
}

//json-iterator测试
func TestJsoniter(t *testing.T) {

	var s = S{1, "1"}
	str, _ := jsoniter.Marshal(&s)
	fmt.Println(string(str))
}
