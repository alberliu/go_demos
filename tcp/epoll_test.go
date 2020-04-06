package tcp

import (
	"log"
	"syscall"
	"testing"
)

func TestEpoll_Add(t *testing.T) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Println(err)
		return
	}

	syscall.Bind(fd, &syscall.SockaddrInet4{Port: 8080})

	syscall.Listen(fd, 1024)
}
