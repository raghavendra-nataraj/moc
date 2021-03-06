// Code generated by protoc-gen-go. DO NOT EDIT.
// source: virtualmachinescaleset.proto

package compute

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	network "github.com/microsoft/moc/rpc/cloudagent/network"
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

type VirtualMachineScaleSetRequest struct {
	VirtualMachineScaleSetSystems []*VirtualMachineScaleSet `protobuf:"bytes,1,rep,name=VirtualMachineScaleSetSystems,proto3" json:"VirtualMachineScaleSetSystems,omitempty"`
	OperationType                 common.Operation          `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                  `json:"-"`
	XXX_unrecognized              []byte                    `json:"-"`
	XXX_sizecache                 int32                     `json:"-"`
}

func (m *VirtualMachineScaleSetRequest) Reset()         { *m = VirtualMachineScaleSetRequest{} }
func (m *VirtualMachineScaleSetRequest) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineScaleSetRequest) ProtoMessage()    {}
func (*VirtualMachineScaleSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2ef1a2fddc3eb6a, []int{0}
}

func (m *VirtualMachineScaleSetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineScaleSetRequest.Unmarshal(m, b)
}
func (m *VirtualMachineScaleSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineScaleSetRequest.Marshal(b, m, deterministic)
}
func (m *VirtualMachineScaleSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineScaleSetRequest.Merge(m, src)
}
func (m *VirtualMachineScaleSetRequest) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineScaleSetRequest.Size(m)
}
func (m *VirtualMachineScaleSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineScaleSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineScaleSetRequest proto.InternalMessageInfo

func (m *VirtualMachineScaleSetRequest) GetVirtualMachineScaleSetSystems() []*VirtualMachineScaleSet {
	if m != nil {
		return m.VirtualMachineScaleSetSystems
	}
	return nil
}

func (m *VirtualMachineScaleSetRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type VirtualMachineScaleSetResponse struct {
	VirtualMachineScaleSetSystems []*VirtualMachineScaleSet `protobuf:"bytes,1,rep,name=VirtualMachineScaleSetSystems,proto3" json:"VirtualMachineScaleSetSystems,omitempty"`
	Result                        *wrappers.BoolValue       `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                         string                    `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                  `json:"-"`
	XXX_unrecognized              []byte                    `json:"-"`
	XXX_sizecache                 int32                     `json:"-"`
}

func (m *VirtualMachineScaleSetResponse) Reset()         { *m = VirtualMachineScaleSetResponse{} }
func (m *VirtualMachineScaleSetResponse) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineScaleSetResponse) ProtoMessage()    {}
func (*VirtualMachineScaleSetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2ef1a2fddc3eb6a, []int{1}
}

func (m *VirtualMachineScaleSetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineScaleSetResponse.Unmarshal(m, b)
}
func (m *VirtualMachineScaleSetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineScaleSetResponse.Marshal(b, m, deterministic)
}
func (m *VirtualMachineScaleSetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineScaleSetResponse.Merge(m, src)
}
func (m *VirtualMachineScaleSetResponse) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineScaleSetResponse.Size(m)
}
func (m *VirtualMachineScaleSetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineScaleSetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineScaleSetResponse proto.InternalMessageInfo

func (m *VirtualMachineScaleSetResponse) GetVirtualMachineScaleSetSystems() []*VirtualMachineScaleSet {
	if m != nil {
		return m.VirtualMachineScaleSetSystems
	}
	return nil
}

func (m *VirtualMachineScaleSetResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *VirtualMachineScaleSetResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Sku struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Capacity             int64    `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sku) Reset()         { *m = Sku{} }
func (m *Sku) String() string { return proto.CompactTextString(m) }
func (*Sku) ProtoMessage()    {}
func (*Sku) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2ef1a2fddc3eb6a, []int{2}
}

func (m *Sku) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sku.Unmarshal(m, b)
}
func (m *Sku) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sku.Marshal(b, m, deterministic)
}
func (m *Sku) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sku.Merge(m, src)
}
func (m *Sku) XXX_Size() int {
	return xxx_messageInfo_Sku.Size(m)
}
func (m *Sku) XXX_DiscardUnknown() {
	xxx_messageInfo_Sku.DiscardUnknown(m)
}

