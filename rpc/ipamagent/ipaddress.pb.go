// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipaddress/ipaddress.proto

package ipamagent

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

type IPAddressFamily int32

const (
	IPAddressFamily_IPV4 IPAddressFamily = 0
	IPAddressFamily_IPV6 IPAddressFamily = 1
)

var IPAddressFamily_name = map[int32]string{
	0: "IPV4",
	1: "IPV6",
}

var IPAddressFamily_value = map[string]int32{
	"IPV4": 0,
	"IPV6": 1,
}

func (x IPAddressFamily) String() string {
	return proto.EnumName(IPAddressFamily_name, int32(x))
}

func (IPAddressFamily) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d3ba8c2fb5244de7, []int{0}
}

type IPAddressRequest struct {
	IPAddresss           []*IPAddress     `protobuf:"bytes,1,rep,name=IPAddresss,proto3" json:"IPAddresss,omitempty"`
	OperationType        common.Operation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *IPAddressRequest) Reset()         { *m = IPAddressRequest{} }
func (m *IPAddressRequest) String() string { return proto.CompactTextString(m) }
func (*IPAddressRequest) ProtoMessage()    {}
func (*IPAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3ba8c2fb5244de7, []int{0}
}

func (m *IPAddressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPAddressRequest.Unmarshal(m, b)
}
func (m *IPAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPAddressRequest.Marshal(b, m, deterministic)
}
func (m *IPAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPAddressRequest.Merge(m, src)
}
func (m *IPAddressRequest) XXX_Size() int {
	return xxx_messageInfo_IPAddressRequest.Size(m)
}
func (m *IPAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IPAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IPAddressRequest proto.InternalMessageInfo

func (m *IPAddressRequest) GetIPAddresss() []*IPAddress {
	if m != nil {
		return m.IPAddresss
	}
	return nil
}

func (m *IPAddressRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type IPAddressResponse struct {
	IPAddresss           []*IPAddress        `protobuf:"bytes,1,rep,name=IPAddresss,proto3" json:"IPAddresss,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *IPAddressResponse) Reset()         { *m = IPAddressResponse{} }
func (m *IPAddressResponse) String() string { return proto.CompactTextString(m) }
func (*IPAddressResponse) ProtoMessage()    {}
func (*IPAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3ba8c2fb5244de7, []int{1}
}

func (m *IPAddressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPAddressResponse.Unmarshal(m, b)
}
func (m *IPAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPAddressResponse.Marshal(b, m, deterministic)
}
func (m *IPAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPAddressResponse.Merge(m, src)
}
func (m *IPAddressResponse) XXX_Size() int {
	return xxx_messageInfo_IPAddressResponse.Size(m)
}
func (m *IPAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IPAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IPAddressResponse proto.InternalMessageInfo

func (m *IPAddressResponse) GetIPAddresss() []*IPAddress {
	if m != nil {
		return m.IPAddresss
	}
	return nil
}

func (m *IPAddressResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *IPAddressResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type IPAddress struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string         `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Address              string         `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Cidr                 string         `protobuf:"bytes,4,opt,name=cidr,proto3" json:"cidr,omitempty"`
	PrefixLength         int32          `protobuf:"varint,5,opt,name=prefixLength,proto3" json:"prefixLength,omitempty"`
	Status               *common.Status `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *IPAddress) Reset()         { *m = IPAddress{} }
func (m *IPAddress) String() string { return proto.CompactTextString(m) }
func (*IPAddress) ProtoMessage()    {}
func (*IPAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3ba8c2fb5244de7, []int{2}
}

func (m *IPAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPAddress.Unmarshal(m, b)
}
func (m *IPAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPAddress.Marshal(b, m, deterministic)
}
func (m *IPAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPAddress.Merge(m, src)
}
func (m *IPAddress) XXX_Size() int {
	return xxx_messageInfo_IPAddress.Size(m)
}
func (m *IPAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_IPAddress.DiscardUnknown(m)
}

var xxx_messageInfo_IPAddress proto.InternalMessageInfo

func (m *IPAddress) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IPAddress) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *IPAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *IPAddress) GetCidr() string {
	if m != nil {
		return m.Cidr
	}
	return ""
}

func (m *IPAddress) GetPrefixLength() int32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *IPAddress) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterEnum("moc.ipaddress.IPAddressFamily", IPAddressFamily_name, IPAddressFamily_value)
	proto.RegisterType((*IPAddressRequest)(nil), "moc.ipaddress.IPAddressRequest")
	proto.RegisterType((*IPAddressResponse)(nil), "moc.ipaddress.IPAddressResponse")
	proto.RegisterType((*IPAddress)(nil), "moc.ipaddress.IPAddress")
}

