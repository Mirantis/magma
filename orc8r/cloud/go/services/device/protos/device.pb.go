// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"
import protos "magma/orc8r/cloud/go/protos"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PhysicalEntity struct {
	// Globally unique identifier per type (MAC/SN)
	DeviceID string `protobuf:"bytes,1,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	// Used to deserialize info
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Any other information (manufacturer, location, owner, etc)
	Info                 []byte   `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PhysicalEntity) Reset()         { *m = PhysicalEntity{} }
func (m *PhysicalEntity) String() string { return proto.CompactTextString(m) }
func (*PhysicalEntity) ProtoMessage()    {}
func (*PhysicalEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_5445f69aa365892f, []int{0}
}
func (m *PhysicalEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhysicalEntity.Unmarshal(m, b)
}
func (m *PhysicalEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhysicalEntity.Marshal(b, m, deterministic)
}
func (dst *PhysicalEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhysicalEntity.Merge(dst, src)
}
func (m *PhysicalEntity) XXX_Size() int {
	return xxx_messageInfo_PhysicalEntity.Size(m)
}
func (m *PhysicalEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_PhysicalEntity.DiscardUnknown(m)
}

var xxx_messageInfo_PhysicalEntity proto.InternalMessageInfo

func (m *PhysicalEntity) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func (m *PhysicalEntity) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PhysicalEntity) GetInfo() []byte {
	if m != nil {
		return m.Info
	}
	return nil
}

type RegisterDevicesRequest struct {
	NetworkID            string            `protobuf:"bytes,1,opt,name=networkID,proto3" json:"networkID,omitempty"`
	Entities             []*PhysicalEntity `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RegisterDevicesRequest) Reset()         { *m = RegisterDevicesRequest{} }
func (m *RegisterDevicesRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterDevicesRequest) ProtoMessage()    {}
func (*RegisterDevicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_5445f69aa365892f, []int{1}
}
func (m *RegisterDevicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterDevicesRequest.Unmarshal(m, b)
}
func (m *RegisterDevicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterDevicesRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterDevicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterDevicesRequest.Merge(dst, src)
}
func (m *RegisterDevicesRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterDevicesRequest.Size(m)
}
func (m *RegisterDevicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterDevicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterDevicesRequest proto.InternalMessageInfo

func (m *RegisterDevicesRequest) GetNetworkID() string {
	if m != nil {
		return m.NetworkID
	}
	return ""
}

func (m *RegisterDevicesRequest) GetEntities() []*PhysicalEntity {
	if m != nil {
		return m.Entities
	}
	return nil
}

type DeviceID struct {
	DeviceID             string   `protobuf:"bytes,1,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceID) Reset()         { *m = DeviceID{} }
func (m *DeviceID) String() string { return proto.CompactTextString(m) }
func (*DeviceID) ProtoMessage()    {}
func (*DeviceID) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_5445f69aa365892f, []int{2}
}
func (m *DeviceID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceID.Unmarshal(m, b)
}
func (m *DeviceID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceID.Marshal(b, m, deterministic)
}
func (dst *DeviceID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceID.Merge(dst, src)
}
func (m *DeviceID) XXX_Size() int {
	return xxx_messageInfo_DeviceID.Size(m)
}
func (m *DeviceID) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceID.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceID proto.InternalMessageInfo

