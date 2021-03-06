// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: src/proto/tfc-cap-updater.proto

package tfc_cap_updater

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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
	AccountID            string   `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	Value                int32    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	Money                int32    `protobuf:"varint,3,opt,name=money,proto3" json:"money,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CapDownscale) Reset()         { *m = CapDownscale{} }
func (m *CapDownscale) String() string { return proto.CompactTextString(m) }
func (*CapDownscale) ProtoMessage()    {}
func (*CapDownscale) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6539d754e1ee526, []int{0}
}
func (m *CapDownscale) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CapDownscale) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CapDownscale.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CapDownscale) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CapDownscale.Merge(m, src)
}
func (m *CapDownscale) XXX_Size() int {
	return m.Size()
}
func (m *CapDownscale) XXX_DiscardUnknown() {
	xxx_messageInfo_CapDownscale.DiscardUnknown(m)
}

var xxx_messageInfo_CapDownscale proto.InternalMessageInfo

func (m *CapDownscale) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *CapDownscale) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *CapDownscale) GetMoney() int32 {
	if m != nil {
		return m.Money
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
	return fileDescriptor_f6539d754e1ee526, []int{1}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return m.Size()
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
	return fileDescriptor_f6539d754e1ee526, []int{2}
}
func (m *DownscaleResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DownscaleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DownscaleResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DownscaleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownscaleResponse.Merge(m, src)
}
func (m *DownscaleResponse) XXX_Size() int {
	return m.Size()
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

func init() { proto.RegisterFile("src/proto/tfc-cap-updater.proto", fileDescriptor_f6539d754e1ee526) }

var fileDescriptor_f6539d754e1ee526 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4a, 0xc3, 0x40,
	0x18, 0x85, 0x3b, 0x86, 0x4a, 0xf3, 0x1b, 0xd0, 0x0e, 0x2e, 0x42, 0xd1, 0x18, 0xb2, 0xea, 0xa6,
	0x29, 0xd4, 0xa5, 0xb8, 0x31, 0x05, 0x71, 0x3b, 0x45, 0x70, 0x3b, 0xfe, 0xf9, 0x8b, 0x42, 0xcd,
	0x8c, 0x99, 0x49, 0xc5, 0x9b, 0x78, 0x07, 0x2f, 0xe2, 0xd2, 0x23, 0x48, 0xbc, 0x88, 0x38, 0xd1,
	0xb4, 0x58, 0xb0, 0xcb, 0xf7, 0xcd, 0x7b, 0xf3, 0xde, 0x30, 0x70, 0x62, 0x4a, 0x1c, 0xeb, 0x52,
	0x59, 0x35, 0xb6, 0x73, 0x1c, 0xa1, 0xd4, 0xa3, 0x4a, 0xe7, 0xd2, 0x52, 0x99, 0x3a, 0xca, 0xf7,
	0xed, 0x1c, 0x53, 0x94, 0x3a, 0xfd, 0xc1, 0xc9, 0x0d, 0x04, 0x99, 0xd4, 0x53, 0xf5, 0x54, 0x18,
	0x94, 0x0b, 0xe2, 0x47, 0xe0, 0x4b, 0x44, 0x55, 0x15, 0xf6, 0x6a, 0x1a, 0xb2, 0x98, 0x0d, 0x7d,
	0xb1, 0x02, 0xfc, 0x10, 0xba, 0x4b, 0xb9, 0xa8, 0x28, 0xdc, 0x89, 0xd9, 0xb0, 0x2b, 0x1a, 0xf1,
	0x4d, 0x1f, 0x54, 0x41, 0xcf, 0xa1, 0xd7, 0x50, 0x27, 0x92, 0x00, 0xe0, 0x92, 0xac, 0xa0, 0xc7,
	0x8a, 0x8c, 0x4d, 0x5e, 0x19, 0xf4, 0xdb, 0x16, 0x41, 0x46, 0xab, 0xc2, 0x10, 0x1f, 0x40, 0x4f,
	0x22, 0x92, 0xb6, 0x94, 0xbb, 0xb2, 0x9e, 0x68, 0x35, 0x3f, 0x03, 0x3f, 0xff, 0x0d, 0xb8, 0xbe,
	0xbd, 0xc9, 0x71, 0xfa, 0x67, 0x7e, 0xba, 0xbe, 0x5d, 0xac, 0xfc, 0xfc, 0x1c, 0xa0, 0x15, 0x26,
	0xf4, 0x62, 0x6f, 0x7b, 0x7a, 0x2d, 0x30, 0xb9, 0x83, 0x7e, 0x26, 0xf5, 0x75, 0x63, 0x9b, 0x51,
	0xb9, 0xbc, 0x47, 0xe2, 0x33, 0x08, 0x5a, 0x77, 0x26, 0x35, 0xff, 0xff, 0xbe, 0x41, 0xb2, 0x71,
	0xbc, 0xf1, 0xfe, 0xa4, 0x73, 0x71, 0xf0, 0x56, 0x47, 0xec, 0xbd, 0x8e, 0xd8, 0x47, 0x1d, 0xb1,
	0x97, 0xcf, 0xa8, 0x73, 0xbb, 0xeb, 0x7e, 0xea, 0xf4, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xb2,
	0x93, 0x10, 0xcc, 0x01, 0x00, 0x00,
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
	Metadata: "src/proto/tfc-cap-updater.proto",
}

func (m *CapDownscale) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CapDownscale) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CapDownscale) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Money != 0 {
		i = encodeVarintTfcCapUpdater(dAtA, i, uint64(m.Money))
		i--
		dAtA[i] = 0x18
	}
	if m.Value != 0 {
		i = encodeVarintTfcCapUpdater(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x10
	}
	if len(m.AccountID) > 0 {
		i -= len(m.AccountID)
		copy(dAtA[i:], m.AccountID)
		i = encodeVarintTfcCapUpdater(dAtA, i, uint64(len(m.AccountID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *DownscaleResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DownscaleResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DownscaleResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Downscales) > 0 {
		for iNdEx := len(m.Downscales) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Downscales[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTfcCapUpdater(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Downscale != nil {
		{
			size, err := m.Downscale.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTfcCapUpdater(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Accepted {
		i--
		if m.Accepted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTfcCapUpdater(dAtA []byte, offset int, v uint64) int {
	offset -= sovTfcCapUpdater(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CapDownscale) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AccountID)
	if l > 0 {
		n += 1 + l + sovTfcCapUpdater(uint64(l))
	}
	if m.Value != 0 {
		n += 1 + sovTfcCapUpdater(uint64(m.Value))
	}
	if m.Money != 0 {
		n += 1 + sovTfcCapUpdater(uint64(m.Money))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DownscaleResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Accepted {
		n += 2
	}
	if m.Downscale != nil {
		l = m.Downscale.Size()
		n += 1 + l + sovTfcCapUpdater(uint64(l))
	}
	if len(m.Downscales) > 0 {
		for _, e := range m.Downscales {
			l = e.Size()
			n += 1 + l + sovTfcCapUpdater(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTfcCapUpdater(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTfcCapUpdater(x uint64) (n int) {
	return sovTfcCapUpdater(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CapDownscale) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTfcCapUpdater
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CapDownscale: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CapDownscale: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTfcCapUpdater
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTfcCapUpdater
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTfcCapUpdater
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTfcCapUpdater
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Money", wireType)
			}
			m.Money = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTfcCapUpdater
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Money |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTfcCapUpdater(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTfcCapUpdater
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTfcCapUpdater
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTfcCapUpdater
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTfcCapUpdater(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTfcCapUpdater
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTfcCapUpdater
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DownscaleResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTfcCapUpdater
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DownscaleResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DownscaleResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accepted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTfcCapUpdater
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Accepted = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Downscale", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTfcCapUpdater
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTfcCapUpdater
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTfcCapUpdater
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Downscale == nil {
				m.Downscale = &CapDownscale{}
			}
			if err := m.Downscale.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Downscales", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTfcCapUpdater
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTfcCapUpdater
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTfcCapUpdater
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Downscales = append(m.Downscales, &CapDownscale{})
			if err := m.Downscales[len(m.Downscales)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTfcCapUpdater(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTfcCapUpdater
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTfcCapUpdater
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTfcCapUpdater(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTfcCapUpdater
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTfcCapUpdater
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTfcCapUpdater
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTfcCapUpdater
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTfcCapUpdater
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTfcCapUpdater
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTfcCapUpdater        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTfcCapUpdater          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTfcCapUpdater = fmt.Errorf("proto: unexpected end of group")
)
