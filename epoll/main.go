package main

import (
	"go_demos/epoll/ge"
	"log"
	"time"
)

type Handler struct {
}

func (Handler) OnConnect(c *ge.Conn) {
	go func() {
		for i := 0; i < 3; i++ {
			go func(i int) {
				if i == 0 {
					for {
						c.Write([]byte("1111111111111111111111111111111111111111111111111111111111111111111111"))
					}
				}
				if i == 1 {
					for {
						c.Write([]byte("22222222222222222222222222222222222222222222222222222222222222222222222"))
					}
				}
				if i == 2 {
					for {
						c.Write([]byte("333333"))
					}

				}

			}(i)
		}
	}()

	log.Println("connect:", c.GetFd())
}
func (Handler) OnMessage(c *ge.Conn, bytes []byte) {
	log.Println("read:", string(bytes))
}
func (Handler) OnClose(c *ge.Conn) {
	log.Println("close:", c.GetFd())
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	server, err := ge.NewServer(":8080", &Handler{})
	if err != nil {
		log.Panicln("err")
		return
	}
	server.SetTimeout(1*time.Minute, 5*time.Minute)
	server.Run()
}
