package main

import (
	"net/http"
	"goweb"
)

func handler400(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(400)
	w.Write([]byte("bad request"))
}
func handler404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Write([]byte("url not found"))
}
func handler405(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(405)
	w.Write([]byte("method not found"))
}

func main() {
	goWeb := goweb.NewGoWeb()
	goWeb.Handler400 = handler400
	goWeb.Handler404 = handler404
	goWeb.Handler405 = handler405

	goweb.ListenAndServe(":8000")
}
