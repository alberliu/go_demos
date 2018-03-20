package main

import "goweb"

func main() {
	group1:=goweb.NewGroup("/group1")
	group1.HandleGet("/handler1",handler)
	group1.HandleGet("/handler2",handler)
	group1.HandleGet("/handler3",handler)

	group2:=goweb.NewGroup("/group2")
	group2.HandleGet("/handler1",handler)
	group2.HandleGet("/handler2",handler)
	group2.HandleGet("/handler3",handler)

	group3:=goweb.NewGroup("/group3")
	group3.HandleGet("/handler1",handler)
	group3.HandleGet("/handler2",handler)
	group3.HandleGet("/handler3",handler)

	goweb.HandleGroup(group1)
	goweb.HandleGroup(group2)
	goweb.HandleGroup(group3)
	goweb.ListenAndServe(":8000")
}