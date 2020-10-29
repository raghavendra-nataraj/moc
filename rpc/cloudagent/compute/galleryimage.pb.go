// Code generated by protoc-gen-go. DO NOT EDIT.
// source: galleryimage.proto

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

type GalleryImageOSType int32

const (
	GalleryImageOSType_UNKNOWN GalleryImageOSType = 0
	GalleryImageOSType_LINUX   GalleryImageOSType = 1
	GalleryImageOSType_WINDOWS GalleryImageOSType = 2
)

var GalleryImageOSType_name = map[int32]string{
	0: "UNKNOWN",
	1: "LINUX",
	2: "WINDOWS",
}

var GalleryImageOSType_value = map[string]int32{
	"UNKNOWN": 0,
	"LINUX":   1,
	"WINDOWS": 2,
}

func (x GalleryImageOSType) String() string {
	return proto.EnumName(GalleryImageOSType_name, int32(x))
}

func (GalleryImageOSType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7523a8dbee37a80a, []int{0}
}

type GalleryImageRequest struct {
	GalleryImages        []*GalleryImage  `protobuf:"bytes,1,rep,name=GalleryImages,proto3" json:"GalleryImages,omitempty"`
	OperationType        common.Operation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GalleryImageRequest) Reset()         { *m = GalleryImageRequest{} }
func (m *GalleryImageRequest) String() string { return proto.CompactTextString(m) }
func (*GalleryImageRequest) ProtoMessage()    {}
func (*GalleryImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7523a8dbee37a80a, []int{0}
}

