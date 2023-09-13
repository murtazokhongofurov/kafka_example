// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat_service.proto

package chat_service

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

type ChatReq struct {
	ChatType             string   `protobuf:"bytes,1,opt,name=chat_type,json=chatType,proto3" json:"chat_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatReq) Reset()         { *m = ChatReq{} }
func (m *ChatReq) String() string { return proto.CompactTextString(m) }
func (*ChatReq) ProtoMessage()    {}
func (*ChatReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e881f93b317ee75c, []int{0}
}

func (m *ChatReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatReq.Unmarshal(m, b)
}
func (m *ChatReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatReq.Marshal(b, m, deterministic)
}
func (m *ChatReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatReq.Merge(m, src)
}
func (m *ChatReq) XXX_Size() int {
	return xxx_messageInfo_ChatReq.Size(m)
}
func (m *ChatReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatReq.DiscardUnknown(m)
}

var xxx_messageInfo_ChatReq proto.InternalMessageInfo

func (m *ChatReq) GetChatType() string {
	if m != nil {
		return m.ChatType
	}
	return ""
}

type ChatRes struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ChatType             string   `protobuf:"bytes,2,opt,name=chat_type,json=chatType,proto3" json:"chat_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatRes) Reset()         { *m = ChatRes{} }
func (m *ChatRes) String() string { return proto.CompactTextString(m) }
func (*ChatRes) ProtoMessage()    {}
func (*ChatRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e881f93b317ee75c, []int{1}
}

func (m *ChatRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatRes.Unmarshal(m, b)
}
func (m *ChatRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatRes.Marshal(b, m, deterministic)
}
func (m *ChatRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatRes.Merge(m, src)
}
func (m *ChatRes) XXX_Size() int {
	return xxx_messageInfo_ChatRes.Size(m)
}
func (m *ChatRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatRes.DiscardUnknown(m)
}

var xxx_messageInfo_ChatRes proto.InternalMessageInfo

func (m *ChatRes) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ChatRes) GetChatType() string {
	if m != nil {
		return m.ChatType
	}
	return ""
}

func init() {
	proto.RegisterType((*ChatReq)(nil), "chat_service.ChatReq")
	proto.RegisterType((*ChatRes)(nil), "chat_service.ChatRes")
}

func init() {
	proto.RegisterFile("chat_service.proto", fileDescriptor_e881f93b317ee75c)
}

var fileDescriptor_e881f93b317ee75c = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xce, 0x48, 0x2c,
	0x89, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x41, 0x16, 0x53, 0x52, 0xe3, 0x62, 0x77, 0xce, 0x48, 0x2c, 0x09, 0x4a, 0x2d, 0x14, 0x92, 0xe6,
	0xe2, 0x04, 0x4b, 0x95, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x71, 0x80,
	0x04, 0x42, 0x2a, 0x0b, 0x52, 0x95, 0xcc, 0x60, 0xea, 0x8a, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53,
	0xc0, 0x0a, 0x98, 0x83, 0x98, 0x32, 0x53, 0x50, 0xf5, 0x31, 0xa1, 0xea, 0x33, 0xf2, 0xe0, 0xe2,
	0x06, 0xe9, 0x0b, 0x86, 0x58, 0x27, 0x64, 0xc9, 0xc5, 0xee, 0x98, 0x92, 0x02, 0x12, 0x11, 0x12,
	0xd5, 0x43, 0x71, 0x1c, 0xd4, 0x15, 0x52, 0x58, 0x85, 0x8b, 0x95, 0x18, 0x9c, 0xc4, 0xa3, 0x44,
	0xd3, 0x53, 0xf3, 0xc0, 0x7e, 0xd0, 0x47, 0x56, 0x92, 0xc4, 0x06, 0x16, 0x33, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x09, 0x16, 0x0c, 0xd7, 0xed, 0x00, 0x00, 0x00,
}
