// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package wire

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type InternalTransaction struct {
	Amount               uint64   `protobuf:"varint,1,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Receiver             string   `protobuf:"bytes,2,opt,name=Receiver,proto3" json:"Receiver,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InternalTransaction) Reset()         { *m = InternalTransaction{} }
func (m *InternalTransaction) String() string { return proto.CompactTextString(m) }
func (*InternalTransaction) ProtoMessage()    {}
func (*InternalTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}

func (m *InternalTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalTransaction.Unmarshal(m, b)
}
func (m *InternalTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalTransaction.Marshal(b, m, deterministic)
}
func (m *InternalTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalTransaction.Merge(m, src)
}
func (m *InternalTransaction) XXX_Size() int {
	return xxx_messageInfo_InternalTransaction.Size(m)
}
func (m *InternalTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_InternalTransaction proto.InternalMessageInfo

func (m *InternalTransaction) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *InternalTransaction) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

type Event struct {
	Index                uint64                 `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	Creator              string                 `protobuf:"bytes,2,opt,name=Creator,proto3" json:"Creator,omitempty"`
	Parents              [][]byte               `protobuf:"bytes,3,rep,name=Parents,proto3" json:"Parents,omitempty"`
	LamportTime          uint64                 `protobuf:"varint,4,opt,name=LamportTime,proto3" json:"LamportTime,omitempty"`
	InternalTransactions []*InternalTransaction `protobuf:"bytes,5,rep,name=InternalTransactions,proto3" json:"InternalTransactions,omitempty"`
	ExternalTransactions [][]byte               `protobuf:"bytes,6,rep,name=ExternalTransactions,proto3" json:"ExternalTransactions,omitempty"`
	Sign                 string                 `protobuf:"bytes,7,opt,name=Sign,proto3" json:"Sign,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{1}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Event) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Event) GetParents() [][]byte {
	if m != nil {
		return m.Parents
	}
	return nil
}

func (m *Event) GetLamportTime() uint64 {
	if m != nil {
		return m.LamportTime
	}
	return 0
}

func (m *Event) GetInternalTransactions() []*InternalTransaction {
	if m != nil {
		return m.InternalTransactions
	}
	return nil
}

func (m *Event) GetExternalTransactions() [][]byte {
	if m != nil {
		return m.ExternalTransactions
	}
	return nil
}

func (m *Event) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

func init() {
	proto.RegisterType((*InternalTransaction)(nil), "wire.InternalTransaction")
	proto.RegisterType((*Event)(nil), "wire.Event")
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_2d17a9d3f0ddf27e) }

var fileDescriptor_2d17a9d3f0ddf27e = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4d, 0x4e, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0xe6, 0xa7, 0x30, 0x61, 0x35, 0x44, 0xc8, 0xb0, 0xb2, 0xba, 0xca, 0x2a, 0x8b,
	0x72, 0x02, 0x84, 0xba, 0x88, 0x04, 0x12, 0x32, 0xbd, 0x80, 0x29, 0x23, 0x64, 0x89, 0x8c, 0x2b,
	0x67, 0x28, 0x3d, 0x03, 0xa7, 0x46, 0x31, 0x0d, 0xea, 0xc2, 0x3b, 0x7f, 0x7e, 0x7e, 0xcf, 0xf3,
	0x06, 0x6a, 0x3a, 0x10, 0x4b, 0xb7, 0x0f, 0x5e, 0x3c, 0x16, 0xdf, 0x2e, 0xd0, 0xaa, 0x87, 0xeb,
	0x9e, 0x85, 0x02, 0xdb, 0xcf, 0x6d, 0xb0, 0x3c, 0xda, 0x9d, 0x38, 0xcf, 0x78, 0x03, 0xd5, 0xc3,
	0xe0, 0xbf, 0x58, 0x54, 0xa6, 0xb3, 0xb6, 0x30, 0x27, 0xc2, 0x3b, 0xb8, 0x30, 0xb4, 0x23, 0x77,
	0xa0, 0xa0, 0x16, 0x3a, 0x6b, 0x2f, 0xcd, 0x3f, 0xaf, 0x7e, 0x16, 0x50, 0x6e, 0xa6, 0x0f, 0xb0,
	0x81, 0xb2, 0xe7, 0x77, 0x3a, 0x9e, 0xcc, 0x7f, 0x80, 0x0a, 0x96, 0x8f, 0x81, 0xac, 0xf8, 0xd9,
	0x3a, 0xe3, 0xa4, 0xbc, 0xd8, 0x40, 0x2c, 0xa3, 0xca, 0x75, 0xde, 0x5e, 0x99, 0x19, 0x51, 0x43,
	0xfd, 0x64, 0x87, 0xbd, 0x0f, 0xb2, 0x75, 0x03, 0xa9, 0x22, 0xe6, 0x9d, 0x5f, 0xe1, 0x33, 0x34,
	0x89, 0x02, 0xa3, 0x2a, 0x75, 0xde, 0xd6, 0xeb, 0xdb, 0x6e, 0x6a, 0xd9, 0x25, 0x5e, 0x98, 0xa4,
	0x0d, 0xd7, 0xd0, 0x6c, 0x8e, 0x89, 0xb8, 0x2a, 0xce, 0x95, 0xd4, 0x10, 0xa1, 0x78, 0x75, 0x1f,
	0xac, 0x96, 0xb1, 0x55, 0x3c, 0xbf, 0x55, 0x71, 0xc9, 0xf7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x21, 0x27, 0x9f, 0xe5, 0x73, 0x01, 0x00, 0x00,
}
