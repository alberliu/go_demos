package gepoll

import (
	"sync"
	"time"
)

var conns sync.Map

type Conn struct {
	fd           int                    // 文件描述符
	lastReadTime int64                  // 最后一次读取数据的时间
	extra        map[string]interface{} // 扩展字段
}

func newConn(fd int) *Conn {
	return &Conn{
		fd:           fd,
		lastReadTime: time.Now().Unix(),
		extra:        make(map[string]interface{}),
	}
}

func (c *Conn) Close() {

}
