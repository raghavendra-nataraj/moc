// Code generated by protoc-gen-go. DO NOT EDIT.
// source: virtualharddisk.proto

package storage

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type VirtualHardDiskType int32

const (
	VirtualHardDiskType_OS_VIRTUALHARDDISK       VirtualHardDiskType = 0
	VirtualHardDiskType_DATADISK_VIRTUALHARDDISK VirtualHardDiskType = 1
)

var VirtualHardDiskType_name = map[int32]string{
	0: "OS_VIRTUALHARDDISK",
	1: "DATADISK_VIRTUALHARDDISK",
}

var VirtualHardDiskType_value = map[string]int32{
	"OS_VIRTUALHARDDISK":       0,
	"DATADISK_VIRTUALHARDDISK": 1,
}

func (x VirtualHardDiskType) String() string {
	return proto.EnumName(VirtualHardDiskType_name, int32(x))
}

func (VirtualHardDiskType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f4b382f86170a6e5, []int{0}
}

type VirtualHardDiskRequest struct {
	VirtualHardDiskSystems []*VirtualHardDisk `protobuf:"bytes,1,rep,name=VirtualHardDiskSystems,proto3" json:"VirtualHardDiskSystems,omitempty"`
	OperationType          common.Operation   `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}           `json:"-"`
	XXX_unrecognized       []byte             `json:"-"`
	XXX_sizecache          int32              `json:"-"`
}

func (m *VirtualHardDiskRequest) Reset()         { *m = VirtualHardDiskRequest{} }
func (m *VirtualHardDiskRequest) String() string { return proto.CompactTextString(m) }
func (*VirtualHardDiskRequest) ProtoMessage()    {}
func (*VirtualHardDiskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b382f86170a6e5, []int{0}
}

func (m *VirtualHardDiskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualHardDiskRequest.Unmarshal(m, b)
}
func (m *VirtualHardDiskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualHardDiskRequest.Marshal(b, m, deterministic)
}
func (m *VirtualHardDiskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualHardDiskRequest.Merge(m, src)
}
func (m *VirtualHardDiskRequest) XXX_Size() int {
	return xxx_messageInfo_VirtualHardDiskRequest.Size(m)
}
func (m *VirtualHardDiskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualHardDiskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualHardDiskRequest proto.InternalMessageInfo

func (m *VirtualHardDiskRequest) GetVirtualHardDiskSystems() []*VirtualHardDisk {
	if m != nil {
		return m.VirtualHardDiskSystems
	}
	return nil
}

func (m *VirtualHardDiskRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type VirtualHardDiskResponse struct {
	VirtualHardDiskSystems []*VirtualHardDisk  `protobuf:"bytes,1,rep,name=VirtualHardDiskSystems,proto3" json:"VirtualHardDiskSystems,omitempty"`
	Result                 *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                  string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}            `json:"-"`
	XXX_unrecognized       []byte              `json:"-"`
	XXX_sizecache          int32               `json:"-"`
}

func (m *VirtualHardDiskResponse) Reset()         { *m = VirtualHardDiskResponse{} }
func (m *VirtualHardDiskResponse) String() string { return proto.CompactTextString(m) }
func (*VirtualHardDiskResponse) ProtoMessage()    {}
func (*VirtualHardDiskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b382f86170a6e5, []int{1}
}

func (m *VirtualHardDiskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualHardDiskResponse.Unmarshal(m, b)
}
func (m *VirtualHardDiskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualHardDiskResponse.Marshal(b, m, deterministic)
}
func (m *VirtualHardDiskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualHardDiskResponse.Merge(m, src)
}
func (m *VirtualHardDiskResponse) XXX_Size() int {
	return xxx_messageInfo_VirtualHardDiskResponse.Size(m)
}
func (m *VirtualHardDiskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualHardDiskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualHardDiskResponse proto.InternalMessageInfo

func (m *VirtualHardDiskResponse) GetVirtualHardDiskSystems() []*VirtualHardDisk {
	if m != nil {
		return m.VirtualHardDiskSystems
	}
	return nil
}

func (m *VirtualHardDiskResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *VirtualHardDiskResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type VirtualHardDisk struct {
	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Path   string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	// Storage container name to hold this vhd
	ContainerName        string              `protobuf:"bytes,5,opt,name=containerName,proto3" json:"containerName,omitempty"`
	Status               *common.Status      `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Size                 int64               `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	Dynamic              bool                `protobuf:"varint,8,opt,name=dynamic,proto3" json:"dynamic,omitempty"`
	Blocksizebytes       int32               `protobuf:"varint,9,opt,name=blocksizebytes,proto3" json:"blocksizebytes,omitempty"`
	Logicalsectorbytes   int32               `protobuf:"varint,10,opt,name=logicalsectorbytes,proto3" json:"logicalsectorbytes,omitempty"`
	Physicalsectorbytes  int32               `protobuf:"varint,11,opt,name=physicalsectorbytes,proto3" json:"physicalsectorbytes,omitempty"`
	Controllernumber     int64               `protobuf:"varint,12,opt,name=controllernumber,proto3" json:"controllernumber,omitempty"`
	Controllerlocation   int64               `protobuf:"varint,13,opt,name=controllerlocation,proto3" json:"controllerlocation,omitempty"`
	Disknumber           int64               `protobuf:"varint,14,opt,name=disknumber,proto3" json:"disknumber,omitempty"`
	VirtualmachineName   string              `protobuf:"bytes,15,opt,name=virtualmachineName,proto3" json:"virtualmachineName,omitempty"`
	Scsipath             string              `protobuf:"bytes,16,opt,name=scsipath,proto3" json:"scsipath,omitempty"`
	Virtualharddisktype  VirtualHardDiskType `protobuf:"varint,17,opt,name=virtualharddisktype,proto3,enum=moc.nodeagent.storage.VirtualHardDiskType" json:"virtualharddisktype,omitempty"`
	Entity               *common.Entity      `protobuf:"bytes,18,opt,name=entity,proto3" json:"entity,omitempty"`
	Tags                 *common.Tags        `protobuf:"bytes,19,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *VirtualHardDisk) Reset()         { *m = VirtualHardDisk{} }
func (m *VirtualHardDisk) String() string { return proto.CompactTextString(m) }
func (*VirtualHardDisk) ProtoMessage()    {}
func (*VirtualHardDisk) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b382f86170a6e5, []int{2}
}

