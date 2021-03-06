// Code generated by protoc-gen-go.
// source: notification.proto
// DO NOT EDIT!

/*
Package notifications is a generated protocol buffer package.

It is generated from these files:
	notification.proto

It has these top-level messages:
	Error
*/
package notifications

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Error_Level int32

const (
	Error_debug   Error_Level = 0
	Error_info    Error_Level = 1
	Error_warning Error_Level = 2
	Error_error   Error_Level = 3
	Error_fatal   Error_Level = 4
)

var Error_Level_name = map[int32]string{
	0: "debug",
	1: "info",
	2: "warning",
	3: "error",
	4: "fatal",
}
var Error_Level_value = map[string]int32{
	"debug":   0,
	"info":    1,
	"warning": 2,
	"error":   3,
	"fatal":   4,
}

func (x Error_Level) String() string {
	return proto.EnumName(Error_Level_name, int32(x))
}
func (Error_Level) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// endpoint: POST /v1/errors
type Error struct {
	Module  string      `protobuf:"bytes,1,opt,name=module" json:"module,omitempty"`
	Id      uint32      `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Type    string      `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Subject string      `protobuf:"bytes,4,opt,name=subject" json:"subject,omitempty"`
	Level   Error_Level `protobuf:"varint,5,opt,name=level,enum=notifications.Error_Level" json:"level,omitempty"`
	Message string      `protobuf:"bytes,6,opt,name=message" json:"message,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Error) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *Error) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Error) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Error) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Error) GetLevel() Error_Level {
	if m != nil {
		return m.Level
	}
	return Error_debug
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "notifications.Error")
	proto.RegisterEnum("notifications.Error_Level", Error_Level_name, Error_Level_value)
}

func init() { proto.RegisterFile("notification.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0xcd, 0x36, 0xd9, 0xda, 0x91, 0x96, 0x30, 0x07, 0x09, 0x9e, 0xa4, 0x27, 0x4f, 0x41,
	0xf4, 0x01, 0x3c, 0x79, 0xf3, 0xb4, 0x6f, 0x90, 0xed, 0xce, 0x2e, 0x91, 0x34, 0x29, 0x49, 0x56,
	0xf1, 0x95, 0x7d, 0x0a, 0x37, 0x69, 0x05, 0xbd, 0xfd, 0xff, 0xc7, 0xc7, 0x3f, 0x30, 0x80, 0x3e,
	0x64, 0x3b, 0xda, 0x83, 0xc9, 0x36, 0x78, 0x7d, 0x8a, 0x21, 0x07, 0xdc, 0xfe, 0x65, 0x69, 0xff,
	0xcd, 0x40, 0xbc, 0xc6, 0x18, 0x22, 0xde, 0x42, 0x7b, 0x0c, 0xc3, 0xec, 0x48, 0xb1, 0x7b, 0xf6,
	0xb0, 0xe9, 0x2e, 0x0d, 0x77, 0xd0, 0xd8, 0x41, 0x35, 0x0b, 0xdb, 0x76, 0x4b, 0x42, 0x04, 0x9e,
	0xbf, 0x4e, 0xa4, 0x56, 0xd5, 0xaa, 0x19, 0x15, 0xac, 0xd3, 0xdc, 0xbf, 0xd3, 0x21, 0x2b, 0x5e,
	0xf1, 0x6f, 0xc5, 0x47, 0x10, 0x8e, 0x3e, 0xc8, 0x29, 0xb1, 0xf0, 0xdd, 0xd3, 0x9d, 0xfe, 0x77,
	0x5e, 0xd7, 0xd3, 0xfa, 0xad, 0x18, 0xdd, 0x59, 0x2c, 0x5b, 0x47, 0x4a, 0xc9, 0x4c, 0xa4, 0xda,
	0xf3, 0xd6, 0xa5, 0xee, 0x5f, 0x40, 0x54, 0x13, 0x37, 0x20, 0x06, 0xea, 0xe7, 0x49, 0x5e, 0xe1,
	0x35, 0x70, 0xeb, 0xc7, 0x20, 0x19, 0xde, 0xc0, 0xfa, 0xd3, 0x44, 0x6f, 0xfd, 0x24, 0x9b, 0x62,
	0x50, 0x99, 0x96, 0xab, 0x12, 0x47, 0x93, 0x8d, 0x93, 0xbc, 0x6f, 0xeb, 0x0b, 0x9e, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x57, 0x0d, 0xdc, 0x0c, 0x18, 0x01, 0x00, 0x00,
}
