package main

import (
	"go_demos/epoll/gepoll"
	"log"
)

type Handler struct {
}

func (Handler) OnConnect(c *gepoll.Conn) {
	log.Println("connect:", c)
}
func (Handler) OnMessage(c *gepoll.Conn, message interface{}) {
	log.Println("read:", c, string(message.([]byte)))
}

func (Handler) OnError(c *gepoll.Conn, err error) {
	log.Println(err)
}

func (Handler) OnClose(c *gepoll.Conn) {
	log.Println("close:", c)
}

func main() {
	server, err := gepoll.NewServer(":8080", &Handler{})
	if err != nil {
		log.Panicln("err")
		return
	}

	server.Run()
}
