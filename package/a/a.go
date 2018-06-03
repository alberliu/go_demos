package a

import "fmt"

type A struct {
	ID   int
	Name string
}

func Print(a A){
	fmt.Println(a)
}
