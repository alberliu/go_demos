package main

import (
	"go_demos/epoll/ge"
	"log"
	"syscall"
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
	n, err := c.Write([]byte("111"))
	if err == syscall.EBADF {
		log.Println("syscall.EBADF")
	}
	log.Println(n, err)

	c.Read()

}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	server, err := ge.NewServer(8080, &Handler{})
	if err != nil {
		log.Panicln("err")
		return
	}
	server.SetTimeout(1*time.Second, 5*time.Second)
	server.Run()
}
