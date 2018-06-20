package md5

import (
	"testing"
	"fmt"
	"giftone/gift-audit/library"
)

func TestMd5(t *testing.T) {
	md5:=library.Md5([]byte("123456"))
	fmt.Println(len(md5))
	fmt.Println(md5)


}

func TestToken(t *testing.T) {
	md5:=library.Md5([]byte("111111"+"gift"))
	fmt.Println(len(md5))
	fmt.Println(md5)


}