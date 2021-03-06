// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/auth/applications/v1beta1/resources.proto

package applicationsv1beta1

import (
	fmt "fmt"
	v1beta1 "github.com/cs3org/go-cs3apis/cs3/auth/provider/v1beta1"
	v1beta11 "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	v1beta12 "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"
	proto "github.com/golang/protobuf/proto"
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

// AppPassword stores information about secondary passwords generated by users
// to be used with third-party applications.
type AppPassword struct {
	// REQUIRED.
	// The generated access password.
	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	// OPTIONAL.
	// The scope of the token to be issued.
	// This would be a list of resources with corresponding role-based access scope.
	TokenScope map[string]*v1beta1.Scope `protobuf:"bytes,2,rep,name=token_scope,json=tokenScope,proto3" json:"token_scope,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// OPTIONAL.
	// A label to be associated with the password.
	Label string `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	// REQUIRED.
	// The user who created the password.
	User *v1beta11.UserId `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	// OPTIONAL.
	// The time when the token will expire.
	Expiration *v1beta12.Timestamp `protobuf:"bytes,5,opt,name=expiration,proto3" json:"expiration,omitempty"`
	// REQUIRED.
	// The creation time of the password.
	Ctime *v1beta12.Timestamp `protobuf:"bytes,6,opt,name=ctime,proto3" json:"ctime,omitempty"`
	// REQUIRED.
	// The last time the password was used.
	Utime                *v1beta12.Timestamp `protobuf:"bytes,7,opt,name=utime,proto3" json:"utime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *AppPassword) Reset()         { *m = AppPassword{} }
func (m *AppPassword) String() string { return proto.CompactTextString(m) }
func (*AppPassword) ProtoMessage()    {}
func (*AppPassword) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b1b72514a65f2c5, []int{0}
}

func (m *AppPassword) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppPassword.Unmarshal(m, b)
}
func (m *AppPassword) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppPassword.Marshal(b, m, deterministic)
}
func (m *AppPassword) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppPassword.Merge(m, src)
}
func (m *AppPassword) XXX_Size() int {
	return xxx_messageInfo_AppPassword.Size(m)
}
func (m *AppPassword) XXX_DiscardUnknown() {
	xxx_messageInfo_AppPassword.DiscardUnknown(m)
}

var xxx_messageInfo_AppPassword proto.InternalMessageInfo

func (m *AppPassword) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AppPassword) GetTokenScope() map[string]*v1beta1.Scope {
	if m != nil {
		return m.TokenScope
	}
	return nil
}

func (m *AppPassword) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *AppPassword) GetUser() *v1beta11.UserId {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *AppPassword) GetExpiration() *v1beta12.Timestamp {
	if m != nil {
		return m.Expiration
	}
	return nil
}

func (m *AppPassword) GetCtime() *v1beta12.Timestamp {
	if m != nil {
		return m.Ctime
	}
	return nil
}

func (m *AppPassword) GetUtime() *v1beta12.Timestamp {
	if m != nil {
		return m.Utime
	}
	return nil
}

func init() {
	proto.RegisterType((*AppPassword)(nil), "cs3.auth.applications.v1beta1.AppPassword")
	proto.RegisterMapType((map[string]*v1beta1.Scope)(nil), "cs3.auth.applications.v1beta1.AppPassword.TokenScopeEntry")
}

func init() {
	proto.RegisterFile("cs3/auth/applications/v1beta1/resources.proto", fileDescriptor_1b1b72514a65f2c5)
}

