package codec

import (
	"encoding/binary"
	"io"
	"log"
	"syscall"
)

const headerLen = 2

const errRTUStr = "resource temporarily unavailable"

// Encode 编码数据
func Encode(bytes []byte) []byte {
	l := len(bytes)
	buffer := make([]byte, l+headerLen)
	// 将消息长度写入buffer
	binary.BigEndian.PutUint16(buffer[0:2], uint16(l))
	// 将消息内容内容写入buffer
	copy(buffer[headerLen:], bytes)
	return buffer[0 : headerLen+l]
}

// Decode 解码
func Decode(fd int) ([]byte, bool, error) {
	var bytes = make([]byte, 1024)
	n, _, err := syscall.Recvfrom(fd, bytes, syscall.MSG_PEEK|syscall.MSG_DONTWAIT)
	if err != nil {
		if err == syscall.EAGAIN {
			return nil, false, nil
		}
		log.Println("recv_from error", err)
		if err == syscall.EBADF {
			log.Println("syscall.EBADF")
		}
		return nil, false, err
	}
	if n == 0 {
		return nil, false, io.EOF
	}

	if n < headerLen {
		return nil, false, nil
	}

	if n > 1024-headerLen {
	}
	valueLen := int(binary.BigEndian.Uint16(bytes[0:headerLen]))
	if n < headerLen+valueLen {
		return nil, false, nil
	}
	n, _, err = syscall.Recvfrom(fd, bytes[0:headerLen+valueLen], syscall.MSG_DONTWAIT)
	if err != nil {
		log.Println("recv_from error", err)
		return nil, false, err
	}
	return bytes[headerLen : headerLen+valueLen], true, nil
}
