package tcp

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"testing"
	"time"

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
		go func() {
			time.Sleep(5 * time.Second)
			err = conn.Close()
			if err != nil {
				fmt.Println("1", err)
			}
			err = conn.Close()
			if err != nil {
				fmt.Println("2", err)
			}
		}()
	}
}

func doServerStuff(conn net.Conn) {
	for {
		conn.SetReadDeadline(time.Now().Add(time.Second * 10))
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			logs.Error(err)
			return
		}
		fmt.Printf("Received data: %v", string(buf))
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
	time.Sleep(10 * time.Second)
}

func TestClient3(t *testing.T) {
	//打开连接:
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		//由于目标计算机积极拒绝而无法创建连接
		log.Println(err)
		return // 终止程序
	}
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(i) * time.Minute)
		conn.Write([]byte(strconv.Itoa(i)))
	}
}