func (m *GalleryImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GalleryImageRequest.Unmarshal(m, b)
}
func (m *GalleryImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GalleryImageRequest.Marshal(b, m, deterministic)
}
func (m *GalleryImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GalleryImageRequest.Merge(m, src)
}
func (m *GalleryImageRequest) XXX_Size() int {
	return xxx_messageInfo_GalleryImageRequest.Size(m)
}
func (m *GalleryImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GalleryImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GalleryImageRequest proto.InternalMessageInfo

func (m *GalleryImageRequest) GetGalleryImages() []*GalleryImage {
	if m != nil {
		return m.GalleryImages
	}
	return nil
}

func (m *GalleryImageRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type GalleryImageResponse struct {
	GalleryImages        []*GalleryImage     `protobuf:"bytes,1,rep,name=GalleryImages,proto3" json:"GalleryImages,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GalleryImageResponse) Reset()         { *m = GalleryImageResponse{} }
func (m *GalleryImageResponse) String() string { return proto.CompactTextString(m) }
func (*GalleryImageResponse) ProtoMessage()    {}
func (*GalleryImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7523a8dbee37a80a, []int{1}
}

func (m *GalleryImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GalleryImageResponse.Unmarshal(m, b)
}
func (m *GalleryImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GalleryImageResponse.Marshal(b, m, deterministic)
}
func (m *GalleryImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GalleryImageResponse.Merge(m, src)
}
func (m *GalleryImageResponse) XXX_Size() int {
	return xxx_messageInfo_GalleryImageResponse.Size(m)
}
func (m *GalleryImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GalleryImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GalleryImageResponse proto.InternalMessageInfo

func (m *GalleryImageResponse) GetGalleryImages() []*GalleryImage {
	if m != nil {
		return m.GalleryImages
	}
	return nil
}

func (m *GalleryImageResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GalleryImageResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type GalleryImage struct {
	Name        string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id          string             `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	ImageOSType GalleryImageOSType `protobuf:"varint,3,opt,name=imageOSType,proto3,enum=moc.cloudagent.compute.GalleryImageOSType" json:"imageOSType,omitempty"`
	// Path of the image on the cloud
	Path          string         `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Status        *common.Status `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	LocationName  string         `protobuf:"bytes,6,opt,name=locationName,proto3" json:"locationName,omitempty"`
	Publisher     string         `protobuf:"bytes,7,opt,name=publisher,proto3" json:"publisher,omitempty"`
	Sku           string         `protobuf:"bytes,8,opt,name=sku,proto3" json:"sku,omitempty"`
	Offer         string         `protobuf:"bytes,9,opt,name=offer,proto3" json:"offer,omitempty"`
	ContainerName string         `protobuf:"bytes,10,opt,name=containerName,proto3" json:"containerName,omitempty"`
	// Source of the GalleryImage from where we can copy the image from.
	SourcePath           string       `protobuf:"bytes,11,opt,name=sourcePath,proto3" json:"sourcePath,omitempty"`
	Tags                 *common.Tags `protobuf:"bytes,12,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GalleryImage) Reset()         { *m = GalleryImage{} }
func (m *GalleryImage) String() string { return proto.CompactTextString(m) }
func (*GalleryImage) ProtoMessage()    {}
func (*GalleryImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7523a8dbee37a80a, []int{2}
}

func (m *GalleryImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GalleryImage.Unmarshal(m, b)
}
func (m *GalleryImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GalleryImage.Marshal(b, m, deterministic)
}
func (m *GalleryImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GalleryImage.Merge(m, src)
}
func (m *GalleryImage) XXX_Size() int {
	return xxx_messageInfo_GalleryImage.Size(m)
}
func (m *GalleryImage) XXX_DiscardUnknown() {
	xxx_messageInfo_GalleryImage.DiscardUnknown(m)
}

var xxx_messageInfo_GalleryImage proto.InternalMessageInfo

func (m *GalleryImage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GalleryImage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GalleryImage) GetImageOSType() GalleryImageOSType {
	if m != nil {
		return m.ImageOSType
	}
	return GalleryImageOSType_UNKNOWN
}

func (m *GalleryImage) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *GalleryImage) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GalleryImage) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *GalleryImage) GetPublisher() string {
	if m != nil {
		return m.Publisher
	}
	return ""
}

func (m *GalleryImage) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *GalleryImage) GetOffer() string {
	if m != nil {
		return m.Offer
	}
	return ""
}

func (m *GalleryImage) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

func (m *GalleryImage) GetSourcePath() string {
	if m != nil {
		return m.SourcePath
	}
	return ""
}

func (m *GalleryImage) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterEnum("moc.cloudagent.compute.GalleryImageOSType", GalleryImageOSType_name, GalleryImageOSType_value)
	proto.RegisterType((*GalleryImageRequest)(nil), "moc.cloudagent.compute.GalleryImageRequest")
	proto.RegisterType((*GalleryImageResponse)(nil), "moc.cloudagent.compute.GalleryImageResponse")
	proto.RegisterType((*GalleryImage)(nil), "moc.cloudagent.compute.GalleryImage")
}

func init() { proto.RegisterFile("galleryimage.proto", fileDescriptor_7523a8dbee37a80a) }

var fileDescriptor_7523a8dbee37a80a = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0x93, 0x26, 0xc5, 0xe3, 0x24, 0x0a, 0x4b, 0x85, 0xac, 0x08, 0xaa, 0x28, 0xf4, 0x10,
	0x15, 0x64, 0x8b, 0xc0, 0x85, 0x23, 0x15, 0x08, 0x05, 0x2a, 0x07, 0x39, 0x2d, 0x41, 0xdc, 0x36,
	0xce, 0xc6, 0xb1, 0x6a, 0x7b, 0xcc, 0xfe, 0x80, 0xca, 0x13, 0xf0, 0x06, 0xbc, 0x04, 0x0f, 0x89,
	0x3c, 0x0e, 0x34, 0x16, 0x1c, 0x72, 0xe0, 0xb6, 0xf3, 0xcd, 0x37, 0xdf, 0xcc, 0x7c, 0x1e, 0x03,
	0x8b, 0x79, 0x9a, 0x0a, 0x79, 0x93, 0x64, 0x3c, 0x16, 0x5e, 0x21, 0x51, 0x23, 0xbb, 0x9f, 0x61,
	0xe4, 0x45, 0x29, 0x9a, 0x15, 0x8f, 0x45, 0xae, 0xbd, 0x08, 0xb3, 0xc2, 0x68, 0x31, 0x38, 0x89,
	0x11, 0xe3, 0x54, 0xf8, 0xc4, 0x5a, 0x9a, 0xb5, 0xff, 0x55, 0xf2, 0xa2, 0x10, 0x52, 0x55, 0x75,
	0x83, 0x4e, 0x84, 0x59, 0x86, 0x79, 0x15, 0x8d, 0x7e, 0x58, 0x70, 0xef, 0x4d, 0x25, 0x3e, 0x2d,
	0xc5, 0x43, 0xf1, 0xd9, 0x08, 0xa5, 0xd9, 0x5b, 0xe8, 0xee, 0xc2, 0xca, 0xb5, 0x86, 0xcd, 0xb1,
	0x33, 0x39, 0xf5, 0xfe, 0xdd, 0xd5, 0xab, 0x69, 0xd4, 0x4b, 0xd9, 0x73, 0xe8, 0xce, 0x0a, 0x21,
	0xb9, 0x4e, 0x30, 0xbf, 0xbc, 0x29, 0x84, 0xdb, 0x18, 0x5a, 0xe3, 0xde, 0xa4, 0x47, 0x5a, 0x7f,
	0x32, 0x61, 0x9d, 0x34, 0xfa, 0x69, 0xc1, 0x71, 0x7d, 0x32, 0x55, 0x60, 0xae, 0xc4, 0x7f, 0x1d,
	0x6d, 0x02, 0xed, 0x50, 0x28, 0x93, 0x6a, 0x9a, 0xc9, 0x99, 0x0c, 0xbc, 0xca, 0x3d, 0xef, 0xb7,
	0x7b, 0xde, 0x39, 0x62, 0xfa, 0x81, 0xa7, 0x46, 0x84, 0x5b, 0x26, 0x3b, 0x86, 0xd6, 0x6b, 0x29,
	0x51, 0xba, 0xcd, 0xa1, 0x35, 0xb6, 0xc3, 0x2a, 0x18, 0x7d, 0x6f, 0x42, 0x67, 0x57, 0x9b, 0x31,
	0x38, 0xcc, 0x79, 0x26, 0x5c, 0x8b, 0x58, 0xf4, 0x66, 0x3d, 0x68, 0x24, 0x2b, 0x6a, 0x65, 0x87,
	0x8d, 0x64, 0xc5, 0x2e, 0xc0, 0xa1, 0x4f, 0x3a, 0x9b, 0x93, 0x2f, 0x4d, 0xf2, 0xe5, 0x6c, 0x9f,
	0x45, 0xaa, 0x8a, 0x70, 0xb7, 0xbc, 0xec, 0x58, 0x70, 0xbd, 0x71, 0x0f, 0xab, 0x8e, 0xe5, 0x9b,
	0x3d, 0x82, 0xb6, 0xd2, 0x5c, 0x1b, 0xe5, 0xb6, 0x68, 0x41, 0x87, 0xc4, 0xe7, 0x04, 0x85, 0xdb,
	0x14, 0x1b, 0x41, 0x27, 0xc5, 0x88, 0xac, 0x0f, 0xca, 0x91, 0xdb, 0x24, 0x50, 0xc3, 0xd8, 0x03,
	0xb0, 0x0b, 0xb3, 0x4c, 0x13, 0xb5, 0x11, 0xd2, 0x3d, 0x22, 0xc2, 0x2d, 0xc0, 0xfa, 0xd0, 0x54,
	0xd7, 0xc6, 0xbd, 0x43, 0x78, 0xf9, 0x2c, 0x5d, 0xc2, 0xf5, 0x5a, 0x48, 0xd7, 0xae, 0x5c, 0xa2,
	0x80, 0x9d, 0x42, 0x37, 0xc2, 0x5c, 0xf3, 0x24, 0x17, 0x92, 0x5a, 0x01, 0x65, 0xeb, 0x20, 0x3b,
	0x01, 0x50, 0x68, 0x64, 0x24, 0xde, 0x97, 0xeb, 0x38, 0x44, 0xd9, 0x41, 0xd8, 0x43, 0x38, 0xd4,
	0x3c, 0x56, 0x6e, 0x87, 0x56, 0xb2, 0x69, 0xa5, 0x4b, 0x1e, 0xab, 0x90, 0xe0, 0xb3, 0x17, 0xc0,
	0xfe, 0xb6, 0x8a, 0x39, 0x70, 0x74, 0x15, 0xbc, 0x0b, 0x66, 0x8b, 0xa0, 0x7f, 0xc0, 0x6c, 0x68,
	0x5d, 0x4c, 0x83, 0xab, 0x8f, 0x7d, 0xab, 0xc4, 0x17, 0xd3, 0xe0, 0xd5, 0x6c, 0x31, 0xef, 0x37,
	0x26, 0xdf, 0xe0, 0xee, 0x6e, 0xe9, 0xcb, 0xd2, 0x7f, 0x26, 0xa0, 0x3d, 0xcd, 0xbf, 0xe0, 0xb5,
	0x60, 0x8f, 0xf7, 0xba, 0xb1, 0xea, 0x17, 0x1a, 0x3c, 0xd9, 0x8f, 0x5c, 0x5d, 0xf5, 0xe8, 0xe0,
	0xfc, 0xe9, 0x27, 0x3f, 0x4e, 0xf4, 0xc6, 0x2c, 0x4b, 0xa2, 0x9f, 0x25, 0x91, 0x44, 0x85, 0x6b,
	0xed, 0x67, 0x18, 0xf9, 0xb2, 0x88, 0xfc, 0x5b, 0x25, 0x7f, 0xab, 0xb4, 0x6c, 0xd3, 0x99, 0x3e,
	0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0x63, 0xbe, 0x62, 0xeb, 0x20, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GalleryImageAgentClient is the client API for GalleryImageAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GalleryImageAgentClient interface {
	Invoke(ctx context.Context, in *GalleryImageRequest, opts ...grpc.CallOption) (*GalleryImageResponse, error)
}

type galleryImageAgentClient struct {
	cc *grpc.ClientConn
}

func NewGalleryImageAgentClient(cc *grpc.ClientConn) GalleryImageAgentClient {
	return &galleryImageAgentClient{cc}
}

func (c *galleryImageAgentClient) Invoke(ctx context.Context, in *GalleryImageRequest, opts ...grpc.CallOption) (*GalleryImageResponse, error) {
	out := new(GalleryImageResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.compute.GalleryImageAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GalleryImageAgentServer is the server API for GalleryImageAgent service.
type GalleryImageAgentServer interface {
	Invoke(context.Context, *GalleryImageRequest) (*GalleryImageResponse, error)
}

// UnimplementedGalleryImageAgentServer can be embedded to have forward compatible implementations.
type UnimplementedGalleryImageAgentServer struct {
}

func (*UnimplementedGalleryImageAgentServer) Invoke(ctx context.Context, req *GalleryImageRequest) (*GalleryImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterGalleryImageAgentServer(s *grpc.Server, srv GalleryImageAgentServer) {
	s.RegisterService(&_GalleryImageAgent_serviceDesc, srv)
}

func _GalleryImageAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GalleryImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GalleryImageAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.compute.GalleryImageAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GalleryImageAgentServer).Invoke(ctx, req.(*GalleryImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GalleryImageAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.compute.GalleryImageAgent",
	HandlerType: (*GalleryImageAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _GalleryImageAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "galleryimage.proto",
}
