package main

import (
	"go_demos/epoll/gepoll"
	"log"
)

type Handler struct {
}

func (Handler) OnConnect(fd int) {

}
func (Handler) OnMessage(fd int, message interface{}) {
	log.Println("read:", fd, string(message.([]byte)))
}

func (Handler) OnError(fd int, err error) {
	log.Println(err)
}

func (Handler) OnClose(fd int) {

}

func main() {
	server, err := gepoll.NewServer(":8080", &Handler{})
	if err != nil {
		log.Panicln("err")
		return
	}

	server.Run()
}