func init() { proto.RegisterFile("ipaddress/ipaddress.proto", fileDescriptor_d3ba8c2fb5244de7) }

var fileDescriptor_d3ba8c2fb5244de7 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4f, 0x6f, 0xd4, 0x3e,
	0x10, 0xad, 0xf7, 0xdf, 0xaf, 0x3b, 0xdb, 0xee, 0x6f, 0xb1, 0x38, 0x98, 0x3d, 0x40, 0x14, 0x84,
	0xb4, 0x42, 0xe0, 0x48, 0xa1, 0x42, 0x5c, 0x5b, 0xfe, 0xa9, 0x08, 0x44, 0x15, 0x50, 0x0f, 0xdc,
	0xbc, 0xc9, 0x6c, 0x6a, 0x29, 0x8e, 0x8d, 0xed, 0x08, 0x7a, 0xe5, 0x4b, 0xf0, 0x0d, 0xf8, 0x90,
	0x9c, 0xd0, 0x3a, 0x69, 0xe8, 0x1e, 0x7a, 0xa2, 0xb7, 0x99, 0x37, 0xef, 0xcd, 0x4c, 0x5e, 0xc6,
	0x70, 0x4f, 0x1a, 0x51, 0x14, 0x16, 0x9d, 0x4b, 0xfa, 0x88, 0x1b, 0xab, 0xbd, 0xa6, 0x87, 0x4a,
	0xe7, 0xbc, 0x07, 0x97, 0xf7, 0x4b, 0xad, 0xcb, 0x0a, 0x93, 0x50, 0x5c, 0x37, 0x9b, 0xe4, 0x9b,
	0x15, 0xc6, 0xa0, 0xed, 0xe8, 0xcb, 0x83, 0x5c, 0x2b, 0xa5, 0xeb, 0x36, 0x8b, 0x7f, 0x10, 0x58,
	0x9c, 0x9e, 0x1d, 0xb7, 0xda, 0x0c, 0xbf, 0x36, 0xe8, 0x3c, 0x7d, 0x01, 0xd0, 0x63, 0x8e, 0x91,
	0x68, 0xb8, 0x9a, 0xa5, 0x8c, 0xef, 0x8c, 0xe1, 0x7f, 0x45, 0xd7, 0xb8, 0xf4, 0x08, 0x0e, 0x3f,
	0x1a, 0xb4, 0xc2, 0x4b, 0x5d, 0x7f, 0xbe, 0x34, 0xc8, 0x06, 0x11, 0x59, 0xcd, 0xd3, 0x79, 0x10,
	0xf7, 0x95, 0x6c, 0x97, 0x14, 0xff, 0x24, 0x70, 0xe7, 0xda, 0x12, 0xce, 0xe8, 0xda, 0xe1, 0x3f,
	0x6c, 0x91, 0xc2, 0x24, 0x43, 0xd7, 0x54, 0x3e, 0x8c, 0x9f, 0xa5, 0x4b, 0xde, 0x7a, 0xc2, 0xaf,
	0x3c, 0xe1, 0x27, 0x5a, 0x57, 0xe7, 0xa2, 0x6a, 0x30, 0xeb, 0x98, 0xf4, 0x2e, 0x8c, 0x5f, 0x5b,
	0xab, 0x2d, 0x1b, 0x46, 0x64, 0x35, 0xcd, 0xda, 0x24, 0xfe, 0x45, 0x60, 0xda, 0x37, 0xa6, 0x14,
	0x46, 0xb5, 0x50, 0xc8, 0x48, 0xa0, 0x84, 0x98, 0xce, 0x61, 0x20, 0x8b, 0x30, 0x67, 0x9a, 0x0d,
	0x64, 0x41, 0x19, 0xfc, 0xd7, 0x2d, 0xd7, 0x75, 0xba, 0x4a, 0xb7, 0xea, 0x5c, 0x16, 0x96, 0x8d,
	0x5a, 0xf5, 0x36, 0xa6, 0x31, 0x1c, 0x18, 0x8b, 0x1b, 0xf9, 0xfd, 0x3d, 0xd6, 0xa5, 0xbf, 0x60,
	0xe3, 0x88, 0xac, 0xc6, 0xd9, 0x0e, 0x46, 0x1f, 0xc2, 0xc4, 0x79, 0xe1, 0x1b, 0xc7, 0xf6, 0xc3,
	0xd7, 0xcc, 0x82, 0x07, 0x9f, 0x02, 0x94, 0x75, 0xa5, 0xc7, 0x8f, 0xe0, 0xff, 0x7e, 0xcf, 0x37,
	0x42, 0xc9, 0xea, 0x92, 0xee, 0xc3, 0xe8, 0xf4, 0xec, 0xfc, 0x68, 0xb1, 0xd7, 0x45, 0xcf, 0x17,
	0x24, 0xfd, 0x4d, 0x60, 0xde, 0xf3, 0x8e, 0x4b, 0xac, 0x3d, 0x7d, 0x07, 0xc3, 0xb7, 0xe8, 0xe9,
	0x83, 0x1b, 0x9d, 0x6d, 0x8f, 0x62, 0x19, 0xdd, 0x4c, 0x68, 0x7f, 0x58, 0xbc, 0x47, 0x3f, 0xc0,
	0xe4, 0xa5, 0x45, 0xe1, 0xf1, 0xd6, 0xda, 0xbd, 0xc2, 0x0a, 0x6f, 0xa9, 0xdd, 0x09, 0xff, 0xf2,
	0xa4, 0x94, 0xfe, 0xa2, 0x59, 0xf3, 0x5c, 0xab, 0x44, 0xc9, 0xdc, 0x6a, 0xa7, 0x37, 0x3e, 0x51,
	0x3a, 0x7f, 0x1a, 0x2e, 0x23, 0xb1, 0x26, 0xdf, 0x3e, 0x2f, 0x25, 0xb6, 0xce, 0xac, 0x27, 0x01,
	0x7c, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x97, 0xac, 0x2c, 0xc1, 0x7c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IPAddressAgentClient is the client API for IPAddressAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IPAddressAgentClient interface {
	Get(ctx context.Context, in *IPAddressRequest, opts ...grpc.CallOption) (*IPAddressResponse, error)
	Create(ctx context.Context, in *IPAddressRequest, opts ...grpc.CallOption) (*IPAddressResponse, error)
	Delete(ctx context.Context, in *IPAddressRequest, opts ...grpc.CallOption) (*IPAddressResponse, error)
}

