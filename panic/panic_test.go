package panic

import (
	"testing"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func TestPanic(t *testing.T){
	defer recoverPanic()

	err:=myPanic()
	if err != nil {
		logs.Error(err)
		return err
	}
	if err != nil {
		logs.Error(err)
		return err
	}
}

func recoverPanic() {
	if p := recover(); p != nil {
		fmt.Println(p)
	}

}

func myPanic()error{
	panic("panic")
}