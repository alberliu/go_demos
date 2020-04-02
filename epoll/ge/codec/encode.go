package codec

import (
	"encoding/binary"
)

const headerLen = 2

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
