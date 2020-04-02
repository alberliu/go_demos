package main

import (
	"go_demos/epoll/ge"
	"log"
	"time"
)

type Handler struct {
}

func (Handler) OnConnect(c *ge.Conn) {
	log.Println("connect:", c.GetFd())
}
func (Handler) OnMessage(c *ge.Conn, bytes []byte) {
	log.Println("read:", string(bytes))
}
func (Handler) OnClose(c *ge.Conn) {
	log.Println("close:", c.GetFd())
}

func main() {
	server, err := ge.NewServer(":8080", &Handler{})
	if err != nil {
		log.Panicln("err")
		return
	}
	server.SetTimeout(1*time.Second, 5*time.Second)
	server.Run()
}
