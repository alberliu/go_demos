package main

import (
	"github.com/astaxie/beego/logs"
)

func main(){
	logs.SetLogger("console")
	logs.EnableFuncCallDepth(true)
	logs.Info("hello")
}
