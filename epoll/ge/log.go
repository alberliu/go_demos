package ge

import "log"

var Log logger = &defaultLog{}

type logger interface {
	Error(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
}

type defaultLog struct{}

func (*defaultLog) Error(args ...interface{}) {
	log.Println(args...)
}
func (*defaultLog) Info(args ...interface{}) {
	log.Println(args...)
}
func (*defaultLog) Debug(args ...interface{}) {
	log.Println(args...)
}
