// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Person struct {
	Id                   *int32   `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Name                 *string  `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Address              *string  `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_fe6c136661490b64, []int{0}
}
func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (dst *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(dst, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Person) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Person) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_fe6c136661490b64, []int{1}
}
func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (dst *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(dst, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Person)(nil), "service.Person")
	proto.RegisterType((*Void)(nil), "service.Void")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataClient is the client API for Data service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataClient interface {
	GetPerson(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Person, error)
}

type dataClient struct {
	cc *grpc.ClientConn
}

func NewDataClient(cc *grpc.ClientConn) DataClient {
	return &dataClient{cc}
}

func (c *dataClient) GetPerson(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Person, error) {
	out := new(Person)
	err := c.cc.Invoke(ctx, "/service.Data/GetPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataServer is the server API for Data service.
type DataServer interface {
	GetPerson(context.Context, *Void) (*Person, error)
}

func RegisterDataServer(s *grpc.Server, srv DataServer) {
	s.RegisterService(&_Data_serviceDesc, srv)
}

func _Data_GetPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).GetPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Data/GetPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).GetPerson(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _Data_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Data",
	HandlerType: (*DataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPerson",
			Handler:    _Data_GetPerson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_fe6c136661490b64) }

var fileDescriptor_service_fe6c136661490b64 = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0xdc, 0xb8,
	0xd8, 0x02, 0x52, 0x8b, 0x8a, 0xf3, 0xf3, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15,
	0x98, 0x34, 0x58, 0x83, 0x98, 0x32, 0x53, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25,
	0x98, 0x14, 0x98, 0x34, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x09, 0x2e, 0xf6, 0xc4, 0x94, 0x94, 0xa2,
	0xd4, 0xe2, 0x62, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x89, 0x8d, 0x8b, 0x25,
	0x2c, 0x3f, 0x33, 0xc5, 0xc8, 0x98, 0x8b, 0xc5, 0x25, 0xb1, 0x24, 0x51, 0x48, 0x9b, 0x8b, 0xd3,
	0x3d, 0xb5, 0x04, 0x6a, 0x34, 0xaf, 0x1e, 0xcc, 0x76, 0x90, 0x1a, 0x29, 0x7e, 0x38, 0x17, 0x22,
	0x0f, 0x08, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x3a, 0x21, 0x76, 0x9d, 0x00, 0x00, 0x00,
}