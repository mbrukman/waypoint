// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testproto.proto

package testproto

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

// Data is just some data, used for tests so meant to be meaningless.
type Data struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Number               int32    `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_testproto_a6708ad74d5692de, []int{0}
}
func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (dst *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(dst, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Data) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

// Other types that we can use. There isn't any purpose for this other
// than to provide message types that can be used for tests.
type A struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *A) Reset()         { *m = A{} }
func (m *A) String() string { return proto.CompactTextString(m) }
func (*A) ProtoMessage()    {}
func (*A) Descriptor() ([]byte, []int) {
	return fileDescriptor_testproto_a6708ad74d5692de, []int{1}
}
func (m *A) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_A.Unmarshal(m, b)
}
func (m *A) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_A.Marshal(b, m, deterministic)
}
func (dst *A) XXX_Merge(src proto.Message) {
	xxx_messageInfo_A.Merge(dst, src)
}
func (m *A) XXX_Size() int {
	return xxx_messageInfo_A.Size(m)
}
func (m *A) XXX_DiscardUnknown() {
	xxx_messageInfo_A.DiscardUnknown(m)
}

var xxx_messageInfo_A proto.InternalMessageInfo

func (m *A) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type B struct {
	Value                int32    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *B) Reset()         { *m = B{} }
func (m *B) String() string { return proto.CompactTextString(m) }
func (*B) ProtoMessage()    {}
func (*B) Descriptor() ([]byte, []int) {
	return fileDescriptor_testproto_a6708ad74d5692de, []int{2}
}
func (m *B) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_B.Unmarshal(m, b)
}
func (m *B) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_B.Marshal(b, m, deterministic)
}
func (dst *B) XXX_Merge(src proto.Message) {
	xxx_messageInfo_B.Merge(dst, src)
}
func (m *B) XXX_Size() int {
	return xxx_messageInfo_B.Size(m)
}
func (m *B) XXX_DiscardUnknown() {
	xxx_messageInfo_B.DiscardUnknown(m)
}

var xxx_messageInfo_B proto.InternalMessageInfo

func (m *B) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*Data)(nil), "testproto.Data")
	proto.RegisterType((*A)(nil), "testproto.A")
	proto.RegisterType((*B)(nil), "testproto.B")
}

func init() { proto.RegisterFile("testproto.proto", fileDescriptor_testproto_a6708ad74d5692de) }

var fileDescriptor_testproto_a6708ad74d5692de = []byte{
	// 109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x49, 0x2d, 0x2e,
	0x29, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x03, 0x93, 0x42, 0x9c, 0x70, 0x01, 0x25, 0x13, 0x2e, 0x16,
	0x97, 0xc4, 0x92, 0x44, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0x48, 0x8c, 0x8b, 0x2d, 0xaf, 0x34, 0x37, 0x29, 0xb5, 0x48,
	0x82, 0x49, 0x81, 0x51, 0x83, 0x35, 0x08, 0xca, 0x53, 0x92, 0xe4, 0x62, 0x74, 0x44, 0xd5, 0xc2,
	0x0a, 0xd5, 0x02, 0x92, 0x72, 0x42, 0x48, 0x31, 0x21, 0x49, 0x25, 0xb1, 0x81, 0xad, 0x34, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x43, 0xa5, 0xe9, 0x90, 0x00, 0x00, 0x00,
}