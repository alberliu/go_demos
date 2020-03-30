package gepoll

import (
	"golang.org/x/sys/unix"
	"log"
	"net"
	"syscall"
)

type server struct {
	epoll   *epoll
	handler Handler
}

// NewServer 创建server服务器
func NewServer(address string, handler Handler) (*server, error) {
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	file, err := listener.File()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("listener_fd:", file.Fd())

	e, err := EpollCreate()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	e.AddListener(int(file.Fd()))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &server{
		epoll:   e,
		handler: handler,
	}, nil
}

// Run 启动服务
func (s *server) Run() {
	for {
		events, err := s.epoll.EpollWait()
		if err != nil {
			log.Println(err)
			continue
		}

		for i := range events {
			log.Println(events[i].Fd, events[i].Events, events[i].Pad)
			if events[i].Fd == int32(s.epoll.lfd) {
				log.Println("accept", events[i].Fd)
				nfd, _, err := unix.Accept(int(events[i].Fd))
				if err != nil {
					log.Println(err)
					continue
				}

				s.epoll.AddRead(nfd)
				s.handler.OnConnect(nfd)
			} else {
				bytes := make([]byte, 100)
				n, err := syscall.Read(int(events[i].Fd), bytes)
				if n == 0 || err != nil {
					log.Println("read_error:", n, err)

					err := s.epoll.Remove(int(events[i].Fd))
					if err != nil {
						log.Println(err)
					}

					s.handler.OnClose(int(events[i].Fd))
					continue
				}
				s.handler.OnMessage(int(events[i].Fd), bytes[0:n])
			}
			// time.Sleep(2*time.Second)
		}
	}
}
