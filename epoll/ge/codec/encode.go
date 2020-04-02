package codec

import (
	"encoding/binary"
)

// Encode 编码数据
func Encode(bytes []byte) []byte {
	buffer := make([]byte, 1024)
	// 将消息长度写入buffer
	binary.BigEndian.PutUint16(buffer[0:2], uint16(len(bytes)))
	// 将消息内容内容写入buffer
	copy(buffer[2:], bytes)
	return buffer[0 : 2+len(bytes)]
}