func (m *VirtualHardDisk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualHardDisk.Unmarshal(m, b)
}
func (m *VirtualHardDisk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualHardDisk.Marshal(b, m, deterministic)
}
func (m *VirtualHardDisk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualHardDisk.Merge(m, src)
}
func (m *VirtualHardDisk) XXX_Size() int {
	return xxx_messageInfo_VirtualHardDisk.Size(m)
}
func (m *VirtualHardDisk) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualHardDisk.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualHardDisk proto.InternalMessageInfo

func (m *VirtualHardDisk) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VirtualHardDisk) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VirtualHardDisk) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *VirtualHardDisk) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *VirtualHardDisk) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

func (m *VirtualHardDisk) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VirtualHardDisk) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *VirtualHardDisk) GetDynamic() bool {
	if m != nil {
		return m.Dynamic
	}
	return false
}

func (m *VirtualHardDisk) GetBlocksizebytes() int32 {
	if m != nil {
		return m.Blocksizebytes
	}
	return 0
}

func (m *VirtualHardDisk) GetLogicalsectorbytes() int32 {
	if m != nil {
		return m.Logicalsectorbytes
	}
	return 0
}

func (m *VirtualHardDisk) GetPhysicalsectorbytes() int32 {
	if m != nil {
		return m.Physicalsectorbytes
	}
	return 0
}

func (m *VirtualHardDisk) GetControllernumber() int64 {
	if m != nil {
		return m.Controllernumber
	}
	return 0
}

func (m *VirtualHardDisk) GetControllerlocation() int64 {
	if m != nil {
		return m.Controllerlocation
	}
	return 0
}

func (m *VirtualHardDisk) GetDisknumber() int64 {
	if m != nil {
		return m.Disknumber
	}
	return 0
}

func (m *VirtualHardDisk) GetVirtualmachineName() string {
	if m != nil {
		return m.VirtualmachineName
	}
	return ""
}

func (m *VirtualHardDisk) GetScsipath() string {
	if m != nil {
		return m.Scsipath
	}
	return ""
}

func (m *VirtualHardDisk) GetVirtualharddisktype() VirtualHardDiskType {
	if m != nil {
		return m.Virtualharddisktype
	}
	return VirtualHardDiskType_OS_VIRTUALHARDDISK
}

