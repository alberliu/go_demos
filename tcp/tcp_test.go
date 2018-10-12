package tcp

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"strings"
	"testing"
	"time"
	"strconv"
	"log"
	"github.com/astaxie/beego/logs"
)

func TestServer(t *testing.T) {
	fmt.Println("Starting the server ...")
	// 创建 listener
	listener, err := net.Listen("tcp", "localhost:50000")
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
		go doServerStuff(conn)
		go func(){
			time.Sleep(5*time.Second)
			conn.Close()
		}()
	}
}

func doServerStuff(conn net.Conn) {
	for {
		conn.SetReadDeadline(time.Now().Add(time.Second*10))
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			logs.Error(err)
			return
		}
		fmt.Printf("Received data: %v", string(buf))
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
	clientName, _ := inputReader.ReadString('')
	trimmedClient := strings.Trim(clientName, "\r")
	// 给服务器发送信息直到程序退出：
	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('')
		trimmedInput := strings.Trim(input, "\r")
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
	}
}

func TestClient2(t *testing.T) {
	//打开连接:
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		//由于目标计算机积极拒绝而无法创建连接
		fmt.Println("Error dialing", err.Error())
		return // 终止程序
	}
	fmt.Println(conn)
	//conn.Close()
}

func TestClient3(t *testing.T) {
	//打开连接:
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		//由于目标计算机积极拒绝而无法创建连接
		log.Println(err)
		return // 终止程序
	}
	for i:=0;i<20;i++{
		time.Sleep( time.Duration(i) * time.Minute)
		conn.Write([]byte(strconv.Itoa(i)))
	}
}
