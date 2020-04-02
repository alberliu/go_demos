package ge

import (
	"encoding/binary"
	"errors"
	"log"
	"sync"
	"syscall"
	"time"
)

type Conn struct {
	m            sync.RWMutex           // 锁
	s            *server                // 服务器引用
	fd           int32                  // 文件描述符
	readBuffer   buffer                 // 读缓存区
	lastReadTime int64                  // 最后一次读取数据的时间
	extra        map[string]interface{} // 扩展字段
}

func newConn(fd int32, s *server) *Conn {
	return &Conn{
		s:            s,
		fd:           fd,
		readBuffer:   newBuffer(make([]byte, 1000)),
		lastReadTime: time.Now().Unix(),
		extra:        make(map[string]interface{}),
	}
}

// 关闭连接
func (c *Conn) GetFd() int32 {
	return c.fd
}

func (c *Conn) Read() error {
	c.m.Lock()
	defer c.m.Unlock()

	err := c.readBuffer.readFromFile(c.fd)
	if err != nil {
		return err
	}
	c.lastReadTime = time.Now().Unix()

	// 读取数据长度
	lenBuf, err := c.readBuffer.seek(0, 2)
	if err != nil {
		return nil
	}
	log.Println("seek")
	// 读取数据内容
	valueLen := int(binary.BigEndian.Uint16(lenBuf))
	log.Println("seek:len:", valueLen)
	// 数据的字节数组长度大于buffer的长度，返回错误
	if valueLen > 1024 {
		return errors.New("illegal len")
	}
	log.Println("seek:len:1", valueLen)
	valueBuf, err := c.readBuffer.read(2, valueLen)
	if err != nil {
		log.Println("readBuffer.read:", err)
		return nil
	}
	log.Println("seek:len:2", valueLen)
	c.s.handler.OnMessage(c, valueBuf)
	log.Println("seek:len:3", valueLen)
	return nil
}

func (c *Conn) Write(bytes []byte) (int, error) {
	c.m.Lock()
	defer c.m.Unlock()

	return syscall.Write(int(c.fd), bytes)
}

// Close 关闭连接
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