var xxx_messageInfo_Sku proto.InternalMessageInfo

func (m *Sku) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Sku) GetCapacity() int64 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

type NetworkConfigurationScaleSet struct {
	Interfaces           []*network.NetworkInterface `protobuf:"bytes,1,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *NetworkConfigurationScaleSet) Reset()         { *m = NetworkConfigurationScaleSet{} }
func (m *NetworkConfigurationScaleSet) String() string { return proto.CompactTextString(m) }
func (*NetworkConfigurationScaleSet) ProtoMessage()    {}
func (*NetworkConfigurationScaleSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2ef1a2fddc3eb6a, []int{3}
}

func (m *NetworkConfigurationScaleSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkConfigurationScaleSet.Unmarshal(m, b)
}
func (m *NetworkConfigurationScaleSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkConfigurationScaleSet.Marshal(b, m, deterministic)
}
func (m *NetworkConfigurationScaleSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkConfigurationScaleSet.Merge(m, src)
}
func (m *NetworkConfigurationScaleSet) XXX_Size() int {
	return xxx_messageInfo_NetworkConfigurationScaleSet.Size(m)
}
func (m *NetworkConfigurationScaleSet) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkConfigurationScaleSet.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkConfigurationScaleSet proto.InternalMessageInfo

func (m *NetworkConfigurationScaleSet) GetInterfaces() []*network.NetworkInterface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

type VirtualMachineProfile struct {
	Vmprefix             string                        `protobuf:"bytes,1,opt,name=vmprefix,proto3" json:"vmprefix,omitempty"`
	Network              *NetworkConfigurationScaleSet `protobuf:"bytes,2,opt,name=network,proto3" json:"network,omitempty"`
	Storage              *StorageConfiguration         `protobuf:"bytes,3,opt,name=storage,proto3" json:"storage,omitempty"`
	Os                   *OperatingSystemConfiguration `protobuf:"bytes,4,opt,name=os,proto3" json:"os,omitempty"`
	Hardware             *HardwareConfiguration        `protobuf:"bytes,5,opt,name=hardware,proto3" json:"hardware,omitempty"`
	Security             *SecurityConfiguration        `protobuf:"bytes,6,opt,name=security,proto3" json:"security,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *VirtualMachineProfile) Reset()         { *m = VirtualMachineProfile{} }
func (m *VirtualMachineProfile) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineProfile) ProtoMessage()    {}
func (*VirtualMachineProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2ef1a2fddc3eb6a, []int{4}
}

