package tcp

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
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
	buf := make([]byte, 10)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Received data: %v \n", string(buf[:n]))

		n, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func TestClient(t *testing.T) {
	//打开连接:
	conn, err := net.Dial("tcp", "127.0.0.1:50000")
	if err != nil {
		//由于目标计算机积极拒绝而无法创建连接
		fmt.Println("Error dialing", err.Error())
		return // 终止程序
	}

	fmt.Println("write before")
	_, err = conn.Write([]byte("hello hello hello hello hello hello hello "))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("write after")
	go func() {
		var buf = make([]byte, 10)
		for {

			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(buf[:n]))
		}
	}()
	select {}
}

func TestVirtualIPClient(t *testing.T) {
	local, err := net.ResolveTCPAddr("tcp", "192.168.41.211:50001")
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
	io.Copy(ioutil.Discard, conn)
}
