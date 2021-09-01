// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cronos/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgConvertVouchers represents a message to convert ibc voucher coins to cronos evm coins.
type MsgConvertVouchers struct {
	Address string                                   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Coins   github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *MsgConvertVouchers) Reset()         { *m = MsgConvertVouchers{} }
func (m *MsgConvertVouchers) String() string { return proto.CompactTextString(m) }
func (*MsgConvertVouchers) ProtoMessage()    {}
func (*MsgConvertVouchers) Descriptor() ([]byte, []int) {
	return fileDescriptor_28e09e4eabb18884, []int{0}
}
func (m *MsgConvertVouchers) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertVouchers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertVouchers.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertVouchers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertVouchers.Merge(m, src)
}
func (m *MsgConvertVouchers) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertVouchers) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertVouchers.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertVouchers proto.InternalMessageInfo

func (m *MsgConvertVouchers) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgConvertVouchers) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

// MsgTransferTokens represents a message to transfer cronos evm coins through ibc.
type MsgTransferTokens struct {
	From  string                                   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To    string                                   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Coins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *MsgTransferTokens) Reset()         { *m = MsgTransferTokens{} }
func (m *MsgTransferTokens) String() string { return proto.CompactTextString(m) }
func (*MsgTransferTokens) ProtoMessage()    {}
func (*MsgTransferTokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_28e09e4eabb18884, []int{1}
}
func (m *MsgTransferTokens) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransferTokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransferTokens.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransferTokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransferTokens.Merge(m, src)
}
func (m *MsgTransferTokens) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransferTokens) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransferTokens.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransferTokens proto.InternalMessageInfo

func (m *MsgTransferTokens) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *MsgTransferTokens) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *MsgTransferTokens) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

// MsgConvertResponse defines the MsgConvert response type.
type MsgConvertResponse struct {
}

func (m *MsgConvertResponse) Reset()         { *m = MsgConvertResponse{} }
func (m *MsgConvertResponse) String() string { return proto.CompactTextString(m) }
func (*MsgConvertResponse) ProtoMessage()    {}
func (*MsgConvertResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_28e09e4eabb18884, []int{2}
}
func (m *MsgConvertResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertResponse.Merge(m, src)
}
func (m *MsgConvertResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgConvertVouchers)(nil), "cryptoorgchain.cronos.cronos.MsgConvertVouchers")
	proto.RegisterType((*MsgTransferTokens)(nil), "cryptoorgchain.cronos.cronos.MsgTransferTokens")
	proto.RegisterType((*MsgConvertResponse)(nil), "cryptoorgchain.cronos.cronos.MsgConvertResponse")
}

func init() { proto.RegisterFile("cronos/tx.proto", fileDescriptor_28e09e4eabb18884) }

