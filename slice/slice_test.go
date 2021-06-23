package slice

import (
	"fmt"
	"strings"
	"testing"
)

func TestSlice(t *testing.T) {
	data := fmt.Sprintf("%s%5s", "1", "15")
	data = strings.Replace(data, " ", "0", -1)
	fmt.Println(data)
}
