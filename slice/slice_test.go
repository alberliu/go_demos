package slice

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(slice[2:1])
}
