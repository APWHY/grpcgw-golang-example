// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/eventService.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type PostEventRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostEventRequest) Reset()         { *m = PostEventRequest{} }
func (m *PostEventRequest) String() string { return proto.CompactTextString(m) }
func (*PostEventRequest) ProtoMessage()    {}
func (*PostEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a46713ac5409664b, []int{0}
}

func (m *PostEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostEventRequest.Unmarshal(m, b)
}
func (m *PostEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostEventRequest.Marshal(b, m, deterministic)
}
func (m *PostEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostEventRequest.Merge(m, src)
}
func (m *PostEventRequest) XXX_Size() int {
	return xxx_messageInfo_PostEventRequest.Size(m)
}
func (m *PostEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostEventRequest proto.InternalMessageInfo

func (m *PostEventRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*PostEventRequest)(nil), "proto.PostEventRequest")
}

func init() { proto.RegisterFile("proto/eventService.proto", fileDescriptor_a46713ac5409664b) }

var fileDescriptor_a46713ac5409664b = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0x09, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x03,
	0x0b, 0x09, 0xb1, 0x82, 0x29, 0x29, 0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0xfd, 0xc4, 0x82,
	0x4c, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88, 0x22, 0x29,
	0x69, 0xa8, 0x2c, 0x98, 0x97, 0x54, 0x9a, 0xa6, 0x9f, 0x9a, 0x5b, 0x50, 0x52, 0x09, 0x91, 0x54,
	0x52, 0xe3, 0x12, 0x08, 0xc8, 0x2f, 0x2e, 0x71, 0x05, 0x99, 0x1d, 0x94, 0x5a, 0x58, 0x9a, 0x5a,
	0x5c, 0x22, 0x24, 0xc4, 0xc5, 0x92, 0x92, 0x58, 0x92, 0x28, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19,
	0x04, 0x66, 0x1b, 0xa5, 0x72, 0xf1, 0xb8, 0x22, 0xd9, 0x2f, 0x14, 0xca, 0xc5, 0x09, 0xd7, 0x27,
	0x24, 0x0e, 0x31, 0x4c, 0x0f, 0xdd, 0x24, 0x29, 0x31, 0x3d, 0x88, 0xdd, 0x7a, 0x30, 0xbb, 0xf5,
	0x5c, 0x41, 0x76, 0x2b, 0x89, 0x36, 0x5d, 0x7e, 0x32, 0x99, 0x89, 0x5f, 0x89, 0x0b, 0xec, 0x64,
	0xb0, 0xc7, 0xac, 0x18, 0xb5, 0x34, 0x18, 0x93, 0xd8, 0xc0, 0x0a, 0x8d, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xfe, 0xc9, 0x72, 0x14, 0xf3, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventServiceClient interface {
	PostEvent(ctx context.Context, opts ...grpc.CallOption) (EventService_PostEventClient, error)
}

type eventServiceClient struct {
	cc *grpc.ClientConn
}

func NewEventServiceClient(cc *grpc.ClientConn) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) PostEvent(ctx context.Context, opts ...grpc.CallOption) (EventService_PostEventClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EventService_serviceDesc.Streams[0], "/proto.EventService/PostEvent", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventServicePostEventClient{stream}
	return x, nil
}

type EventService_PostEventClient interface {
	Send(*PostEventRequest) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type eventServicePostEventClient struct {
	grpc.ClientStream
}

func (x *eventServicePostEventClient) Send(m *PostEventRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventServicePostEventClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventServiceServer is the server API for EventService service.
type EventServiceServer interface {
	PostEvent(EventService_PostEventServer) error
}

func RegisterEventServiceServer(s *grpc.Server, srv EventServiceServer) {
	s.RegisterService(&_EventService_serviceDesc, srv)
}

func _EventService_PostEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventServiceServer).PostEvent(&eventServicePostEventServer{stream})
}

type EventService_PostEventServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*PostEventRequest, error)
	grpc.ServerStream
}

type eventServicePostEventServer struct {
	grpc.ServerStream
}

func (x *eventServicePostEventServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventServicePostEventServer) Recv() (*PostEventRequest, error) {
	m := new(PostEventRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EventService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PostEvent",
			Handler:       _EventService_PostEvent_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/eventService.proto",
}
