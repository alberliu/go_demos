package main

import (
	"fmt"
	"net"
)

func main() {
	local, err := net.ResolveTCPAddr("tcp", "192.168.6.100:1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(local)
	remote, err := net.ResolveTCPAddr("tcp", "192.168.43.211:50000")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.DialTCP("tcp", local, remote)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for {
			var buf = make([]byte, 100)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(buf[:n]))
		}
	}()
	fmt.Println("write before")
	_, err = conn.Write([]byte("hello"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("write after")
	select {}
}
