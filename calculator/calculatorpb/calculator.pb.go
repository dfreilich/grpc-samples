// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculatorpb/calculator.proto

package calculatorpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SumRequest struct {
	Nums                 []int32  `protobuf:"zigzag32,1,rep,packed,name=nums,proto3" json:"nums,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{0}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetNums() []int32 {
	if m != nil {
		return m.Nums
	}
	return nil
}

type SumResponse struct {
	SumResult            int32    `protobuf:"zigzag32,1,opt,name=sum_result,json=sumResult,proto3" json:"sum_result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{1}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetSumResult() int32 {
	if m != nil {
		return m.SumResult
	}
	return 0
}

type PrimeNumberDecompositionRequest struct {
	Num                  int32    `protobuf:"zigzag32,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberDecompositionRequest) Reset()         { *m = PrimeNumberDecompositionRequest{} }
func (m *PrimeNumberDecompositionRequest) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDecompositionRequest) ProtoMessage()    {}
func (*PrimeNumberDecompositionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{2}
}

func (m *PrimeNumberDecompositionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberDecompositionRequest.Unmarshal(m, b)
}
func (m *PrimeNumberDecompositionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberDecompositionRequest.Marshal(b, m, deterministic)
}
func (m *PrimeNumberDecompositionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberDecompositionRequest.Merge(m, src)
}
func (m *PrimeNumberDecompositionRequest) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberDecompositionRequest.Size(m)
}
func (m *PrimeNumberDecompositionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberDecompositionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberDecompositionRequest proto.InternalMessageInfo

func (m *PrimeNumberDecompositionRequest) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type PrimeNumberDecompositionResponse struct {
	PrimeFactor          int32    `protobuf:"zigzag32,1,opt,name=prime_factor,json=primeFactor,proto3" json:"prime_factor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberDecompositionResponse) Reset()         { *m = PrimeNumberDecompositionResponse{} }
func (m *PrimeNumberDecompositionResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDecompositionResponse) ProtoMessage()    {}
func (*PrimeNumberDecompositionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{3}
}

func (m *PrimeNumberDecompositionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberDecompositionResponse.Unmarshal(m, b)
}
func (m *PrimeNumberDecompositionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberDecompositionResponse.Marshal(b, m, deterministic)
}
func (m *PrimeNumberDecompositionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberDecompositionResponse.Merge(m, src)
}
func (m *PrimeNumberDecompositionResponse) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberDecompositionResponse.Size(m)
}
func (m *PrimeNumberDecompositionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberDecompositionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberDecompositionResponse proto.InternalMessageInfo

func (m *PrimeNumberDecompositionResponse) GetPrimeFactor() int32 {
	if m != nil {
		return m.PrimeFactor
	}
	return 0
}

type ComputeAverageRequest struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAverageRequest) Reset()         { *m = ComputeAverageRequest{} }
func (m *ComputeAverageRequest) String() string { return proto.CompactTextString(m) }
func (*ComputeAverageRequest) ProtoMessage()    {}
func (*ComputeAverageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{4}
}

