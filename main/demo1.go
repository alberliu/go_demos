package main

import "goweb"

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func handler(id int64, name string) User {
	return User{id, name}
}

func main() {
	goweb := goweb.NewGoWeb();
	goweb.HandleGet("/test/{id}/{name}", handler)
	goweb.ListenAndServe(":8000")
}
