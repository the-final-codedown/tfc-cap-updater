// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/tfc/cap/updater/tfc-cap-updater.proto

package tfc_cap_updater

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

type CapDownscale struct {
	AccountID            int64    `protobuf:"varint,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	Delta                int32    `protobuf:"varint,2,opt,name=delta,proto3" json:"delta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CapDownscale) Reset()         { *m = CapDownscale{} }
func (m *CapDownscale) String() string { return proto.CompactTextString(m) }
func (*CapDownscale) ProtoMessage()    {}
func (*CapDownscale) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4143ba0c3384ac8, []int{0}
}

func (m *CapDownscale) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CapDownscale.Unmarshal(m, b)
}
func (m *CapDownscale) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CapDownscale.Marshal(b, m, deterministic)
}
func (m *CapDownscale) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CapDownscale.Merge(m, src)
}
func (m *CapDownscale) XXX_Size() int {
	return xxx_messageInfo_CapDownscale.Size(m)
}
func (m *CapDownscale) XXX_DiscardUnknown() {
	xxx_messageInfo_CapDownscale.DiscardUnknown(m)
}

var xxx_messageInfo_CapDownscale proto.InternalMessageInfo

func (m *CapDownscale) GetAccountID() int64 {
	if m != nil {
		return m.AccountID
	}
	return 0
}

func (m *CapDownscale) GetDelta() int32 {
	if m != nil {
		return m.Delta
	}
	return 0
}

