package md5

import (
	"testing"
	"fmt"
	"giftone/ico-audit/library"
)

func TestMd5(t *testing.T) {
	md5:=library.Md5("111111")
	fmt.Println(len(md5))
	fmt.Println(md5)


}

func TestToken(t *testing.T) {
	md5:=library.Md5("18829291351"+"96e79218965eb72c92a549dd5a330112")// 96e79218965eb72c92a549dd5a330112
	fmt.Println(len(md5))
	fmt.Println(md5)
}