package panic

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	myPanic()
}

func recoverPanic() {
	if p := recover(); p != nil {
		fmt.Println(p)
	}

}

func myPanic() error {
	panic("panic")
}
