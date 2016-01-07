// Code generated by protoc-gen-go.
// source: TextMessage.proto
// DO NOT EDIT!

/*
Package gochat is a generated protocol buffer package.

It is generated from these files:
	TextMessage.proto

It has these top-level messages:
	TextMessage
*/
package gochat

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type TextMessage struct {
	MessageUid string `protobuf:"bytes,1,opt,name=message_uid" json:"message_uid,omitempty"`
	SenderId   int32  `protobuf:"varint,2,opt,name=sender_id" json:"sender_id,omitempty"`
	SenderAddr string `protobuf:"bytes,3,opt,name=sender_addr" json:"sender_addr,omitempty"`
	Timestamp  int64  `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	Text       string `protobuf:"bytes,5,opt,name=text" json:"text,omitempty"`
}

func (m *TextMessage) Reset()                    { *m = TextMessage{} }
func (m *TextMessage) String() string            { return proto.CompactTextString(m) }
func (*TextMessage) ProtoMessage()               {}
func (*TextMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*TextMessage)(nil), "gochat.TextMessage")
}

var fileDescriptor0 = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x0c, 0x49, 0xad, 0x28,
	0xf1, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b,
	0xcf, 0x4f, 0xce, 0x48, 0x2c, 0x51, 0x9a, 0xc7, 0xc8, 0xc5, 0x8d, 0x24, 0x2b, 0x24, 0xcf, 0xc5,
	0x9d, 0x0b, 0x61, 0xc6, 0x97, 0x66, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x71, 0x41,
	0x85, 0x42, 0x33, 0x53, 0x84, 0xa4, 0xb9, 0x38, 0x8b, 0x53, 0xf3, 0x52, 0x52, 0x8b, 0xe2, 0x81,
	0xd2, 0x4c, 0x40, 0x69, 0xd6, 0x20, 0x0e, 0x88, 0x80, 0x67, 0x0a, 0x48, 0x37, 0x54, 0x32, 0x31,
	0x25, 0xa5, 0x48, 0x82, 0x19, 0xa2, 0x1b, 0x22, 0xe4, 0x08, 0x14, 0x11, 0x92, 0xe1, 0xe2, 0x2c,
	0xc9, 0x04, 0x9a, 0x56, 0x92, 0x98, 0x5b, 0x20, 0xc1, 0x02, 0x94, 0x66, 0x0e, 0x42, 0x08, 0x08,
	0x09, 0x71, 0xb1, 0x94, 0x00, 0xdd, 0x22, 0xc1, 0x0a, 0xd6, 0x07, 0x66, 0x27, 0xb1, 0x81, 0xdd,
	0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x21, 0xed, 0x5d, 0x33, 0xc4, 0x00, 0x00, 0x00,
}