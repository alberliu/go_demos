// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Pb struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pb) Reset()         { *m = Pb{} }
func (m *Pb) String() string { return proto.CompactTextString(m) }
func (*Pb) ProtoMessage()    {}
func (*Pb) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *Pb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pb.Unmarshal(m, b)
}
func (m *Pb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pb.Marshal(b, m, deterministic)
}
func (m *Pb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pb.Merge(m, src)
}
func (m *Pb) XXX_Size() int {
	return xxx_messageInfo_Pb.Size(m)
}
func (m *Pb) XXX_DiscardUnknown() {
	xxx_messageInfo_Pb.DiscardUnknown(m)
}

var xxx_messageInfo_Pb proto.InternalMessageInfo

func (m *Pb) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func init() {
	proto.RegisterType((*Pb)(nil), "pb.Pb")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 65 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x12, 0xe2, 0x62, 0x0a, 0x48,
	0x12, 0xe2, 0xe1, 0x62, 0x4c, 0x94, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62, 0x4c, 0x4c, 0x62,
	0x03, 0x4b, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x45, 0xef, 0x50, 0x2c, 0x00, 0x00,
	0x00,
}