func (m *ComputeAverageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAverageRequest.Unmarshal(m, b)
}
func (m *ComputeAverageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAverageRequest.Marshal(b, m, deterministic)
}
func (m *ComputeAverageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAverageRequest.Merge(m, src)
}
func (m *ComputeAverageRequest) XXX_Size() int {
	return xxx_messageInfo_ComputeAverageRequest.Size(m)
}
func (m *ComputeAverageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAverageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAverageRequest proto.InternalMessageInfo

func (m *ComputeAverageRequest) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type ComputeAverageResponse struct {
	Average              float64  `protobuf:"fixed64,1,opt,name=average,proto3" json:"average,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAverageResponse) Reset()         { *m = ComputeAverageResponse{} }
func (m *ComputeAverageResponse) String() string { return proto.CompactTextString(m) }
func (*ComputeAverageResponse) ProtoMessage()    {}
func (*ComputeAverageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{5}
}

func (m *ComputeAverageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAverageResponse.Unmarshal(m, b)
}
func (m *ComputeAverageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAverageResponse.Marshal(b, m, deterministic)
}
func (m *ComputeAverageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAverageResponse.Merge(m, src)
}
func (m *ComputeAverageResponse) XXX_Size() int {
	return xxx_messageInfo_ComputeAverageResponse.Size(m)
}
func (m *ComputeAverageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAverageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAverageResponse proto.InternalMessageInfo

func (m *ComputeAverageResponse) GetAverage() float64 {
	if m != nil {
		return m.Average
	}
	return 0
}

type FindMaximumRequest struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaximumRequest) Reset()         { *m = FindMaximumRequest{} }
func (m *FindMaximumRequest) String() string { return proto.CompactTextString(m) }
func (*FindMaximumRequest) ProtoMessage()    {}
func (*FindMaximumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{6}
}

func (m *FindMaximumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaximumRequest.Unmarshal(m, b)
}
func (m *FindMaximumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaximumRequest.Marshal(b, m, deterministic)
}
func (m *FindMaximumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaximumRequest.Merge(m, src)
}
func (m *FindMaximumRequest) XXX_Size() int {
	return xxx_messageInfo_FindMaximumRequest.Size(m)
}
func (m *FindMaximumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaximumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaximumRequest proto.InternalMessageInfo

func (m *FindMaximumRequest) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type FindMaximumResponse struct {
	CurrentMaximum       int32    `protobuf:"varint,1,opt,name=current_maximum,json=currentMaximum,proto3" json:"current_maximum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaximumResponse) Reset()         { *m = FindMaximumResponse{} }
func (m *FindMaximumResponse) String() string { return proto.CompactTextString(m) }
func (*FindMaximumResponse) ProtoMessage()    {}
func (*FindMaximumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{7}
}

func (m *FindMaximumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaximumResponse.Unmarshal(m, b)
}
func (m *FindMaximumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaximumResponse.Marshal(b, m, deterministic)
}
func (m *FindMaximumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaximumResponse.Merge(m, src)
}
func (m *FindMaximumResponse) XXX_Size() int {
	return xxx_messageInfo_FindMaximumResponse.Size(m)
}
func (m *FindMaximumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaximumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaximumResponse proto.InternalMessageInfo

func (m *FindMaximumResponse) GetCurrentMaximum() int32 {
	if m != nil {
		return m.CurrentMaximum
	}
	return 0
}

type SquareRootRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareRootRequest) Reset()         { *m = SquareRootRequest{} }
func (m *SquareRootRequest) String() string { return proto.CompactTextString(m) }
func (*SquareRootRequest) ProtoMessage()    {}
func (*SquareRootRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{8}
}

func (m *SquareRootRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareRootRequest.Unmarshal(m, b)
}
func (m *SquareRootRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareRootRequest.Marshal(b, m, deterministic)
}
func (m *SquareRootRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareRootRequest.Merge(m, src)
}
func (m *SquareRootRequest) XXX_Size() int {
	return xxx_messageInfo_SquareRootRequest.Size(m)
}
func (m *SquareRootRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareRootRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SquareRootRequest proto.InternalMessageInfo

func (m *SquareRootRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type SquareRootResponse struct {
	NumberRoot           float64  `protobuf:"fixed64,1,opt,name=number_root,json=numberRoot,proto3" json:"number_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareRootResponse) Reset()         { *m = SquareRootResponse{} }
func (m *SquareRootResponse) String() string { return proto.CompactTextString(m) }
func (*SquareRootResponse) ProtoMessage()    {}
func (*SquareRootResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{9}
}

func (m *SquareRootResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareRootResponse.Unmarshal(m, b)
}
func (m *SquareRootResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareRootResponse.Marshal(b, m, deterministic)
}
func (m *SquareRootResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareRootResponse.Merge(m, src)
}
func (m *SquareRootResponse) XXX_Size() int {
	return xxx_messageInfo_SquareRootResponse.Size(m)
}
func (m *SquareRootResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareRootResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SquareRootResponse proto.InternalMessageInfo

func (m *SquareRootResponse) GetNumberRoot() float64 {
	if m != nil {
		return m.NumberRoot
	}
	return 0
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "calculator.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calculator.SumResponse")
	proto.RegisterType((*PrimeNumberDecompositionRequest)(nil), "calculator.PrimeNumberDecompositionRequest")
	proto.RegisterType((*PrimeNumberDecompositionResponse)(nil), "calculator.PrimeNumberDecompositionResponse")
	proto.RegisterType((*ComputeAverageRequest)(nil), "calculator.ComputeAverageRequest")
	proto.RegisterType((*ComputeAverageResponse)(nil), "calculator.ComputeAverageResponse")
	proto.RegisterType((*FindMaximumRequest)(nil), "calculator.FindMaximumRequest")
	proto.RegisterType((*FindMaximumResponse)(nil), "calculator.FindMaximumResponse")
	proto.RegisterType((*SquareRootRequest)(nil), "calculator.SquareRootRequest")
	proto.RegisterType((*SquareRootResponse)(nil), "calculator.SquareRootResponse")
}

func init() {
	proto.RegisterFile("calculator/calculatorpb/calculator.proto", fileDescriptor_7f42938f8c8365cf)
}

var fileDescriptor_7f42938f8c8365cf = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x5d, 0x6b, 0xd4, 0x40,
	0x14, 0x35, 0x44, 0x2b, 0xde, 0x94, 0xd5, 0x5c, 0x71, 0x0d, 0x81, 0x76, 0xd3, 0x79, 0xd0, 0x48,
	0x4b, 0x2d, 0x2d, 0x82, 0x4f, 0x82, 0x56, 0xfb, 0x56, 0x91, 0xc4, 0x27, 0x7d, 0x08, 0xd9, 0x38,
	0x4a, 0x20, 0x93, 0x49, 0xe7, 0xa3, 0xe8, 0x5f, 0xf4, 0x57, 0xc9, 0x4e, 0x26, 0x66, 0xb6, 0xbb,
	0xd1, 0xbe, 0xcd, 0x9c, 0x7b, 0x3e, 0xc8, 0xdc, 0x13, 0x48, 0xab, 0xb2, 0xa9, 0x74, 0x53, 0x2a,
	0x2e, 0x5e, 0x8e, 0xc7, 0x6e, 0xe9, 0x5c, 0x8e, 0x3b, 0xc1, 0x15, 0x47, 0x18, 0x11, 0x92, 0x00,
	0xe4, 0x9a, 0x65, 0xf4, 0x4a, 0x53, 0xa9, 0x10, 0xe1, 0x6e, 0xab, 0x99, 0x8c, 0xbc, 0xc4, 0x4f,
	0xc3, 0xcc, 0x9c, 0xc9, 0x11, 0x04, 0x86, 0x21, 0x3b, 0xde, 0x4a, 0x8a, 0x7b, 0x00, 0x52, 0xb3,
	0x42, 0x50, 0xa9, 0x1b, 0x15, 0x79, 0x89, 0x97, 0x86, 0xd9, 0x03, 0x69, 0x08, 0xba, 0x51, 0xe4,
	0x0c, 0x16, 0x9f, 0x44, 0xcd, 0xe8, 0x47, 0xcd, 0x96, 0x54, 0xbc, 0xa7, 0x15, 0x67, 0x1d, 0x97,
	0xb5, 0xaa, 0x79, 0x3b, 0x84, 0x3c, 0x02, 0xbf, 0xd5, 0xcc, 0x4a, 0x57, 0x47, 0xf2, 0x01, 0x92,
	0x69, 0x91, 0xcd, 0x3d, 0x80, 0xdd, 0x6e, 0xc5, 0x29, 0xbe, 0x97, 0x95, 0xe2, 0xc2, 0xca, 0x03,
	0x83, 0x5d, 0x18, 0x88, 0xbc, 0x80, 0x27, 0xe7, 0x9c, 0x75, 0x5a, 0xd1, 0xb7, 0xd7, 0x54, 0x94,
	0x3f, 0xe8, 0x96, 0xc4, 0x7b, 0x7d, 0xe2, 0x29, 0xcc, 0x6f, 0x52, 0x6d, 0x4e, 0x04, 0xf7, 0xcb,
	0x1e, 0x32, 0x7c, 0x2f, 0x1b, 0xae, 0xe4, 0x19, 0xe0, 0x45, 0xdd, 0x7e, 0xbb, 0x2c, 0x7f, 0xd6,
	0x6c, 0x7c, 0xb2, 0x4d, 0xef, 0x37, 0xf0, 0x78, 0x8d, 0x67, 0x8d, 0x9f, 0xc3, 0xc3, 0x4a, 0x0b,
	0x41, 0x5b, 0x55, 0xb0, 0x7e, 0x64, 0x45, 0x33, 0x0b, 0x5b, 0x01, 0x39, 0x84, 0x30, 0xbf, 0xd2,
	0xa5, 0xa0, 0x19, 0xe7, 0x6a, 0x88, 0x99, 0xc3, 0x4e, 0x6b, 0x5e, 0xc7, 0x8a, 0xec, 0x8d, 0xbc,
	0x02, 0x74, 0xc9, 0x36, 0x6b, 0x01, 0x41, 0x3f, 0x2f, 0x04, 0xe7, 0xca, 0x7e, 0x08, 0xf4, 0xd0,
	0x8a, 0x78, 0xfa, 0xdb, 0x87, 0xf0, 0xfc, 0x6f, 0x0b, 0x72, 0x2a, 0xae, 0xeb, 0x8a, 0xe2, 0x6b,
	0xf0, 0x73, 0xcd, 0x70, 0x7e, 0xec, 0x54, 0x66, 0x6c, 0x47, 0xfc, 0x74, 0x03, 0xef, 0xe3, 0xc8,
	0x1d, 0xfc, 0x05, 0xd1, 0xd4, 0x06, 0xf1, 0xd0, 0x95, 0xfd, 0xa7, 0x1c, 0xf1, 0xd1, 0xed, 0xc8,
	0x43, 0xf0, 0x89, 0x87, 0x5f, 0x61, 0xb6, 0xbe, 0x4a, 0x3c, 0x70, 0x3d, 0xb6, 0x36, 0x22, 0x26,
	0xff, 0xa2, 0x0c, 0xe6, 0xa9, 0x87, 0x9f, 0x21, 0x70, 0x76, 0x89, 0xfb, 0xae, 0x6c, 0xb3, 0x0c,
	0xf1, 0x62, 0x72, 0x3e, 0x7a, 0x9e, 0x78, 0x78, 0x09, 0x30, 0x2e, 0x0d, 0xf7, 0xd6, 0x9e, 0xf5,
	0xe6, 0xe6, 0xe3, 0xfd, 0xa9, 0xf1, 0x60, 0xf9, 0x6e, 0xf6, 0x65, 0xd7, 0xfd, 0xe1, 0x97, 0x3b,
	0xe6, 0x37, 0x3f, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0x44, 0xc2, 0x51, 0xd5, 0x12, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	// Unary
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompositionClient, error)
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error)
	FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaximumClient, error)
	SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/PrimeNumberDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeNumberDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeNumberDecompositionClient interface {
	Recv() (*PrimeNumberDecompositionResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeNumberDecompositionClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeNumberDecompositionClient) Recv() (*PrimeNumberDecompositionResponse, error) {
	m := new(PrimeNumberDecompositionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[1], "/calculator.CalculatorService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceComputeAverageClient{stream}
	return x, nil
}

type CalculatorService_ComputeAverageClient interface {
	Send(*ComputeAverageRequest) error
	CloseAndRecv() (*ComputeAverageResponse, error)
	grpc.ClientStream
}

type calculatorServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceComputeAverageClient) Send(m *ComputeAverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageClient) CloseAndRecv() (*ComputeAverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputeAverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[2], "/calculator.CalculatorService/FindMaximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceFindMaximumClient{stream}
	return x, nil
}

type CalculatorService_FindMaximumClient interface {
	Send(*FindMaximumRequest) error
	Recv() (*FindMaximumResponse, error)
	grpc.ClientStream
}

type calculatorServiceFindMaximumClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceFindMaximumClient) Send(m *FindMaximumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceFindMaximumClient) Recv() (*FindMaximumResponse, error) {
	m := new(FindMaximumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error) {
	out := new(SquareRootResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/SquareRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	// Unary
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	PrimeNumberDecomposition(*PrimeNumberDecompositionRequest, CalculatorService_PrimeNumberDecompositionServer) error
	ComputeAverage(CalculatorService_ComputeAverageServer) error
	FindMaximum(CalculatorService_FindMaximumServer) error
	SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error)
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Sum(ctx context.Context, req *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedCalculatorServiceServer) PrimeNumberDecomposition(req *PrimeNumberDecompositionRequest, srv CalculatorService_PrimeNumberDecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeNumberDecomposition not implemented")
}
func (*UnimplementedCalculatorServiceServer) ComputeAverage(srv CalculatorService_ComputeAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAverage not implemented")
}
func (*UnimplementedCalculatorServiceServer) FindMaximum(srv CalculatorService_FindMaximumServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMaximum not implemented")
}
func (*UnimplementedCalculatorServiceServer) SquareRoot(ctx context.Context, req *SquareRootRequest) (*SquareRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SquareRoot not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_PrimeNumberDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberDecompositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeNumberDecomposition(m, &calculatorServicePrimeNumberDecompositionServer{stream})
}

type CalculatorService_PrimeNumberDecompositionServer interface {
	Send(*PrimeNumberDecompositionResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeNumberDecompositionServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeNumberDecompositionServer) Send(m *PrimeNumberDecompositionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).ComputeAverage(&calculatorServiceComputeAverageServer{stream})
}

type CalculatorService_ComputeAverageServer interface {
	SendAndClose(*ComputeAverageResponse) error
	Recv() (*ComputeAverageRequest, error)
	grpc.ServerStream
}

type calculatorServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceComputeAverageServer) SendAndClose(m *ComputeAverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageServer) Recv() (*ComputeAverageRequest, error) {
	m := new(ComputeAverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_FindMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).FindMaximum(&calculatorServiceFindMaximumServer{stream})
}

type CalculatorService_FindMaximumServer interface {
	Send(*FindMaximumResponse) error
	Recv() (*FindMaximumRequest, error)
	grpc.ServerStream
}

type calculatorServiceFindMaximumServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceFindMaximumServer) Send(m *FindMaximumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceFindMaximumServer) Recv() (*FindMaximumRequest, error) {
	m := new(FindMaximumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_SquareRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).SquareRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/SquareRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).SquareRoot(ctx, req.(*SquareRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
		{
			MethodName: "SquareRoot",
			Handler:    _CalculatorService_SquareRoot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumberDecomposition",
			Handler:       _CalculatorService_PrimeNumberDecomposition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _CalculatorService_ComputeAverage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMaximum",
			Handler:       _CalculatorService_FindMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calculatorpb/calculator.proto",
}
