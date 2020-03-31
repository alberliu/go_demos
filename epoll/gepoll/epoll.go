// +build linux

package gepoll

import (
	"golang.org/x/sys/unix"
	"log"
	"os"
	"syscall"
)

const PollAll = unix.POLLIN | unix.POLLPRI | unix.POLLERR | unix.POLLHUP | unix.POLLNVAL

type epoll struct {
	fd  int
	lfd int
}

func EpollCreate() (*epoll, error) {
	fd, err := unix.EpollCreate1(0)
	if err != nil {
		return nil, err
	}
	return &epoll{
		fd: fd,
	}, nil
}

func (e *epoll) AddListener(fd int) error {
	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_ADD, int(fd), &unix.EpollEvent{Events: unix.POLLIN | unix.POLLHUP, Fd: int32(fd)})
	if err != nil {
		return err
	}
	e.lfd = fd
	return nil
}

func (e *epoll) AddRead(fd int) error {
	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_ADD, int(fd), &unix.EpollEvent{
		Events: PollAll,
		Fd:     int32(fd),
	})
	if err != nil {
		return err
	}
	return nil
}

func (e *epoll) Remove(fd int) error {
	// 移除文件描述符的监听
	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_DEL, fd, nil)
	if err != nil {
		return err
	}

	// 关闭文件描述符
	err = os.NewFile(uintptr(fd), "").Close()
	if err != nil {
		log.Println(err)
	}

	return nil
}

func (e *epoll) EpollWait() {
	log.Println("wait start")
	events := make([]unix.EpollEvent, 100)
	n, err := unix.EpollWait(e.fd, events, -1)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("wait end")

	for i := 0; i < n; i++ {
		eventQueue <- events[i]
	}
}
