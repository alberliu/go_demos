package codec

import (
	"encoding/binary"
	"errors"
)

var (
	ErrIllegalHeader = errors.New("illegal header len")
)

func Decode(b *Buffer) ([]byte, bool, error) {
	// 读取数据长度
	lenBuf, err := b.Seek(0, 2)
	if err != nil {
		return nil, false, nil
	}
	// 读取数据内容
	valueLen := int(binary.BigEndian.Uint16(lenBuf))
	// 数据的字节数组长度大于buffer的长度，返回错误
	if valueLen > 1024 {
		return nil, false, ErrIllegalHeader
	}
	valueBuf, err := b.Read(2, valueLen)
	if err != nil {
		return nil, false, nil
	}
	return valueBuf, true, nil
}
