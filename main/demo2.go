package main

import "goweb"

func handler10(ctx goweb.Context) {

}

func handler0(ctx goweb.Context) {

}
func handler1(ctx goweb.Context) User {
	return User{}
}

func handler2(user User) User {
	return User{}
}

func handler3(ctx goweb.Context, user User) User {
	return User{}
}

func handler4(name string, id int64) User {
	return User{}
}

func handler5(ctx goweb.Context, name string, id int64) User {
	return User{}
}