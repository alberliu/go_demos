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
	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_DEL, fd, nil)
	if err != nil {
		return err
	}

	return nil
}

func (e *epoll) EpollWait() ([]unix.EpollEvent, error) {
	log.Println("wait start")
	events := make([]unix.EpollEvent, 100)
	n, err := unix.EpollWait(e.fd, events, -1)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("wait end")
	return events[0:n], nil

	for i := 0; i < n; i++ {
		log.Println(events[i].Fd, events[i].Events, events[i].Pad)
		if events[i].Fd == int32(e.lfd) {
			log.Println("accept", events[i].Fd)
			nfd, _, err := unix.Accept(int(events[i].Fd))
			if err != nil {
				log.Println(err)
				continue
			}

			e.AddRead(nfd)

			// 可以确定，是阻塞调用的
			n, err := unix.Write(nfd, make([]byte, 1024*1024*1024))
			if err != nil {
				log.Println(err)
				return
			}
			log.Println("write:", n)

		} else {
			bytes := make([]byte, 100)
			n, err := unix.Read(int(events[i].Fd), bytes)
			if n == 0 || err != nil {
				log.Println("read_error:", n, err)
				err := e.Remove(int(events[i].Fd))
				if err != nil {
					log.Println(err)
				}

				err = os.NewFile(uintptr(events[i].Fd), "").Close()
				if err != nil {
					log.Println(err)
				}
				continue
			}
			log.Println("read:", string(bytes[0:n]))
		}
	}
}
