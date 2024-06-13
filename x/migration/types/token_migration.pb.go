// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: selfchain/migration/token_migration.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type TokenMigration struct {
	MsgHash   string `protobuf:"bytes,1,opt,name=msgHash,proto3" json:"msgHash,omitempty"`
	Processed bool   `protobuf:"varint,2,opt,name=processed,proto3" json:"processed,omitempty"`
}

func (m *TokenMigration) Reset()         { *m = TokenMigration{} }
func (m *TokenMigration) String() string { return proto.CompactTextString(m) }
func (*TokenMigration) ProtoMessage()    {}
func (*TokenMigration) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4c85e2c2274004d, []int{0}
}
func (m *TokenMigration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenMigration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenMigration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenMigration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenMigration.Merge(m, src)
}
func (m *TokenMigration) XXX_Size() int {
	return m.Size()
}
func (m *TokenMigration) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenMigration.DiscardUnknown(m)
}

var xxx_messageInfo_TokenMigration proto.InternalMessageInfo

func (m *TokenMigration) GetMsgHash() string {
	if m != nil {
		return m.MsgHash
	}
	return ""
}

func (m *TokenMigration) GetProcessed() bool {
	if m != nil {
		return m.Processed
	}
	return false
}

func init() {
	proto.RegisterType((*TokenMigration)(nil), "selfchain.migration.TokenMigration")
}

func init() {
	proto.RegisterFile("selfchain/migration/token_migration.proto", fileDescriptor_b4c85e2c2274004d)
}

var fileDescriptor_b4c85e2c2274004d = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2c, 0x4e, 0xcd, 0x49,
	0x4b, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0xcf, 0xcd, 0x4c, 0x2f, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3,
	0x2f, 0xc9, 0xcf, 0x4e, 0xcd, 0x8b, 0x87, 0xf3, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x84,
	0xe1, 0x4a, 0xf5, 0xe0, 0x52, 0x4a, 0x1e, 0x5c, 0x7c, 0x21, 0x20, 0xd5, 0xbe, 0x30, 0x11, 0x21,
	0x09, 0x2e, 0xf6, 0xdc, 0xe2, 0x74, 0x8f, 0xc4, 0xe2, 0x0c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x18, 0x57, 0x48, 0x86, 0x8b, 0xb3, 0xa0, 0x28, 0x3f, 0x39, 0xb5, 0xb8, 0x38, 0x35, 0x45,
	0x82, 0x49, 0x81, 0x51, 0x83, 0x23, 0x08, 0x21, 0xe0, 0x64, 0x7a, 0xe2, 0x91, 0x1c, 0xe3, 0x85,
	0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3,
	0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xd2, 0x08, 0x37, 0x56, 0x20, 0xbb, 0xb2, 0xb2, 0x20, 0xb5, 0x38,
	0x89, 0x0d, 0xec, 0x38, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x12, 0x81, 0x09, 0x81, 0xc9,
	0x00, 0x00, 0x00,
}

func (m *TokenMigration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenMigration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenMigration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Processed {
		i--
		if m.Processed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.MsgHash) > 0 {
		i -= len(m.MsgHash)
		copy(dAtA[i:], m.MsgHash)
		i = encodeVarintTokenMigration(dAtA, i, uint64(len(m.MsgHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTokenMigration(dAtA []byte, offset int, v uint64) int {
	offset -= sovTokenMigration(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenMigration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MsgHash)
	if l > 0 {
		n += 1 + l + sovTokenMigration(uint64(l))
	}
	if m.Processed {
		n += 2
	}
	return n
}

func sovTokenMigration(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTokenMigration(x uint64) (n int) {
	return sovTokenMigration(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenMigration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokenMigration
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
			return fmt.Errorf("proto: TokenMigration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenMigration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenMigration
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
				return ErrInvalidLengthTokenMigration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTokenMigration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Processed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenMigration
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
			m.Processed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTokenMigration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTokenMigration
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
func skipTokenMigration(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTokenMigration
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
					return 0, ErrIntOverflowTokenMigration
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
					return 0, ErrIntOverflowTokenMigration
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
				return 0, ErrInvalidLengthTokenMigration
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTokenMigration
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTokenMigration
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTokenMigration        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTokenMigration          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTokenMigration = fmt.Errorf("proto: unexpected end of group")
)