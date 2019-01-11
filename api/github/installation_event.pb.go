// Code generated by protoc-gen-go. DO NOT EDIT.
// source: installation_event.proto

package github_pb // import "github.com/izumin5210/ghrsync/api/github"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type InstallationEvent struct {
	InstallationEventId  string   `protobuf:"bytes,1,opt,name=installation_event_id,json=installationEventId,proto3" json:"installation_event_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallationEvent) Reset()         { *m = InstallationEvent{} }
func (m *InstallationEvent) String() string { return proto.CompactTextString(m) }
func (*InstallationEvent) ProtoMessage()    {}
func (*InstallationEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_installation_event_e5d80bc794d41c1f, []int{0}
}
func (m *InstallationEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallationEvent.Unmarshal(m, b)
}
func (m *InstallationEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallationEvent.Marshal(b, m, deterministic)
}
func (dst *InstallationEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallationEvent.Merge(dst, src)
}
func (m *InstallationEvent) XXX_Size() int {
	return xxx_messageInfo_InstallationEvent.Size(m)
}
func (m *InstallationEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallationEvent.DiscardUnknown(m)
}

var xxx_messageInfo_InstallationEvent proto.InternalMessageInfo

func (m *InstallationEvent) GetInstallationEventId() string {
	if m != nil {
		return m.InstallationEventId
	}
	return ""
}

type CreateInstallationEventRequest struct {
	InstallationEvent    *InstallationEvent `protobuf:"bytes,1,opt,name=installation_event,json=installationEvent,proto3" json:"installation_event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateInstallationEventRequest) Reset()         { *m = CreateInstallationEventRequest{} }
func (m *CreateInstallationEventRequest) String() string { return proto.CompactTextString(m) }
func (*CreateInstallationEventRequest) ProtoMessage()    {}
func (*CreateInstallationEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_installation_event_e5d80bc794d41c1f, []int{1}
}
func (m *CreateInstallationEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateInstallationEventRequest.Unmarshal(m, b)
}
func (m *CreateInstallationEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateInstallationEventRequest.Marshal(b, m, deterministic)
}
func (dst *CreateInstallationEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateInstallationEventRequest.Merge(dst, src)
}
func (m *CreateInstallationEventRequest) XXX_Size() int {
	return xxx_messageInfo_CreateInstallationEventRequest.Size(m)
}
func (m *CreateInstallationEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateInstallationEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateInstallationEventRequest proto.InternalMessageInfo

func (m *CreateInstallationEventRequest) GetInstallationEvent() *InstallationEvent {
	if m != nil {
		return m.InstallationEvent
	}
	return nil
}

func init() {
	proto.RegisterType((*InstallationEvent)(nil), "izumin5210.ghrsync.github.InstallationEvent")
	proto.RegisterType((*CreateInstallationEventRequest)(nil), "izumin5210.ghrsync.github.CreateInstallationEventRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InstallationEventServiceClient is the client API for InstallationEventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InstallationEventServiceClient interface {
	CreateInstallationEvent(ctx context.Context, in *CreateInstallationEventRequest, opts ...grpc.CallOption) (*InstallationEvent, error)
}

type installationEventServiceClient struct {
	cc *grpc.ClientConn
}

func NewInstallationEventServiceClient(cc *grpc.ClientConn) InstallationEventServiceClient {
	return &installationEventServiceClient{cc}
}

func (c *installationEventServiceClient) CreateInstallationEvent(ctx context.Context, in *CreateInstallationEventRequest, opts ...grpc.CallOption) (*InstallationEvent, error) {
	out := new(InstallationEvent)
	err := c.cc.Invoke(ctx, "/izumin5210.ghrsync.github.InstallationEventService/CreateInstallationEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InstallationEventServiceServer is the server API for InstallationEventService service.
type InstallationEventServiceServer interface {
	CreateInstallationEvent(context.Context, *CreateInstallationEventRequest) (*InstallationEvent, error)
}

func RegisterInstallationEventServiceServer(s *grpc.Server, srv InstallationEventServiceServer) {
	s.RegisterService(&_InstallationEventService_serviceDesc, srv)
}

func _InstallationEventService_CreateInstallationEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInstallationEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstallationEventServiceServer).CreateInstallationEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/izumin5210.ghrsync.github.InstallationEventService/CreateInstallationEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstallationEventServiceServer).CreateInstallationEvent(ctx, req.(*CreateInstallationEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InstallationEventService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "izumin5210.ghrsync.github.InstallationEventService",
	HandlerType: (*InstallationEventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInstallationEvent",
			Handler:    _InstallationEventService_CreateInstallationEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "installation_event.proto",
}

func init() {
	proto.RegisterFile("installation_event.proto", fileDescriptor_installation_event_e5d80bc794d41c1f)
}

var fileDescriptor_installation_event_e5d80bc794d41c1f = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0xcc, 0x2b, 0x2e,
	0x49, 0xcc, 0xc9, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x8b, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xcc, 0xac, 0x2a, 0xcd, 0xcd, 0xcc, 0x33, 0x35, 0x32, 0x34,
	0xd0, 0x4b, 0xcf, 0x28, 0x2a, 0xae, 0xcc, 0x4b, 0xd6, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0x92,
	0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb,
	0x2f, 0x01, 0xeb, 0x2e, 0x86, 0x68, 0x54, 0x72, 0xe7, 0x12, 0xf4, 0x44, 0x32, 0xd4, 0x15, 0x64,
	0xa6, 0x90, 0x11, 0x97, 0x28, 0xa6, 0x4d, 0xf1, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0xc2, 0x99, 0xe8, 0x3a, 0x3c, 0x53, 0x94, 0x6a, 0xb9, 0xe4, 0x9c, 0x8b, 0x52, 0x13, 0x4b,
	0x52, 0x31, 0x8c, 0x0b, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x8a, 0xe6, 0x12, 0xc2, 0x34,
	0x15, 0x6c, 0x24, 0xb7, 0x91, 0x8e, 0x1e, 0x4e, 0x0f, 0xe8, 0x61, 0x1a, 0x28, 0x88, 0xe1, 0x00,
	0xa3, 0x8b, 0x8c, 0x5c, 0x12, 0x18, 0x0a, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0xb6,
	0x30, 0x72, 0x89, 0xe3, 0x70, 0x9c, 0x90, 0x25, 0x1e, 0x9b, 0xf1, 0x7b, 0x48, 0x8a, 0x24, 0x47,
	0x2b, 0x19, 0x34, 0x5d, 0x7e, 0x32, 0x99, 0x49, 0x4b, 0x49, 0x44, 0x1f, 0x33, 0x14, 0x8a, 0xad,
	0xb0, 0x04, 0x8d, 0x93, 0x49, 0x94, 0x11, 0xd4, 0xb4, 0xe4, 0xfc, 0x5c, 0x7d, 0x84, 0x5d, 0xfa,
	0x50, 0xbb, 0xc0, 0x51, 0x0a, 0x51, 0x61, 0x0d, 0xa1, 0xe2, 0x0b, 0x92, 0x92, 0xd8, 0xc0, 0x11,
	0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x93, 0xac, 0x9e, 0x37, 0x2d, 0x02, 0x00, 0x00,
}