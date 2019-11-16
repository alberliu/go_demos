package pb

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"testing"
)

func TestPb_Descriptor(t *testing.T) {
	p := Pb{A: 25674589}
	bytes, err := proto.Marshal(&p)
	fmt.Println(len(bytes), err)
}
