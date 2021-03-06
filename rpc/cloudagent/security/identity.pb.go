// Code generated by protoc-gen-go. DO NOT EDIT.
// source: identity.proto

package security

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
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

type IdentityRequest struct {
	Identitys            []*Identity      `protobuf:"bytes,1,rep,name=Identitys,proto3" json:"Identitys,omitempty"`
	OperationType        common.Operation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *IdentityRequest) Reset()         { *m = IdentityRequest{} }
func (m *IdentityRequest) String() string { return proto.CompactTextString(m) }
func (*IdentityRequest) ProtoMessage()    {}
func (*IdentityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61c7956abb761639, []int{0}
}

func (m *IdentityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentityRequest.Unmarshal(m, b)
}
func (m *IdentityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentityRequest.Marshal(b, m, deterministic)
}
func (m *IdentityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityRequest.Merge(m, src)
}
func (m *IdentityRequest) XXX_Size() int {
	return xxx_messageInfo_IdentityRequest.Size(m)
}
func (m *IdentityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityRequest proto.InternalMessageInfo

func (m *IdentityRequest) GetIdentitys() []*Identity {
	if m != nil {
		return m.Identitys
	}
	return nil
}

func (m *IdentityRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type IdentityResponse struct {
	Identitys            []*Identity         `protobuf:"bytes,1,rep,name=Identitys,proto3" json:"Identitys,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *IdentityResponse) Reset()         { *m = IdentityResponse{} }
func (m *IdentityResponse) String() string { return proto.CompactTextString(m) }
func (*IdentityResponse) ProtoMessage()    {}
func (*IdentityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61c7956abb761639, []int{1}
}

func (m *IdentityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentityResponse.Unmarshal(m, b)
}
func (m *IdentityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentityResponse.Marshal(b, m, deterministic)
}
func (m *IdentityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityResponse.Merge(m, src)
}
func (m *IdentityResponse) XXX_Size() int {
	return xxx_messageInfo_IdentityResponse.Size(m)
}
func (m *IdentityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityResponse proto.InternalMessageInfo

func (m *IdentityResponse) GetIdentitys() []*Identity {
	if m != nil {
		return m.Identitys
	}
	return nil
}

func (m *IdentityResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *IdentityResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Identity struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string         `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	ResourceGroup        string         `protobuf:"bytes,3,opt,name=resourceGroup,proto3" json:"resourceGroup,omitempty"`
	Password             string         `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Token                string         `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
	Certificate          []byte         `protobuf:"bytes,6,opt,name=certificate,proto3" json:"certificate,omitempty"`
	Status               *common.Status `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	LocationName         string         `protobuf:"bytes,10,opt,name=locationName,proto3" json:"locationName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Identity) Reset()         { *m = Identity{} }
func (m *Identity) String() string { return proto.CompactTextString(m) }
func (*Identity) ProtoMessage()    {}
func (*Identity) Descriptor() ([]byte, []int) {
	return fileDescriptor_61c7956abb761639, []int{2}
}

func (m *Identity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identity.Unmarshal(m, b)
}
func (m *Identity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identity.Marshal(b, m, deterministic)
}
func (m *Identity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identity.Merge(m, src)
}
func (m *Identity) XXX_Size() int {
	return xxx_messageInfo_Identity.Size(m)
}
func (m *Identity) XXX_DiscardUnknown() {
	xxx_messageInfo_Identity.DiscardUnknown(m)
}

var xxx_messageInfo_Identity proto.InternalMessageInfo

func (m *Identity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Identity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Identity) GetResourceGroup() string {
	if m != nil {
		return m.ResourceGroup
	}
	return ""
}

func (m *Identity) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Identity) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Identity) GetCertificate() []byte {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func (m *Identity) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Identity) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func init() {
	proto.RegisterType((*IdentityRequest)(nil), "moc.cloudagent.security.IdentityRequest")
	proto.RegisterType((*IdentityResponse)(nil), "moc.cloudagent.security.IdentityResponse")
	proto.RegisterType((*Identity)(nil), "moc.cloudagent.security.Identity")
}

func init() { proto.RegisterFile("identity.proto", fileDescriptor_61c7956abb761639) }

var fileDescriptor_61c7956abb761639 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4d, 0x8b, 0xd4, 0x40,
	0x10, 0x35, 0xb3, 0xbb, 0x71, 0xa7, 0xe6, 0x43, 0x69, 0x04, 0x43, 0x0e, 0x12, 0xa3, 0x87, 0x78,
	0xb0, 0x03, 0x51, 0xc4, 0x9b, 0xb8, 0x20, 0xb2, 0x17, 0x85, 0x28, 0x1e, 0xbc, 0x48, 0xa6, 0x53,
	0x33, 0x36, 0x9b, 0xa4, 0xda, 0xfe, 0x70, 0x99, 0x7f, 0xe0, 0x9f, 0xf0, 0x47, 0xfa, 0x0f, 0x64,
	0x3a, 0xc9, 0x8c, 0x73, 0x10, 0x85, 0xbd, 0xa5, 0x5e, 0xbf, 0x7a, 0xf5, 0xea, 0x55, 0x60, 0x29,
	0x6b, 0xec, 0xac, 0xb4, 0x5b, 0xae, 0x34, 0x59, 0x62, 0xf7, 0x5b, 0x12, 0x5c, 0x34, 0xe4, 0xea,
	0x6a, 0x83, 0x9d, 0xe5, 0x06, 0x85, 0xd3, 0xd2, 0x6e, 0xe3, 0x07, 0x1b, 0xa2, 0x4d, 0x83, 0xb9,
	0xa7, 0xad, 0xdc, 0x3a, 0xbf, 0xd6, 0x95, 0x52, 0xa8, 0x4d, 0xdf, 0x18, 0xcf, 0x05, 0xb5, 0x2d,
	0x75, 0x7d, 0x95, 0xfe, 0x08, 0xe0, 0xce, 0xe5, 0xa0, 0x5c, 0xe2, 0x37, 0x87, 0xc6, 0xb2, 0x57,
	0x30, 0x1d, 0x21, 0x13, 0x05, 0xc9, 0x49, 0x36, 0x2b, 0x1e, 0xf2, 0xbf, 0x8c, 0xe3, 0xfb, 0xe6,
	0x43, 0x0f, 0x7b, 0x0e, 0x8b, 0xf7, 0x0a, 0x75, 0x65, 0x25, 0x75, 0x1f, 0xb7, 0x0a, 0xa3, 0x49,
	0x12, 0x64, 0xcb, 0x62, 0xe9, 0x45, 0xf6, 0x2f, 0xe5, 0x31, 0x29, 0xfd, 0x19, 0xc0, 0xdd, 0x83,
	0x15, 0xa3, 0xa8, 0x33, 0x78, 0x73, 0x2f, 0x05, 0x84, 0x25, 0x1a, 0xd7, 0x58, 0x6f, 0x62, 0x56,
	0xc4, 0xbc, 0xcf, 0x87, 0x8f, 0xf9, 0xf0, 0x0b, 0xa2, 0xe6, 0x53, 0xd5, 0x38, 0x2c, 0x07, 0x26,
	0xbb, 0x07, 0x67, 0x6f, 0xb4, 0x26, 0x1d, 0x9d, 0x24, 0x41, 0x36, 0x2d, 0xfb, 0x22, 0xfd, 0x15,
	0xc0, 0xf9, 0xa8, 0xcb, 0x18, 0x9c, 0x76, 0x55, 0x8b, 0x51, 0xe0, 0x19, 0xfe, 0x9b, 0x2d, 0x61,
	0x22, 0x6b, 0x3f, 0x66, 0x5a, 0x4e, 0x64, 0xcd, 0x1e, 0xc3, 0x42, 0xa3, 0x21, 0xa7, 0x05, 0xbe,
	0xd5, 0xe4, 0xd4, 0x20, 0x77, 0x0c, 0xb2, 0x18, 0xce, 0x55, 0x65, 0xcc, 0x35, 0xe9, 0x3a, 0x3a,
	0xf5, 0x84, 0x7d, 0xbd, 0x33, 0x62, 0xe9, 0x0a, 0xbb, 0xe8, 0xac, 0x37, 0xe2, 0x0b, 0x96, 0xc0,
	0x4c, 0xa0, 0xb6, 0x72, 0x2d, 0x45, 0x65, 0x31, 0x0a, 0x93, 0x20, 0x9b, 0x97, 0x7f, 0x42, 0xec,
	0x11, 0x84, 0xc6, 0x56, 0xd6, 0x99, 0xe8, 0xb6, 0x5f, 0x7a, 0xe6, 0x23, 0xfb, 0xe0, 0xa1, 0x72,
	0x78, 0x62, 0x29, 0xcc, 0x1b, 0x12, 0x3e, 0xff, 0x77, 0xbb, 0x55, 0xc0, 0xcf, 0x38, 0xc2, 0x0a,
	0x05, 0x8b, 0x71, 0xe5, 0xd7, 0xbb, 0xac, 0xd9, 0x17, 0x08, 0x2f, 0xbb, 0xef, 0x74, 0x85, 0x2c,
	0xfb, 0xf7, 0x19, 0xfa, 0xff, 0x29, 0x7e, 0xf2, 0x1f, 0xcc, 0xfe, 0xdc, 0xe9, 0xad, 0x8b, 0x97,
	0x9f, 0x5f, 0x6c, 0xa4, 0xfd, 0xea, 0x56, 0x5c, 0x50, 0x9b, 0xb7, 0x52, 0x68, 0x32, 0xb4, 0xb6,
	0x79, 0x4b, 0xe2, 0xa9, 0x3f, 0x59, 0xae, 0x95, 0xc8, 0x0f, 0x62, 0xf9, 0x28, 0xb6, 0x0a, 0xfd,
	0xf3, 0xb3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x44, 0xe5, 0x5b, 0xd4, 0x2a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IdentityAgentClient is the client API for IdentityAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdentityAgentClient interface {
	Invoke(ctx context.Context, in *IdentityRequest, opts ...grpc.CallOption) (*IdentityResponse, error)
}

type identityAgentClient struct {
	cc *grpc.ClientConn
}

func NewIdentityAgentClient(cc *grpc.ClientConn) IdentityAgentClient {
	return &identityAgentClient{cc}
}

func (c *identityAgentClient) Invoke(ctx context.Context, in *IdentityRequest, opts ...grpc.CallOption) (*IdentityResponse, error) {
	out := new(IdentityResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.security.IdentityAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityAgentServer is the server API for IdentityAgent service.
type IdentityAgentServer interface {
	Invoke(context.Context, *IdentityRequest) (*IdentityResponse, error)
}

// UnimplementedIdentityAgentServer can be embedded to have forward compatible implementations.
type UnimplementedIdentityAgentServer struct {
}

func (*UnimplementedIdentityAgentServer) Invoke(ctx context.Context, req *IdentityRequest) (*IdentityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterIdentityAgentServer(s *grpc.Server, srv IdentityAgentServer) {
	s.RegisterService(&_IdentityAgent_serviceDesc, srv)
}

func _IdentityAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.security.IdentityAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityAgentServer).Invoke(ctx, req.(*IdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IdentityAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.security.IdentityAgent",
	HandlerType: (*IdentityAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _IdentityAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity.proto",
}
