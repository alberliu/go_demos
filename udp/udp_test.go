package udp

import (
	"bytes"
	"fmt"
	"net"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("listen")
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	fmt.Println("read before")
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	fmt.Println("read after")
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}

func TestClient(t *testing.T) {
	udpAddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:1200")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = conn.Write([]byte("anything"))
	if err != nil {
		fmt.Println(err)
		return
	}
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf[0:n]))
}

func portInUse(portNumber int) int {
	res := -1
	var outBytes bytes.Buffer
	cmdStr := fmt.Sprintf("netstat -ano -p tcp | findstr %d", portNumber)
	cmd := exec.Command("cmd", "/c", cmdStr)
	cmd.Stdout = &outBytes
	cmd.Run()
	resStr := outBytes.String()
	r := regexp.MustCompile(`\s\d+\s`).FindAllString(resStr, -1)
	if len(r) > 0 {
		pid, err := strconv.Atoi(strings.TrimSpace(r[0]))
		if err != nil {
			res = -1
		} else {
			res = pid
		}
	}
	return res
}

func TestPort(t *testing.T) {
	fmt.Println(portInUse(80))
}
