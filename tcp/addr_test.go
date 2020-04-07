package tcp

import (
	"fmt"
	"log"
	"net"
	"testing"
)

func TestIP(t *testing.T) {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(addr.Port)
	fmt.Println(([]byte)(addr.IP))
}
