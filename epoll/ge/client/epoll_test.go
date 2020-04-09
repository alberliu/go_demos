package client

import (
	"fmt"
	"go_demos/epoll/ge/codec"
	"log"
	"net"
	"testing"
	"time"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func TestEpoll_Client(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Println("Error dialing", err.Error())
		return // 终止程序
	}

	n, err := conn.Write(codec.Encode([]byte("hello")))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("write:", n)

	time.Sleep(2 * time.Second)
	n, err = conn.Write(codec.Encode([]byte("hello")))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("write:", n)

	time.Sleep(2 * time.Second)
	n, err = conn.Write(codec.Encode([]byte("hello")))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("write:", n)

	go func() {
		for {
			var bytes = make([]byte, 100)
			n, err := conn.Read(bytes)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println(string(bytes[0:n]))
		}
	}()

	select {}
}

func TestBenchEpoll_Client(t *testing.T) {
	var conns []net.Conn
	for i := 0; i < 10000; i++ {
		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			log.Println("Error dialing", err.Error())
			continue
		}
		conns = append(conns, conn)
		if i%100 == 0 {
			fmt.Println(i)
		}
		time.Sleep(time.Millisecond * 10)
	}

	select {}
}