var fileDescriptor_1b1b72514a65f2c5 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0xca, 0xd3, 0x30,
	0x18, 0xc6, 0x69, 0xbb, 0x4e, 0x4d, 0x41, 0x25, 0x7a, 0x50, 0x86, 0x83, 0xce, 0xa3, 0x79, 0x60,
	0xca, 0x56, 0x14, 0x19, 0x9e, 0x74, 0xc3, 0x03, 0xcf, 0x46, 0x9d, 0x1e, 0xe8, 0x60, 0x64, 0x6d,
	0xc0, 0xb0, 0xb6, 0x09, 0x49, 0x3a, 0xed, 0x5d, 0x78, 0x0d, 0x1e, 0x7a, 0x29, 0x5e, 0x82, 0x57,
	0x23, 0x49, 0xff, 0xac, 0x7c, 0xf0, 0x8d, 0x9d, 0x25, 0x7d, 0x7e, 0xcf, 0xf3, 0xe6, 0x7d, 0xfb,
	0x82, 0xd7, 0xa9, 0x8c, 0x42, 0x5c, 0xa9, 0xef, 0x21, 0xe6, 0x3c, 0xa7, 0x29, 0x56, 0x94, 0x95,
	0x32, 0x3c, 0x2f, 0x8e, 0x44, 0xe1, 0x45, 0x28, 0x88, 0x64, 0x95, 0x48, 0x89, 0x44, 0x5c, 0x30,
	0xc5, 0xe0, 0x34, 0x95, 0x11, 0xd2, 0x38, 0x1a, 0xe2, 0xa8, 0xc5, 0x27, 0xaf, 0xfa, 0x34, 0x2e,
	0xd8, 0x99, 0x66, 0x44, 0xdc, 0x97, 0xd4, 0xa0, 0x34, 0x23, 0xa5, 0xa2, 0xaa, 0x0e, 0x2b, 0x79,
	0x05, 0xd5, 0x45, 0x43, 0x55, 0x73, 0x72, 0x79, 0x97, 0xb9, 0x35, 0xf2, 0xcb, 0x7f, 0x0e, 0xf0,
	0x62, 0xce, 0xb7, 0x58, 0xca, 0x1f, 0x4c, 0x64, 0x70, 0x02, 0x1e, 0xf2, 0xf6, 0xec, 0x5b, 0x81,
	0x35, 0x7f, 0x94, 0xf4, 0x77, 0xf8, 0x0d, 0x78, 0x8a, 0x9d, 0x48, 0x79, 0x90, 0x29, 0xe3, 0xc4,
	0xb7, 0x03, 0x67, 0xee, 0x2d, 0x57, 0xe8, 0x6a, 0x57, 0x68, 0x10, 0x8e, 0x76, 0xda, 0xfd, 0x49,
	0x9b, 0x3f, 0x94, 0x4a, 0xd4, 0x09, 0x50, 0xfd, 0x07, 0xf8, 0x1c, 0xb8, 0x39, 0x3e, 0x92, 0xdc,
	0x77, 0x4c, 0xd5, 0xe6, 0x02, 0xdf, 0x80, 0x91, 0xee, 0xce, 0x1f, 0x05, 0xd6, 0xdc, 0x5b, 0xce,
	0x4c, 0xad, 0xae, 0x6f, 0xa4, 0x95, 0xbe, 0xce, 0x67, 0x49, 0xc4, 0xc7, 0x2c, 0x31, 0x38, 0x7c,
	0x0f, 0x00, 0xf9, 0xc9, 0xa9, 0x30, 0x6f, 0xf1, 0x5d, 0x63, 0x7e, 0x61, 0xcc, 0x4d, 0xef, 0x9d,
	0x69, 0x47, 0x0b, 0x22, 0x15, 0x2e, 0x78, 0x32, 0xe0, 0xe1, 0x12, 0xb8, 0xa9, 0xa2, 0x05, 0xf1,
	0xc7, 0x37, 0x18, 0x1b, 0x54, 0x7b, 0x2a, 0xe3, 0x79, 0x70, 0x8b, 0xc7, 0xa0, 0x93, 0x03, 0x78,
	0x72, 0x67, 0x22, 0xf0, 0x29, 0x70, 0x4e, 0xa4, 0x6e, 0x27, 0xaf, 0x8f, 0xf0, 0x2d, 0x70, 0xcf,
	0x38, 0xaf, 0xf4, 0xb8, 0x75, 0x70, 0x70, 0x19, 0x77, 0xb7, 0x25, 0x7d, 0x01, 0x93, 0x93, 0x34,
	0xf8, 0xca, 0x7e, 0x67, 0xad, 0x7f, 0x59, 0x60, 0x96, 0xb2, 0xe2, 0xfa, 0x1f, 0x5a, 0x3f, 0x4e,
	0xba, 0x95, 0xd9, 0xea, 0x95, 0xd8, 0x5a, 0x5f, 0x9f, 0x0d, 0xb9, 0x16, 0xfb, 0x6d, 0x3b, 0x9b,
	0x38, 0xfe, 0x63, 0x4f, 0x37, 0x32, 0x42, 0xb1, 0x0e, 0x8b, 0x87, 0x61, 0x5f, 0x16, 0x6b, 0x4d,
	0xfd, 0x35, 0xfa, 0x5e, 0xeb, 0xfb, 0xa1, 0xbe, 0x6f, 0xf5, 0xe3, 0xd8, 0xac, 0x5d, 0xf4, 0x3f,
	0x00, 0x00, 0xff, 0xff, 0x4f, 0x72, 0x99, 0xb3, 0x3b, 0x03, 0x00, 0x00,
}
