// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: selfchain/selfvesting/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_931c644e99d2a099, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_931c644e99d2a099, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetVestingPositionsRequest struct {
	Beneficiary string `protobuf:"bytes,1,opt,name=beneficiary,proto3" json:"beneficiary,omitempty"`
}

func (m *QueryGetVestingPositionsRequest) Reset()         { *m = QueryGetVestingPositionsRequest{} }
func (m *QueryGetVestingPositionsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetVestingPositionsRequest) ProtoMessage()    {}
func (*QueryGetVestingPositionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_931c644e99d2a099, []int{2}
}
func (m *QueryGetVestingPositionsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetVestingPositionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetVestingPositionsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetVestingPositionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetVestingPositionsRequest.Merge(m, src)
}
func (m *QueryGetVestingPositionsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetVestingPositionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetVestingPositionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetVestingPositionsRequest proto.InternalMessageInfo

func (m *QueryGetVestingPositionsRequest) GetBeneficiary() string {
	if m != nil {
		return m.Beneficiary
	}
	return ""
}

type QueryGetVestingPositionsResponse struct {
	VestingPositions VestingPositions `protobuf:"bytes,1,opt,name=vestingPositions,proto3" json:"vestingPositions"`
}

func (m *QueryGetVestingPositionsResponse) Reset()         { *m = QueryGetVestingPositionsResponse{} }
func (m *QueryGetVestingPositionsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetVestingPositionsResponse) ProtoMessage()    {}
func (*QueryGetVestingPositionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_931c644e99d2a099, []int{3}
}
func (m *QueryGetVestingPositionsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetVestingPositionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetVestingPositionsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetVestingPositionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetVestingPositionsResponse.Merge(m, src)
}
func (m *QueryGetVestingPositionsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetVestingPositionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetVestingPositionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetVestingPositionsResponse proto.InternalMessageInfo

func (m *QueryGetVestingPositionsResponse) GetVestingPositions() VestingPositions {
	if m != nil {
		return m.VestingPositions
	}
	return VestingPositions{}
}

type QueryAllVestingPositionsRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllVestingPositionsRequest) Reset()         { *m = QueryAllVestingPositionsRequest{} }
func (m *QueryAllVestingPositionsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllVestingPositionsRequest) ProtoMessage()    {}
func (*QueryAllVestingPositionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_931c644e99d2a099, []int{4}
}
func (m *QueryAllVestingPositionsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllVestingPositionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllVestingPositionsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllVestingPositionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllVestingPositionsRequest.Merge(m, src)
}
func (m *QueryAllVestingPositionsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllVestingPositionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllVestingPositionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllVestingPositionsRequest proto.InternalMessageInfo

func (m *QueryAllVestingPositionsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllVestingPositionsResponse struct {
	VestingPositions []VestingPositions  `protobuf:"bytes,1,rep,name=vestingPositions,proto3" json:"vestingPositions"`
	Pagination       *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllVestingPositionsResponse) Reset()         { *m = QueryAllVestingPositionsResponse{} }
func (m *QueryAllVestingPositionsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllVestingPositionsResponse) ProtoMessage()    {}
func (*QueryAllVestingPositionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_931c644e99d2a099, []int{5}
}
func (m *QueryAllVestingPositionsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllVestingPositionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllVestingPositionsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllVestingPositionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllVestingPositionsResponse.Merge(m, src)
}
func (m *QueryAllVestingPositionsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllVestingPositionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllVestingPositionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllVestingPositionsResponse proto.InternalMessageInfo

func (m *QueryAllVestingPositionsResponse) GetVestingPositions() []VestingPositions {
	if m != nil {
		return m.VestingPositions
	}
	return nil
}

func (m *QueryAllVestingPositionsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "selfchain.selfvesting.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "selfchain.selfvesting.QueryParamsResponse")
	proto.RegisterType((*QueryGetVestingPositionsRequest)(nil), "selfchain.selfvesting.QueryGetVestingPositionsRequest")
	proto.RegisterType((*QueryGetVestingPositionsResponse)(nil), "selfchain.selfvesting.QueryGetVestingPositionsResponse")
	proto.RegisterType((*QueryAllVestingPositionsRequest)(nil), "selfchain.selfvesting.QueryAllVestingPositionsRequest")
	proto.RegisterType((*QueryAllVestingPositionsResponse)(nil), "selfchain.selfvesting.QueryAllVestingPositionsResponse")
}

func init() { proto.RegisterFile("selfchain/selfvesting/query.proto", fileDescriptor_931c644e99d2a099) }

var fileDescriptor_931c644e99d2a099 = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcf, 0x6b, 0x53, 0x31,
	0x1c, 0x6f, 0xe6, 0x2c, 0x98, 0x5d, 0x46, 0x36, 0x41, 0x8a, 0x7d, 0xad, 0x01, 0xb5, 0x16, 0x7c,
	0x71, 0x13, 0x36, 0x41, 0x10, 0x36, 0xc1, 0x5d, 0xeb, 0x3b, 0x08, 0x7a, 0x91, 0xb4, 0xa4, 0xcf,
	0xc0, 0x5b, 0xf2, 0xd6, 0x64, 0xc5, 0x22, 0xbb, 0x78, 0xf0, 0xac, 0xf8, 0xcf, 0x78, 0x17, 0x64,
	0xc7, 0x81, 0x17, 0x4f, 0x22, 0xad, 0x7f, 0x88, 0x34, 0xf9, 0xce, 0xbd, 0xb6, 0xef, 0x3d, 0x2b,
	0xec, 0xd4, 0x90, 0x7c, 0xbe, 0x9f, 0x1f, 0xe9, 0x27, 0x0f, 0xdf, 0x32, 0x22, 0xe9, 0xf7, 0xde,
	0x70, 0xa9, 0xd8, 0x74, 0x35, 0x14, 0xc6, 0x4a, 0x15, 0xb3, 0xa3, 0x63, 0x31, 0x18, 0x85, 0xe9,
	0x40, 0x5b, 0x4d, 0xae, 0xff, 0x85, 0x84, 0x19, 0x48, 0x6d, 0x33, 0xd6, 0xb1, 0x76, 0x08, 0x36,
	0x5d, 0x79, 0x70, 0xed, 0x66, 0xac, 0x75, 0x9c, 0x08, 0xc6, 0x53, 0xc9, 0xb8, 0x52, 0xda, 0x72,
	0x2b, 0xb5, 0x32, 0x70, 0xda, 0xee, 0x69, 0x73, 0xa8, 0x0d, 0xeb, 0x72, 0x23, 0xbc, 0x06, 0x1b,
	0x6e, 0x75, 0x85, 0xe5, 0x5b, 0x2c, 0xe5, 0xb1, 0x54, 0x0e, 0x0c, 0x58, 0x9a, 0xef, 0x2c, 0xe5,
	0x03, 0x7e, 0x78, 0xce, 0x77, 0x3f, 0x1f, 0x03, 0xbf, 0xaf, 0x53, 0x6d, 0x64, 0x56, 0xbe, 0x55,
	0x0e, 0x97, 0xaa, 0x0f, 0x31, 0xe8, 0x26, 0x26, 0xcf, 0xa7, 0xf6, 0x3a, 0x4e, 0x2d, 0x12, 0x47,
	0xc7, 0xc2, 0x58, 0x1a, 0xe1, 0x8d, 0x99, 0x5d, 0x93, 0x6a, 0x65, 0x04, 0x79, 0x8c, 0xab, 0xde,
	0xd5, 0x0d, 0xd4, 0x44, 0xad, 0xb5, 0xed, 0x7a, 0x98, 0x7b, 0x63, 0xa1, 0x1f, 0xdb, 0x5f, 0x3d,
	0xfd, 0xd9, 0xa8, 0x44, 0x30, 0x42, 0x9f, 0xe2, 0x86, 0xe3, 0x3c, 0x10, 0xf6, 0x85, 0xc7, 0x75,
	0xce, 0x5d, 0x83, 0x2c, 0x69, 0xe2, 0xb5, 0xae, 0x50, 0xa2, 0x2f, 0x7b, 0x92, 0x0f, 0x46, 0x4e,
	0xe4, 0x5a, 0x94, 0xdd, 0xa2, 0x27, 0xb8, 0x59, 0x4c, 0x02, 0x2e, 0x5f, 0xe2, 0xf5, 0xe1, 0xdc,
	0x19, 0xf8, 0xbd, 0x5b, 0xe0, 0x77, 0x9e, 0x0a, 0x9c, 0x2f, 0xd0, 0x50, 0x09, 0x19, 0xf6, 0x92,
	0xa4, 0x28, 0xc3, 0x33, 0x8c, 0x2f, 0xfe, 0x61, 0xd0, 0xbd, 0x13, 0xfa, 0x3a, 0x84, 0xd3, 0x3a,
	0x84, 0xbe, 0x72, 0x50, 0x87, 0xb0, 0xc3, 0x63, 0x01, 0xb3, 0x51, 0x66, 0x92, 0x7e, 0x43, 0x10,
	0x35, 0x57, 0xab, 0x34, 0xea, 0x95, 0x4b, 0x88, 0x4a, 0x0e, 0x66, 0x72, 0xac, 0xc0, 0xfd, 0xfd,
	0x2b, 0x87, 0xf7, 0x95, 0x0d, 0xb2, 0xfd, 0x69, 0x15, 0x5f, 0x75, 0x41, 0xc8, 0x07, 0x84, 0xab,
	0xbe, 0x1a, 0xe4, 0x5e, 0x81, 0xbd, 0xc5, 0x2e, 0xd6, 0xda, 0xcb, 0x40, 0xbd, 0x2e, 0xbd, 0xfd,
	0xfe, 0xfb, 0xef, 0xcf, 0x2b, 0x0d, 0x52, 0x67, 0x65, 0x6f, 0x8a, 0x7c, 0x45, 0x78, 0x7d, 0xfe,
	0x22, 0xc8, 0x4e, 0x99, 0x4e, 0x71, 0x69, 0x6b, 0xbb, 0xff, 0x3d, 0x07, 0x66, 0x9f, 0x38, 0xb3,
	0x8f, 0xc8, 0x0e, 0x5b, 0xf2, 0x71, 0xb3, 0x77, 0x99, 0xa7, 0x70, 0x42, 0xbe, 0x20, 0xbc, 0x31,
	0x4f, 0xbe, 0x97, 0x24, 0xe5, 0x41, 0x8a, 0x9b, 0x5b, 0x1e, 0xa4, 0xa4, 0x85, 0xf4, 0x81, 0x0b,
	0xd2, 0x26, 0xad, 0x65, 0x83, 0xec, 0xef, 0x9e, 0x8e, 0x03, 0x74, 0x36, 0x0e, 0xd0, 0xaf, 0x71,
	0x80, 0x3e, 0x4e, 0x82, 0xca, 0xd9, 0x24, 0xa8, 0xfc, 0x98, 0x04, 0x95, 0x57, 0xf5, 0x0b, 0x8a,
	0xb7, 0x33, 0x24, 0x76, 0x94, 0x0a, 0xd3, 0xad, 0xba, 0xaf, 0xd6, 0xc3, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x8c, 0xa1, 0x28, 0x5b, 0xce, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of VestingPositions items.
	VestingPositions(ctx context.Context, in *QueryGetVestingPositionsRequest, opts ...grpc.CallOption) (*QueryGetVestingPositionsResponse, error)
	VestingPositionsAll(ctx context.Context, in *QueryAllVestingPositionsRequest, opts ...grpc.CallOption) (*QueryAllVestingPositionsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/selfchain.selfvesting.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) VestingPositions(ctx context.Context, in *QueryGetVestingPositionsRequest, opts ...grpc.CallOption) (*QueryGetVestingPositionsResponse, error) {
	out := new(QueryGetVestingPositionsResponse)
	err := c.cc.Invoke(ctx, "/selfchain.selfvesting.Query/VestingPositions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) VestingPositionsAll(ctx context.Context, in *QueryAllVestingPositionsRequest, opts ...grpc.CallOption) (*QueryAllVestingPositionsResponse, error) {
	out := new(QueryAllVestingPositionsResponse)
	err := c.cc.Invoke(ctx, "/selfchain.selfvesting.Query/VestingPositionsAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of VestingPositions items.
	VestingPositions(context.Context, *QueryGetVestingPositionsRequest) (*QueryGetVestingPositionsResponse, error)
	VestingPositionsAll(context.Context, *QueryAllVestingPositionsRequest) (*QueryAllVestingPositionsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) VestingPositions(ctx context.Context, req *QueryGetVestingPositionsRequest) (*QueryGetVestingPositionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VestingPositions not implemented")
}
func (*UnimplementedQueryServer) VestingPositionsAll(ctx context.Context, req *QueryAllVestingPositionsRequest) (*QueryAllVestingPositionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VestingPositionsAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/selfchain.selfvesting.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_VestingPositions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetVestingPositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).VestingPositions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/selfchain.selfvesting.Query/VestingPositions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).VestingPositions(ctx, req.(*QueryGetVestingPositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_VestingPositionsAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllVestingPositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).VestingPositionsAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/selfchain.selfvesting.Query/VestingPositionsAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).VestingPositionsAll(ctx, req.(*QueryAllVestingPositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "selfchain.selfvesting.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "VestingPositions",
			Handler:    _Query_VestingPositions_Handler,
		},
		{
			MethodName: "VestingPositionsAll",
			Handler:    _Query_VestingPositionsAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "selfchain/selfvesting/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetVestingPositionsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetVestingPositionsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetVestingPositionsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Beneficiary) > 0 {
		i -= len(m.Beneficiary)
		copy(dAtA[i:], m.Beneficiary)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Beneficiary)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetVestingPositionsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetVestingPositionsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetVestingPositionsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.VestingPositions.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllVestingPositionsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllVestingPositionsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllVestingPositionsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllVestingPositionsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllVestingPositionsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllVestingPositionsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.VestingPositions) > 0 {
		for iNdEx := len(m.VestingPositions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestingPositions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetVestingPositionsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Beneficiary)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetVestingPositionsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.VestingPositions.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllVestingPositionsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllVestingPositionsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.VestingPositions) > 0 {
		for _, e := range m.VestingPositions {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetVestingPositionsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetVestingPositionsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetVestingPositionsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Beneficiary", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Beneficiary = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetVestingPositionsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetVestingPositionsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetVestingPositionsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPositions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VestingPositions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllVestingPositionsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllVestingPositionsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllVestingPositionsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllVestingPositionsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllVestingPositionsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllVestingPositionsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPositions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingPositions = append(m.VestingPositions, VestingPositions{})
			if err := m.VestingPositions[len(m.VestingPositions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