var fileDescriptor_28e09e4eabb18884 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0x3d, 0x8e, 0x1a, 0x31,
	0x14, 0x1e, 0x0f, 0xf9, 0x51, 0x1c, 0x09, 0x14, 0x8b, 0x62, 0x82, 0xa2, 0x01, 0x51, 0xd1, 0x60,
	0x03, 0xb9, 0x01, 0x94, 0x11, 0x0d, 0x42, 0x29, 0xd2, 0xcd, 0x0c, 0xc6, 0x8c, 0x10, 0x7e, 0x23,
	0x3f, 0x83, 0xe0, 0x16, 0x49, 0x9b, 0x23, 0xe4, 0x24, 0x94, 0x94, 0xa9, 0x76, 0x57, 0x70, 0x82,
	0xbd, 0xc1, 0x0a, 0xcf, 0xa0, 0x65, 0x59, 0x69, 0x7f, 0x8a, 0xad, 0xde, 0xb3, 0x9f, 0xbf, 0x1f,
	0x7d, 0x7e, 0xb4, 0x92, 0x18, 0xd0, 0x80, 0xc2, 0xae, 0x79, 0x66, 0xc0, 0x02, 0xfb, 0x96, 0x98,
	0x4d, 0x66, 0x01, 0x8c, 0x4a, 0x66, 0x51, 0xaa, 0x79, 0x3e, 0x2f, 0x4a, 0xad, 0xaa, 0x40, 0x81,
	0x7b, 0x28, 0x8e, 0x5d, 0x8e, 0xa9, 0x85, 0x09, 0xe0, 0x02, 0x50, 0xc4, 0x11, 0x4a, 0xb1, 0xea,
	0xc6, 0xd2, 0x46, 0x5d, 0x91, 0x40, 0xaa, 0xf3, 0x79, 0xf3, 0x0f, 0xa1, 0x6c, 0x88, 0x6a, 0x00,
	0x7a, 0x25, 0x8d, 0xfd, 0x09, 0xcb, 0x64, 0x26, 0x0d, 0xb2, 0x80, 0x7e, 0x8c, 0x26, 0x13, 0x23,
	0x11, 0x03, 0xd2, 0x20, 0xad, 0x4f, 0xa3, 0xd3, 0x91, 0x45, 0xf4, 0xfd, 0x11, 0x8e, 0x81, 0xdf,
	0x28, 0xb5, 0x3e, 0xf7, 0xbe, 0xf2, 0x5c, 0x80, 0x1f, 0x05, 0x78, 0x21, 0xc0, 0x07, 0x90, 0xea,
	0x7e, 0x67, 0x7b, 0x55, 0xf7, 0xfe, 0x5d, 0xd7, 0x5b, 0x2a, 0xb5, 0xb3, 0x65, 0xcc, 0x13, 0x58,
	0x88, 0xc2, 0x4d, 0x5e, 0xda, 0x38, 0x99, 0x0b, 0xbb, 0xc9, 0x24, 0x3a, 0x00, 0x8e, 0x72, 0xe6,
	0xe6, 0x5f, 0x42, 0xbf, 0x0c, 0x51, 0x8d, 0x4d, 0xa4, 0x71, 0x2a, 0xcd, 0x18, 0xe6, 0x52, 0x23,
	0x63, 0xf4, 0xdd, 0xd4, 0xc0, 0xa2, 0xf0, 0xe3, 0x7a, 0x56, 0xa6, 0xbe, 0x85, 0xc0, 0x77, 0x37,
	0xbe, 0x85, 0x7b, 0x73, 0xa5, 0x37, 0x33, 0x57, 0x3d, 0xcf, 0x6b, 0x24, 0x31, 0x03, 0x8d, 0xb2,
	0x77, 0x4b, 0x68, 0x69, 0x88, 0x8a, 0x2d, 0x69, 0xe5, 0x32, 0xca, 0x0e, 0x7f, 0xea, 0xdb, 0xf8,
	0xe3, 0xf0, 0x6b, 0x2f, 0x46, 0x9c, 0xe4, 0x19, 0xd2, 0xf2, 0x45, 0x5a, 0xe2, 0x59, 0x8e, 0x87,
	0x80, 0xd7, 0x8b, 0xf6, 0x7f, 0x6c, 0xf7, 0x21, 0xd9, 0xed, 0x43, 0x72, 0xb3, 0x0f, 0xc9, 0xef,
	0x43, 0xe8, 0xed, 0x0e, 0xa1, 0xf7, 0xff, 0x10, 0x7a, 0xbf, 0xba, 0xe7, 0xa1, 0x3a, 0xd6, 0x36,
	0x18, 0xd5, 0x76, 0xbc, 0xa2, 0xd8, 0xea, 0xf5, 0xa9, 0x71, 0x19, 0xc7, 0x1f, 0xdc, 0x3a, 0x7e,
	0xbf, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x38, 0xcb, 0x9e, 0xf5, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// ConvertVouchers defines a method for converting ibc voucher to cronos evm coins.
	ConvertVouchers(ctx context.Context, in *MsgConvertVouchers, opts ...grpc.CallOption) (*MsgConvertResponse, error)
	// TransferTokens defines a method to transfer cronos evm coins to another chain through IBC
	TransferTokens(ctx context.Context, in *MsgTransferTokens, opts ...grpc.CallOption) (*MsgConvertResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ConvertVouchers(ctx context.Context, in *MsgConvertVouchers, opts ...grpc.CallOption) (*MsgConvertResponse, error) {
	out := new(MsgConvertResponse)
	err := c.cc.Invoke(ctx, "/cryptoorgchain.cronos.cronos.Msg/ConvertVouchers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TransferTokens(ctx context.Context, in *MsgTransferTokens, opts ...grpc.CallOption) (*MsgConvertResponse, error) {
	out := new(MsgConvertResponse)
	err := c.cc.Invoke(ctx, "/cryptoorgchain.cronos.cronos.Msg/TransferTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// ConvertVouchers defines a method for converting ibc voucher to cronos evm coins.
	ConvertVouchers(context.Context, *MsgConvertVouchers) (*MsgConvertResponse, error)
	// TransferTokens defines a method to transfer cronos evm coins to another chain through IBC
	TransferTokens(context.Context, *MsgTransferTokens) (*MsgConvertResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) ConvertVouchers(ctx context.Context, req *MsgConvertVouchers) (*MsgConvertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertVouchers not implemented")
}
func (*UnimplementedMsgServer) TransferTokens(ctx context.Context, req *MsgTransferTokens) (*MsgConvertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferTokens not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_ConvertVouchers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConvertVouchers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConvertVouchers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cryptoorgchain.cronos.cronos.Msg/ConvertVouchers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConvertVouchers(ctx, req.(*MsgConvertVouchers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_TransferTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransferTokens)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TransferTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cryptoorgchain.cronos.cronos.Msg/TransferTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TransferTokens(ctx, req.(*MsgTransferTokens))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cryptoorgchain.cronos.cronos.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConvertVouchers",
			Handler:    _Msg_ConvertVouchers_Handler,
		},
		{
			MethodName: "TransferTokens",
			Handler:    _Msg_TransferTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cronos/tx.proto",
}

func (m *MsgConvertVouchers) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertVouchers) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertVouchers) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgTransferTokens) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransferTokens) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransferTokens) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintTx(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgConvertResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgConvertVouchers) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgTransferTokens) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgConvertResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgConvertVouchers) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgConvertVouchers: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertVouchers: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgTransferTokens) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgTransferTokens: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransferTokens: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgConvertResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgConvertResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
