// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package hello

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

type HelloReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReq) Reset()         { *m = HelloReq{} }
func (m *HelloReq) String() string { return proto.CompactTextString(m) }
func (*HelloReq) ProtoMessage()    {}
func (*HelloReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_2750a4f156640bbf, []int{0}
}
func (m *HelloReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReq.Unmarshal(m, b)
}
func (m *HelloReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReq.Marshal(b, m, deterministic)
}
func (dst *HelloReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReq.Merge(dst, src)
}
func (m *HelloReq) XXX_Size() int {
	return xxx_messageInfo_HelloReq.Size(m)
}
func (m *HelloReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReq.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReq proto.InternalMessageInfo

func (m *HelloReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloRsp struct {
	Message              string   `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRsp) Reset()         { *m = HelloRsp{} }
func (m *HelloRsp) String() string { return proto.CompactTextString(m) }
func (*HelloRsp) ProtoMessage()    {}
func (*HelloRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_2750a4f156640bbf, []int{1}
}
func (m *HelloRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRsp.Unmarshal(m, b)
}
func (m *HelloRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRsp.Marshal(b, m, deterministic)
}
func (dst *HelloRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRsp.Merge(dst, src)
}
func (m *HelloRsp) XXX_Size() int {
	return xxx_messageInfo_HelloRsp.Size(m)
}
func (m *HelloRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRsp.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRsp proto.InternalMessageInfo

func (m *HelloRsp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloReq)(nil), "hello.HelloReq")
	proto.RegisterType((*HelloRsp)(nil), "hello.HelloRsp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRsp, error)
	SayHelloStream(ctx context.Context, opts ...grpc.CallOption) (HelloService_SayHelloStreamClient, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRsp, error) {
	out := new(HelloRsp)
	err := c.cc.Invoke(ctx, "/hello.HelloService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) SayHelloStream(ctx context.Context, opts ...grpc.CallOption) (HelloService_SayHelloStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HelloService_serviceDesc.Streams[0], "/hello.HelloService/SayHelloStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceSayHelloStreamClient{stream}
	return x, nil
}

type HelloService_SayHelloStreamClient interface {
	Send(*HelloReq) error
	Recv() (*HelloRsp, error)
	grpc.ClientStream
}

type helloServiceSayHelloStreamClient struct {
	grpc.ClientStream
}

func (x *helloServiceSayHelloStreamClient) Send(m *HelloReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceSayHelloStreamClient) Recv() (*HelloRsp, error) {
	m := new(HelloRsp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	SayHello(context.Context, *HelloReq) (*HelloRsp, error)
	SayHelloStream(HelloService_SayHelloStreamServer) error
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.HelloService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).SayHello(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_SayHelloStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).SayHelloStream(&helloServiceSayHelloStreamServer{stream})
}

type HelloService_SayHelloStreamServer interface {
	Send(*HelloRsp) error
	Recv() (*HelloReq, error)
	grpc.ServerStream
}

type helloServiceSayHelloStreamServer struct {
	grpc.ServerStream
}

func (x *helloServiceSayHelloStreamServer) Send(m *HelloRsp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceSayHelloStreamServer) Recv() (*HelloReq, error) {
	m := new(HelloReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloStream",
			Handler:       _HelloService_SayHelloStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hello.proto",
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_hello_2750a4f156640bbf) }

var fileDescriptor_hello_2750a4f156640bbf = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0xe4, 0xb8, 0x38, 0x3c,
	0x40, 0x8c, 0xa0, 0xd4, 0x42, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49, 0x05, 0x26, 0x5f, 0x5c, 0x20, 0x24, 0xc1, 0xc5, 0x9e,
	0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x0e, 0x53, 0x02, 0xe3, 0x1a, 0x55, 0x70, 0xf1, 0x80, 0x55, 0x05,
	0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0xe9, 0x71, 0x71, 0x04, 0x27, 0x56, 0x82, 0x85, 0x84,
	0xf8, 0xf5, 0x20, 0xd6, 0xc2, 0xac, 0x91, 0x42, 0x15, 0x28, 0x2e, 0x50, 0x62, 0x10, 0xb2, 0xe0,
	0xe2, 0x83, 0xa9, 0x0f, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0x25, 0x46, 0x97, 0x06, 0xa3, 0x01, 0xa3,
	0x13, 0x77, 0x14, 0x67, 0x41, 0x92, 0xb1, 0x3e, 0x58, 0x2e, 0x89, 0x0d, 0xec, 0x35, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x31, 0x75, 0x1e, 0xe9, 0x00, 0x00, 0x00,
}