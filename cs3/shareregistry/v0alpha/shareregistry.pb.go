// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/shareregistry/v0alpha/shareregistry.proto

package shareregistryv0alphapb

import (
	context "context"
	fmt "fmt"
	rpc "github.com/cernbox/go-cs3apis/cs3/rpc"
	sharetypes "github.com/cernbox/go-cs3apis/cs3/sharetypes"
	types "github.com/cernbox/go-cs3apis/cs3/types"
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

type GetShareProviderRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque *types.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The type of share.
	ShareType            sharetypes.ShareType `protobuf:"varint,2,opt,name=share_type,json=shareType,proto3,enum=cs3.sharetypes.ShareType" json:"share_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetShareProviderRequest) Reset()         { *m = GetShareProviderRequest{} }
func (m *GetShareProviderRequest) String() string { return proto.CompactTextString(m) }
func (*GetShareProviderRequest) ProtoMessage()    {}
func (*GetShareProviderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4371bf047f9b0345, []int{0}
}

func (m *GetShareProviderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShareProviderRequest.Unmarshal(m, b)
}
func (m *GetShareProviderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShareProviderRequest.Marshal(b, m, deterministic)
}
func (m *GetShareProviderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShareProviderRequest.Merge(m, src)
}
func (m *GetShareProviderRequest) XXX_Size() int {
	return xxx_messageInfo_GetShareProviderRequest.Size(m)
}
func (m *GetShareProviderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShareProviderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetShareProviderRequest proto.InternalMessageInfo

func (m *GetShareProviderRequest) GetOpaque() *types.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *GetShareProviderRequest) GetShareType() sharetypes.ShareType {
	if m != nil {
		return m.ShareType
	}
	return sharetypes.ShareType_SHARE_TYPE_INVALID
}

type GetShareProviderResponse struct {
	// REQUIRED.
	// The response status.
	Status *rpc.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *types.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The share provider handling the requested share resource.
	Provider             *sharetypes.ProviderInfo `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *GetShareProviderResponse) Reset()         { *m = GetShareProviderResponse{} }
func (m *GetShareProviderResponse) String() string { return proto.CompactTextString(m) }
func (*GetShareProviderResponse) ProtoMessage()    {}
func (*GetShareProviderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4371bf047f9b0345, []int{1}
}

func (m *GetShareProviderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShareProviderResponse.Unmarshal(m, b)
}
func (m *GetShareProviderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShareProviderResponse.Marshal(b, m, deterministic)
}
func (m *GetShareProviderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShareProviderResponse.Merge(m, src)
}
func (m *GetShareProviderResponse) XXX_Size() int {
	return xxx_messageInfo_GetShareProviderResponse.Size(m)
}
func (m *GetShareProviderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShareProviderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetShareProviderResponse proto.InternalMessageInfo

func (m *GetShareProviderResponse) GetStatus() *rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetShareProviderResponse) GetOpaque() *types.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *GetShareProviderResponse) GetProvider() *sharetypes.ProviderInfo {
	if m != nil {
		return m.Provider
	}
	return nil
}

type ListShareProvidersRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque               *types.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListShareProvidersRequest) Reset()         { *m = ListShareProvidersRequest{} }
func (m *ListShareProvidersRequest) String() string { return proto.CompactTextString(m) }
func (*ListShareProvidersRequest) ProtoMessage()    {}
func (*ListShareProvidersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4371bf047f9b0345, []int{2}
}

