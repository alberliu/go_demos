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

func Intersection(a, b []int64) []int64 {
	c := make([]int64, 0)
	for _, i := range a {
		for _, j := range b {
			if i == j {
				c = append(c, i)
			}
		}
	}
	return c
}

func TestMap2(t *testing.T) {
	var a = []int64{1, 2, 3}
	var b = []int64{3, 4, 5}
	fmt.Println(Intersection(a, b))

}
