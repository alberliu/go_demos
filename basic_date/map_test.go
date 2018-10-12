package basic_date

import (
	"fmt"
	"testing"
)

func TestMap1(t *testing.T) {
	a := []interface{}{1}
	b := []interface{}{1}
	a = append(a, b...)
	fmt.Println(a)
}
