// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/app/provider/v1beta1/provider_api.proto

package providerv1beta1

import (
	context "context"
	fmt "fmt"
	v1beta12 "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	v1beta11 "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	v1beta1 "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"
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

type OpenRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The resource reference.
	Ref *v1beta11.Reference `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty"`
	// REQUIRED.
	// The access token this application provider will use when contacting
	// the storage provider to read and write.
	// Service implementors MUST make sure that the access token only grants
	// access to the requested resource.
	// Service implementors should use a ResourceId rather than a filename to grant access, as
	// ResourceIds MUST NOT change when a resource is renamed.
	// The access token MUST be short-lived.
	// TODO(labkode): investigate token derivation techniques.
	AccessToken          string   `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenRequest) Reset()         { *m = OpenRequest{} }
func (m *OpenRequest) String() string { return proto.CompactTextString(m) }
func (*OpenRequest) ProtoMessage()    {}
func (*OpenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c007b70b037097fe, []int{0}
}

func (m *OpenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenRequest.Unmarshal(m, b)
}
func (m *OpenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenRequest.Marshal(b, m, deterministic)
}
func (m *OpenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenRequest.Merge(m, src)
}
func (m *OpenRequest) XXX_Size() int {
	return xxx_messageInfo_OpenRequest.Size(m)
}
func (m *OpenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OpenRequest proto.InternalMessageInfo

func (m *OpenRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *OpenRequest) GetRef() *v1beta11.Reference {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *OpenRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type OpenResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta12.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The url that user agents will render to clients.
	// Usually the rendering happens by using HTML iframes,
	// at least, Office 365, Collabora, OnlyOffice do like that.
	IframeUrl            string   `protobuf:"bytes,3,opt,name=iframe_url,json=iframeUrl,proto3" json:"iframe_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenResponse) Reset()         { *m = OpenResponse{} }
func (m *OpenResponse) String() string { return proto.CompactTextString(m) }
func (*OpenResponse) ProtoMessage()    {}
func (*OpenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c007b70b037097fe, []int{1}
}

func (m *OpenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenResponse.Unmarshal(m, b)
}
func (m *OpenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenResponse.Marshal(b, m, deterministic)
}
func (m *OpenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenResponse.Merge(m, src)
}
func (m *OpenResponse) XXX_Size() int {
	return xxx_messageInfo_OpenResponse.Size(m)
}
func (m *OpenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OpenResponse proto.InternalMessageInfo

func (m *OpenResponse) GetStatus() *v1beta12.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *OpenResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *OpenResponse) GetIframeUrl() string {
	if m != nil {
		return m.IframeUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*OpenRequest)(nil), "cs3.app.provider.v1beta1.OpenRequest")
	proto.RegisterType((*OpenResponse)(nil), "cs3.app.provider.v1beta1.OpenResponse")
}

func init() {
	proto.RegisterFile("cs3/app/provider/v1beta1/provider_api.proto", fileDescriptor_c007b70b037097fe)
}