func (m *ListShareProvidersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListShareProvidersRequest.Unmarshal(m, b)
}
func (m *ListShareProvidersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListShareProvidersRequest.Marshal(b, m, deterministic)
}
func (m *ListShareProvidersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListShareProvidersRequest.Merge(m, src)
}
func (m *ListShareProvidersRequest) XXX_Size() int {
	return xxx_messageInfo_ListShareProvidersRequest.Size(m)
}
func (m *ListShareProvidersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListShareProvidersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListShareProvidersRequest proto.InternalMessageInfo

func (m *ListShareProvidersRequest) GetOpaque() *types.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

type ListShareProvidersResponse struct {
	// REQUIRED.
	// The response status.
	Status *rpc.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *types.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The list of share providers this registry knows about.
	Providers            []*sharetypes.ProviderInfo `protobuf:"bytes,3,rep,name=providers,proto3" json:"providers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ListShareProvidersResponse) Reset()         { *m = ListShareProvidersResponse{} }
func (m *ListShareProvidersResponse) String() string { return proto.CompactTextString(m) }
func (*ListShareProvidersResponse) ProtoMessage()    {}
func (*ListShareProvidersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4371bf047f9b0345, []int{3}
}

func (m *ListShareProvidersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListShareProvidersResponse.Unmarshal(m, b)
}
func (m *ListShareProvidersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListShareProvidersResponse.Marshal(b, m, deterministic)
}
func (m *ListShareProvidersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListShareProvidersResponse.Merge(m, src)
}
func (m *ListShareProvidersResponse) XXX_Size() int {
	return xxx_messageInfo_ListShareProvidersResponse.Size(m)
}
func (m *ListShareProvidersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListShareProvidersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListShareProvidersResponse proto.InternalMessageInfo

func (m *ListShareProvidersResponse) GetStatus() *rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListShareProvidersResponse) GetOpaque() *types.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *ListShareProvidersResponse) GetProviders() []*sharetypes.ProviderInfo {
	if m != nil {
		return m.Providers
	}
	return nil
}

func init() {
	proto.RegisterType((*GetShareProviderRequest)(nil), "cs3.shareregistryv0alpha.GetShareProviderRequest")
	proto.RegisterType((*GetShareProviderResponse)(nil), "cs3.shareregistryv0alpha.GetShareProviderResponse")
	proto.RegisterType((*ListShareProvidersRequest)(nil), "cs3.shareregistryv0alpha.ListShareProvidersRequest")
	proto.RegisterType((*ListShareProvidersResponse)(nil), "cs3.shareregistryv0alpha.ListShareProvidersResponse")
}

func init() {
	proto.RegisterFile("cs3/shareregistry/v0alpha/shareregistry.proto", fileDescriptor_4371bf047f9b0345)
}

var fileDescriptor_4371bf047f9b0345 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcf, 0x6b, 0xe2, 0x40,
	0x18, 0x25, 0x11, 0x64, 0x9d, 0x85, 0xfd, 0x31, 0xeb, 0xee, 0xc6, 0x20, 0xac, 0xe4, 0xb2, 0xee,
	0x61, 0x47, 0x9b, 0xf4, 0x50, 0x7a, 0x53, 0xb1, 0x22, 0x14, 0x94, 0x89, 0x94, 0xb6, 0x14, 0x4a,
	0x4c, 0xa7, 0x35, 0xd0, 0x9a, 0x71, 0x26, 0x0a, 0xb9, 0xb4, 0xff, 0x43, 0x4f, 0x3d, 0xf7, 0x54,
	0x4a, 0xff, 0x92, 0xfe, 0x55, 0x25, 0x93, 0xd1, 0x26, 0x6a, 0x10, 0x0f, 0xbd, 0x0d, 0xdf, 0x7b,
	0xef, 0x9b, 0xf7, 0xbe, 0xf9, 0x06, 0xfc, 0x77, 0xb9, 0x55, 0xe3, 0x23, 0x87, 0x11, 0x46, 0xae,
	0x3c, 0x1e, 0xb0, 0xb0, 0x36, 0xab, 0x3b, 0xd7, 0x74, 0xe4, 0xa4, 0xab, 0x88, 0x32, 0x3f, 0xf0,
	0xa1, 0xe6, 0x72, 0x0b, 0xa5, 0x00, 0xc9, 0xd6, 0x8b, 0x51, 0x23, 0x46, 0xdd, 0x1a, 0x0f, 0x9c,
	0x60, 0xca, 0x63, 0xbe, 0xfe, 0x67, 0xd1, 0x3e, 0x08, 0x29, 0xe1, 0x89, 0xa3, 0x24, 0xfc, 0x8c,
	0x08, 0x31, 0x96, 0x28, 0x1b, 0xb7, 0xe0, 0x77, 0x87, 0x04, 0x76, 0xc4, 0xee, 0x33, 0x7f, 0xe6,
	0x5d, 0x10, 0x86, 0xc9, 0x64, 0x4a, 0x78, 0x00, 0xff, 0x81, 0xbc, 0x4f, 0x9d, 0xc9, 0x94, 0x68,
	0x4a, 0x45, 0xa9, 0x7e, 0x36, 0xbf, 0xa3, 0xc8, 0x53, 0x2c, 0xee, 0x09, 0x00, 0x4b, 0x02, 0xdc,
	0x03, 0x40, 0x5c, 0x78, 0x1e, 0xa1, 0x9a, 0x5a, 0x51, 0xaa, 0x5f, 0xcc, 0x12, 0x5a, 0x44, 0x88,
	0x35, 0xe2, 0x92, 0x41, 0x48, 0x09, 0x2e, 0xf0, 0xf9, 0xd1, 0x78, 0x52, 0x80, 0xb6, 0x6a, 0x80,
	0x53, 0x7f, 0xcc, 0x09, 0xfc, 0x0b, 0xf2, 0x71, 0x48, 0xe9, 0xe0, 0xab, 0x68, 0xc9, 0xa8, 0x8b,
	0x6c, 0x51, 0xc6, 0x12, 0x4e, 0x58, 0x55, 0x37, 0x5b, 0xfd, 0x44, 0xe5, 0x3d, 0x5a, 0x4e, 0x90,
	0xcb, 0xcb, 0x46, 0xe7, 0x3e, 0xba, 0xe3, 0x4b, 0x1f, 0x2f, 0xd8, 0xc6, 0x01, 0x28, 0x1d, 0x7a,
	0x3c, 0x6d, 0x95, 0x6f, 0x3f, 0x2c, 0xe3, 0x45, 0x01, 0xfa, 0xba, 0x46, 0x1f, 0x18, 0x7a, 0x1f,
	0x14, 0xe6, 0x31, 0xb8, 0x96, 0xab, 0xe4, 0x36, 0xa6, 0x7e, 0xa7, 0x9b, 0xf7, 0x2a, 0x28, 0x0a,
	0xab, 0x58, 0x2e, 0xa2, 0x4d, 0xd8, 0xcc, 0x73, 0x09, 0x0c, 0xc1, 0xb7, 0xe5, 0x97, 0x83, 0x3b,
	0x28, 0x6b, 0x6f, 0x51, 0xc6, 0x9a, 0xe9, 0xe6, 0x36, 0x12, 0x39, 0xa3, 0x3b, 0x00, 0x57, 0x27,
	0x08, 0xad, 0xec, 0x4e, 0x99, 0x0f, 0xa7, 0xef, 0x6e, 0x27, 0x8a, 0x0d, 0x34, 0x1f, 0x14, 0x50,
	0x76, 0xfd, 0x9b, 0x4c, 0x6d, 0x13, 0xda, 0xc9, 0x6a, 0x3f, 0xfa, 0x6b, 0x7d, 0xe5, 0xf4, 0xd7,
	0x3a, 0x2e, 0x1d, 0x3e, 0xaa, 0x3f, 0x5a, 0xcd, 0xde, 0xb1, 0x3d, 0xe8, 0xe1, 0x46, 0xa7, 0x8d,
	0xdb, 0x9d, 0xae, 0x3d, 0xc0, 0x27, 0xcf, 0xaa, 0xd6, 0xb2, 0x2d, 0x94, 0x9a, 0xfd, 0x51, 0xbd,
	0x11, 0x49, 0x5e, 0x05, 0x74, 0xb6, 0x0e, 0x1a, 0xe6, 0xc5, 0xc7, 0xb6, 0xde, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xb9, 0x12, 0xf3, 0xfc, 0x71, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShareRegistryServiceClient is the client API for ShareRegistryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShareRegistryServiceClient interface {
	// Returns the share provider that is reponsible for the given
	// resource reference.
	// MUST return CODE_NOT_FOUND if the reference does not exist.
	GetShareProvider(ctx context.Context, in *GetShareProviderRequest, opts ...grpc.CallOption) (*GetShareProviderResponse, error)
	// Returns a list of the available share providers known by this registry.
	ListShareProviders(ctx context.Context, in *ListShareProvidersRequest, opts ...grpc.CallOption) (*ListShareProvidersResponse, error)
}

type shareRegistryServiceClient struct {
	cc *grpc.ClientConn
}

func NewShareRegistryServiceClient(cc *grpc.ClientConn) ShareRegistryServiceClient {
	return &shareRegistryServiceClient{cc}
}

func (c *shareRegistryServiceClient) GetShareProvider(ctx context.Context, in *GetShareProviderRequest, opts ...grpc.CallOption) (*GetShareProviderResponse, error) {
	out := new(GetShareProviderResponse)
	err := c.cc.Invoke(ctx, "/cs3.shareregistryv0alpha.ShareRegistryService/GetShareProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shareRegistryServiceClient) ListShareProviders(ctx context.Context, in *ListShareProvidersRequest, opts ...grpc.CallOption) (*ListShareProvidersResponse, error) {
	out := new(ListShareProvidersResponse)
	err := c.cc.Invoke(ctx, "/cs3.shareregistryv0alpha.ShareRegistryService/ListShareProviders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShareRegistryServiceServer is the server API for ShareRegistryService service.
type ShareRegistryServiceServer interface {
	// Returns the share provider that is reponsible for the given
	// resource reference.
	// MUST return CODE_NOT_FOUND if the reference does not exist.
	GetShareProvider(context.Context, *GetShareProviderRequest) (*GetShareProviderResponse, error)
	// Returns a list of the available share providers known by this registry.
	ListShareProviders(context.Context, *ListShareProvidersRequest) (*ListShareProvidersResponse, error)
}

// UnimplementedShareRegistryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedShareRegistryServiceServer struct {
}

func (*UnimplementedShareRegistryServiceServer) GetShareProvider(ctx context.Context, req *GetShareProviderRequest) (*GetShareProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShareProvider not implemented")
}
func (*UnimplementedShareRegistryServiceServer) ListShareProviders(ctx context.Context, req *ListShareProvidersRequest) (*ListShareProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListShareProviders not implemented")
}

func RegisterShareRegistryServiceServer(s *grpc.Server, srv ShareRegistryServiceServer) {
	s.RegisterService(&_ShareRegistryService_serviceDesc, srv)
}

func _ShareRegistryService_GetShareProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShareProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShareRegistryServiceServer).GetShareProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.shareregistryv0alpha.ShareRegistryService/GetShareProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShareRegistryServiceServer).GetShareProvider(ctx, req.(*GetShareProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShareRegistryService_ListShareProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListShareProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShareRegistryServiceServer).ListShareProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.shareregistryv0alpha.ShareRegistryService/ListShareProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShareRegistryServiceServer).ListShareProviders(ctx, req.(*ListShareProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShareRegistryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.shareregistryv0alpha.ShareRegistryService",
	HandlerType: (*ShareRegistryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShareProvider",
			Handler:    _ShareRegistryService_GetShareProvider_Handler,
		},
		{
			MethodName: "ListShareProviders",
			Handler:    _ShareRegistryService_ListShareProviders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/shareregistry/v0alpha/shareregistry.proto",
}