package basic_date

import (
	"testing"
	"fmt"
)

func TestSlice(t *testing.T){
	var slice =make([]int,0,3)
	fmt.Println(len(slice),cap(slice))
	slice=append(slice,1)
	fmt.Println(len(slice),cap(slice))
	slice=append(slice,1)
	fmt.Println(len(slice),cap(slice))
	slice=append(slice,1)
	fmt.Println(len(slice),cap(slice))
	slice=append(slice,1)
	fmt.Println(len(slice),cap(slice))
	slice=append(slice,1)
	fmt.Println(len(slice),cap(slice))
}