type iPAddressAgentClient struct {
	cc *grpc.ClientConn
}

func NewIPAddressAgentClient(cc *grpc.ClientConn) IPAddressAgentClient {
	return &iPAddressAgentClient{cc}
}

func (c *iPAddressAgentClient) Get(ctx context.Context, in *IPAddressRequest, opts ...grpc.CallOption) (*IPAddressResponse, error) {
	out := new(IPAddressResponse)
	err := c.cc.Invoke(ctx, "/moc.ipaddress.IPAddressAgent/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPAddressAgentClient) Create(ctx context.Context, in *IPAddressRequest, opts ...grpc.CallOption) (*IPAddressResponse, error) {
	out := new(IPAddressResponse)
	err := c.cc.Invoke(ctx, "/moc.ipaddress.IPAddressAgent/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPAddressAgentClient) Delete(ctx context.Context, in *IPAddressRequest, opts ...grpc.CallOption) (*IPAddressResponse, error) {
	out := new(IPAddressResponse)
	err := c.cc.Invoke(ctx, "/moc.ipaddress.IPAddressAgent/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IPAddressAgentServer is the server API for IPAddressAgent service.
type IPAddressAgentServer interface {
	Get(context.Context, *IPAddressRequest) (*IPAddressResponse, error)
	Create(context.Context, *IPAddressRequest) (*IPAddressResponse, error)
	Delete(context.Context, *IPAddressRequest) (*IPAddressResponse, error)
}

// UnimplementedIPAddressAgentServer can be embedded to have forward compatible implementations.
type UnimplementedIPAddressAgentServer struct {
}

func (*UnimplementedIPAddressAgentServer) Get(ctx context.Context, req *IPAddressRequest) (*IPAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedIPAddressAgentServer) Create(ctx context.Context, req *IPAddressRequest) (*IPAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedIPAddressAgentServer) Delete(ctx context.Context, req *IPAddressRequest) (*IPAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterIPAddressAgentServer(s *grpc.Server, srv IPAddressAgentServer) {
	s.RegisterService(&_IPAddressAgent_serviceDesc, srv)
}

func _IPAddressAgent_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPAddressAgentServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.ipaddress.IPAddressAgent/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPAddressAgentServer).Get(ctx, req.(*IPAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPAddressAgent_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPAddressAgentServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.ipaddress.IPAddressAgent/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPAddressAgentServer).Create(ctx, req.(*IPAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPAddressAgent_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPAddressAgentServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.ipaddress.IPAddressAgent/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPAddressAgentServer).Delete(ctx, req.(*IPAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IPAddressAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.ipaddress.IPAddressAgent",
	HandlerType: (*IPAddressAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _IPAddressAgent_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _IPAddressAgent_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _IPAddressAgent_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ipaddress/ipaddress.proto",
}