func (m *DeviceID) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func (m *DeviceID) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type GetDeviceInfoRequest struct {
	NetworkID            string      `protobuf:"bytes,1,opt,name=networkID,proto3" json:"networkID,omitempty"`
	DeviceIDs            []*DeviceID `protobuf:"bytes,2,rep,name=deviceIDs,proto3" json:"deviceIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetDeviceInfoRequest) Reset()         { *m = GetDeviceInfoRequest{} }
func (m *GetDeviceInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetDeviceInfoRequest) ProtoMessage()    {}
func (*GetDeviceInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_5445f69aa365892f, []int{3}
}
func (m *GetDeviceInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceInfoRequest.Unmarshal(m, b)
}
func (m *GetDeviceInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceInfoRequest.Marshal(b, m, deterministic)
}
func (dst *GetDeviceInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceInfoRequest.Merge(dst, src)
}
func (m *GetDeviceInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetDeviceInfoRequest.Size(m)
}
func (m *GetDeviceInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceInfoRequest proto.InternalMessageInfo

func (m *GetDeviceInfoRequest) GetNetworkID() string {
	if m != nil {
		return m.NetworkID
	}
	return ""
}

func (m *GetDeviceInfoRequest) GetDeviceIDs() []*DeviceID {
	if m != nil {
		return m.DeviceIDs
	}
	return nil
}

type GetDeviceInfoResponse struct {
	// A map of device IDs to corresponding PhysicalEntity structure
	DeviceMap            map[string]*PhysicalEntity `protobuf:"bytes,1,rep,name=deviceMap,proto3" json:"deviceMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *GetDeviceInfoResponse) Reset()         { *m = GetDeviceInfoResponse{} }
func (m *GetDeviceInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetDeviceInfoResponse) ProtoMessage()    {}
func (*GetDeviceInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_5445f69aa365892f, []int{4}
}
func (m *GetDeviceInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceInfoResponse.Unmarshal(m, b)
}
func (m *GetDeviceInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceInfoResponse.Marshal(b, m, deterministic)
}
func (dst *GetDeviceInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceInfoResponse.Merge(dst, src)
}
func (m *GetDeviceInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetDeviceInfoResponse.Size(m)
}
func (m *GetDeviceInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceInfoResponse proto.InternalMessageInfo

func (m *GetDeviceInfoResponse) GetDeviceMap() map[string]*PhysicalEntity {
	if m != nil {
		return m.DeviceMap
	}
	return nil
}

type DeleteDevicesRequest struct {
	NetworkID            string      `protobuf:"bytes,1,opt,name=networkID,proto3" json:"networkID,omitempty"`
	DeviceIDs            []*DeviceID `protobuf:"bytes,2,rep,name=deviceIDs,proto3" json:"deviceIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DeleteDevicesRequest) Reset()         { *m = DeleteDevicesRequest{} }
func (m *DeleteDevicesRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteDevicesRequest) ProtoMessage()    {}
func (*DeleteDevicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_5445f69aa365892f, []int{5}
}
func (m *DeleteDevicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDevicesRequest.Unmarshal(m, b)
}
func (m *DeleteDevicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDevicesRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteDevicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDevicesRequest.Merge(dst, src)
}
func (m *DeleteDevicesRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteDevicesRequest.Size(m)
}
func (m *DeleteDevicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDevicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDevicesRequest proto.InternalMessageInfo

func (m *DeleteDevicesRequest) GetNetworkID() string {
	if m != nil {
		return m.NetworkID
	}
	return ""
}

func (m *DeleteDevicesRequest) GetDeviceIDs() []*DeviceID {
	if m != nil {
		return m.DeviceIDs
	}
	return nil
}

func init() {
	proto.RegisterType((*PhysicalEntity)(nil), "magma.orc8r.device.PhysicalEntity")
	proto.RegisterType((*RegisterDevicesRequest)(nil), "magma.orc8r.device.RegisterDevicesRequest")
	proto.RegisterType((*DeviceID)(nil), "magma.orc8r.device.DeviceID")
	proto.RegisterType((*GetDeviceInfoRequest)(nil), "magma.orc8r.device.GetDeviceInfoRequest")
	proto.RegisterType((*GetDeviceInfoResponse)(nil), "magma.orc8r.device.GetDeviceInfoResponse")
	proto.RegisterMapType((map[string]*PhysicalEntity)(nil), "magma.orc8r.device.GetDeviceInfoResponse.DeviceMapEntry")
	proto.RegisterType((*DeleteDevicesRequest)(nil), "magma.orc8r.device.DeleteDevicesRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DeviceClient is the client API for Device service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeviceClient interface {
	RegisterDevices(ctx context.Context, in *RegisterDevicesRequest, opts ...grpc.CallOption) (*protos.Void, error)
	GetDeviceInfo(ctx context.Context, in *GetDeviceInfoRequest, opts ...grpc.CallOption) (*GetDeviceInfoResponse, error)
	DeleteDevices(ctx context.Context, in *DeleteDevicesRequest, opts ...grpc.CallOption) (*protos.Void, error)
}

type deviceClient struct {
	cc *grpc.ClientConn
}

func NewDeviceClient(cc *grpc.ClientConn) DeviceClient {
	return &deviceClient{cc}
}

func (c *deviceClient) RegisterDevices(ctx context.Context, in *RegisterDevicesRequest, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.device.Device/RegisterDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) GetDeviceInfo(ctx context.Context, in *GetDeviceInfoRequest, opts ...grpc.CallOption) (*GetDeviceInfoResponse, error) {
	out := new(GetDeviceInfoResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.device.Device/GetDeviceInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) DeleteDevices(ctx context.Context, in *DeleteDevicesRequest, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.device.Device/DeleteDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceServer is the server API for Device service.
type DeviceServer interface {
	RegisterDevices(context.Context, *RegisterDevicesRequest) (*protos.Void, error)
	GetDeviceInfo(context.Context, *GetDeviceInfoRequest) (*GetDeviceInfoResponse, error)
	DeleteDevices(context.Context, *DeleteDevicesRequest) (*protos.Void, error)
}

func RegisterDeviceServer(s *grpc.Server, srv DeviceServer) {
	s.RegisterService(&_Device_serviceDesc, srv)
}

func _Device_RegisterDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).RegisterDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.device.Device/RegisterDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).RegisterDevices(ctx, req.(*RegisterDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_GetDeviceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).GetDeviceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.device.Device/GetDeviceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).GetDeviceInfo(ctx, req.(*GetDeviceInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_DeleteDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).DeleteDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.device.Device/DeleteDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).DeleteDevices(ctx, req.(*DeleteDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Device_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.device.Device",
	HandlerType: (*DeviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDevices",
			Handler:    _Device_RegisterDevices_Handler,
		},
		{
			MethodName: "GetDeviceInfo",
			Handler:    _Device_GetDeviceInfo_Handler,
		},
		{
			MethodName: "DeleteDevices",
			Handler:    _Device_DeleteDevices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "device.proto",
}

func init() { proto.RegisterFile("device.proto", fileDescriptor_device_5445f69aa365892f) }

var fileDescriptor_device_5445f69aa365892f = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x1d, 0xa8, 0x9c, 0xa1, 0x2d, 0xb0, 0x02, 0xe4, 0x5a, 0x3d, 0x44, 0x7b, 0x32, 0x1c,
	0x1c, 0xa9, 0x5c, 0xac, 0x1c, 0x38, 0x20, 0x57, 0x88, 0x43, 0x11, 0x5a, 0xa1, 0x1e, 0x38, 0xe1,
	0x38, 0x63, 0x63, 0xc5, 0xde, 0x35, 0xde, 0x4d, 0x90, 0xff, 0x83, 0x1f, 0xe3, 0x8f, 0x90, 0xbd,
	0x71, 0x22, 0x87, 0x3d, 0x38, 0xea, 0x29, 0x93, 0x99, 0x79, 0x6f, 0xde, 0xbc, 0x59, 0xc3, 0xc5,
	0x0a, 0xb7, 0x79, 0x82, 0x41, 0x55, 0x0b, 0x25, 0x08, 0x29, 0xe3, 0xac, 0x8c, 0x03, 0x51, 0x27,
	0x61, 0x1d, 0xe8, 0x8a, 0x77, 0x9d, 0x09, 0x91, 0x15, 0x38, 0xef, 0x3a, 0x96, 0x9b, 0x74, 0x1e,
	0xf3, 0x46, 0xb7, 0x7b, 0xd7, 0x5d, 0xa3, 0xae, 0xc8, 0x79, 0x22, 0xca, 0x52, 0x70, 0x5d, 0xa2,
	0xdf, 0xe0, 0xea, 0xeb, 0xcf, 0x46, 0xe6, 0x49, 0x5c, 0xdc, 0x71, 0x95, 0xab, 0x86, 0x78, 0xe0,
	0x68, 0xc6, 0xcf, 0x91, 0x6b, 0xcd, 0x2c, 0x7f, 0xca, 0xf6, 0xff, 0x09, 0x81, 0x27, 0xaa, 0xa9,
	0xd0, 0xb5, 0xbb, 0x7c, 0x17, 0xb7, 0xb9, 0x9c, 0xa7, 0xc2, 0x9d, 0xcc, 0x2c, 0xff, 0x82, 0x75,
	0x31, 0xdd, 0xc2, 0x1b, 0x86, 0x59, 0x2e, 0x15, 0xd6, 0x51, 0x87, 0x95, 0x0c, 0x7f, 0x6d, 0x50,
	0x2a, 0x72, 0x03, 0x53, 0x8e, 0xea, 0xb7, 0xa8, 0xd7, 0x7b, 0xfa, 0x43, 0x82, 0x7c, 0x00, 0x07,
	0x5b, 0x15, 0x39, 0x4a, 0xd7, 0x9e, 0x4d, 0xfc, 0x67, 0xb7, 0x34, 0xf8, 0x7f, 0xd5, 0x60, 0xa8,
	0x98, 0xed, 0x31, 0x74, 0x01, 0x4e, 0xd4, 0x6b, 0x3d, 0x71, 0x0f, 0x5a, 0xc1, 0xab, 0x4f, 0xa8,
	0x76, 0x70, 0x9e, 0x8a, 0x71, 0x8a, 0x17, 0x30, 0xed, 0x59, 0x7b, 0xc9, 0x37, 0x26, 0xc9, 0xbd,
	0x2c, 0x76, 0x68, 0xa7, 0x7f, 0x2d, 0x78, 0x7d, 0x34, 0x52, 0x56, 0x82, 0x4b, 0x24, 0x0f, 0x3d,
	0xeb, 0x7d, 0x5c, 0xb9, 0x56, 0xc7, 0x1a, 0x9a, 0x58, 0x8d, 0xe8, 0xdd, 0xac, 0xfb, 0xb8, 0xba,
	0xe3, 0xaa, 0x6e, 0xd8, 0x81, 0xca, 0xfb, 0x01, 0x57, 0xc3, 0x22, 0x79, 0x01, 0x93, 0x35, 0x36,
	0xbb, 0xbd, 0xda, 0x90, 0x84, 0xf0, 0x74, 0x1b, 0x17, 0x1b, 0x6d, 0xce, 0xb8, 0x03, 0x68, 0xc0,
	0xc2, 0x0e, 0xad, 0xd6, 0xc5, 0x08, 0x0b, 0x54, 0x78, 0xd2, 0xdd, 0x1f, 0xe1, 0xe2, 0xed, 0x1f,
	0x1b, 0xce, 0x75, 0x9e, 0x30, 0x78, 0x7e, 0xf4, 0xec, 0xc8, 0x3b, 0x13, 0x8d, 0xf9, 0x6d, 0x7a,
	0x2f, 0x07, 0xbd, 0x0f, 0x22, 0x5f, 0xd1, 0x33, 0x92, 0xc2, 0xe5, 0xc0, 0x65, 0xe2, 0x8f, 0x38,
	0x84, 0xe6, 0x7b, 0x3b, 0xfa, 0x64, 0xf4, 0x8c, 0x7c, 0x81, 0xcb, 0x81, 0x71, 0xe6, 0x39, 0x26,
	0x6f, 0x8d, 0xba, 0x3f, 0x3a, 0xdf, 0xcf, 0xf5, 0xf7, 0xbe, 0xd4, 0xbf, 0xef, 0xff, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x62, 0x4c, 0x5a, 0x74, 0x43, 0x04, 0x00, 0x00,
}
