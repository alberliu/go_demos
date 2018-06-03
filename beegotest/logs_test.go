package beegotest

import (
	"testing"
	"github.com/astaxie/beego/logs"
)

func TestLog(t *testing.T){
	logs.SetLogger("console")
	logs.EnableFuncCallDepth(true)

	logs.Info("hello")
}