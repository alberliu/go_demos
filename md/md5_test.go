package md

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	data := []byte("a")
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制

	fmt.Println(md5str1)
}
