package main

import (
	"fmt"
	"sort"
	"syscall"
)

// docker run -v $(pwd)/:/main -d alpine /main

func main() {
	syscall.Listen()
	s := []int{1, 2, 7, 4, 3}
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	fmt.Println(s)
}