func (m *VirtualMachineProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineProfile.Unmarshal(m, b)
}
func (m *VirtualMachineProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineProfile.Marshal(b, m, deterministic)
}
func (m *VirtualMachineProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineProfile.Merge(m, src)
}
func (m *VirtualMachineProfile) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineProfile.Size(m)
}
func (m *VirtualMachineProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineProfile.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineProfile proto.InternalMessageInfo

func (m *VirtualMachineProfile) GetVmprefix() string {
	if m != nil {
		return m.Vmprefix
	}
	return ""
}

func (m *VirtualMachineProfile) GetNetwork() *NetworkConfigurationScaleSet {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *VirtualMachineProfile) GetStorage() *StorageConfiguration {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *VirtualMachineProfile) GetOs() *OperatingSystemConfiguration {
	if m != nil {
		return m.Os
	}
	return nil
}

func (m *VirtualMachineProfile) GetHardware() *HardwareConfiguration {
	if m != nil {
		return m.Hardware
	}
	return nil
}

func (m *VirtualMachineProfile) GetSecurity() *SecurityConfiguration {
	if m != nil {
		return m.Security
	}
	return nil
}

type VirtualMachineScaleSet struct {
	Name                    string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                      string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Sku                     *Sku                   `protobuf:"bytes,3,opt,name=sku,proto3" json:"sku,omitempty"`
	Virtualmachineprofile   *VirtualMachineProfile `protobuf:"bytes,4,opt,name=virtualmachineprofile,proto3" json:"virtualmachineprofile,omitempty"`
	VirtualMachineSystems   []*VirtualMachine      `protobuf:"bytes,5,rep,name=VirtualMachineSystems,proto3" json:"VirtualMachineSystems,omitempty"`
	Nodefqdn                string                 `protobuf:"bytes,6,opt,name=nodefqdn,proto3" json:"nodefqdn,omitempty"`
	GroupName               string                 `protobuf:"bytes,7,opt,name=groupName,proto3" json:"groupName,omitempty"`
	LocationName            string                 `protobuf:"bytes,8,opt,name=locationName,proto3" json:"locationName,omitempty"`
	Status                  *common.Status         `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	DisableHighAvailability bool                   `protobuf:"varint,11,opt,name=disableHighAvailability,proto3" json:"disableHighAvailability,omitempty"`
	AllowedOwnerNodes       []string               `protobuf:"bytes,12,rep,name=allowedOwnerNodes,proto3" json:"allowedOwnerNodes,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}               `json:"-"`
	XXX_unrecognized        []byte                 `json:"-"`
	XXX_sizecache           int32                  `json:"-"`
}

func (m *VirtualMachineScaleSet) Reset()         { *m = VirtualMachineScaleSet{} }
func (m *VirtualMachineScaleSet) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineScaleSet) ProtoMessage()    {}
func (*VirtualMachineScaleSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2ef1a2fddc3eb6a, []int{5}
}

func (m *VirtualMachineScaleSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineScaleSet.Unmarshal(m, b)
}
func (m *VirtualMachineScaleSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineScaleSet.Marshal(b, m, deterministic)
}
func (m *VirtualMachineScaleSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineScaleSet.Merge(m, src)
}
func (m *VirtualMachineScaleSet) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineScaleSet.Size(m)
}
func (m *VirtualMachineScaleSet) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineScaleSet.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineScaleSet proto.InternalMessageInfo

func (m *VirtualMachineScaleSet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VirtualMachineScaleSet) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VirtualMachineScaleSet) GetSku() *Sku {
	if m != nil {
		return m.Sku
	}
	return nil
}

func (m *VirtualMachineScaleSet) GetVirtualmachineprofile() *VirtualMachineProfile {
	if m != nil {
		return m.Virtualmachineprofile
	}
	return nil
}

func (m *VirtualMachineScaleSet) GetVirtualMachineSystems() []*VirtualMachine {
	if m != nil {
		return m.VirtualMachineSystems
	}
	return nil
}

func (m *VirtualMachineScaleSet) GetNodefqdn() string {
	if m != nil {
		return m.Nodefqdn
	}
	return ""
}

func (m *VirtualMachineScaleSet) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *VirtualMachineScaleSet) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *VirtualMachineScaleSet) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VirtualMachineScaleSet) GetDisableHighAvailability() bool {
	if m != nil {
		return m.DisableHighAvailability
	}
	return false
}

func (m *VirtualMachineScaleSet) GetAllowedOwnerNodes() []string {
	if m != nil {
		return m.AllowedOwnerNodes
	}
	return nil
}

func init() {
	proto.RegisterType((*VirtualMachineScaleSetRequest)(nil), "moc.cloudagent.compute.VirtualMachineScaleSetRequest")
	proto.RegisterType((*VirtualMachineScaleSetResponse)(nil), "moc.cloudagent.compute.VirtualMachineScaleSetResponse")
	proto.RegisterType((*Sku)(nil), "moc.cloudagent.compute.Sku")
	proto.RegisterType((*NetworkConfigurationScaleSet)(nil), "moc.cloudagent.compute.NetworkConfigurationScaleSet")
	proto.RegisterType((*VirtualMachineProfile)(nil), "moc.cloudagent.compute.VirtualMachineProfile")
	proto.RegisterType((*VirtualMachineScaleSet)(nil), "moc.cloudagent.compute.VirtualMachineScaleSet")
}

func init() { proto.RegisterFile("virtualmachinescaleset.proto", fileDescriptor_d2ef1a2fddc3eb6a) }

