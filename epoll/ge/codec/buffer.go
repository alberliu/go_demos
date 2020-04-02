package codec

import (
	"errors"
	"io"
	"syscall"
)

var (
	ErrNotEnough = errors.New("not enough")
)

// Buffer 读缓冲区,每个tcp长连接对应一个读缓冲区
type Buffer struct {
	buf   []byte // 应用内缓存区
	start int    // 有效字节开始位置
	end   int    // 有效字节结束位置
}

// NewBuffer 创建一个缓存区
func NewBuffer(bytes []byte) *Buffer {
	return &Buffer{bytes, 0, 0}
}

func (b *Buffer) Len() int {
	return b.end - b.start
}

// grow 将有效的字节前移
func (b *Buffer) grow() {
	if b.start == 0 {
		return
	}
	copy(b.buf, b.buf[b.start:b.end])
	b.end -= b.start
	b.start = 0
}

// readFromFile 从文件描述符里面读取数据，如果reader阻塞，会发生阻塞
func (b *Buffer) ReadFromFile(fd int32) error {
	b.grow()
	n, err := syscall.Read(int(fd), b.buf[b.end:])
	if n == 0 || err != nil {
		if err == syscall.EAGAIN {
			return nil
		}
		if err != nil {
			return err
		}
		return io.EOF
	}

	b.end += n
	return nil
}

// seek 返回n个字节，而不产生移位，如果没有足够字节，返回错误
func (b *Buffer) Seek(start, end int) ([]byte, error) {
	if b.end-b.start >= end-start {
		buf := b.buf[b.start+start : b.start+end]
		return buf, nil
	}
	return nil, ErrNotEnough
}

// read 舍弃offset个字段，读取n个字段,如果没有足够的字节，返回错误
func (b *Buffer) Read(offset, limit int) ([]byte, error) {
	if b.Len() < offset+limit {
		return nil, ErrNotEnough
	}
	b.start += offset
	buf := b.buf[b.start : b.start+limit]
	b.start += limit
	return buf, nil
}

// Get 获取buffer中的字节数组，不改变Buffer中的字节流
func (b *Buffer) Get() []byte {
	return b.buf[b.start:b.end]
}
