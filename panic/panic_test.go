package panic

import (
	"testing"
	"fmt"
	"errors"
)

func TestPanic(t *testing.T){
	defer func() {
		if p := recover(); p != nil {
			err := fmt.Errorf("internal error: %v", p)
			fmt.Println(err)
		}
	}()

	errors.New("error")
	//panic("panic")
	myPanic()
}

func myPanic(){
	panic("panic")
}