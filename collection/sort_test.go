package collection

import (
	"testing"
	"sort"
	"fmt"
)

func TestSort(t *testing.T){
	x := 11
	s := []int{3, 6, 8, 11, 45} //注意已经升序排序
	pos := sort.Search(len(s), func(i int) bool { return s[i] >= x })
	if pos < len(s) && s[pos] == x {
		fmt.Println(x, "在s中的位置为：", pos)
	} else {
		fmt.Println("s不包含元素", x)
	}


}


