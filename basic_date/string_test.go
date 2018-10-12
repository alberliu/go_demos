package basic_date

import (
	"testing"
	"strings"
	"fmt"
)

func TestStringBuild(t *testing.T){
	var builder strings.Builder
	builder.WriteString("hello")
	fmt.Println(builder.String())
}