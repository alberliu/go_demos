package reflect

import (
	"fmt"
	"strings"
	"testing"
)

type Geter interface {
	get(a int) int
}

func TestInterface(t *testing.T) {
	text := "jjjfjjsdjf \n fskfjlsajd \n"
	fmt.Println(text)

	fmt.Println(strings.Count(text, "\n"))

}
