package main

import (
	"goweb"
	"net/http"
)

func interceptor1(http.ResponseWriter, *http.Request) bool {
	return true
}
func interceptor2(http.ResponseWriter, *http.Request) bool {
	return true
}
func interceptor3(http.ResponseWriter, *http.Request) bool {
	return true
}

func main() {
	goweb := goweb.NewGoWeb();
	goweb.AddInterceptor(interceptor1)
	goweb.AddInterceptor(interceptor2)
	goweb.AddInterceptor(interceptor3)
	goweb.ListenAndServe(":8000")
}
