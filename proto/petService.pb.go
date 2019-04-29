// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/petService.proto

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

type Pet struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Owner                string   `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Species              string   `protobuf:"bytes,4,opt,name=species,proto3" json:"species,omitempty"`
	Sex                  string   `protobuf:"bytes,5,opt,name=sex,proto3" json:"sex,omitempty"`
	Birth                string   `protobuf:"bytes,6,opt,name=birth,proto3" json:"birth,omitempty"`
	Death                string   `protobuf:"bytes,7,opt,name=death,proto3" json:"death,omitempty"`
	Created              string   `protobuf:"bytes,8,opt,name=created,proto3" json:"created,omitempty"`
	Updated              string   `protobuf:"bytes,9,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pet) Reset()         { *m = Pet{} }
func (m *Pet) String() string { return proto.CompactTextString(m) }
func (*Pet) ProtoMessage()    {}
func (*Pet) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeb53027c03a7ba6, []int{0}
}

func (m *Pet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pet.Unmarshal(m, b)
}
func (m *Pet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pet.Marshal(b, m, deterministic)
}
func (m *Pet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pet.Merge(m, src)
}
func (m *Pet) XXX_Size() int {
	return xxx_messageInfo_Pet.Size(m)
}
func (m *Pet) XXX_DiscardUnknown() {
	xxx_messageInfo_Pet.DiscardUnknown(m)
}

var xxx_messageInfo_Pet proto.InternalMessageInfo

func (m *Pet) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Pet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Pet) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Pet) GetSpecies() string {
	if m != nil {
		return m.Species
	}
	return ""
}

func (m *Pet) GetSex() string {
	if m != nil {
		return m.Sex
	}
	return ""
}

func (m *Pet) GetBirth() string {
	if m != nil {
		return m.Birth
	}
	return ""
}

func (m *Pet) GetDeath() string {
	if m != nil {
		return m.Death
	}
	return ""
}

func (m *Pet) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *Pet) GetUpdated() string {
	if m != nil {
		return m.Updated
	}
	return ""
}

type GetPetsResponse struct {
	Pets                 []*Pet           `protobuf:"bytes,1,rep,name=pets,proto3" json:"pets,omitempty"`
	Length               int32            `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
	Meta                 *ResponsePayload `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetPetsResponse) Reset()         { *m = GetPetsResponse{} }
func (m *GetPetsResponse) String() string { return proto.CompactTextString(m) }
func (*GetPetsResponse) ProtoMessage()    {}
func (*GetPetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeb53027c03a7ba6, []int{1}
}

func (m *GetPetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPetsResponse.Unmarshal(m, b)
}
func (m *GetPetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPetsResponse.Marshal(b, m, deterministic)
}
func (m *GetPetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPetsResponse.Merge(m, src)
}
func (m *GetPetsResponse) XXX_Size() int {
	return xxx_messageInfo_GetPetsResponse.Size(m)
}
func (m *GetPetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPetsResponse proto.InternalMessageInfo

func (m *GetPetsResponse) GetPets() []*Pet {
	if m != nil {
		return m.Pets
	}
	return nil
}

func (m *GetPetsResponse) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *GetPetsResponse) GetMeta() *ResponsePayload {
	if m != nil {
		return m.Meta
	}
	return nil
}

type PostPetRequest struct {
	Pet                  *Pet     `protobuf:"bytes,1,opt,name=pet,proto3" json:"pet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostPetRequest) Reset()         { *m = PostPetRequest{} }
func (m *PostPetRequest) String() string { return proto.CompactTextString(m) }
func (*PostPetRequest) ProtoMessage()    {}
func (*PostPetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeb53027c03a7ba6, []int{2}
}

func (m *PostPetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostPetRequest.Unmarshal(m, b)
}
func (m *PostPetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostPetRequest.Marshal(b, m, deterministic)
}
func (m *PostPetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostPetRequest.Merge(m, src)
}
func (m *PostPetRequest) XXX_Size() int {
	return xxx_messageInfo_PostPetRequest.Size(m)
}
func (m *PostPetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostPetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostPetRequest proto.InternalMessageInfo

func (m *PostPetRequest) GetPet() *Pet {
	if m != nil {
		return m.Pet
	}
	return nil
}

type PostPetResponse struct {
	Pet                  *Pet             `protobuf:"bytes,1,opt,name=pet,proto3" json:"pet,omitempty"`
	Meta                 *ResponsePayload `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PostPetResponse) Reset()         { *m = PostPetResponse{} }