func (m *VirtualHardDisk) GetEntity() *common.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *VirtualHardDisk) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterEnum("moc.nodeagent.storage.VirtualHardDiskType", VirtualHardDiskType_name, VirtualHardDiskType_value)
	proto.RegisterType((*VirtualHardDiskRequest)(nil), "moc.nodeagent.storage.VirtualHardDiskRequest")
	proto.RegisterType((*VirtualHardDiskResponse)(nil), "moc.nodeagent.storage.VirtualHardDiskResponse")
	proto.RegisterType((*VirtualHardDisk)(nil), "moc.nodeagent.storage.VirtualHardDisk")
}

func init() { proto.RegisterFile("virtualharddisk.proto", fileDescriptor_f4b382f86170a6e5) }

var fileDescriptor_f4b382f86170a6e5 = []byte{
	// 701 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xdd, 0x4e, 0x1b, 0x39,
	0x14, 0x66, 0x80, 0x04, 0xe2, 0x40, 0x80, 0x13, 0xc8, 0x7a, 0xb3, 0xbb, 0x28, 0xca, 0xae, 0x50,
	0x84, 0xb4, 0x13, 0x94, 0xdd, 0x17, 0x08, 0x0d, 0x12, 0x88, 0x0a, 0x24, 0x27, 0xe5, 0xa2, 0xaa,
	0x5a, 0x39, 0x8e, 0x99, 0x58, 0x99, 0x19, 0x4f, 0x6d, 0x0f, 0x55, 0xfa, 0x50, 0x7d, 0x89, 0x5e,
	0x56, 0xea, 0x33, 0x55, 0xf6, 0x0c, 0x21, 0x24, 0xa9, 0xc4, 0x4d, 0xef, 0x7c, 0xce, 0xf7, 0x9d,
	0xcf, 0x9f, 0x7f, 0xce, 0x41, 0x47, 0x0f, 0x42, 0x99, 0x94, 0x86, 0x63, 0xaa, 0x46, 0x23, 0xa1,
	0x27, 0x7e, 0xa2, 0xa4, 0x91, 0x70, 0x14, 0x49, 0xe6, 0xc7, 0x72, 0xc4, 0x69, 0xc0, 0x63, 0xe3,
	0x6b, 0x23, 0x15, 0x0d, 0x78, 0xfd, 0x38, 0x90, 0x32, 0x08, 0x79, 0xdb, 0x91, 0x86, 0xe9, 0x7d,
	0xfb, 0x93, 0xa2, 0x49, 0xc2, 0x95, 0xce, 0xca, 0xea, 0x3b, 0x4c, 0x46, 0x91, 0x8c, 0xf3, 0x08,
	0x62, 0x69, 0xc4, 0xbd, 0x60, 0xd4, 0x88, 0x59, 0xee, 0x8f, 0x45, 0x05, 0x1e, 0x25, 0x66, 0x9a,
	0x81, 0xcd, 0x2f, 0x1e, 0xaa, 0xdd, 0x65, 0x7e, 0x2e, 0xa9, 0x1a, 0xf5, 0x84, 0x9e, 0x10, 0xfe,
	0x31, 0xe5, 0xda, 0xc0, 0xfb, 0x25, 0xa4, 0x3f, 0xd5, 0x86, 0x47, 0x1a, 0x7b, 0x8d, 0x8d, 0x56,
	0xb9, 0x73, 0xe2, 0xaf, 0x74, 0xec, 0x2f, 0xca, 0xfd, 0x44, 0x05, 0xfe, 0x47, 0xbb, 0xb7, 0x09,
	0x57, 0xce, 0xea, 0x60, 0x9a, 0x70, 0xbc, 0xde, 0xf0, 0x5a, 0x95, 0x4e, 0xc5, 0xc9, 0xce, 0x10,
	0xf2, 0x9c, 0xd4, 0xfc, 0xea, 0xa1, 0xdf, 0x96, 0x0c, 0xeb, 0x44, 0xc6, 0x9a, 0xff, 0x72, 0xc7,
	0x1d, 0x54, 0x24, 0x5c, 0xa7, 0xa1, 0x71, 0x56, 0xcb, 0x9d, 0xba, 0x9f, 0x5d, 0xad, 0xff, 0x78,
	0xb5, 0xfe, 0xb9, 0x94, 0xe1, 0x1d, 0x0d, 0x53, 0x4e, 0x72, 0x26, 0x1c, 0xa2, 0xc2, 0x85, 0x52,
	0x52, 0xe1, 0x8d, 0x86, 0xd7, 0x2a, 0x91, 0x2c, 0x68, 0x7e, 0x2f, 0xa0, 0xbd, 0x85, 0x4d, 0x00,
	0xd0, 0x66, 0x4c, 0x23, 0x8e, 0x3d, 0x47, 0x74, 0x6b, 0xa8, 0xa0, 0x75, 0x31, 0x72, 0xbb, 0x95,
	0xc8, 0xba, 0x18, 0x41, 0x0d, 0x15, 0xb5, 0x4c, 0x15, 0xe3, 0xb9, 0x5c, 0x1e, 0xd9, 0xda, 0x84,
	0x9a, 0x31, 0xde, 0xcc, 0x6a, 0xed, 0x1a, 0xfe, 0x41, 0xbb, 0x4c, 0xc6, 0x86, 0x8a, 0x98, 0xab,
	0x1b, 0x2b, 0x5c, 0x70, 0xe0, 0xf3, 0x24, 0xfc, 0x8d, 0x8a, 0xda, 0x50, 0x93, 0x6a, 0x5c, 0x74,
	0x67, 0x2a, 0xbb, 0x3b, 0xea, 0xbb, 0x14, 0xc9, 0x21, 0x2b, 0xaf, 0xc5, 0x67, 0x8e, 0xb7, 0x1a,
	0x5e, 0x6b, 0x83, 0xb8, 0x35, 0x60, 0xb4, 0x35, 0x9a, 0xc6, 0x34, 0x12, 0x0c, 0x6f, 0x37, 0xbc,
	0xd6, 0x36, 0x79, 0x0c, 0xe1, 0x04, 0x55, 0x86, 0xa1, 0x64, 0x13, 0x4b, 0x1b, 0x4e, 0x0d, 0xd7,
	0xb8, 0xd4, 0xf0, 0x5a, 0x05, 0xb2, 0x90, 0x05, 0x1f, 0x41, 0x28, 0x03, 0xc1, 0x68, 0xa8, 0x39,
	0x33, 0x52, 0x65, 0x5c, 0xe4, 0xb8, 0x2b, 0x10, 0x38, 0x43, 0xd5, 0x64, 0x3c, 0xd5, 0x8b, 0x05,
	0x65, 0x57, 0xb0, 0x0a, 0x82, 0x53, 0xb4, 0x6f, 0x4f, 0xab, 0x64, 0x18, 0x72, 0x15, 0xa7, 0xd1,
	0x90, 0x2b, 0xbc, 0xe3, 0xce, 0xb0, 0x94, 0xb7, 0x6e, 0x9e, 0x72, 0xa1, 0xcc, 0x5a, 0x08, 0xef,
	0x3a, 0xf6, 0x0a, 0x04, 0x8e, 0x11, 0xb2, 0xdd, 0x9b, 0xab, 0x56, 0x1c, 0x6f, 0x2e, 0x63, 0xf5,
	0xf2, 0x46, 0x8f, 0x28, 0x1b, 0x8b, 0x98, 0xbb, 0x37, 0xd8, 0x73, 0x6f, 0xb0, 0x02, 0x81, 0x3a,
	0xda, 0xd6, 0x4c, 0x0b, 0xf7, 0x8c, 0xfb, 0x8e, 0x35, 0x8b, 0xe1, 0x1d, 0xaa, 0x2e, 0x0c, 0x0d,
	0x63, 0x1b, 0xe6, 0xc0, 0x35, 0xcc, 0xe9, 0xcb, 0x7e, 0xb5, 0xed, 0x1e, 0xb2, 0x4a, 0xc6, 0x7e,
	0x01, 0x1e, 0x1b, 0x61, 0xa6, 0x18, 0xe6, 0xbe, 0xc0, 0x85, 0x4b, 0x91, 0x1c, 0x82, 0xbf, 0xd0,
	0xa6, 0xa1, 0x81, 0xc6, 0x55, 0x47, 0x29, 0x39, 0xca, 0x80, 0x06, 0x9a, 0xb8, 0xf4, 0xe9, 0x35,
	0xaa, 0xae, 0xd8, 0x0f, 0x6a, 0x08, 0x6e, 0xfb, 0x1f, 0xee, 0xae, 0xc8, 0xe0, 0x4d, 0xf7, 0xf5,
	0x65, 0x97, 0xf4, 0x7a, 0x57, 0xfd, 0xeb, 0xfd, 0x35, 0xf8, 0x13, 0xe1, 0x5e, 0x77, 0xd0, 0xb5,
	0xd1, 0x12, 0xea, 0x75, 0xbe, 0x79, 0xe8, 0x70, 0x41, 0xad, 0x6b, 0x8f, 0x06, 0x02, 0x15, 0xaf,
	0xe2, 0x07, 0x39, 0xe1, 0xf0, 0xef, 0x0b, 0x5b, 0x39, 0x9b, 0x65, 0x75, 0xff, 0xa5, 0xf4, 0x6c,
	0x92, 0x34, 0xd7, 0xe0, 0x12, 0x1d, 0xbc, 0x1a, 0x73, 0x36, 0xb9, 0x99, 0x1b, 0xa8, 0x50, 0x5b,
	0x6a, 0xf8, 0x0b, 0x3b, 0x4b, 0xeb, 0xbf, 0x3b, 0xf9, 0x79, 0xea, 0x93, 0xd2, 0xf9, 0xd9, 0x5b,
	0x3f, 0x10, 0x66, 0x9c, 0x0e, 0x7d, 0x26, 0xa3, 0x76, 0x24, 0x98, 0x92, 0x5a, 0xde, 0x9b, 0x76,
	0x24, 0x59, 0x5b, 0x25, 0xac, 0x3d, 0x73, 0xd5, 0xce, 0x5d, 0x0d, 0x8b, 0x4e, 0xfe, 0xbf, 0x1f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x52, 0x30, 0xe0, 0x0b, 0x2a, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VirtualHardDiskAgentClient is the client API for VirtualHardDiskAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VirtualHardDiskAgentClient interface {
	Invoke(ctx context.Context, in *VirtualHardDiskRequest, opts ...grpc.CallOption) (*VirtualHardDiskResponse, error)
	CheckNotification(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.NotificationResponse, error)
}

type virtualHardDiskAgentClient struct {
	cc *grpc.ClientConn
}

func NewVirtualHardDiskAgentClient(cc *grpc.ClientConn) VirtualHardDiskAgentClient {
	return &virtualHardDiskAgentClient{cc}
}

func (c *virtualHardDiskAgentClient) Invoke(ctx context.Context, in *VirtualHardDiskRequest, opts ...grpc.CallOption) (*VirtualHardDiskResponse, error) {
	out := new(VirtualHardDiskResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.storage.VirtualHardDiskAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualHardDiskAgentClient) CheckNotification(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.NotificationResponse, error) {
	out := new(common.NotificationResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.storage.VirtualHardDiskAgent/CheckNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtualHardDiskAgentServer is the server API for VirtualHardDiskAgent service.
type VirtualHardDiskAgentServer interface {
	Invoke(context.Context, *VirtualHardDiskRequest) (*VirtualHardDiskResponse, error)
	CheckNotification(context.Context, *empty.Empty) (*common.NotificationResponse, error)
}

// UnimplementedVirtualHardDiskAgentServer can be embedded to have forward compatible implementations.
type UnimplementedVirtualHardDiskAgentServer struct {
}

func (*UnimplementedVirtualHardDiskAgentServer) Invoke(ctx context.Context, req *VirtualHardDiskRequest) (*VirtualHardDiskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}
func (*UnimplementedVirtualHardDiskAgentServer) CheckNotification(ctx context.Context, req *empty.Empty) (*common.NotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckNotification not implemented")
}

func RegisterVirtualHardDiskAgentServer(s *grpc.Server, srv VirtualHardDiskAgentServer) {
	s.RegisterService(&_VirtualHardDiskAgent_serviceDesc, srv)
}

func _VirtualHardDiskAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualHardDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualHardDiskAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.storage.VirtualHardDiskAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualHardDiskAgentServer).Invoke(ctx, req.(*VirtualHardDiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualHardDiskAgent_CheckNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualHardDiskAgentServer).CheckNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.storage.VirtualHardDiskAgent/CheckNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualHardDiskAgentServer).CheckNotification(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _VirtualHardDiskAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.nodeagent.storage.VirtualHardDiskAgent",
	HandlerType: (*VirtualHardDiskAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _VirtualHardDiskAgent_Invoke_Handler,
		},
		{
			MethodName: "CheckNotification",
			Handler:    _VirtualHardDiskAgent_CheckNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "virtualharddisk.proto",
}
