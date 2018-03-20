package err

import "testing"

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
