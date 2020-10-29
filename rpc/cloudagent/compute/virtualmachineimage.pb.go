// Code generated by protoc-gen-go. DO NOT EDIT.
// source: virtualmachineimage.proto

package compute

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

type VirtualMachineImageRequest struct {
	VirtualMachineImages []*VirtualMachineImage `protobuf:"bytes,1,rep,name=VirtualMachineImages,proto3" json:"VirtualMachineImages,omitempty"`
	OperationType        common.Operation       `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *VirtualMachineImageRequest) Reset()         { *m = VirtualMachineImageRequest{} }
func (m *VirtualMachineImageRequest) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineImageRequest) ProtoMessage()    {}
func (*VirtualMachineImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8061d387a83e475c, []int{0}
}

func (m *VirtualMachineImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineImageRequest.Unmarshal(m, b)
}
func (m *VirtualMachineImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineImageRequest.Marshal(b, m, deterministic)
}
func (m *VirtualMachineImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineImageRequest.Merge(m, src)
}
func (m *VirtualMachineImageRequest) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineImageRequest.Size(m)
}
func (m *VirtualMachineImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineImageRequest proto.InternalMessageInfo

func (m *VirtualMachineImageRequest) GetVirtualMachineImages() []*VirtualMachineImage {
	if m != nil {
		return m.VirtualMachineImages
	}
	return nil
}

func (m *VirtualMachineImageRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type VirtualMachineImageResponse struct {
	VirtualMachineImages []*VirtualMachineImage `protobuf:"bytes,1,rep,name=VirtualMachineImages,proto3" json:"VirtualMachineImages,omitempty"`
	Result               *wrappers.BoolValue    `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string                 `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *VirtualMachineImageResponse) Reset()         { *m = VirtualMachineImageResponse{} }
func (m *VirtualMachineImageResponse) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineImageResponse) ProtoMessage()    {}
func (*VirtualMachineImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8061d387a83e475c, []int{1}
}

func (m *VirtualMachineImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineImageResponse.Unmarshal(m, b)
}
func (m *VirtualMachineImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineImageResponse.Marshal(b, m, deterministic)
}
func (m *VirtualMachineImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineImageResponse.Merge(m, src)
}
func (m *VirtualMachineImageResponse) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineImageResponse.Size(m)
}
func (m *VirtualMachineImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineImageResponse proto.InternalMessageInfo

func (m *VirtualMachineImageResponse) GetVirtualMachineImages() []*VirtualMachineImage {
	if m != nil {
		return m.VirtualMachineImages
	}
	return nil
}

func (m *VirtualMachineImageResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *VirtualMachineImageResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type VirtualMachineImage struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string         `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	ImageReference       string         `protobuf:"bytes,3,opt,name=imageReference,proto3" json:"imageReference,omitempty"`
	Path                 string         `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Status               *common.Status `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	ContainerName        string         `protobuf:"bytes,6,opt,name=containerName,proto3" json:"containerName,omitempty"`
	GroupName            string         `protobuf:"bytes,18,opt,name=groupName,proto3" json:"groupName,omitempty"`
	LocationName         string         `protobuf:"bytes,19,opt,name=locationName,proto3" json:"locationName,omitempty"`
	Tags                 *common.Tags   `protobuf:"bytes,20,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *VirtualMachineImage) Reset()         { *m = VirtualMachineImage{} }
func (m *VirtualMachineImage) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineImage) ProtoMessage()    {}
func (*VirtualMachineImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8061d387a83e475c, []int{2}
}

func (m *VirtualMachineImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineImage.Unmarshal(m, b)
}
func (m *VirtualMachineImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineImage.Marshal(b, m, deterministic)
}
func (m *VirtualMachineImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineImage.Merge(m, src)
}
func (m *VirtualMachineImage) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineImage.Size(m)
}
func (m *VirtualMachineImage) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineImage.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineImage proto.InternalMessageInfo

func (m *VirtualMachineImage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VirtualMachineImage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VirtualMachineImage) GetImageReference() string {
	if m != nil {
		return m.ImageReference
	}
	return ""
}

func (m *VirtualMachineImage) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *VirtualMachineImage) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VirtualMachineImage) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

func (m *VirtualMachineImage) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *VirtualMachineImage) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *VirtualMachineImage) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*VirtualMachineImageRequest)(nil), "moc.cloudagent.compute.VirtualMachineImageRequest")
	proto.RegisterType((*VirtualMachineImageResponse)(nil), "moc.cloudagent.compute.VirtualMachineImageResponse")
	proto.RegisterType((*VirtualMachineImage)(nil), "moc.cloudagent.compute.VirtualMachineImage")
}

func init() { proto.RegisterFile("virtualmachineimage.proto", fileDescriptor_8061d387a83e475c) }

var fileDescriptor_8061d387a83e475c = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x69, 0x1a, 0x29, 0x9b, 0x36, 0x87, 0x6d, 0x84, 0x8c, 0xf9, 0x51, 0x64, 0x10, 0x8a,
	0x84, 0x64, 0x0b, 0x97, 0x17, 0xa0, 0x12, 0x87, 0x1e, 0x00, 0xc9, 0x54, 0x3d, 0x70, 0x41, 0x9b,
	0xcd, 0xc4, 0x59, 0x61, 0xef, 0x2c, 0xfb, 0x53, 0xc4, 0x53, 0xf0, 0x0c, 0x3c, 0x03, 0x2f, 0xc1,
	0x63, 0x21, 0xcf, 0x06, 0xaa, 0x20, 0x5f, 0x7a, 0xe0, 0xb6, 0xfb, 0xcd, 0xb7, 0xdf, 0xf7, 0xed,
	0xcc, 0xb0, 0x07, 0x37, 0xca, 0xfa, 0x20, 0xda, 0x4e, 0xc8, 0x9d, 0xd2, 0xa0, 0x3a, 0xd1, 0x40,
	0x61, 0x2c, 0x7a, 0xe4, 0xf7, 0x3b, 0x94, 0x85, 0x6c, 0x31, 0x6c, 0x44, 0x03, 0xda, 0x17, 0x12,
	0x3b, 0x13, 0x3c, 0x64, 0x4f, 0x1a, 0xc4, 0xa6, 0x85, 0x92, 0x58, 0xeb, 0xb0, 0x2d, 0xbf, 0x5a,
	0x61, 0x0c, 0x58, 0x17, 0xdf, 0x65, 0x27, 0x12, 0xbb, 0x0e, 0x75, 0xbc, 0xe5, 0x3f, 0x13, 0x96,
	0x5d, 0x47, 0x8f, 0xb7, 0xd1, 0xe3, 0xb2, 0xf7, 0xa8, 0xe1, 0x4b, 0x00, 0xe7, 0xf9, 0x27, 0xb6,
	0x18, 0xa8, 0xba, 0x34, 0x59, 0x1e, 0xad, 0x66, 0xd5, 0x8b, 0x62, 0x38, 0x43, 0x31, 0xa4, 0x38,
	0x28, 0xc4, 0x5f, 0xb1, 0xd3, 0xf7, 0x06, 0xac, 0xf0, 0x0a, 0xf5, 0xd5, 0x37, 0x03, 0xe9, 0x68,
	0x99, 0xac, 0xe6, 0xd5, 0x9c, 0x94, 0xff, 0x56, 0xea, 0x43, 0x52, 0xfe, 0x2b, 0x61, 0x0f, 0x07,
	0x53, 0x3b, 0x83, 0xda, 0xc1, 0xff, 0x8f, 0x5d, 0xb1, 0x49, 0x0d, 0x2e, 0xb4, 0x9e, 0xf2, 0xce,
	0xaa, 0xac, 0x88, 0x5d, 0x2f, 0xfe, 0x74, 0xbd, 0xb8, 0x40, 0x6c, 0xaf, 0x45, 0x1b, 0xa0, 0xde,
	0x33, 0xf9, 0x82, 0x1d, 0xbf, 0xb1, 0x16, 0x6d, 0x7a, 0xb4, 0x4c, 0x56, 0xd3, 0x3a, 0x5e, 0xf2,
	0x1f, 0x23, 0x76, 0x36, 0x60, 0xc1, 0x39, 0x1b, 0x6b, 0xd1, 0x41, 0x9a, 0x10, 0x99, 0xce, 0x7c,
	0xce, 0x46, 0x6a, 0x43, 0x8e, 0xd3, 0x7a, 0xa4, 0x36, 0xfc, 0x39, 0x9b, 0xab, 0xf8, 0xef, 0x2d,
	0x58, 0xd0, 0x12, 0xf6, 0xd2, 0xff, 0xa0, 0xbd, 0x96, 0x11, 0x7e, 0x97, 0x8e, 0xa3, 0x56, 0x7f,
	0xe6, 0x4f, 0xd9, 0xc4, 0x79, 0xe1, 0x83, 0x4b, 0x8f, 0xe9, 0x07, 0x33, 0x6a, 0xca, 0x07, 0x82,
	0xea, 0x7d, 0x89, 0x3f, 0x63, 0xa7, 0x12, 0xb5, 0x17, 0x4a, 0x83, 0x7d, 0xd7, 0xa7, 0x99, 0x90,
	0xc2, 0x21, 0xc8, 0x1f, 0xb1, 0x69, 0x63, 0x31, 0x18, 0x62, 0x70, 0x62, 0xdc, 0x02, 0x3c, 0x67,
	0x27, 0x2d, 0x4a, 0x9a, 0x1d, 0x11, 0xce, 0x88, 0x70, 0x80, 0xf1, 0xc7, 0x6c, 0xec, 0x45, 0xe3,
	0xd2, 0x05, 0x45, 0x99, 0x52, 0x94, 0x2b, 0xd1, 0xb8, 0x9a, 0xe0, 0xea, 0x7b, 0xc2, 0xd2, 0x81,
	0x1e, 0xbd, 0xee, 0x87, 0xc7, 0x1d, 0x9b, 0x5c, 0xea, 0x1b, 0xfc, 0x0c, 0xbc, 0xba, 0xcb, 0x5c,
	0xe3, 0x82, 0x67, 0xe7, 0x77, 0x7a, 0x13, 0xd7, 0x2b, 0xbf, 0x77, 0xf1, 0xf2, 0x63, 0xd9, 0x28,
	0xbf, 0x0b, 0xeb, 0x9e, 0x5f, 0x76, 0x4a, 0x5a, 0x74, 0xb8, 0xf5, 0x65, 0x87, 0xb2, 0xb4, 0x46,
	0x96, 0xb7, 0x82, 0xe5, 0x5e, 0x70, 0x3d, 0xa1, 0xd5, 0x38, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff,
	0x57, 0x72, 0x0e, 0xac, 0xd3, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VirtualMachineImageAgentClient is the client API for VirtualMachineImageAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VirtualMachineImageAgentClient interface {
	Invoke(ctx context.Context, in *VirtualMachineImageRequest, opts ...grpc.CallOption) (*VirtualMachineImageResponse, error)
}

type virtualMachineImageAgentClient struct {
	cc *grpc.ClientConn
}

func NewVirtualMachineImageAgentClient(cc *grpc.ClientConn) VirtualMachineImageAgentClient {
	return &virtualMachineImageAgentClient{cc}
}

func (c *virtualMachineImageAgentClient) Invoke(ctx context.Context, in *VirtualMachineImageRequest, opts ...grpc.CallOption) (*VirtualMachineImageResponse, error) {
	out := new(VirtualMachineImageResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.compute.VirtualMachineImageAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtualMachineImageAgentServer is the server API for VirtualMachineImageAgent service.
type VirtualMachineImageAgentServer interface {
	Invoke(context.Context, *VirtualMachineImageRequest) (*VirtualMachineImageResponse, error)
}

// UnimplementedVirtualMachineImageAgentServer can be embedded to have forward compatible implementations.
type UnimplementedVirtualMachineImageAgentServer struct {
}

func (*UnimplementedVirtualMachineImageAgentServer) Invoke(ctx context.Context, req *VirtualMachineImageRequest) (*VirtualMachineImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterVirtualMachineImageAgentServer(s *grpc.Server, srv VirtualMachineImageAgentServer) {
	s.RegisterService(&_VirtualMachineImageAgent_serviceDesc, srv)
}

func _VirtualMachineImageAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualMachineImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineImageAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.compute.VirtualMachineImageAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineImageAgentServer).Invoke(ctx, req.(*VirtualMachineImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VirtualMachineImageAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.compute.VirtualMachineImageAgent",
	HandlerType: (*VirtualMachineImageAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _VirtualMachineImageAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "virtualmachineimage.proto",
}
