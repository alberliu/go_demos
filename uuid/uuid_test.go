package uuid

import (
	"testing"
	"github.com/satori/go.uuid"
	"fmt"
)

func TestUUID(t *testing.T){
	for i:=0;i<10;i++{
		u2, _ := uuid.NewV4()
		fmt.Println(u2)
	}
}
