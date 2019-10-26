package service

import (
	"fmt"
	"testing"
)

func TestMediaService_Get(t *testing.T) {
	InitConf()
	conf, err := MediaService.Get("huawei", "2.0.1")
	fmt.Println(conf, err)
	fmt.Printf("%+v", conf)
}

func TestHandleVersion(t *testing.T) {
	fmt.Println(HandleVersion("1"))
}
