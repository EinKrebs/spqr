// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/operation.proto

package proto

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

type OperationStatus int32

const (
	OperationStatus_PLANNED OperationStatus = 0
	OperationStatus_RUNNING OperationStatus = 1
	OperationStatus_DONE    OperationStatus = 2
)

var OperationStatus_name = map[int32]string{
	0: "PLANNED",
	1: "RUNNING",
	2: "DONE",
}

var OperationStatus_value = map[string]int32{
	"PLANNED": 0,
	"RUNNING": 1,
	"DONE":    2,
}

func (x OperationStatus) String() string {
	return proto.EnumName(OperationStatus_name, int32(x))
}

func (OperationStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3bd593fb4dbc6b9a, []int{0}
}

type Operation struct {
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status               OperationStatus `protobuf:"varint,2,opt,name=status,proto3,enum=spqr.OperationStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Operation) Reset()         { *m = Operation{} }
func (m *Operation) String() string { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()    {}
func (*Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bd593fb4dbc6b9a, []int{0}
}

func (m *Operation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operation.Unmarshal(m, b)
}
func (m *Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operation.Marshal(b, m, deterministic)
}
func (m *Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operation.Merge(m, src)
}
func (m *Operation) XXX_Size() int {
	return xxx_messageInfo_Operation.Size(m)
}
func (m *Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_Operation proto.InternalMessageInfo

func (m *Operation) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Operation) GetStatus() OperationStatus {
	if m != nil {
		return m.Status
	}
	return OperationStatus_PLANNED
}

type GetOperationRequest struct {
	OperationId          string   `protobuf:"bytes,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOperationRequest) Reset()         { *m = GetOperationRequest{} }
func (m *GetOperationRequest) String() string { return proto.CompactTextString(m) }
func (*GetOperationRequest) ProtoMessage()    {}
func (*GetOperationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bd593fb4dbc6b9a, []int{1}
}

func (m *GetOperationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOperationRequest.Unmarshal(m, b)
}
func (m *GetOperationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOperationRequest.Marshal(b, m, deterministic)
}
func (m *GetOperationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOperationRequest.Merge(m, src)
}
func (m *GetOperationRequest) XXX_Size() int {
	return xxx_messageInfo_GetOperationRequest.Size(m)
}
func (m *GetOperationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOperationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOperationRequest proto.InternalMessageInfo

func (m *GetOperationRequest) GetOperationId() string {
	if m != nil {
		return m.OperationId
	}
	return ""
}

type GetOperationReply struct {
	Operation            *Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetOperationReply) Reset()         { *m = GetOperationReply{} }
func (m *GetOperationReply) String() string { return proto.CompactTextString(m) }
func (*GetOperationReply) ProtoMessage()    {}
func (*GetOperationReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bd593fb4dbc6b9a, []int{2}
}

func (m *GetOperationReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOperationReply.Unmarshal(m, b)
}
func (m *GetOperationReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOperationReply.Marshal(b, m, deterministic)
}
func (m *GetOperationReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOperationReply.Merge(m, src)
}
func (m *GetOperationReply) XXX_Size() int {
	return xxx_messageInfo_GetOperationReply.Size(m)
}
func (m *GetOperationReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOperationReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetOperationReply proto.InternalMessageInfo

func (m *GetOperationReply) GetOperation() *Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func init() {
	proto.RegisterEnum("spqr.OperationStatus", OperationStatus_name, OperationStatus_value)
	proto.RegisterType((*Operation)(nil), "spqr.Operation")
	proto.RegisterType((*GetOperationRequest)(nil), "spqr.GetOperationRequest")
	proto.RegisterType((*GetOperationReply)(nil), "spqr.GetOperationReply")
}

func init() { proto.RegisterFile("protos/operation.proto", fileDescriptor_3bd593fb4dbc6b9a) }

var fileDescriptor_3bd593fb4dbc6b9a = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0xcf, 0x2f, 0x48, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x03, 0x0b, 0x08,
	0xb1, 0x14, 0x17, 0x14, 0x16, 0x29, 0x79, 0x71, 0x71, 0xfa, 0xc3, 0x24, 0x84, 0xf8, 0xb8, 0x98,
	0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x98, 0x32, 0x53, 0x84, 0x74, 0xb9, 0xd8,
	0x8a, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0x25, 0x98, 0x14, 0x18, 0x35, 0xf8, 0x8c, 0x44, 0xf5, 0x40,
	0x7a, 0xf4, 0xe0, 0x1a, 0x82, 0xc1, 0x92, 0x41, 0x50, 0x45, 0x4a, 0x16, 0x5c, 0xc2, 0xee, 0xa9,
	0x25, 0x70, 0xd9, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x45, 0x2e, 0x1e, 0xb8, 0xdd,
	0xf1, 0x70, 0xf3, 0xb9, 0xe1, 0x62, 0x9e, 0x29, 0x4a, 0x4e, 0x5c, 0x82, 0xa8, 0x3a, 0x0b, 0x72,
	0x2a, 0x85, 0x74, 0xb9, 0x38, 0xe1, 0x6a, 0xc0, 0x9a, 0xb8, 0x8d, 0xf8, 0xd1, 0x1c, 0x10, 0x84,
	0x50, 0xa1, 0x65, 0xca, 0xc5, 0x8f, 0xe6, 0x30, 0x21, 0x6e, 0x2e, 0xf6, 0x00, 0x1f, 0x47, 0x3f,
	0x3f, 0x57, 0x17, 0x01, 0x06, 0x10, 0x27, 0x28, 0xd4, 0xcf, 0xcf, 0xd3, 0xcf, 0x5d, 0x80, 0x51,
	0x88, 0x83, 0x8b, 0xc5, 0xc5, 0xdf, 0xcf, 0x55, 0x80, 0xc9, 0x28, 0x82, 0x4b, 0x00, 0xa1, 0x2d,
	0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0xc8, 0x85, 0x8b, 0x07, 0xd9, 0x39, 0x42, 0x92, 0x10, 0x6b,
	0xb1, 0x78, 0x4e, 0x4a, 0x1c, 0x9b, 0x54, 0x41, 0x4e, 0xa5, 0x12, 0x83, 0x13, 0x4f, 0x14, 0x17,
	0x48, 0x4e, 0x1f, 0x1c, 0xdc, 0x49, 0x6c, 0x60, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xc9,
	0x40, 0x9c, 0x7f, 0x8f, 0x01, 0x00, 0x00,
}