var fileDescriptor_c007b70b037097fe = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0x49, 0x2a, 0x85, 0x3b, 0xb9, 0xa0, 0x64, 0x63, 0x2c, 0xf7, 0xc2, 0xb5, 0xa0, 0x16,
	0x94, 0x09, 0x69, 0x56, 0x2e, 0xd3, 0xae, 0x5c, 0x35, 0x44, 0xeb, 0x42, 0x0a, 0x65, 0x3a, 0x3d,
	0x95, 0x60, 0x9b, 0x39, 0x9d, 0x99, 0x14, 0x5c, 0xf9, 0x0c, 0xbe, 0x80, 0x0b, 0x97, 0x3e, 0x8a,
	0x4f, 0x25, 0xf3, 0x27, 0xa1, 0x50, 0x8a, 0x77, 0x99, 0xf3, 0xfd, 0xbe, 0x9c, 0xef, 0x9b, 0x19,
	0xf2, 0x96, 0xab, 0x3c, 0x65, 0x88, 0x29, 0x4a, 0x71, 0xaa, 0xb7, 0x20, 0xd3, 0x53, 0xb6, 0x01,
	0xcd, 0xb2, 0x7e, 0xb0, 0x66, 0x58, 0x53, 0x94, 0x42, 0x8b, 0x38, 0xe1, 0x2a, 0xa7, 0x0c, 0x91,
	0x76, 0x1a, 0xf5, 0xf0, 0xe8, 0xce, 0xfc, 0x46, 0x22, 0xef, 0xdd, 0x4a, 0x33, 0xdd, 0x2a, 0xe7,
	0x1b, 0xbd, 0x33, 0xaa, 0xd2, 0x42, 0xb2, 0xaf, 0x70, 0xb9, 0x48, 0x82, 0x12, 0xad, 0xe4, 0xd0,
	0xd1, 0xf7, 0x86, 0xd6, 0xdf, 0x11, 0x54, 0x8f, 0xd8, 0x2f, 0x27, 0x8f, 0x7f, 0x05, 0x24, 0x5a,
	0x20, 0x34, 0x15, 0x1c, 0x5b, 0x50, 0x3a, 0xce, 0xc8, 0x50, 0x20, 0x3b, 0xb6, 0x90, 0x04, 0x0f,
	0xc1, 0x24, 0x9a, 0xbe, 0xa0, 0x26, 0xa5, 0x73, 0x78, 0x3f, 0x5d, 0x58, 0xa0, 0xf2, 0x60, 0xfc,
	0x9e, 0x0c, 0x24, 0xec, 0x92, 0xd0, 0xf2, 0x6f, 0x2c, 0xef, 0xd3, 0x5d, 0x34, 0xa3, 0x15, 0xec,
	0x40, 0x42, 0xc3, 0xa1, 0x32, 0x9e, 0xf8, 0x25, 0xb9, 0x65, 0x9c, 0x83, 0x52, 0x6b, 0x2d, 0xbe,
	0x41, 0x93, 0x0c, 0x1e, 0x82, 0xc9, 0x4d, 0x15, 0xb9, 0xd9, 0x27, 0x33, 0x1a, 0xff, 0x0c, 0xc8,
	0xad, 0x0b, 0xa8, 0x50, 0x34, 0x0a, 0xe2, 0x94, 0x0c, 0xdd, 0x71, 0xf8, 0x84, 0xcf, 0xed, 0x46,
	0x89, 0xbc, 0x5f, 0xf2, 0xd1, 0xca, 0x95, 0xc7, 0xce, 0x2a, 0x85, 0x8f, 0xad, 0x74, 0x4f, 0x48,
	0xbd, 0x93, 0xec, 0x00, 0xeb, 0x56, 0xee, 0x7d, 0xaa, 0x1b, 0x37, 0x59, 0xca, 0xfd, 0x74, 0x4b,
	0xa2, 0xd2, 0x37, 0x2b, 0xca, 0x0f, 0xf1, 0x92, 0x3c, 0x31, 0x09, 0xe3, 0x57, 0xf4, 0xda, 0x8d,
	0xd2, 0xb3, 0x23, 0x1e, 0xbd, 0xfe, 0x1f, 0xe6, 0x8a, 0xce, 0x7e, 0x90, 0x3b, 0x2e, 0x0e, 0x57,
	0xe1, 0xd9, 0xb3, 0x3e, 0x03, 0xd6, 0xa5, 0xb9, 0xcc, 0x32, 0xf8, 0xf2, 0xb4, 0xa3, 0x3c, 0xf4,
	0x3b, 0x1c, 0xcc, 0x8b, 0xf2, 0x4f, 0x98, 0xcc, 0x55, 0x4e, 0x0b, 0x44, 0xda, 0x79, 0xe8, 0xe7,
	0x6c, 0x66, 0x80, 0xbf, 0x56, 0x5a, 0x15, 0x88, 0xab, 0x4e, 0x5a, 0x79, 0x69, 0x33, 0xb4, 0x4f,
	0x24, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0xad, 0xe8, 0xa5, 0x62, 0xd6, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProviderAPIClient is the client API for ProviderAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProviderAPIClient interface {
	// Returns the iframe url
	// MUST return CODE_NOT_FOUND if the resource does not exist.
	Open(ctx context.Context, in *OpenRequest, opts ...grpc.CallOption) (*OpenResponse, error)
}

type providerAPIClient struct {
	cc *grpc.ClientConn
}

func NewProviderAPIClient(cc *grpc.ClientConn) ProviderAPIClient {
	return &providerAPIClient{cc}
}

func (c *providerAPIClient) Open(ctx context.Context, in *OpenRequest, opts ...grpc.CallOption) (*OpenResponse, error) {
	out := new(OpenResponse)
	err := c.cc.Invoke(ctx, "/cs3.app.provider.v1beta1.ProviderAPI/Open", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderAPIServer is the server API for ProviderAPI service.
type ProviderAPIServer interface {
	// Returns the iframe url
	// MUST return CODE_NOT_FOUND if the resource does not exist.
	Open(context.Context, *OpenRequest) (*OpenResponse, error)
}

// UnimplementedProviderAPIServer can be embedded to have forward compatible implementations.
type UnimplementedProviderAPIServer struct {
}

func (*UnimplementedProviderAPIServer) Open(ctx context.Context, req *OpenRequest) (*OpenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Open not implemented")
}

func RegisterProviderAPIServer(s *grpc.Server, srv ProviderAPIServer) {
	s.RegisterService(&_ProviderAPI_serviceDesc, srv)
}

func _ProviderAPI_Open_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderAPIServer).Open(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.app.provider.v1beta1.ProviderAPI/Open",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderAPIServer).Open(ctx, req.(*OpenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProviderAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.app.provider.v1beta1.ProviderAPI",
	HandlerType: (*ProviderAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Open",
			Handler:    _ProviderAPI_Open_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/app/provider/v1beta1/provider_api.proto",
}