func (m *PostPetResponse) String() string { return proto.CompactTextString(m) }
func (*PostPetResponse) ProtoMessage()    {}
func (*PostPetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeb53027c03a7ba6, []int{3}
}

func (m *PostPetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostPetResponse.Unmarshal(m, b)
}
func (m *PostPetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostPetResponse.Marshal(b, m, deterministic)
}
func (m *PostPetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostPetResponse.Merge(m, src)
}
func (m *PostPetResponse) XXX_Size() int {
	return xxx_messageInfo_PostPetResponse.Size(m)
}
func (m *PostPetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostPetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostPetResponse proto.InternalMessageInfo

func (m *PostPetResponse) GetPet() *Pet {
	if m != nil {
		return m.Pet
	}
	return nil
}

func (m *PostPetResponse) GetMeta() *ResponsePayload {
	if m != nil {
		return m.Meta
	}
	return nil
}

func init() {
	proto.RegisterType((*Pet)(nil), "proto.Pet")
	proto.RegisterType((*GetPetsResponse)(nil), "proto.GetPetsResponse")
	proto.RegisterType((*PostPetRequest)(nil), "proto.PostPetRequest")
	proto.RegisterType((*PostPetResponse)(nil), "proto.PostPetResponse")
}

func init() { proto.RegisterFile("proto/petService.proto", fileDescriptor_eeb53027c03a7ba6) }

var fileDescriptor_eeb53027c03a7ba6 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x55, 0x92, 0xa6, 0xdd, 0x9d, 0x15, 0xbb, 0x95, 0x81, 0xca, 0x2a, 0x2b, 0x54, 0xe5, 0x54,
	0xed, 0x21, 0x95, 0xca, 0x8d, 0x3b, 0xe2, 0x80, 0x90, 0xa2, 0x70, 0xe4, 0xe4, 0x36, 0x43, 0x6b,
	0xa9, 0x8d, 0x4d, 0x3c, 0x01, 0xf6, 0xca, 0x2f, 0xf0, 0x13, 0xfc, 0x0b, 0x47, 0x7e, 0x81, 0x0f,
	0x41, 0x1e, 0x3b, 0x91, 0x76, 0x0f, 0x70, 0xb2, 0xdf, 0x9b, 0xf1, 0xf3, 0x9b, 0xd1, 0x83, 0x85,
	0xed, 0x0c, 0x99, 0x8d, 0x45, 0xfa, 0x80, 0xdd, 0x17, 0xbd, 0xc7, 0x92, 0x09, 0x91, 0xf3, 0xb1,
	0xbc, 0x3d, 0x18, 0x73, 0x38, 0xe1, 0x46, 0x59, 0xbd, 0x51, 0x6d, 0x6b, 0x48, 0x91, 0x36, 0xad,
	0x0b, 0x4d, 0xcb, 0x17, 0xb1, 0xca, 0x68, 0xd7, 0x7f, 0xda, 0xe0, 0xd9, 0xd2, 0x7d, 0x2c, 0xce,
	0x83, 0x72, 0x4f, 0xfa, 0x14, 0x98, 0xe2, 0x57, 0x02, 0x59, 0x85, 0x24, 0xae, 0x21, 0xd5, 0x8d,
	0x4c, 0x56, 0xc9, 0x3a, 0xaf, 0x53, 0xdd, 0x08, 0x01, 0x93, 0x56, 0x9d, 0x51, 0xa6, 0xab, 0x64,
	0x7d, 0x59, 0xf3, 0x5d, 0x3c, 0x83, 0xdc, 0x7c, 0x6d, 0xb1, 0x93, 0x19, 0x93, 0x01, 0x08, 0x09,
	0x33, 0x67, 0x71, 0xaf, 0xd1, 0xc9, 0x09, 0xf3, 0x03, 0x14, 0x73, 0xc8, 0x1c, 0x7e, 0x93, 0x39,
	0xb3, 0xfe, 0xea, 0x15, 0x76, 0xba, 0xa3, 0xa3, 0x9c, 0x06, 0x05, 0x06, 0x9e, 0x6d, 0x50, 0xd1,
	0x51, 0xce, 0x02, 0xcb, 0xc0, 0xeb, 0xee, 0x3b, 0x54, 0x84, 0x8d, 0xbc, 0x08, 0xba, 0x11, 0xfa,
	0x4a, 0x6f, 0x1b, 0xae, 0x5c, 0x86, 0x4a, 0x84, 0x45, 0x0f, 0x37, 0x6f, 0x91, 0x2a, 0x24, 0x57,
	0xa3, 0xb3, 0xa6, 0x75, 0x28, 0x5e, 0xc2, 0xc4, 0x22, 0x39, 0x99, 0xac, 0xb2, 0xf5, 0xd5, 0x16,
	0xc2, 0xd8, 0x65, 0x85, 0x54, 0x33, 0x2f, 0x16, 0x30, 0x3d, 0x61, 0x7b, 0xa0, 0x23, 0x8f, 0x9a,
	0xd7, 0x11, 0x89, 0x3b, 0x98, 0x9c, 0x91, 0x14, 0xcf, 0x7a, 0xb5, 0x5d, 0xc4, 0x77, 0x83, 0x6c,
	0xa5, 0xee, 0x4f, 0x46, 0x35, 0x35, 0xf7, 0x14, 0x25, 0x5c, 0x57, 0xc6, 0xf9, 0x7f, 0x6b, 0xfc,
	0xdc, 0xa3, 0x23, 0x71, 0x0b, 0x99, 0x45, 0xe2, 0x7d, 0x3e, 0xfc, 0xd4, 0xd3, 0xc5, 0x47, 0xb8,
	0x19, 0xfb, 0xa3, 0xcd, 0x7f, 0x3e, 0x18, 0xcd, 0xa4, 0xff, 0x37, 0xb3, 0xfd, 0x99, 0x00, 0x54,
	0x63, 0x74, 0xc4, 0x3b, 0x98, 0xc5, 0x95, 0x88, 0x45, 0x19, 0xb2, 0x51, 0x0e, 0xd9, 0x28, 0xdf,
	0xf8, 0x6c, 0x2c, 0x07, 0xbd, 0x47, 0xab, 0x2b, 0xe6, 0xdf, 0x7f, 0xff, 0xf9, 0x91, 0x82, 0xb8,
	0xe0, 0xa8, 0x79, 0x1f, 0xef, 0x61, 0x16, 0x8d, 0x8b, 0xe7, 0x83, 0xc7, 0x07, 0x83, 0x8f, 0x5a,
	0x8f, 0xe6, 0x2b, 0x9e, 0xb2, 0xd6, 0x93, 0x62, 0xd4, 0x7a, 0x9d, 0xdc, 0xed, 0xa6, 0xdc, 0xfb,
	0xea, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f, 0xa5, 0x4b, 0x6f, 0xf1, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PetServiceClient is the client API for PetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PetServiceClient interface {
	GetPets(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetPetsResponse, error)
	PostPet(ctx context.Context, in *PostPetRequest, opts ...grpc.CallOption) (*PostPetResponse, error)
}

type petServiceClient struct {
	cc *grpc.ClientConn
}

func NewPetServiceClient(cc *grpc.ClientConn) PetServiceClient {
	return &petServiceClient{cc}
}

func (c *petServiceClient) GetPets(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetPetsResponse, error) {
	out := new(GetPetsResponse)
	err := c.cc.Invoke(ctx, "/proto.PetService/GetPets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petServiceClient) PostPet(ctx context.Context, in *PostPetRequest, opts ...grpc.CallOption) (*PostPetResponse, error) {
	out := new(PostPetResponse)
	err := c.cc.Invoke(ctx, "/proto.PetService/PostPet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PetServiceServer is the server API for PetService service.
type PetServiceServer interface {
	GetPets(context.Context, *empty.Empty) (*GetPetsResponse, error)
	PostPet(context.Context, *PostPetRequest) (*PostPetResponse, error)
}

func RegisterPetServiceServer(s *grpc.Server, srv PetServiceServer) {
	s.RegisterService(&_PetService_serviceDesc, srv)
}

func _PetService_GetPets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetServiceServer).GetPets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PetService/GetPets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetServiceServer).GetPets(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetService_PostPet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostPetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetServiceServer).PostPet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PetService/PostPet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetServiceServer).PostPet(ctx, req.(*PostPetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PetService",
	HandlerType: (*PetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPets",
			Handler:    _PetService_GetPets_Handler,
		},
		{
			MethodName: "PostPet",
			Handler:    _PetService_PostPet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/petService.proto",
}