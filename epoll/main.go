package main

import (
	"go_demos/epoll/gepoll"
	"log"
	"net"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Println(err)
		return
	}
	file, err := listener.File()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("listener_fd:", file.Fd())

	e, err := gepoll.EpollCreate()
	if err != nil {
		log.Println(err)
		return
	}

	e.AddListener(int(file.Fd()))
	if err != nil {
		log.Println(err)
		return
	}

	for {
		e.EpollWait()
	}
}
