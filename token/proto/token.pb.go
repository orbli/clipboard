// Code generated by protoc-gen-go. DO NOT EDIT.
// source: token/proto/token.proto

package orbli_micro_token

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Token struct {
	Token                []byte               `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Secret               []byte               `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	Parent               string               `protobuf:"bytes,3,opt,name=parent,proto3" json:"parent,omitempty"`
	Data                 []byte               `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	ExpireAt             *timestamp.Timestamp `protobuf:"bytes,8,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_713d5ae1cbb77fde, []int{0}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *Token) GetSecret() []byte {
	if m != nil {
		return m.Secret
	}
	return nil
}

func (m *Token) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *Token) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Token) GetExpireAt() *timestamp.Timestamp {
	if m != nil {
		return m.ExpireAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Token)(nil), "orbli.micro.token.Token")
}

func init() { proto.RegisterFile("token/proto/token.proto", fileDescriptor_713d5ae1cbb77fde) }

var fileDescriptor_713d5ae1cbb77fde = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x50, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x25, 0x25, 0x8d, 0xda, 0xa3, 0x0b, 0x56, 0x05, 0x56, 0x16, 0xa2, 0x4e, 0x99, 0x1c, 0xa9,
	0x0c, 0x2c, 0x28, 0x12, 0x82, 0x91, 0x01, 0x85, 0x32, 0x23, 0x27, 0x39, 0x2a, 0x8b, 0xa4, 0xb6,
	0x9c, 0x03, 0xf1, 0x2d, 0x7c, 0x1a, 0x5f, 0x83, 0x7a, 0x6e, 0x27, 0xc4, 0x92, 0xed, 0xbd, 0x7b,
	0xef, 0xd9, 0xef, 0x0e, 0x2e, 0xc9, 0xbe, 0xe3, 0xae, 0x70, 0xde, 0x92, 0x2d, 0x18, 0x2b, 0xc6,
	0xe2, 0xdc, 0xfa, 0xba, 0x33, 0xaa, 0x37, 0x8d, 0xb7, 0x8a, 0x85, 0xf4, 0x6a, 0x6b, 0xed, 0xb6,
	0xc3, 0x60, 0xae, 0x3f, 0xde, 0x0a, 0x32, 0x3d, 0x0e, 0xa4, 0x7b, 0x17, 0x32, 0xab, 0xef, 0x08,
	0xa6, 0x9b, 0xbd, 0x55, 0x2c, 0x61, 0xca, 0x19, 0x19, 0x65, 0x51, 0xbe, 0xa8, 0x02, 0x11, 0x17,
	0x90, 0x0c, 0xd8, 0x78, 0x24, 0x39, 0xe1, 0xf1, 0x81, 0xed, 0xe7, 0x4e, 0x7b, 0xdc, 0x91, 0x3c,
	0xcd, 0xa2, 0x7c, 0x5e, 0x1d, 0x98, 0x10, 0x10, 0xb7, 0x9a, 0xb4, 0x8c, 0xd9, 0xcd, 0x58, 0xdc,
	0xc0, 0x1c, 0xbf, 0x9c, 0xf1, 0xf8, 0xaa, 0x49, 0xce, 0xb2, 0x28, 0x3f, 0x5b, 0xa7, 0x2a, 0x14,
	0x53, 0xc7, 0x62, 0x6a, 0x73, 0x2c, 0x56, 0xcd, 0x82, 0xf9, 0x8e, 0xd6, 0x3f, 0x13, 0x58, 0x70,
	0xb9, 0x67, 0xf4, 0x9f, 0xa6, 0x41, 0x51, 0x42, 0x72, 0xef, 0x51, 0x13, 0x0a, 0xa9, 0xfe, 0x2c,
	0xab, 0xd8, 0x9a, 0xfe, 0xab, 0xac, 0x4e, 0xc4, 0x2d, 0xc4, 0x15, 0xea, 0x76, 0x64, 0xba, 0x84,
	0xe4, 0xc5, 0xb5, 0xe3, 0x7f, 0x2f, 0x21, 0x79, 0xc0, 0x0e, 0x47, 0xe7, 0x1f, 0x61, 0x19, 0xf2,
	0x4f, 0x7c, 0x6b, 0x6c, 0x59, 0x18, 0xc6, 0xbd, 0x56, 0x27, 0x7c, 0xfa, 0xeb, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x02, 0x2c, 0xbd, 0xa1, 0x4f, 0x02, 0x00, 0x00,
}
