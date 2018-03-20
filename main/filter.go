package main

import (
	"net/http"
	"goweb"
)

func filter(w http.ResponseWriter, r *http.Request, f func(http.ResponseWriter, *http.Request)) {
	f(w, r)
}

func main() {
	goweb := goweb.NewGoWeb();
	goweb.Filter = filter
	goweb.ListenAndServe(":8000")
}
