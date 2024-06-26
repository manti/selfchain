// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: selfchain/selfvesting/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

type MsgRelease struct {
	Creator  string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	PosIndex uint64 `protobuf:"varint,2,opt,name=posIndex,proto3" json:"posIndex,omitempty"`
}

func (m *MsgRelease) Reset()         { *m = MsgRelease{} }
func (m *MsgRelease) String() string { return proto.CompactTextString(m) }
func (*MsgRelease) ProtoMessage()    {}
func (*MsgRelease) Descriptor() ([]byte, []int) {
	return fileDescriptor_70a0f46b1e8b78ab, []int{0}
}
func (m *MsgRelease) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRelease) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRelease.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRelease) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRelease.Merge(m, src)
}
func (m *MsgRelease) XXX_Size() int {
	return m.Size()
}
func (m *MsgRelease) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRelease.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRelease proto.InternalMessageInfo

func (m *MsgRelease) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgRelease) GetPosIndex() uint64 {
	if m != nil {
		return m.PosIndex
	}
	return 0
}

type MsgReleaseResponse struct {
	PeriodToVest uint64 `protobuf:"varint,1,opt,name=periodToVest,proto3" json:"periodToVest,omitempty"`
	AmountToVest string `protobuf:"bytes,2,opt,name=amountToVest,proto3" json:"amountToVest,omitempty"`
}

func (m *MsgReleaseResponse) Reset()         { *m = MsgReleaseResponse{} }
func (m *MsgReleaseResponse) String() string { return proto.CompactTextString(m) }
func (*MsgReleaseResponse) ProtoMessage()    {}
func (*MsgReleaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_70a0f46b1e8b78ab, []int{1}
}
func (m *MsgReleaseResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgReleaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgReleaseResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgReleaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgReleaseResponse.Merge(m, src)
}
func (m *MsgReleaseResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgReleaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgReleaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgReleaseResponse proto.InternalMessageInfo

func (m *MsgReleaseResponse) GetPeriodToVest() uint64 {
	if m != nil {
		return m.PeriodToVest
	}
	return 0
}

func (m *MsgReleaseResponse) GetAmountToVest() string {
	if m != nil {
		return m.AmountToVest
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgRelease)(nil), "selfchain.selfvesting.MsgRelease")
	proto.RegisterType((*MsgReleaseResponse)(nil), "selfchain.selfvesting.MsgReleaseResponse")
}

func init() { proto.RegisterFile("selfchain/selfvesting/tx.proto", fileDescriptor_70a0f46b1e8b78ab) }

var fileDescriptor_70a0f46b1e8b78ab = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x4e, 0xcd, 0x49,
	0x4b, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x07, 0xb1, 0xca, 0x52, 0x8b, 0x4b, 0x32, 0xf3, 0xd2, 0xf5,
	0x4b, 0x2a, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x44, 0xe1, 0xf2, 0x7a, 0x48, 0xf2, 0x4a,
	0x4e, 0x5c, 0x5c, 0xbe, 0xc5, 0xe9, 0x41, 0xa9, 0x39, 0xa9, 0x89, 0xc5, 0xa9, 0x42, 0x12, 0x5c,
	0xec, 0xc9, 0x45, 0xa9, 0x89, 0x25, 0xf9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30,
	0xae, 0x90, 0x14, 0x17, 0x47, 0x41, 0x7e, 0xb1, 0x67, 0x5e, 0x4a, 0x6a, 0x85, 0x04, 0x93, 0x02,
	0xa3, 0x06, 0x4b, 0x10, 0x9c, 0xaf, 0x14, 0xc3, 0x25, 0x84, 0x30, 0x23, 0x28, 0xb5, 0xb8, 0x20,
	0x3f, 0xaf, 0x38, 0x55, 0x48, 0x89, 0x8b, 0xa7, 0x20, 0xb5, 0x28, 0x33, 0x3f, 0x25, 0x24, 0x3f,
	0x2c, 0xb5, 0xb8, 0x04, 0x6c, 0x20, 0x4b, 0x10, 0x8a, 0x18, 0x48, 0x4d, 0x62, 0x6e, 0x7e, 0x69,
	0x5e, 0x09, 0x54, 0x0d, 0x13, 0xd8, 0x52, 0x14, 0x31, 0xa3, 0x38, 0x2e, 0x66, 0xdf, 0xe2, 0x74,
	0xa1, 0x70, 0x2e, 0x76, 0x98, 0x2b, 0x15, 0xf5, 0xb0, 0xfa, 0x45, 0x0f, 0xe1, 0x08, 0x29, 0x4d,
	0x82, 0x4a, 0x60, 0xee, 0x74, 0x32, 0x3f, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07,
	0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86,
	0x28, 0x59, 0x44, 0x90, 0x56, 0xa0, 0x06, 0x6a, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0x60,
	0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x10, 0x5b, 0xbe, 0x32, 0x7a, 0x01, 0x00, 0x00,
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
	Release(ctx context.Context, in *MsgRelease, opts ...grpc.CallOption) (*MsgReleaseResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Release(ctx context.Context, in *MsgRelease, opts ...grpc.CallOption) (*MsgReleaseResponse, error) {
	out := new(MsgReleaseResponse)
	err := c.cc.Invoke(ctx, "/selfchain.selfvesting.Msg/Release", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Release(context.Context, *MsgRelease) (*MsgReleaseResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Release(ctx context.Context, req *MsgRelease) (*MsgReleaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Release not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Release_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRelease)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Release(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/selfchain.selfvesting.Msg/Release",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Release(ctx, req.(*MsgRelease))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "selfchain.selfvesting.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Release",
			Handler:    _Msg_Release_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "selfchain/selfvesting/tx.proto",
}

func (m *MsgRelease) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRelease) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRelease) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PosIndex != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.PosIndex))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgReleaseResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgReleaseResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgReleaseResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AmountToVest) > 0 {
		i -= len(m.AmountToVest)
		copy(dAtA[i:], m.AmountToVest)
		i = encodeVarintTx(dAtA, i, uint64(len(m.AmountToVest)))
		i--
		dAtA[i] = 0x12
	}
	if m.PeriodToVest != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.PeriodToVest))
		i--
		dAtA[i] = 0x8
	}
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
func (m *MsgRelease) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.PosIndex != 0 {
		n += 1 + sovTx(uint64(m.PosIndex))
	}
	return n
}

func (m *MsgReleaseResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PeriodToVest != 0 {
		n += 1 + sovTx(uint64(m.PeriodToVest))
	}
	l = len(m.AmountToVest)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRelease) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRelease: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRelease: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PosIndex", wireType)
			}
			m.PosIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PosIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *MsgReleaseResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgReleaseResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgReleaseResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodToVest", wireType)
			}
			m.PeriodToVest = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeriodToVest |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountToVest", wireType)
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
			m.AmountToVest = string(dAtA[iNdEx:postIndex])
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
