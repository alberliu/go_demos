package err

import (
	"testing"
	"errors"
)

func Test_Err(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string))
		}
	}()
	go func() {
		panic("panic error!")
	}()
}

func Test_Err2(t *testing.T){
	A()

}

func A()(int,error){
	err:=B()
	if err!=nil{
		return 0,err
	}
	return 1,nil
}

func B()error{
	return errors.New("error")
}
