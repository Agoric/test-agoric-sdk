// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: agoric/vibc/msgs.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
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

// MsgSendPacket is an SDK message for sending an outgoing IBC packet
type MsgSendPacket struct {
	Packet types.Packet                                  `protobuf:"bytes,1,opt,name=packet,proto3" json:"packet" yaml:"packet"`
	Sender github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"submitter" yaml:"submitter"`
}

func (m *MsgSendPacket) Reset()         { *m = MsgSendPacket{} }
func (m *MsgSendPacket) String() string { return proto.CompactTextString(m) }
func (*MsgSendPacket) ProtoMessage()    {}
func (*MsgSendPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_78e9bb7be62a4c00, []int{0}
}
func (m *MsgSendPacket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendPacket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendPacket.Merge(m, src)
}
func (m *MsgSendPacket) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendPacket.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendPacket proto.InternalMessageInfo

func (m *MsgSendPacket) GetPacket() types.Packet {
	if m != nil {
		return m.Packet
	}
	return types.Packet{}
}

func (m *MsgSendPacket) GetSender() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Sender
	}
	return nil
}

// Empty response for SendPacket.
type MsgSendPacketResponse struct {
}

func (m *MsgSendPacketResponse) Reset()         { *m = MsgSendPacketResponse{} }
func (m *MsgSendPacketResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSendPacketResponse) ProtoMessage()    {}
func (*MsgSendPacketResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_78e9bb7be62a4c00, []int{1}
}
func (m *MsgSendPacketResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendPacketResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendPacketResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendPacketResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendPacketResponse.Merge(m, src)
}
func (m *MsgSendPacketResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendPacketResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendPacketResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendPacketResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSendPacket)(nil), "agoric.vibc.MsgSendPacket")
	proto.RegisterType((*MsgSendPacketResponse)(nil), "agoric.vibc.MsgSendPacketResponse")
}

func init() { proto.RegisterFile("agoric/vibc/msgs.proto", fileDescriptor_78e9bb7be62a4c00) }

var fileDescriptor_78e9bb7be62a4c00 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xbf, 0x6e, 0xea, 0x30,
	0x14, 0xc6, 0x93, 0x7b, 0xaf, 0x90, 0xae, 0xb9, 0x48, 0x57, 0x51, 0xff, 0xa0, 0x54, 0x8a, 0x69,
	0x26, 0x16, 0x6c, 0x41, 0x87, 0x4a, 0x6c, 0x30, 0x17, 0xa9, 0x0d, 0x9d, 0xba, 0x25, 0x8e, 0x65,
	0x52, 0x48, 0x1c, 0xe5, 0x04, 0x54, 0xde, 0xa2, 0x8f, 0xd0, 0xc7, 0x61, 0x64, 0x6b, 0xa7, 0xa8,
	0x82, 0xa5, 0x62, 0x64, 0xec, 0x54, 0x25, 0x4e, 0x2a, 0x18, 0x3a, 0xf9, 0xf8, 0x3b, 0xe7, 0x3b,
	0xd6, 0xcf, 0x1f, 0x3a, 0x73, 0x85, 0x4c, 0x02, 0x46, 0x17, 0x81, 0xc7, 0x68, 0x08, 0x02, 0x48,
	0x9c, 0xc8, 0x54, 0x1a, 0x75, 0xa5, 0x93, 0x5c, 0x37, 0x4f, 0x84, 0x14, 0xb2, 0xd0, 0x69, 0x5e,
	0xa9, 0x11, 0xf3, 0x32, 0xb7, 0x30, 0x99, 0x70, 0xca, 0x26, 0x6e, 0x14, 0xf1, 0x19, 0x5d, 0x74,
	0xab, 0x52, 0x8d, 0xd8, 0xaf, 0x3a, 0x6a, 0x8c, 0x40, 0x8c, 0x79, 0xe4, 0xdf, 0xba, 0x6c, 0xca,
	0x53, 0xe3, 0x1e, 0xd5, 0xe2, 0xa2, 0x6a, 0xea, 0x2d, 0xbd, 0x5d, 0xef, 0x5d, 0x90, 0xc0, 0x63,
	0x24, 0xdf, 0x42, 0x2a, 0xeb, 0xa2, 0x4b, 0xd4, 0xf0, 0x10, 0xaf, 0x32, 0xac, 0xed, 0x32, 0x5c,
	0x5a, 0xf6, 0x19, 0x6e, 0x2c, 0xdd, 0x70, 0xd6, 0xb7, 0xd5, 0xdd, 0x76, 0xca, 0x86, 0xf1, 0x88,
	0x6a, 0xc0, 0x23, 0x9f, 0x27, 0xcd, 0x5f, 0x2d, 0xbd, 0xfd, 0x6f, 0xe8, 0xec, 0x32, 0xfc, 0x17,
	0xe6, 0x5e, 0x18, 0xa4, 0x29, 0x4f, 0xf6, 0x19, 0xfe, 0xaf, 0x7c, 0xdf, 0x92, 0xfd, 0x99, 0xe1,
	0x8e, 0x08, 0xd2, 0xc9, 0xdc, 0x23, 0x4c, 0x86, 0x94, 0x49, 0x08, 0x25, 0x94, 0x47, 0x07, 0xfc,
	0x29, 0x4d, 0x97, 0x31, 0x07, 0x32, 0x60, 0x6c, 0xe0, 0xfb, 0x09, 0x07, 0x70, 0xca, 0x17, 0xfa,
	0x7f, 0x3e, 0x5e, 0xb0, 0x66, 0x9f, 0xa3, 0xd3, 0x23, 0x30, 0x87, 0x43, 0x2c, 0x23, 0xe0, 0xbd,
	0x31, 0xfa, 0x3d, 0x02, 0x61, 0xdc, 0x20, 0x74, 0x40, 0x6d, 0x92, 0x83, 0xef, 0x24, 0x47, 0x46,
	0xd3, 0xfe, 0xb9, 0x57, 0x2d, 0x1d, 0xde, 0xad, 0x36, 0x96, 0xbe, 0xde, 0x58, 0xfa, 0xfb, 0xc6,
	0xd2, 0x9f, 0xb7, 0x96, 0xb6, 0xde, 0x5a, 0xda, 0xdb, 0xd6, 0xd2, 0x1e, 0xae, 0x0f, 0x20, 0x06,
	0x2a, 0x4a, 0xb5, 0xae, 0x80, 0x10, 0x72, 0xe6, 0x46, 0xa2, 0xa2, 0x7b, 0x52, 0x29, 0x17, 0x64,
	0x5e, 0xad, 0x48, 0xe8, 0xea, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x45, 0x01, 0xfd, 0x45, 0x01, 0x02,
	0x00, 0x00,
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
	// Force sending an arbitrary packet on a channel.
	SendPacket(ctx context.Context, in *MsgSendPacket, opts ...grpc.CallOption) (*MsgSendPacketResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SendPacket(ctx context.Context, in *MsgSendPacket, opts ...grpc.CallOption) (*MsgSendPacketResponse, error) {
	out := new(MsgSendPacketResponse)
	err := c.cc.Invoke(ctx, "/agoric.vibc.Msg/SendPacket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Force sending an arbitrary packet on a channel.
	SendPacket(context.Context, *MsgSendPacket) (*MsgSendPacketResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SendPacket(ctx context.Context, req *MsgSendPacket) (*MsgSendPacketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPacket not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SendPacket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendPacket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendPacket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agoric.vibc.Msg/SendPacket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendPacket(ctx, req.(*MsgSendPacket))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "agoric.vibc.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendPacket",
			Handler:    _Msg_SendPacket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agoric/vibc/msgs.proto",
}

func (m *MsgSendPacket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Packet.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgs(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgSendPacketResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendPacketResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendPacketResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsgs(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSendPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Packet.Size()
	n += 1 + l + sovMsgs(uint64(l))
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	return n
}

func (m *MsgSendPacketResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsgs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgs(x uint64) (n int) {
	return sovMsgs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSendPacket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgs
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
			return fmt.Errorf("proto: MsgSendPacket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendPacket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Packet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Packet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgs
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
func (m *MsgSendPacketResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgs
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
			return fmt.Errorf("proto: MsgSendPacketResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendPacketResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgs
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
func skipMsgs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgs
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
					return 0, ErrIntOverflowMsgs
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
					return 0, ErrIntOverflowMsgs
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
				return 0, ErrInvalidLengthMsgs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgs = fmt.Errorf("proto: unexpected end of group")
)
