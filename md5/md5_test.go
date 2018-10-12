package md5

import (
	"testing"
	"fmt"
	"giftone/ico-audit/lib"
	lib2 "step-wx/lib"
)

func TestMd5(t *testing.T) {
	fmt.Println(lib2.RandString(18))
}

func TestToken(t *testing.T) {
	md5:= lib.Md5("123456") // 96e79218965eb72c92a549dd5a330112
	fmt.Println(len(md5))
	fmt.Println(md5)
}