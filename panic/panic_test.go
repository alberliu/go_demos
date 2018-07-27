package panic

import (
	"testing"
	"fmt"
)

func TestPanic(t *testing.T){
	defer recoverPanic()

	myPanic()
}

func recoverPanic() {
	if p := recover(); p != nil {
		fmt.Println(p)
	}
}

func myPanic(){
	panic("panic")
}