var fileDescriptor_d2ef1a2fddc3eb6a = []byte{
	// 729 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0x25, 0xed, 0xd6, 0xb5, 0xee, 0x98, 0x84, 0xb5, 0x8d, 0xa8, 0x1b, 0x53, 0x15, 0x24, 0xd4,
	0x8b, 0x35, 0x95, 0xca, 0x06, 0xdc, 0x6e, 0xfc, 0x68, 0xbb, 0xa0, 0x43, 0x2e, 0xda, 0x05, 0xe2,
	0xc6, 0x4d, 0xdc, 0xd4, 0x6a, 0x12, 0x67, 0xfe, 0x69, 0xe9, 0x9b, 0xc0, 0x13, 0xf0, 0x30, 0xbc,
	0x04, 0x8f, 0x82, 0xe2, 0x38, 0x1d, 0x19, 0x4d, 0xb5, 0x5d, 0x71, 0x57, 0xfb, 0x3b, 0xe7, 0xf8,
	0xf3, 0xf9, 0x4e, 0x5c, 0x70, 0x38, 0xa3, 0x5c, 0x2a, 0x1c, 0x46, 0xd8, 0x9b, 0xd0, 0x98, 0x08,
	0x0f, 0x87, 0x44, 0x10, 0xe9, 0x26, 0x9c, 0x49, 0x06, 0xf7, 0x23, 0xe6, 0xb9, 0x5e, 0xc8, 0x94,
	0x8f, 0x03, 0x12, 0x4b, 0xd7, 0x63, 0x51, 0xa2, 0x24, 0x69, 0x1d, 0x05, 0x8c, 0x05, 0x21, 0xe9,
	0x69, 0xd4, 0x48, 0x8d, 0x7b, 0x73, 0x8e, 0x93, 0x84, 0x70, 0x91, 0xf1, 0x5a, 0xbb, 0x45, 0x55,
	0xb3, 0xbb, 0xed, 0xb1, 0x28, 0x62, 0xb1, 0x59, 0xed, 0xc7, 0x44, 0xce, 0x19, 0x9f, 0xd2, 0x58,
	0x12, 0x3e, 0xc6, 0x9e, 0x41, 0x39, 0xbf, 0x2c, 0xf0, 0xec, 0x3a, 0xa3, 0x7f, 0xcc, 0xe8, 0xc3,
	0xb4, 0xa9, 0x21, 0x91, 0x88, 0xdc, 0x28, 0x22, 0x24, 0x94, 0x65, 0x80, 0xe1, 0x42, 0x48, 0x12,
	0x09, 0xdb, 0x6a, 0x57, 0x3b, 0xcd, 0xbe, 0xeb, 0xae, 0xee, 0xde, 0x2d, 0x51, 0x5f, 0x2f, 0x0a,
	0x4f, 0xc0, 0xe3, 0xab, 0x84, 0x70, 0x2c, 0x29, 0x8b, 0x3f, 0x2f, 0x12, 0x62, 0x57, 0xda, 0x56,
	0x67, 0xa7, 0xbf, 0xa3, 0x4f, 0x59, 0x56, 0x50, 0x11, 0xe4, 0xfc, 0xb6, 0xc0, 0x51, 0xd9, 0x6d,
	0x44, 0xc2, 0x62, 0x41, 0xfe, 0xd3, 0x75, 0xfa, 0xa0, 0x86, 0x88, 0x50, 0xa1, 0xd4, 0xf7, 0x68,
	0xf6, 0x5b, 0x6e, 0x36, 0x53, 0x37, 0x9f, 0xa9, 0x7b, 0xce, 0x58, 0x78, 0x8d, 0x43, 0x45, 0x90,
	0x41, 0xc2, 0x5d, 0xb0, 0xf9, 0x9e, 0x73, 0xc6, 0xed, 0x6a, 0xdb, 0xea, 0x34, 0x50, 0xb6, 0x70,
	0x4e, 0x41, 0x75, 0x38, 0x55, 0x10, 0x82, 0x8d, 0x18, 0x47, 0xc4, 0xb6, 0x74, 0x4d, 0xff, 0x86,
	0x2d, 0x50, 0xf7, 0x70, 0x82, 0x3d, 0x2a, 0x17, 0xfa, 0x98, 0x2a, 0x5a, 0xae, 0x9d, 0x09, 0x38,
	0x1c, 0x64, 0x09, 0x78, 0xcb, 0xe2, 0x31, 0x0d, 0x54, 0xe6, 0x5a, 0xde, 0x27, 0xbc, 0x00, 0x60,
	0x19, 0x8d, 0xdc, 0x83, 0xce, 0x5d, 0x0f, 0x4c, 0x86, 0x5c, 0xa3, 0x74, 0x99, 0x13, 0xd0, 0x5f,
	0x5c, 0xe7, 0x67, 0x15, 0xec, 0x15, 0xcd, 0xf8, 0xc4, 0xd9, 0x98, 0x86, 0xba, 0xbf, 0x59, 0x94,
	0x70, 0x32, 0xa6, 0xdf, 0x4c, 0xdf, 0xcb, 0x35, 0x1c, 0x80, 0x2d, 0xa3, 0x6e, 0x1c, 0x3a, 0x29,
	0x1b, 0xc0, 0xba, 0x6b, 0xa0, 0x5c, 0x04, 0x7e, 0x00, 0x5b, 0x42, 0x32, 0x8e, 0x03, 0xa2, 0xed,
	0x6b, 0xf6, 0x8f, 0xcb, 0xf4, 0x86, 0x19, 0xac, 0xa0, 0x87, 0x72, 0x32, 0x7c, 0x07, 0x2a, 0x4c,
	0xd8, 0x1b, 0xeb, 0x5b, 0x32, 0x21, 0x8c, 0x83, 0x6c, 0xdc, 0x45, 0xa9, 0x0a, 0x13, 0xf0, 0x12,
	0xd4, 0x27, 0x98, 0xfb, 0x73, 0xcc, 0x89, 0xbd, 0xa9, 0xb5, 0xba, 0x65, 0x5a, 0x17, 0x06, 0x57,
	0x14, 0x59, 0xd2, 0x53, 0x29, 0x41, 0x3c, 0xc5, 0xd3, 0x21, 0xd7, 0xd6, 0x4b, 0x0d, 0x0d, 0xee,
	0x8e, 0x54, 0x4e, 0x77, 0x7e, 0x6c, 0x80, 0xfd, 0xd5, 0xb1, 0x5d, 0x19, 0xaf, 0x1d, 0x50, 0xa1,
	0xbe, 0x9e, 0x4e, 0x03, 0x55, 0xa8, 0x0f, 0xbb, 0xa0, 0x2a, 0xa6, 0xca, 0xd8, 0x7b, 0x50, 0xda,
	0xc4, 0x54, 0xa1, 0x14, 0x07, 0x3d, 0xb0, 0x57, 0x7c, 0xa7, 0x92, 0x2c, 0x16, 0xc6, 0xdc, 0xee,
	0xfd, 0x3e, 0x38, 0x93, 0x25, 0xb4, 0x5a, 0x0b, 0x7e, 0xbd, 0x9b, 0xbd, 0xfc, 0xab, 0xde, 0xd4,
	0x89, 0x7e, 0x71, 0xbf, 0x43, 0xd0, 0x6a, 0x91, 0x34, 0xc0, 0x31, 0xf3, 0xc9, 0xf8, 0xc6, 0x8f,
	0xb5, 0xf7, 0x0d, 0xb4, 0x5c, 0xc3, 0x43, 0xd0, 0x08, 0x38, 0x53, 0xc9, 0x20, 0xb5, 0x6d, 0x4b,
	0x17, 0x6f, 0x37, 0xa0, 0x03, 0xb6, 0x43, 0xe6, 0xe9, 0x01, 0x68, 0x40, 0x5d, 0x03, 0x0a, 0x7b,
	0xf0, 0x39, 0xa8, 0x09, 0x89, 0xa5, 0x12, 0x76, 0x43, 0x3b, 0xd2, 0xd4, 0xcd, 0x0e, 0xf5, 0x16,
	0x32, 0x25, 0xf8, 0x06, 0x3c, 0xf5, 0xa9, 0xc0, 0xa3, 0x90, 0x5c, 0xd0, 0x60, 0x72, 0x36, 0xc3,
	0x34, 0xc4, 0x23, 0x1a, 0xa6, 0x69, 0x68, 0xb6, 0xad, 0x4e, 0x1d, 0x95, 0x95, 0xe1, 0x31, 0x78,
	0x82, 0xc3, 0x90, 0xcd, 0x89, 0x7f, 0x35, 0x8f, 0x09, 0x1f, 0x30, 0x9f, 0x08, 0x7b, 0xbb, 0x5d,
	0xed, 0x34, 0xd0, 0xbf, 0x85, 0xfe, 0x77, 0x0b, 0x1c, 0xac, 0xce, 0xc6, 0x59, 0x6a, 0x1c, 0x5c,
	0x80, 0xda, 0x65, 0x3c, 0x63, 0x53, 0x02, 0x4f, 0x1f, 0xf8, 0x52, 0x66, 0x7f, 0x2b, 0xad, 0x57,
	0x0f, 0xa5, 0x65, 0xef, 0xb7, 0xf3, 0xe8, 0xfc, 0xf5, 0x97, 0xd3, 0x80, 0xca, 0x89, 0x1a, 0xa5,
	0x94, 0x5e, 0x44, 0x3d, 0xce, 0x04, 0x1b, 0xcb, 0x5e, 0xc4, 0xbc, 0xae, 0x7e, 0x4e, 0x7b, 0x3c,
	0xf1, 0x7a, 0xb7, 0xca, 0x3d, 0xa3, 0x3c, 0xaa, 0xe9, 0xea, 0xcb, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xbf, 0x03, 0x14, 0x01, 0x86, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VirtualMachineScaleSetAgentClient is the client API for VirtualMachineScaleSetAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VirtualMachineScaleSetAgentClient interface {
	Invoke(ctx context.Context, in *VirtualMachineScaleSetRequest, opts ...grpc.CallOption) (*VirtualMachineScaleSetResponse, error)
}

type virtualMachineScaleSetAgentClient struct {
	cc *grpc.ClientConn
}

func NewVirtualMachineScaleSetAgentClient(cc *grpc.ClientConn) VirtualMachineScaleSetAgentClient {
	return &virtualMachineScaleSetAgentClient{cc}
}

func (c *virtualMachineScaleSetAgentClient) Invoke(ctx context.Context, in *VirtualMachineScaleSetRequest, opts ...grpc.CallOption) (*VirtualMachineScaleSetResponse, error) {
	out := new(VirtualMachineScaleSetResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.compute.VirtualMachineScaleSetAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtualMachineScaleSetAgentServer is the server API for VirtualMachineScaleSetAgent service.
type VirtualMachineScaleSetAgentServer interface {
	Invoke(context.Context, *VirtualMachineScaleSetRequest) (*VirtualMachineScaleSetResponse, error)
}

// UnimplementedVirtualMachineScaleSetAgentServer can be embedded to have forward compatible implementations.
type UnimplementedVirtualMachineScaleSetAgentServer struct {
}

func (*UnimplementedVirtualMachineScaleSetAgentServer) Invoke(ctx context.Context, req *VirtualMachineScaleSetRequest) (*VirtualMachineScaleSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterVirtualMachineScaleSetAgentServer(s *grpc.Server, srv VirtualMachineScaleSetAgentServer) {
	s.RegisterService(&_VirtualMachineScaleSetAgent_serviceDesc, srv)
}

func _VirtualMachineScaleSetAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualMachineScaleSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineScaleSetAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.compute.VirtualMachineScaleSetAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineScaleSetAgentServer).Invoke(ctx, req.(*VirtualMachineScaleSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VirtualMachineScaleSetAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.compute.VirtualMachineScaleSetAgent",
	HandlerType: (*VirtualMachineScaleSetAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _VirtualMachineScaleSetAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "virtualmachinescaleset.proto",
}
