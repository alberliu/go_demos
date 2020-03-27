package gepoll

import (
	"log"
	"net"
	"testing"
	"time"
)

func TestEpoll_Add(t *testing.T) {
	conn, err := net.Dial("tcp", "172.16.4.136:8080")
	if err != nil {
		log.Println("Error dialing", err.Error())
		return // 终止程序
	}

	go func() {
		time.Sleep(time.Second)
		bytes := make([]byte, 1024*1024*1024)
		n, err := conn.Write(bytes)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("write:", n)
	}()

	go func() {
		for {
			var bytes = make([]byte, 100)
			n, err := conn.Read(bytes)
			if err != nil {
				log.Println(err)
				return
			}
			log.Println("read:", string(bytes[0:n]))
		}
	}()

	select {}
}
