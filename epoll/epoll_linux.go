package epoll

import (
	"golang.org/x/sys/unix"
	"log"
	"net"
	"sync"
	"syscall"
)

type epoll struct {
	fd    int
	conns sync.Map
}

func EpollCreate() (*epoll, error) {
	fd, err := unix.EpollCreate1(0)
	if err != nil {
		return nil, err
	}

	return &epoll{
		fd:    fd,
		conns: make(map[int]net.Conn),
	}, nil
}

func (e *epoll) Add(conn net.TCPConn) error {
	file, err := conn.File()
	if err != nil {
		return err
	}

	fd := file.Fd()
	err = unix.EpollCtl(e.fd, syscall.EPOLL_CTL_ADD, int(fd), &unix.EpollEvent{Events: unix.POLLIN | unix.POLLHUP, Fd: int32(file.Fd())})
	if err != nil {
		return err
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	e.conns[fd] = conn
	if len(e.conns)%100 == 0 {
		log.Printf("total number of connections: %v", len(e.connections))
	}
	return nil
}

func (e *epoll) Remove(conn net.TCPConn) error {
	fd := socketFD(conn)
	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_DEL, fd, nil)
	if err != nil {
		return err
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	delete(e.conns, fd)
	if len(e.conns)%100 == 0 {
		log.Printf("total number of connections: %v", len(e.connections))
	}
	return nil
}

func (e *epoll) Wait() ([]net.TCPConn, error) {
	events := make([]unix.EpollEvent, 100)
	n, err := unix.EpollWait(e.fd, events, 100)
	if err != nil {
		return nil, err
	}

	var connections []net.Conn
	for i := 0; i < n; i++ {
		conn := e.connections[int(events[i].Fd)]
		connections = append(connections, conn)
	}
	return connections, nil
}
