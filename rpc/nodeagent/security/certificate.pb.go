// Code generated by protoc-gen-go. DO NOT EDIT.
// source: certificate.proto

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

type CertificateType int32

const (
	CertificateType_Client CertificateType = 0
	CertificateType_Server CertificateType = 1
)

var CertificateType_name = map[int32]string{
	0: "Client",
	1: "Server",
}

var CertificateType_value = map[string]int32{
	"Client": 0,
	"Server": 1,
}

func (x CertificateType) String() string {
	return proto.EnumName(CertificateType_name, int32(x))
}

func (CertificateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c0d34c34dd33be4b, []int{0}
}

type CertificateRequest struct {
	Certificates         []*Certificate `protobuf:"bytes,1,rep,name=Certificates,proto3" json:"Certificates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CertificateRequest) Reset()         { *m = CertificateRequest{} }
func (m *CertificateRequest) String() string { return proto.CompactTextString(m) }
func (*CertificateRequest) ProtoMessage()    {}
func (*CertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0d34c34dd33be4b, []int{0}
}

func (m *CertificateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateRequest.Unmarshal(m, b)
}
func (m *CertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateRequest.Marshal(b, m, deterministic)
}
func (m *CertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateRequest.Merge(m, src)
}
func (m *CertificateRequest) XXX_Size() int {
	return xxx_messageInfo_CertificateRequest.Size(m)
}
func (m *CertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateRequest proto.InternalMessageInfo

func (m *CertificateRequest) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

type CertificateResponse struct {
	Certificates         []*Certificate      `protobuf:"bytes,1,rep,name=Certificates,proto3" json:"Certificates,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CertificateResponse) Reset()         { *m = CertificateResponse{} }
func (m *CertificateResponse) String() string { return proto.CompactTextString(m) }
func (*CertificateResponse) ProtoMessage()    {}
func (*CertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0d34c34dd33be4b, []int{1}
}

func (m *CertificateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateResponse.Unmarshal(m, b)
}
func (m *CertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateResponse.Marshal(b, m, deterministic)
}
func (m *CertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateResponse.Merge(m, src)
}
func (m *CertificateResponse) XXX_Size() int {
	return xxx_messageInfo_CertificateResponse.Size(m)
}
func (m *CertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateResponse proto.InternalMessageInfo

func (m *CertificateResponse) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

func (m *CertificateResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *CertificateResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Certificate struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string          `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	NotBefore            int64           `protobuf:"varint,3,opt,name=notBefore,proto3" json:"notBefore,omitempty"`
	NotAfter             int64           `protobuf:"varint,4,opt,name=notAfter,proto3" json:"notAfter,omitempty"`
	Certificate          string          `protobuf:"bytes,5,opt,name=certificate,proto3" json:"certificate,omitempty"`
	Status               *common.Status  `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Type                 CertificateType `protobuf:"varint,7,opt,name=type,proto3,enum=moc.nodeagent.security.CertificateType" json:"type,omitempty"`
	Entity               *common.Entity  `protobuf:"bytes,8,opt,name=entity,proto3" json:"entity,omitempty"`
	Tags                 *common.Tags    `protobuf:"bytes,9,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Certificate) Reset()         { *m = Certificate{} }
func (m *Certificate) String() string { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()    {}
func (*Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0d34c34dd33be4b, []int{2}
}

func (m *Certificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificate.Unmarshal(m, b)
}
func (m *Certificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificate.Marshal(b, m, deterministic)
}
func (m *Certificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificate.Merge(m, src)
}
func (m *Certificate) XXX_Size() int {
	return xxx_messageInfo_Certificate.Size(m)
}
func (m *Certificate) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificate.DiscardUnknown(m)
}

var xxx_messageInfo_Certificate proto.InternalMessageInfo

func (m *Certificate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Certificate) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Certificate) GetNotBefore() int64 {
	if m != nil {
		return m.NotBefore
	}
	return 0
}

func (m *Certificate) GetNotAfter() int64 {
	if m != nil {
		return m.NotAfter
	}
	return 0
}

func (m *Certificate) GetCertificate() string {
	if m != nil {
		return m.Certificate
	}
	return ""
}

func (m *Certificate) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Certificate) GetType() CertificateType {
	if m != nil {
		return m.Type
	}
	return CertificateType_Client
}

func (m *Certificate) GetEntity() *common.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *Certificate) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterEnum("moc.nodeagent.security.CertificateType", CertificateType_name, CertificateType_value)
	proto.RegisterType((*CertificateRequest)(nil), "moc.nodeagent.security.CertificateRequest")
	proto.RegisterType((*CertificateResponse)(nil), "moc.nodeagent.security.CertificateResponse")
	proto.RegisterType((*Certificate)(nil), "moc.nodeagent.security.Certificate")
}

func init() { proto.RegisterFile("certificate.proto", fileDescriptor_c0d34c34dd33be4b) }

var fileDescriptor_c0d34c34dd33be4b = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x9d, 0xd4, 0xd4, 0xe3, 0xaa, 0x84, 0x05, 0xa1, 0x95, 0x05, 0xc8, 0x4a, 0x0f, 0x84,
	0x22, 0xd9, 0xc2, 0x1c, 0x39, 0x35, 0xa1, 0xea, 0x11, 0x69, 0x5b, 0x38, 0x20, 0x21, 0xe1, 0x6c,
	0xc6, 0xc6, 0xc2, 0xf6, 0x9a, 0xdd, 0x31, 0x28, 0x7f, 0xc4, 0x9d, 0x7f, 0xe0, 0xbb, 0x50, 0xd6,
	0x69, 0xeb, 0x22, 0x0e, 0x39, 0xd0, 0xdb, 0xce, 0x9b, 0x37, 0xef, 0xf9, 0xad, 0x67, 0xe1, 0x81,
	0x44, 0x4d, 0x65, 0x5e, 0xca, 0x8c, 0x30, 0x6e, 0xb5, 0x22, 0xc5, 0x1e, 0xd7, 0x4a, 0xc6, 0x8d,
	0x5a, 0x61, 0x56, 0x60, 0x43, 0xb1, 0x41, 0xd9, 0xe9, 0x92, 0xd6, 0xe1, 0xb3, 0x42, 0xa9, 0xa2,
	0xc2, 0xc4, 0xb2, 0x96, 0x5d, 0x9e, 0xfc, 0xd0, 0x59, 0xdb, 0xa2, 0x36, 0xfd, 0x5c, 0x78, 0x28,
	0x55, 0x5d, 0xab, 0xa6, 0xaf, 0xa6, 0x9f, 0x80, 0x2d, 0x6e, 0xa4, 0x05, 0x7e, 0xeb, 0xd0, 0x10,
	0x3b, 0x87, 0xc3, 0x01, 0x6a, 0xb8, 0x13, 0x8d, 0x66, 0x41, 0x7a, 0x1c, 0xff, 0xdb, 0x32, 0x1e,
	0x2a, 0xdc, 0x1a, 0x9c, 0xfe, 0x74, 0xe0, 0xe1, 0x2d, 0x7d, 0xd3, 0xaa, 0xc6, 0xe0, 0x7f, 0x33,
	0x60, 0x29, 0x78, 0x02, 0x4d, 0x57, 0x11, 0x77, 0x23, 0x67, 0x16, 0xa4, 0x61, 0xdc, 0xc7, 0x8f,
	0xaf, 0xe2, 0xc7, 0x73, 0xa5, 0xaa, 0x0f, 0x59, 0xd5, 0xa1, 0xd8, 0x32, 0xd9, 0x23, 0xd8, 0x3f,
	0xd3, 0x5a, 0x69, 0x3e, 0x8a, 0x9c, 0x99, 0x2f, 0xfa, 0x62, 0xfa, 0xcb, 0x85, 0x60, 0x20, 0xcd,
	0x18, 0x8c, 0x9b, 0xac, 0x46, 0xee, 0x58, 0x92, 0x3d, 0xb3, 0x23, 0x70, 0xcb, 0x95, 0x75, 0xf2,
	0x85, 0x5b, 0xae, 0xd8, 0x13, 0xf0, 0x1b, 0x45, 0x73, 0xcc, 0x95, 0x46, 0xab, 0x36, 0x12, 0x37,
	0x00, 0x0b, 0xe1, 0xa0, 0x51, 0x74, 0x9a, 0x13, 0x6a, 0x3e, 0xb6, 0xcd, 0xeb, 0x9a, 0x45, 0x10,
	0x0c, 0x7e, 0x29, 0xdf, 0xb7, 0x92, 0x43, 0x88, 0x1d, 0x83, 0x67, 0x28, 0xa3, 0xce, 0x70, 0xcf,
	0x26, 0x0b, 0xec, 0xe5, 0x5c, 0x58, 0x48, 0x6c, 0x5b, 0xec, 0x0d, 0x8c, 0x69, 0xdd, 0x22, 0xbf,
	0x17, 0x39, 0xb3, 0xa3, 0xf4, 0xf9, 0x0e, 0xf7, 0x77, 0xb9, 0x6e, 0x51, 0xd8, 0xa1, 0x8d, 0x03,
	0x36, 0x54, 0xd2, 0x9a, 0x1f, 0x0c, 0x1c, 0xce, 0x2c, 0x24, 0xb6, 0x2d, 0xf6, 0x14, 0xc6, 0x94,
	0x15, 0x86, 0xfb, 0x96, 0xe2, 0x5b, 0xca, 0x65, 0x56, 0x18, 0x61, 0xe1, 0x93, 0x17, 0x70, 0xff,
	0x2f, 0x71, 0x06, 0xe0, 0x2d, 0xaa, 0x12, 0x1b, 0x9a, 0xec, 0x6d, 0xce, 0x17, 0xa8, 0xbf, 0xa3,
	0x9e, 0x38, 0xe9, 0x6f, 0x17, 0x26, 0x03, 0xee, 0xe9, 0xe6, 0x13, 0xd9, 0x57, 0x38, 0x5a, 0x68,
	0xcc, 0x08, 0xdf, 0xe9, 0xf7, 0xed, 0x6a, 0x93, 0xfb, 0x64, 0x97, 0x25, 0xe8, 0xf7, 0x34, 0x7c,
	0xb9, 0x13, 0xb7, 0xdf, 0xb9, 0xe9, 0x1e, 0xfb, 0x0c, 0xa3, 0x73, 0xa4, 0xbb, 0x74, 0x90, 0xe0,
	0xbd, 0xc5, 0x0a, 0xef, 0x34, 0xc6, 0xfc, 0xd5, 0xc7, 0xa4, 0x28, 0xe9, 0x4b, 0xb7, 0x8c, 0xa5,
	0xaa, 0x93, 0xba, 0x94, 0x5a, 0x19, 0x95, 0x53, 0x52, 0x2b, 0x99, 0xe8, 0x56, 0x26, 0xd7, 0x42,
	0xc9, 0x95, 0xd0, 0xd2, 0xb3, 0xcf, 0xe1, 0xf5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x52, 0x87,
	0xee, 0xee, 0x48, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CertificateAgentClient is the client API for CertificateAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CertificateAgentClient interface {
	CreateOrUpdate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
	Get(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
	Delete(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
}

type certificateAgentClient struct {
	cc *grpc.ClientConn
}

func NewCertificateAgentClient(cc *grpc.ClientConn) CertificateAgentClient {
	return &certificateAgentClient{cc}
}

func (c *certificateAgentClient) CreateOrUpdate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.security.CertificateAgent/CreateOrUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAgentClient) Get(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.security.CertificateAgent/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAgentClient) Delete(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.security.CertificateAgent/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateAgentServer is the server API for CertificateAgent service.
type CertificateAgentServer interface {
	CreateOrUpdate(context.Context, *CertificateRequest) (*CertificateResponse, error)
	Get(context.Context, *CertificateRequest) (*CertificateResponse, error)
	Delete(context.Context, *CertificateRequest) (*CertificateResponse, error)
}

// UnimplementedCertificateAgentServer can be embedded to have forward compatible implementations.
type UnimplementedCertificateAgentServer struct {
}

func (*UnimplementedCertificateAgentServer) CreateOrUpdate(ctx context.Context, req *CertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdate not implemented")
}
func (*UnimplementedCertificateAgentServer) Get(ctx context.Context, req *CertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCertificateAgentServer) Delete(ctx context.Context, req *CertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterCertificateAgentServer(s *grpc.Server, srv CertificateAgentServer) {
	s.RegisterService(&_CertificateAgent_serviceDesc, srv)
}

func _CertificateAgent_CreateOrUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).CreateOrUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.security.CertificateAgent/CreateOrUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).CreateOrUpdate(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAgent_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.security.CertificateAgent/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).Get(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAgent_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.security.CertificateAgent/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).Delete(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.nodeagent.security.CertificateAgent",
	HandlerType: (*CertificateAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrUpdate",
			Handler:    _CertificateAgent_CreateOrUpdate_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CertificateAgent_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CertificateAgent_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "certificate.proto",
}
