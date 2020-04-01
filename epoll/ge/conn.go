package ge

import (
	"log"
	"sync"
	"syscall"
	"time"
)

type Conn struct {
	m            sync.RWMutex           // 锁
	s            *server                // 服务器引用
	fd           int32                  // 文件描述符
	lastReadTime int64                  // 最后一次读取数据的时间
	extra        map[string]interface{} // 扩展字段
}

func newConn(fd int32, s *server) *Conn {
	return &Conn{
		s:            s,
		fd:           fd,
		lastReadTime: time.Now().Unix(),
		extra:        make(map[string]interface{}),
	}
}

// 关闭连接
func (c *Conn) GetFd() int32 {
	c.m.RLock()
	defer c.m.RUnlock()

	return c.fd
}

func (c *Conn) Read(bytes []byte) (int, error) {
	c.m.Lock()
	defer c.m.Unlock()

	c.lastReadTime = time.Now().Unix()
	return syscall.Read(int(c.fd), bytes)
}

func (c *Conn) Write(bytes []byte) (int, error) {
	c.m.Lock()
	defer c.m.Unlock()

	return syscall.Write(int(c.fd), bytes)
}

// 关闭连接
func (c *Conn) Close() error {
	c.m.Lock()
	defer c.m.Unlock()

	// 从epoll监听的文件描述符中删除
	err := c.s.epoll.RemoveAndClose(int(c.fd))
	if err != nil {
		log.Println(err)
	}

	// 从conns中删除conn
	c.s.conns.Delete(c.fd)
	return nil
}
