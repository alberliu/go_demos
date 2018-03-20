package buf

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"strings"
	"testing"
	"time"
	"bytes"
)

func TestServer(t *testing.T) {
	fmt.Println("Starting the server ...")
	// 创建 listener
	listener, err := net.Listen("tcp", "localhost:50001")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return //终止程序
	}
	// 监听并接受来自客户端的连接
	for {
		conn, err := listener.Accept()
		fmt.Println(conn.RemoteAddr())
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 终止程序
		}
		go doReadBufio(conn)
	}
}
func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 10)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		fmt.Println("Received data:", time.Now(), string(buf))

	}
}
func doReadBuffer(conn net.Conn) {
	buffer := bytes.NewBuffer(make([]byte, 0, 20))
	bytes := make([]byte, 10);
	for {
		rl, err := conn.Read(bytes);
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		fmt.Println(string(bytes))
		fmt.Println(rl)
		wl, err := buffer.Write(bytes[0:rl])
		fmt.Println(wl)
		if buffer.Len() > 15 {
			byts := make([]byte, 15);
			readl, err := buffer.Read(byts)
			if err != nil {
				fmt.Println("Error reading", err.Error())
				return
			}
			fmt.Println(readl)
			fmt.Println(string(byts))
		}
		fmt.Println()
	}
}

func doReadBufio(conn net.Conn) {
	reader := bufio.NewReaderSize(conn, 100)
	bytes:=make([]byte,15)
	for {
		if reader.Buffered() > 15 {
			len,err:=reader.Read(bytes)
			if err!=nil {
				fmt.Println(len)
				fmt.Println(string(bytes))
			}
		}
	}
}

func TestClient(t *testing.T) {
	//打开连接:
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		//由于目标计算机积极拒绝而无法创建连接
		fmt.Println("Error dialing", err.Error())
		return // 终止程序
	}
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\r\n")
	// 给服务器发送信息直到程序退出：
	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
	}
}

func TestClient2(t *testing.T) {
	//打开连接:
	conn, err := net.Dial("tcp", "localhost:50001")
	if err != nil {
		//由于目标计算机积极拒绝而无法创建连接
		fmt.Println("Error dialing", err.Error())
		return // 终止程序
	}

	conn.Write([]byte("12345"))
	conn.Write([]byte("12345"))
	conn.Write([]byte("12345"))
	conn.Write([]byte("12345"))
	conn.Write([]byte("12345"))
	//time.Sleep(10 * time.Second)

}
