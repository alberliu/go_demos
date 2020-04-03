package ge

import (
	"go_demos/epoll/ge/codec"
	"log"
	"sync"
	"syscall"
	"time"
)

const maxBufferLen = 1024

type Conn struct {
	rm           sync.Mutex             // read锁
	wm           sync.Mutex             // write锁
	s            *server                // 服务器引用
	fd           int32                  // 文件描述符
	readBuffer   *codec.Buffer          // 读缓存区
	lastReadTime int64                  // 最后一次读取数据的时间
	extra        map[string]interface{} // 扩展字段
}

func newConn(fd int32, s *server) *Conn {
	return &Conn{
		s:            s,
		fd:           fd,
		readBuffer:   codec.NewBuffer(make([]byte, maxBufferLen)),
		lastReadTime: time.Now().Unix(),
		extra:        make(map[string]interface{}),
	}
}

// 关闭连接
func (c *Conn) GetFd() int32 {
	return c.fd
}

// 关闭连接
func (c *Conn) GetRemoteIP() int32 {
	return c.fd
}

func (c *Conn) Read() error {
	c.rm.Lock()
	defer c.rm.Unlock()

	err := c.readBuffer.ReadFromFile(c.fd)
	if err != nil {
		return err
	}
	c.lastReadTime = time.Now().Unix()

	bytes, ok, err := codec.Decode(c.readBuffer)
	if err != nil {
		return err
	}

	if ok {
		c.s.handler.OnMessage(c, bytes)
	}
	return nil
}

func (c *Conn) Write(bytes []byte) (int, error) {
	/*c.wm.Lock()
	defer c.wm.Unlock()*/

	return syscall.Write(int(c.fd), bytes)
}

// Close 关闭连接
func (c *Conn) Close() error {
	// 从epoll监听的文件描述符中删除
	err := c.s.epoll.RemoveAndClose(int(c.fd))
	if err != nil {
		log.Println(err)
	}

	// 从conns中删除conn
	c.s.conns.Delete(c.fd)
	return nil
}
