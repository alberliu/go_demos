// +build linux

package ge

import (
	"golang.org/x/sys/unix"
	"log"
	"syscall"
)

const PollAll = unix.POLLIN | unix.POLLPRI | unix.POLLERR | unix.POLLHUP | unix.POLLNVAL

type epoll struct {
	fd  int
	lfd int
}

func EpollCreate() (*epoll, error) {
	fd, err := syscall.EpollCreate1(0)
	if err != nil {
		return nil, err
	}
	return &epoll{
		fd: fd,
	}, nil
}

func (e *epoll) AddListener(fd int) error {
	err := syscall.EpollCtl(e.fd, syscall.EPOLL_CTL_ADD, int(fd), &syscall.EpollEvent{Events: unix.POLLIN | unix.POLLHUP, Fd: int32(fd)})
	if err != nil {
		return err
	}
	e.lfd = fd
	return nil
}

func (e *epoll) AddRead(fd int) error {
	err := syscall.EpollCtl(e.fd, syscall.EPOLL_CTL_ADD, int(fd), &syscall.EpollEvent{
		Events: PollAll,
		Fd:     int32(fd),
	})
	if err != nil {
		return err
	}
	return nil
}

func (e *epoll) RemoveAndClose(fd int) error {
	// 移除文件描述符的监听
	err := syscall.EpollCtl(e.fd, syscall.EPOLL_CTL_DEL, fd, nil)
	if err != nil {
		return err
	}

	err = syscall.Close(fd)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (e *epoll) EpollWait(eventQueue chan syscall.EpollEvent) {
	events := make([]syscall.EpollEvent, 100)
	n, err := syscall.EpollWait(e.fd, events, -1)
	if err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < n; i++ {
		eventQueue <- events[i]
	}
}
