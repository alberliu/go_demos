package main

import "github.com/astaxie/beego/logs"

func main() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"log.log","level":6}`)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)

	logs.Trace("hello")
	logs.Debug("hello")
	logs.Info("hello")
	logs.Warn("hello")
	logs.Error("hello")
}