// Created a blank get request
type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4143ba0c3384ac8, []int{1}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type DownscaleResponse struct {
	Accepted             bool            `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
	Downscale            *CapDownscale   `protobuf:"bytes,2,opt,name=downscale,proto3" json:"downscale,omitempty"`
	Downscales           []*CapDownscale `protobuf:"bytes,3,rep,name=downscales,proto3" json:"downscales,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DownscaleResponse) Reset()         { *m = DownscaleResponse{} }
func (m *DownscaleResponse) String() string { return proto.CompactTextString(m) }
func (*DownscaleResponse) ProtoMessage()    {}
func (*DownscaleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4143ba0c3384ac8, []int{2}
}

func (m *DownscaleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownscaleResponse.Unmarshal(m, b)
}
func (m *DownscaleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownscaleResponse.Marshal(b, m, deterministic)
}
func (m *DownscaleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownscaleResponse.Merge(m, src)
}
func (m *DownscaleResponse) XXX_Size() int {
	return xxx_messageInfo_DownscaleResponse.Size(m)
}
func (m *DownscaleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DownscaleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DownscaleResponse proto.InternalMessageInfo

func (m *DownscaleResponse) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

func (m *DownscaleResponse) GetDownscale() *CapDownscale {
	if m != nil {
		return m.Downscale
	}
	return nil
}

func (m *DownscaleResponse) GetDownscales() []*CapDownscale {
	if m != nil {
		return m.Downscales
	}
	return nil
}

func init() {
	proto.RegisterType((*CapDownscale)(nil), "tfc.cap.updater.CapDownscale")
	proto.RegisterType((*GetRequest)(nil), "tfc.cap.updater.GetRequest")
	proto.RegisterType((*DownscaleResponse)(nil), "tfc.cap.updater.DownscaleResponse")
}

func init() {
	proto.RegisterFile("proto/tfc/cap/updater/tfc-cap-updater.proto", fileDescriptor_c4143ba0c3384ac8)
}

var fileDescriptor_c4143ba0c3384ac8 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x8d, 0xa1, 0xd2, 0x8e, 0x01, 0xe9, 0xe2, 0x21, 0x14, 0x85, 0x90, 0x53, 0x41, 0xba,
	0x81, 0x7a, 0x14, 0x2f, 0xa6, 0x20, 0x5e, 0xb7, 0xf8, 0x00, 0xe3, 0x64, 0x8a, 0x42, 0xc9, 0x8e,
	0xd9, 0x89, 0x3e, 0x94, 0x2f, 0x29, 0xa6, 0x35, 0x2d, 0x16, 0xec, 0xf1, 0x9b, 0xfd, 0xff, 0x8f,
	0x59, 0x06, 0x6e, 0xa4, 0xf1, 0xea, 0x0b, 0x5d, 0x51, 0x41, 0x28, 0x45, 0x2b, 0x15, 0x2a, 0x37,
	0x3f, 0x3c, 0x23, 0x94, 0xd9, 0x96, 0x6d, 0x97, 0x32, 0x17, 0xba, 0x22, 0x4b, 0x28, 0x76, 0x3b,
	0xce, 0x1f, 0x20, 0x29, 0x51, 0x16, 0xfe, 0xb3, 0x0e, 0x84, 0x6b, 0x36, 0x57, 0x30, 0x42, 0x22,
	0xdf, 0xd6, 0xfa, 0xb4, 0x48, 0xa3, 0x2c, 0x9a, 0xc6, 0x6e, 0x37, 0x30, 0x97, 0x30, 0xa8, 0x78,
	0xad, 0x98, 0x9e, 0x66, 0xd1, 0x74, 0xe0, 0x36, 0x90, 0x27, 0x00, 0x8f, 0xac, 0x8e, 0xdf, 0x5b,
	0x0e, 0x9a, 0x7f, 0x45, 0x30, 0xee, 0x7d, 0x8e, 0x83, 0xf8, 0x3a, 0xb0, 0x99, 0xc0, 0x10, 0x89,
	0x58, 0x94, 0xab, 0x4e, 0x3b, 0x74, 0x3d, 0x9b, 0x3b, 0x18, 0x55, 0xbf, 0x85, 0xce, 0x7c, 0x3e,
	0xbf, 0xb6, 0x7f, 0x16, 0xb5, 0xfb, 0x5b, 0xba, 0x5d, 0xde, 0xdc, 0x03, 0xf4, 0x10, 0xd2, 0x38,
	0x8b, 0x8f, 0xb7, 0xf7, 0x0a, 0xf3, 0x57, 0x18, 0x97, 0x28, 0xcf, 0x9b, 0xd8, 0x92, 0x9b, 0x8f,
	0x37, 0x62, 0xb3, 0x84, 0xa4, 0x4f, 0x97, 0x28, 0xe6, 0x7f, 0xdf, 0x24, 0x3f, 0x78, 0x3e, 0xf8,
	0x7f, 0x7e, 0xf2, 0x72, 0xd6, 0x5d, 0xe0, 0xf6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x18, 0xa9, 0xe6,
	0x92, 0xb0, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CapUpdaterServiceClient is the client API for CapUpdaterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CapUpdaterServiceClient interface {
	DownscaleCap(ctx context.Context, in *CapDownscale, opts ...grpc.CallOption) (*DownscaleResponse, error)
}

type capUpdaterServiceClient struct {
	cc *grpc.ClientConn
}

func NewCapUpdaterServiceClient(cc *grpc.ClientConn) CapUpdaterServiceClient {
	return &capUpdaterServiceClient{cc}
}

func (c *capUpdaterServiceClient) DownscaleCap(ctx context.Context, in *CapDownscale, opts ...grpc.CallOption) (*DownscaleResponse, error) {
	out := new(DownscaleResponse)
	err := c.cc.Invoke(ctx, "/tfc.cap.updater.CapUpdaterService/DownscaleCap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CapUpdaterServiceServer is the server API for CapUpdaterService service.
type CapUpdaterServiceServer interface {
	DownscaleCap(context.Context, *CapDownscale) (*DownscaleResponse, error)
}

// UnimplementedCapUpdaterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCapUpdaterServiceServer struct {
}

func (*UnimplementedCapUpdaterServiceServer) DownscaleCap(ctx context.Context, req *CapDownscale) (*DownscaleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownscaleCap not implemented")
}

func RegisterCapUpdaterServiceServer(s *grpc.Server, srv CapUpdaterServiceServer) {
	s.RegisterService(&_CapUpdaterService_serviceDesc, srv)
}

func _CapUpdaterService_DownscaleCap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CapDownscale)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CapUpdaterServiceServer).DownscaleCap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfc.cap.updater.CapUpdaterService/DownscaleCap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CapUpdaterServiceServer).DownscaleCap(ctx, req.(*CapDownscale))
	}
	return interceptor(ctx, in, info, handler)
}

var _CapUpdaterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tfc.cap.updater.CapUpdaterService",
	HandlerType: (*CapUpdaterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DownscaleCap",
			Handler:    _CapUpdaterService_DownscaleCap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/tfc/cap/updater/tfc-cap-updater.proto",
}