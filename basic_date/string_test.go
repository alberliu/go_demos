package basic_date

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringBuild(t *testing.T) {
	var builder strings.Builder
	builder.WriteString("hello")
	builder.WriteString("123456789123456789123456789")
	fmt.Println(builder.String())
}

func TestSplit(t *testing.T) {
	fmt.Println(strings.Split("1|2", "|"))
}
