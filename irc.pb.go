// Code generated by protoc-gen-go.
// source: irc.proto
// DO NOT EDIT!

/*
Package lirc is a generated protocol buffer package.

It is generated from these files:
	irc.proto

It has these top-level messages:
	Actor
	ChannelAction
*/
package lirc

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

type ActionType int32

const (
	ActionType_MESSAGE ActionType = 0
	ActionType_JOIN    ActionType = 1
	ActionType_QUIT    ActionType = 2
	ActionType_NICK    ActionType = 3
)

var ActionType_name = map[int32]string{
	0: "MESSAGE",
	1: "JOIN",
	2: "QUIT",
	3: "NICK",
}
var ActionType_value = map[string]int32{
	"MESSAGE": 0,
	"JOIN":    1,
	"QUIT":    2,
	"NICK":    3,
}

func (x ActionType) String() string {
	return proto.EnumName(ActionType_name, int32(x))
}
func (ActionType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Actor struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	User string `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Host string `protobuf:"bytes,3,opt,name=host" json:"host,omitempty"`
}

func (m *Actor) Reset()                    { *m = Actor{} }
func (m *Actor) String() string            { return proto.CompactTextString(m) }
func (*Actor) ProtoMessage()               {}
func (*Actor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Actor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Actor) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Actor) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type ChannelAction struct {
	// timestamp in seconds since Unix epoch
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	// Message.Prefix
	Actor *Actor `protobuf:"bytes,2,opt,name=actor" json:"actor,omitempty"`
	// channel name
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	// Message.Command
	Type ActionType `protobuf:"varint,4,opt,name=type,enum=lirc.ActionType" json:"type,omitempty"`
	// Message.Trailing
	Message string `protobuf:"bytes,5,opt,name=message" json:"message,omitempty"`
}

func (m *ChannelAction) Reset()                    { *m = ChannelAction{} }
func (m *ChannelAction) String() string            { return proto.CompactTextString(m) }
func (*ChannelAction) ProtoMessage()               {}
func (*ChannelAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ChannelAction) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ChannelAction) GetActor() *Actor {
	if m != nil {
		return m.Actor
	}
	return nil
}

func (m *ChannelAction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChannelAction) GetType() ActionType {
	if m != nil {
		return m.Type
	}
	return ActionType_MESSAGE
}

func (m *ChannelAction) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Actor)(nil), "lirc.Actor")
	proto.RegisterType((*ChannelAction)(nil), "lirc.ChannelAction")
	proto.RegisterEnum("lirc.ActionType", ActionType_name, ActionType_value)
}

func init() { proto.RegisterFile("irc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x71, 0xe3, 0x50, 0x72, 0x11, 0xc8, 0xf2, 0xe4, 0x81, 0xa1, 0x54, 0x0c, 0x15, 0x43,
	0x86, 0x32, 0x30, 0x47, 0x51, 0x85, 0x02, 0xa2, 0x88, 0xb4, 0x3c, 0x80, 0x89, 0x2c, 0x6a, 0xa9,
	0xb1, 0x2d, 0xdb, 0x0c, 0x7d, 0x1f, 0x1e, 0x14, 0x9d, 0x23, 0xda, 0x6e, 0x9f, 0xbf, 0xf3, 0xfd,
	0xfe, 0x65, 0x28, 0xb4, 0xef, 0x2b, 0xe7, 0x6d, 0xb4, 0x9c, 0xee, 0xb5, 0xef, 0xe7, 0x0d, 0xe4,
	0x75, 0x1f, 0xad, 0xe7, 0x1c, 0xa8, 0x91, 0x83, 0x12, 0x64, 0x46, 0x16, 0x45, 0x97, 0x18, 0xdd,
	0x4f, 0x50, 0x5e, 0x4c, 0x46, 0x87, 0x8c, 0x6e, 0x67, 0x43, 0x14, 0xd9, 0xe8, 0x90, 0xe7, 0xbf,
	0x04, 0xae, 0x9b, 0x9d, 0x34, 0x46, 0xed, 0xeb, 0x3e, 0x6a, 0x6b, 0xf8, 0x2d, 0x14, 0x51, 0x0f,
	0x2a, 0x44, 0x39, 0xb8, 0x14, 0x99, 0x75, 0x27, 0xc1, 0xef, 0x20, 0x97, 0xf8, 0x68, 0x0a, 0x2e,
	0x97, 0x65, 0x85, 0x55, 0xaa, 0xd4, 0xa3, 0x1b, 0x27, 0xc7, 0x3a, 0xd9, 0x59, 0x9d, 0x7b, 0xa0,
	0xf1, 0xe0, 0x94, 0xa0, 0x33, 0xb2, 0xb8, 0x59, 0xb2, 0xe3, 0x96, 0xb6, 0x66, 0x7b, 0x70, 0xaa,
	0x4b, 0x53, 0x2e, 0x60, 0x3a, 0xa8, 0x10, 0xe4, 0xb7, 0x12, 0x79, 0x5a, 0xfe, 0x3f, 0x3e, 0x3c,
	0x01, 0x9c, 0x6e, 0xf3, 0x12, 0xa6, 0x6f, 0xab, 0xcd, 0xa6, 0x7e, 0x5e, 0xb1, 0x0b, 0x7e, 0x05,
	0xf4, 0xe5, 0xbd, 0x5d, 0x33, 0x82, 0xf4, 0xf1, 0xd9, 0x6e, 0xd9, 0x04, 0x69, 0xdd, 0x36, 0xaf,
	0x2c, 0xfb, 0xba, 0x4c, 0x3f, 0xf6, 0xf8, 0x17, 0x00, 0x00, 0xff, 0xff, 0xce, 0xb8, 0x0a, 0xbb,
	0x3e, 0x01, 0x00, 0x00,
}