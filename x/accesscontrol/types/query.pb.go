// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accesscontrol_x/query.proto

package types

import (
	context "context"
	fmt "fmt"
	accesscontrol "github.com/cosmos/cosmos-sdk/types/accesscontrol"
	_ "github.com/cosmos/cosmos-sdk/types/query"
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

type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d83f2274e13e6a16, []int{0}
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

type QueryParamsResponse struct {
	// params defines the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d83f2274e13e6a16, []int{1}
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

type ResourceDependencyMappingFromMessageKeyRequest struct {
	MessageKey string `protobuf:"bytes,1,opt,name=message_key,json=messageKey,proto3" json:"message_key,omitempty" yaml:"message_key"`
}

func (m *ResourceDependencyMappingFromMessageKeyRequest) Reset() {
	*m = ResourceDependencyMappingFromMessageKeyRequest{}
}
func (m *ResourceDependencyMappingFromMessageKeyRequest) String() string {
	return proto.CompactTextString(m)
}
func (*ResourceDependencyMappingFromMessageKeyRequest) ProtoMessage() {}
func (*ResourceDependencyMappingFromMessageKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d83f2274e13e6a16, []int{2}
}
func (m *ResourceDependencyMappingFromMessageKeyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourceDependencyMappingFromMessageKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResourceDependencyMappingFromMessageKeyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResourceDependencyMappingFromMessageKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceDependencyMappingFromMessageKeyRequest.Merge(m, src)
}
func (m *ResourceDependencyMappingFromMessageKeyRequest) XXX_Size() int {
	return m.Size()
}
func (m *ResourceDependencyMappingFromMessageKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceDependencyMappingFromMessageKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceDependencyMappingFromMessageKeyRequest proto.InternalMessageInfo

func (m *ResourceDependencyMappingFromMessageKeyRequest) GetMessageKey() string {
	if m != nil {
		return m.MessageKey
	}
	return ""
}

type ResourceDependencyMappingFromMessageKeyResponse struct {
	MessageDependencyMapping accesscontrol.MessageDependencyMapping `protobuf:"bytes,1,opt,name=message_dependency_mapping,json=messageDependencyMapping,proto3" json:"message_dependency_mapping" yaml:"message_dependency_mapping"`
}

func (m *ResourceDependencyMappingFromMessageKeyResponse) Reset() {
	*m = ResourceDependencyMappingFromMessageKeyResponse{}
}
func (m *ResourceDependencyMappingFromMessageKeyResponse) String() string {
	return proto.CompactTextString(m)
}
func (*ResourceDependencyMappingFromMessageKeyResponse) ProtoMessage() {}
func (*ResourceDependencyMappingFromMessageKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d83f2274e13e6a16, []int{3}
}
func (m *ResourceDependencyMappingFromMessageKeyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourceDependencyMappingFromMessageKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResourceDependencyMappingFromMessageKeyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResourceDependencyMappingFromMessageKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceDependencyMappingFromMessageKeyResponse.Merge(m, src)
}
func (m *ResourceDependencyMappingFromMessageKeyResponse) XXX_Size() int {
	return m.Size()
}
func (m *ResourceDependencyMappingFromMessageKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceDependencyMappingFromMessageKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceDependencyMappingFromMessageKeyResponse proto.InternalMessageInfo

func (m *ResourceDependencyMappingFromMessageKeyResponse) GetMessageDependencyMapping() accesscontrol.MessageDependencyMapping {
	if m != nil {
		return m.MessageDependencyMapping
	}
	return accesscontrol.MessageDependencyMapping{}
}

type ListResourceDependencyMappingRequest struct {
}

func (m *ListResourceDependencyMappingRequest) Reset()         { *m = ListResourceDependencyMappingRequest{} }
func (m *ListResourceDependencyMappingRequest) String() string { return proto.CompactTextString(m) }
func (*ListResourceDependencyMappingRequest) ProtoMessage()    {}
func (*ListResourceDependencyMappingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d83f2274e13e6a16, []int{4}
}
func (m *ListResourceDependencyMappingRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListResourceDependencyMappingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListResourceDependencyMappingRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListResourceDependencyMappingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResourceDependencyMappingRequest.Merge(m, src)
}
func (m *ListResourceDependencyMappingRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListResourceDependencyMappingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResourceDependencyMappingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListResourceDependencyMappingRequest proto.InternalMessageInfo

type ListResourceDependencyMappingResponse struct {
	MessageDependencyMappingList []accesscontrol.MessageDependencyMapping `protobuf:"bytes,1,rep,name=message_dependency_mapping_list,json=messageDependencyMappingList,proto3" json:"message_dependency_mapping_list" yaml:"message_dependency_mapping_list"`
}

func (m *ListResourceDependencyMappingResponse) Reset()         { *m = ListResourceDependencyMappingResponse{} }
func (m *ListResourceDependencyMappingResponse) String() string { return proto.CompactTextString(m) }
func (*ListResourceDependencyMappingResponse) ProtoMessage()    {}
func (*ListResourceDependencyMappingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d83f2274e13e6a16, []int{5}
}
func (m *ListResourceDependencyMappingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListResourceDependencyMappingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListResourceDependencyMappingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListResourceDependencyMappingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResourceDependencyMappingResponse.Merge(m, src)
}
func (m *ListResourceDependencyMappingResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListResourceDependencyMappingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResourceDependencyMappingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResourceDependencyMappingResponse proto.InternalMessageInfo

func (m *ListResourceDependencyMappingResponse) GetMessageDependencyMappingList() []accesscontrol.MessageDependencyMapping {
	if m != nil {
		return m.MessageDependencyMappingList
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "cosmos.accesscontrol_x.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "cosmos.accesscontrol_x.v1beta1.QueryParamsResponse")
	proto.RegisterType((*ResourceDependencyMappingFromMessageKeyRequest)(nil), "cosmos.accesscontrol_x.v1beta1.ResourceDependencyMappingFromMessageKeyRequest")
	proto.RegisterType((*ResourceDependencyMappingFromMessageKeyResponse)(nil), "cosmos.accesscontrol_x.v1beta1.ResourceDependencyMappingFromMessageKeyResponse")
	proto.RegisterType((*ListResourceDependencyMappingRequest)(nil), "cosmos.accesscontrol_x.v1beta1.ListResourceDependencyMappingRequest")
	proto.RegisterType((*ListResourceDependencyMappingResponse)(nil), "cosmos.accesscontrol_x.v1beta1.ListResourceDependencyMappingResponse")
}

func init() {
	proto.RegisterFile("cosmos/accesscontrol_x/query.proto", fileDescriptor_d83f2274e13e6a16)
}

var fileDescriptor_d83f2274e13e6a16 = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xb3, 0x85, 0x46, 0x62, 0x7b, 0x5b, 0x2a, 0x14, 0x59, 0xc5, 0x29, 0xab, 0x92, 0xb6,
	0x48, 0x78, 0xd5, 0x54, 0x02, 0x89, 0x1b, 0x21, 0x20, 0x21, 0x5a, 0x3e, 0x7c, 0xa4, 0x42, 0xd6,
	0xc6, 0x19, 0x8c, 0xd5, 0xd8, 0xeb, 0x7a, 0x1d, 0x54, 0x0b, 0x71, 0xe1, 0x09, 0x90, 0x78, 0x00,
	0x24, 0xee, 0x88, 0xd7, 0xe8, 0xb1, 0x52, 0x0f, 0xf4, 0x14, 0xa1, 0x84, 0x27, 0xe8, 0x85, 0x2b,
	0x8a, 0xbd, 0x8d, 0xf2, 0xe5, 0x38, 0x55, 0x7b, 0x4a, 0xbc, 0x9e, 0x99, 0xff, 0xfc, 0x66, 0xff,
	0x63, 0x4c, 0x6d, 0x21, 0x3d, 0x21, 0x19, 0xb7, 0x6d, 0x90, 0xd2, 0x16, 0x7e, 0x14, 0x8a, 0x96,
	0x75, 0xc8, 0x0e, 0xda, 0x10, 0xc6, 0x46, 0x10, 0x8a, 0x48, 0x10, 0x3d, 0x8d, 0x31, 0xc6, 0x62,
	0x8c, 0x8f, 0x5b, 0x0d, 0x88, 0xf8, 0x96, 0xb6, 0xec, 0x08, 0x47, 0x24, 0xa1, 0xac, 0xff, 0x2f,
	0xcd, 0xd2, 0x56, 0x1c, 0x21, 0x9c, 0x16, 0x30, 0x1e, 0xb8, 0x8c, 0xfb, 0xbe, 0x88, 0x78, 0xe4,
	0x0a, 0x5f, 0xaa, 0xb7, 0xf7, 0x94, 0x6e, 0x83, 0x4b, 0x48, 0xc5, 0x98, 0x2a, 0xc7, 0x02, 0xee,
	0xb8, 0x7e, 0x12, 0xac, 0x62, 0x37, 0xa6, 0xf5, 0x38, 0xfa, 0xa4, 0x22, 0xd7, 0x32, 0x68, 0x1c,
	0xf0, 0x41, 0xba, 0x4a, 0x9b, 0x2e, 0x63, 0xf2, 0xa6, 0xaf, 0xf8, 0x9a, 0x87, 0xdc, 0x93, 0x26,
	0x1c, 0xb4, 0x41, 0x46, 0x74, 0x0f, 0xdf, 0x1c, 0x39, 0x95, 0x81, 0xf0, 0x25, 0x90, 0x3a, 0x2e,
	0x06, 0xc9, 0x49, 0x09, 0xad, 0xa2, 0x8d, 0xa5, 0x6a, 0xc5, 0x98, 0x3d, 0x0d, 0x23, 0xcd, 0xaf,
	0x5d, 0x3f, 0xea, 0x94, 0x0b, 0xa6, 0xca, 0xa5, 0x2e, 0x36, 0x4c, 0x90, 0xa2, 0x1d, 0xda, 0x50,
	0x87, 0x00, 0xfc, 0x26, 0xf8, 0x76, 0xbc, 0xcb, 0x83, 0xc0, 0xf5, 0x9d, 0x67, 0xa1, 0xf0, 0x76,
	0x41, 0x4a, 0xee, 0xc0, 0x0b, 0x88, 0x55, 0x3b, 0xe4, 0x21, 0x5e, 0xf2, 0xd2, 0x43, 0x6b, 0x1f,
	0xe2, 0x44, 0xfc, 0x46, 0xed, 0xd6, 0x59, 0xa7, 0x4c, 0x62, 0xee, 0xb5, 0x1e, 0xd1, 0xa1, 0x97,
	0xd4, 0xc4, 0xde, 0x20, 0x9f, 0x9e, 0x20, 0xcc, 0xe6, 0xd6, 0x52, 0x90, 0xdf, 0x11, 0xd6, 0xce,
	0x0b, 0x36, 0x07, 0x39, 0x96, 0x97, 0x26, 0x29, 0xf2, 0x07, 0x53, 0xc9, 0x07, 0xdc, 0xaa, 0xec,
	0x84, 0x64, 0x6d, 0xb3, 0x3f, 0x89, 0xb3, 0x4e, 0xf9, 0xce, 0x68, 0xe3, 0x93, 0x3a, 0xd4, 0x2c,
	0x79, 0x19, 0x45, 0x68, 0x05, 0xaf, 0xed, 0xb8, 0x32, 0xca, 0x04, 0x3b, 0xbf, 0xc5, 0xdf, 0x08,
	0xdf, 0xcd, 0x09, 0x54, 0xcc, 0x3f, 0x11, 0x2e, 0x67, 0xf7, 0x62, 0xb5, 0x5c, 0x19, 0x95, 0xd0,
	0xea, 0xb5, 0x4b, 0x80, 0x1b, 0x0a, 0xbc, 0x92, 0x07, 0x9e, 0x88, 0x51, 0x73, 0x25, 0x8b, 0xbe,
	0x0f, 0x54, 0x3d, 0x5d, 0xc4, 0x8b, 0x89, 0x41, 0xc9, 0x0f, 0x84, 0x8b, 0xa9, 0xcb, 0x48, 0x35,
	0xcf, 0x8d, 0x93, 0x46, 0xd7, 0xb6, 0x2f, 0x94, 0x93, 0x4e, 0x8b, 0xb2, 0x2f, 0x27, 0x7f, 0xbf,
	0x2d, 0x6c, 0x92, 0x75, 0xa6, 0x56, 0x2c, 0xfd, 0xb9, 0x2f, 0x9b, 0xfb, 0x63, 0x7b, 0x99, 0x3a,
	0x9e, 0xfc, 0x5a, 0xc0, 0xeb, 0x73, 0xda, 0x90, 0xbc, 0xcc, 0xeb, 0xe8, 0x62, 0xbb, 0xa3, 0xbd,
	0xba, 0xb2, 0x7a, 0x8a, 0xde, 0x4e, 0xe8, 0xdf, 0x91, 0xbd, 0x5c, 0xfa, 0x50, 0x55, 0x9e, 0x76,
	0xcb, 0xef, 0x43, 0xe1, 0x59, 0x43, 0x7b, 0xcb, 0x3e, 0x0d, 0x3d, 0x7c, 0x26, 0xff, 0x10, 0xbe,
	0x3d, 0xd3, 0xba, 0xa4, 0x9e, 0xc7, 0x35, 0xcf, 0x8a, 0x68, 0x4f, 0x2f, 0x59, 0x45, 0xcd, 0xe4,
	0x79, 0x32, 0x93, 0x27, 0xe4, 0x71, 0xee, 0x4c, 0xfa, 0xee, 0xb6, 0x66, 0x0c, 0xa6, 0xb6, 0x73,
	0xd4, 0xd5, 0xd1, 0x71, 0x57, 0x47, 0x7f, 0xba, 0x3a, 0xfa, 0xda, 0xd3, 0x0b, 0xc7, 0x3d, 0xbd,
	0x70, 0xda, 0xd3, 0x0b, 0x6f, 0xab, 0x8e, 0x1b, 0x7d, 0x68, 0x37, 0x0c, 0x5b, 0x78, 0x53, 0x64,
	0x0e, 0xc7, 0x84, 0xa2, 0x38, 0x00, 0xd9, 0x28, 0x26, 0x5f, 0xf9, 0xed, 0xff, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x31, 0x19, 0x2e, 0xf2, 0xdb, 0x06, 0x00, 0x00,
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
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	ResourceDependencyMappingFromMessageKey(ctx context.Context, in *ResourceDependencyMappingFromMessageKeyRequest, opts ...grpc.CallOption) (*ResourceDependencyMappingFromMessageKeyResponse, error)
	ListResourceDependencyMapping(ctx context.Context, in *ListResourceDependencyMappingRequest, opts ...grpc.CallOption) (*ListResourceDependencyMappingResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.accesscontrol_x.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ResourceDependencyMappingFromMessageKey(ctx context.Context, in *ResourceDependencyMappingFromMessageKeyRequest, opts ...grpc.CallOption) (*ResourceDependencyMappingFromMessageKeyResponse, error) {
	out := new(ResourceDependencyMappingFromMessageKeyResponse)
	err := c.cc.Invoke(ctx, "/cosmos.accesscontrol_x.v1beta1.Query/ResourceDependencyMappingFromMessageKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListResourceDependencyMapping(ctx context.Context, in *ListResourceDependencyMappingRequest, opts ...grpc.CallOption) (*ListResourceDependencyMappingResponse, error) {
	out := new(ListResourceDependencyMappingResponse)
	err := c.cc.Invoke(ctx, "/cosmos.accesscontrol_x.v1beta1.Query/ListResourceDependencyMapping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	ResourceDependencyMappingFromMessageKey(context.Context, *ResourceDependencyMappingFromMessageKeyRequest) (*ResourceDependencyMappingFromMessageKeyResponse, error)
	ListResourceDependencyMapping(context.Context, *ListResourceDependencyMappingRequest) (*ListResourceDependencyMappingResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) ResourceDependencyMappingFromMessageKey(ctx context.Context, req *ResourceDependencyMappingFromMessageKeyRequest) (*ResourceDependencyMappingFromMessageKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResourceDependencyMappingFromMessageKey not implemented")
}
func (*UnimplementedQueryServer) ListResourceDependencyMapping(ctx context.Context, req *ListResourceDependencyMappingRequest) (*ListResourceDependencyMappingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListResourceDependencyMapping not implemented")
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
		FullMethod: "/cosmos.accesscontrol_x.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ResourceDependencyMappingFromMessageKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceDependencyMappingFromMessageKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ResourceDependencyMappingFromMessageKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.accesscontrol_x.v1beta1.Query/ResourceDependencyMappingFromMessageKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ResourceDependencyMappingFromMessageKey(ctx, req.(*ResourceDependencyMappingFromMessageKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListResourceDependencyMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResourceDependencyMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListResourceDependencyMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.accesscontrol_x.v1beta1.Query/ListResourceDependencyMapping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListResourceDependencyMapping(ctx, req.(*ListResourceDependencyMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.accesscontrol_x.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "ResourceDependencyMappingFromMessageKey",
			Handler:    _Query_ResourceDependencyMappingFromMessageKey_Handler,
		},
		{
			MethodName: "ListResourceDependencyMapping",
			Handler:    _Query_ListResourceDependencyMapping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/accesscontrol_x/query.proto",
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

func (m *ResourceDependencyMappingFromMessageKeyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceDependencyMappingFromMessageKeyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResourceDependencyMappingFromMessageKeyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MessageKey) > 0 {
		i -= len(m.MessageKey)
		copy(dAtA[i:], m.MessageKey)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.MessageKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResourceDependencyMappingFromMessageKeyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceDependencyMappingFromMessageKeyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResourceDependencyMappingFromMessageKeyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.MessageDependencyMapping.MarshalToSizedBuffer(dAtA[:i])
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

func (m *ListResourceDependencyMappingRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListResourceDependencyMappingRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListResourceDependencyMappingRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ListResourceDependencyMappingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListResourceDependencyMappingResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListResourceDependencyMappingResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MessageDependencyMappingList) > 0 {
		for iNdEx := len(m.MessageDependencyMappingList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MessageDependencyMappingList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *ResourceDependencyMappingFromMessageKeyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MessageKey)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *ResourceDependencyMappingFromMessageKeyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MessageDependencyMapping.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *ListResourceDependencyMappingRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ListResourceDependencyMappingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MessageDependencyMappingList) > 0 {
		for _, e := range m.MessageDependencyMappingList {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
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
func (m *ResourceDependencyMappingFromMessageKeyRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ResourceDependencyMappingFromMessageKeyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceDependencyMappingFromMessageKeyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageKey", wireType)
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
			m.MessageKey = string(dAtA[iNdEx:postIndex])
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
func (m *ResourceDependencyMappingFromMessageKeyResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ResourceDependencyMappingFromMessageKeyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceDependencyMappingFromMessageKeyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageDependencyMapping", wireType)
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
			if err := m.MessageDependencyMapping.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *ListResourceDependencyMappingRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ListResourceDependencyMappingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListResourceDependencyMappingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *ListResourceDependencyMappingResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ListResourceDependencyMappingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListResourceDependencyMappingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageDependencyMappingList", wireType)
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
			m.MessageDependencyMappingList = append(m.MessageDependencyMappingList, accesscontrol.MessageDependencyMapping{})
			if err := m.MessageDependencyMappingList[len(m.MessageDependencyMappingList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
