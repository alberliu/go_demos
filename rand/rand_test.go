package rand

import (
	"testing"
	"math/rand"
	"fmt"
)

func TestRand(t *testing.T){
	for i:=0;i<1000;i++ {
		fmt.Println(10000 - rand.Intn(9000))
	}
}
