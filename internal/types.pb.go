// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

// Package internal holds protobuf types used by the server

package internal

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

// Challenge is the authentication challenge to be used by client
type Challenge struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	IssuedAt             int64    `protobuf:"varint,2,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	Expiration           int64    `protobuf:"varint,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
	Principal            string   `protobuf:"bytes,4,opt,name=principal,proto3" json:"principal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Challenge) Reset()         { *m = Challenge{} }
func (m *Challenge) String() string { return proto.CompactTextString(m) }
func (*Challenge) ProtoMessage()    {}
func (*Challenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0}
}

func (m *Challenge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Challenge.Unmarshal(m, b)
}
func (m *Challenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Challenge.Marshal(b, m, deterministic)
}
func (m *Challenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Challenge.Merge(m, src)
}
func (m *Challenge) XXX_Size() int {
	return xxx_messageInfo_Challenge.Size(m)
}
func (m *Challenge) XXX_DiscardUnknown() {
	xxx_messageInfo_Challenge.DiscardUnknown(m)
}

var xxx_messageInfo_Challenge proto.InternalMessageInfo

func (m *Challenge) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *Challenge) GetIssuedAt() int64 {
	if m != nil {
		return m.IssuedAt
	}
	return 0
}

func (m *Challenge) GetExpiration() int64 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

func (m *Challenge) GetPrincipal() string {
	if m != nil {
		return m.Principal
	}
	return ""
}

func init() {
	proto.RegisterType((*Challenge)(nil), "internal.Challenge")
}

func init() {
	proto.RegisterFile("types.proto", fileDescriptor_d938547f84707355)
}

var fileDescriptor_d938547f84707355 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b,
	0xcc, 0x51, 0x6a, 0x65, 0xe4, 0xe2, 0x74, 0xce, 0x48, 0xcc, 0xc9, 0x49, 0xcd, 0x4b, 0x4f, 0x15,
	0x92, 0xe5, 0xe2, 0x2a, 0x4e, 0x2d, 0x2e, 0xce, 0xcc, 0xcf, 0x8b, 0xcf, 0x4c, 0x91, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x84, 0x8a, 0x78, 0xa6, 0x08, 0x49, 0x73, 0x71, 0x66, 0x16, 0x17,
	0x97, 0xa6, 0xa6, 0xc4, 0x27, 0x96, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x30, 0x07, 0x71, 0x40, 0x04,
	0x1c, 0x4b, 0x84, 0xe4, 0xb8, 0xb8, 0x52, 0x2b, 0x0a, 0x32, 0x8b, 0x12, 0x4b, 0x32, 0xf3, 0xf3,
	0x24, 0x98, 0xc1, 0xb2, 0x48, 0x22, 0x42, 0x32, 0x5c, 0x9c, 0x05, 0x45, 0x99, 0x79, 0xc9, 0x99,
	0x05, 0x89, 0x39, 0x12, 0x2c, 0x10, 0xa3, 0xe1, 0x02, 0x49, 0x6c, 0x60, 0x87, 0x19, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x9e, 0xc4, 0xb0, 0x49, 0xa7, 0x00, 0x00, 0x00,
}
