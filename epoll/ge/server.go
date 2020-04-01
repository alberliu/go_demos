package ge

import (
	"log"
	"net"
	"sync"
	"syscall"
	"time"
)

// Handler Server 注册接口
type Handler interface {
	OnConnect(c *Conn)
	OnMessage(c *Conn, message interface{})
	OnClose(c *Conn)
}

type server struct {
	epoll         *epoll                  // 系统相关网络模型
	handler       Handler                 // 注册的处理
	eventQueue    chan syscall.EpollEvent // 事件队列
	conns         sync.Map                // TCP长连接管理
	timeoutTicker time.Duration           // 超时时间检查间隔
	timeout       int64                   // 超时时间(单位秒)
	stop          chan int                // 服务器关闭信号
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
		epoll:      e,
		handler:    handler,
		eventQueue: make(chan syscall.EpollEvent, 1000),
		timeout:    0,
		stop:       make(chan int),
	}, nil
}

// SetTimeout 设置超时检查时间以及超时时间,默认不进行超时时间检查
func (s *server) SetTimeout(ticker, timeout time.Duration) {
	s.timeoutTicker = ticker
	s.timeout = int64(timeout.Seconds())
}

// Run 启动服务
func (s *server) Run() {
	s.startConsumer()
	s.checkTimeout()
	s.startProducer()
}

// Run 启动服务
func (s *server) Stop() {
	close(s.stop)
	close(s.eventQueue)
}

// StartProducer 启动生产者
func (s *server) startProducer() {
	for {
		select {
		case <-s.stop:
			log.Println("stop producer")
			return
		default:
			s.epoll.EpollWait(s.eventQueue)
		}
	}
}

// StartConsumer 启动消费者
func (s *server) startConsumer() {
	go s.consume()
}

// Consume 消费者
func (s *server) consume() {
	for event := range s.eventQueue {
		log.Println("event:", event.Fd, event.Events, event.Pad)
		// 客户端请求建立连接
		if event.Fd == int32(s.epoll.lfd) {
			log.Println("accept", event.Fd)
			nfd, _, err := syscall.Accept(int(event.Fd))
			if err != nil {
				log.Println(err)
				continue
			}

			s.epoll.AddRead(nfd)
			conn := newConn(int32(nfd), s)
			s.conns.Store(int32(nfd), conn)
			s.handler.OnConnect(conn)
			continue
		}

		v, ok := s.conns.Load(event.Fd)
		if !ok {
			log.Println("not found in conns,", event.Fd)
			continue
		}
		c := v.(*Conn)

		bytes := make([]byte, 100)
		n, err := c.Read(bytes)

		if n == 0 || err != nil {
			log.Println("read_error:", n, err)
			if err == syscall.EAGAIN {
				continue
			}

			c.Close()
			s.handler.OnClose(c)
			continue
		}
		log.Println("can read;", event.Fd, s.conns)
		s.handler.OnMessage(c, bytes[0:n])
	}
}

func (s *server) checkTimeout() {
	if s.timeout == 0 || s.timeoutTicker == 0 {
		return
	}
	go func() {
		ticker := time.NewTicker(s.timeoutTicker)
		for {
			select {
			case <-s.stop:
				return
			case <-ticker.C:
				s.conns.Range(func(key, value interface{}) bool {
					c := value.(*Conn)

					if time.Now().Unix()-c.lastReadTime > s.timeout {
						c.Close()
						s.handler.OnClose(c)
					}
					return true
				})
			}
		}
	}()
}
