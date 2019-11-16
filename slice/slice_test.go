package slice

import (
	"fmt"
	"runtime"
	"testing"
)

func TestSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("%p\n", &slice)
	fmt.Println(len(slice), cap(slice))
	runtime.GOMAXPROCS()

	slice = slice[0:0]
	fmt.Printf("%p\n", &slice)
	fmt.Println(len(slice), cap(slice))
}
