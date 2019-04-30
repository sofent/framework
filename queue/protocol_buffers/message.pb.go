// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package message

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Message struct {
	Hash                 string               `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Retries              uint32               `protobuf:"varint,2,opt,name=retries,proto3" json:"retries,omitempty"`
	Tried                uint32               `protobuf:"varint,3,opt,name=tried,proto3" json:"tried,omitempty"`
	PushedAt             *timestamp.Timestamp `protobuf:"bytes,4,opt,name=pushed_at,json=pushedAt,proto3" json:"pushed_at,omitempty"`
	Delay                *duration.Duration   `protobuf:"bytes,5,opt,name=delay,proto3" json:"delay,omitempty"`
	Param                []byte               `protobuf:"bytes,6,opt,name=param,proto3" json:"param,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Message) GetRetries() uint32 {
	if m != nil {
		return m.Retries
	}
	return 0
}

func (m *Message) GetTried() uint32 {
	if m != nil {
		return m.Tried
	}
	return 0
}

func (m *Message) GetPushedAt() *timestamp.Timestamp {
	if m != nil {
		return m.PushedAt
	}
	return nil
}

func (m *Message) GetDelay() *duration.Duration {
	if m != nil {
		return m.Delay
	}
	return nil
}

func (m *Message) GetParam() []byte {
	if m != nil {
		return m.Param
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "message")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8e, 0xb1, 0x4e, 0xc5, 0x20,
	0x14, 0x40, 0x83, 0xbe, 0xbe, 0x5a, 0xb4, 0x0b, 0x71, 0xc0, 0x0e, 0x4a, 0x9c, 0x98, 0x68, 0xa2,
	0x83, 0xb3, 0x89, 0x5f, 0x40, 0xdc, 0xcd, 0x6d, 0xb8, 0xb6, 0x4d, 0x8a, 0x10, 0xa0, 0x83, 0x1f,
	0xe9, 0x3f, 0x99, 0x42, 0xbb, 0xe8, 0x76, 0x0f, 0x9c, 0x0b, 0x87, 0xb6, 0x16, 0x63, 0x84, 0x11,
	0x95, 0x0f, 0x2e, 0xb9, 0xee, 0x61, 0x74, 0x6e, 0x5c, 0xb0, 0xcf, 0x34, 0xac, 0x9f, 0x7d, 0x9a,
	0x2d, 0xc6, 0x04, 0xd6, 0xef, 0xc2, 0xfd, 0x5f, 0xc1, 0xac, 0x01, 0xd2, 0xec, 0xbe, 0xca, 0xfd,
	0xe3, 0x0f, 0xa1, 0xf5, 0xfe, 0x24, 0x63, 0xf4, 0x34, 0x41, 0x9c, 0x38, 0x11, 0x44, 0x36, 0x3a,
	0xcf, 0x8c, 0xd3, 0x3a, 0x60, 0x0a, 0x33, 0x46, 0x7e, 0x21, 0x88, 0x6c, 0xf5, 0x81, 0xec, 0x96,
	0x56, 0xdb, 0x60, 0xf8, 0x65, 0x3e, 0x2f, 0xc0, 0x5e, 0x68, 0xe3, 0xd7, 0x38, 0xa1, 0xf9, 0x80,
	0xc4, 0x4f, 0x82, 0xc8, 0xeb, 0xa7, 0x4e, 0x95, 0x06, 0x75, 0x34, 0xa8, 0xf7, 0x23, 0x52, 0x5f,
	0x15, 0xf9, 0x35, 0xb1, 0x9e, 0x56, 0x06, 0x17, 0xf8, 0xe6, 0x55, 0x5e, 0xba, 0xfb, 0xb7, 0xf4,
	0xb6, 0x87, 0xeb, 0xe2, 0x6d, 0xff, 0x7b, 0x08, 0x60, 0xf9, 0x59, 0x10, 0x79, 0xa3, 0x0b, 0x0c,
	0xe7, 0xec, 0x3f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xf6, 0xb7, 0x8d, 0x28, 0x01, 0x00,
	0x00,
}
