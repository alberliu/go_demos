package main

import (
	"fmt"
	"log"
	"net"
)

// docker run -v $(pwd)/:/main -d alpine /main

func main() {
	fmt.Println("Starting the server ...")
	// 创建 listener
	listener, err := net.Listen("tcp", ":50000")
	if err != nil {
		log.Println(err)
		return //终止程序
	}
	// 监听并接受来自客户端的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			return // 终止程序
		}
		fmt.Println("listener ok")
		go HandlerConn(conn)
	}
}

func HandlerConn(conn net.Conn) {
	conn.RemoteAddr()
	buf := make([]byte, 10)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.Close()
		fmt.Printf("Received data: %v \n", string(buf[:n]))

		n, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
