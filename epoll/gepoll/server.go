package gepoll

import (
	"golang.org/x/sys/unix"
	"log"
	"net"
	"syscall"
)

// eventQueue 全局事件队列
var eventQueue = make(chan unix.EpollEvent, 1000)

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
	s.StartConsumer()
	s.StartProducer()
}

// StartProducer 启动生产者
func (s *server) StartProducer() {
	for {
		s.epoll.EpollWait()
	}
}

// StartConsumer 启动消费者
func (s *server) StartConsumer() {
	go s.Consume()
}

// Consume 消费者
func (s *server) Consume() {
	for event := range eventQueue {
		log.Println(event.Fd, event.Events, event.Pad)
		if event.Fd == int32(s.epoll.lfd) {
			log.Println("accept", event.Fd)
			nfd, _, err := unix.Accept(int(event.Fd))
			if err != nil {
				log.Println(err)
				return
			}

			s.epoll.AddRead(nfd)
			conn := newConn(nfd)
			conns.Store(nfd, conn)
			s.handler.OnConnect(conn)
			return
		}

		bytes := make([]byte, 100)
		n, err := syscall.Read(int(event.Fd), bytes)
		if n == 0 || err != nil {
			log.Println("read_error:", n, err)

			if err == syscall.EAGAIN {
				return
			}
			err := s.epoll.Remove(int(event.Fd))
			if err != nil {
				log.Println(err)
			}

			c, ok := conns.Load(event.Fd)
			if !ok {
				log.Println("not found in conns,", event.Fd)
				return
			}

			s.handler.OnClose(c.(*Conn))
			return
		}

		c, ok := conns.Load(event.Fd)
		if !ok {
			log.Println("not found in conns,", event.Fd)
			return
		}
		s.handler.OnMessage(c.(*Conn), bytes[0:n])
	}
}
