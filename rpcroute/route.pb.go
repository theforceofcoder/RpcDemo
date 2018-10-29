// Code generated by protoc-gen-go.
// source: route.proto
// DO NOT EDIT!

/*
Package rpcroute is a generated protocol buffer package.

It is generated from these files:
	route.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
*/
package rpcroute

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

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HelloResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*HelloRequest)(nil), "rpcroute.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "rpcroute.HelloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Route service

type RouteClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type routeClient struct {
	cc *grpc.ClientConn
}

func NewRouteClient(cc *grpc.ClientConn) RouteClient {
	return &routeClient{cc}
}

func (c *routeClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := grpc.Invoke(ctx, "/rpcroute.Route/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Route service

type RouteServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterRouteServer(s *grpc.Server, srv RouteServer) {
	s.RegisterService(&_Route_serviceDesc, srv)
}

func _Route_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcroute.Route/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Route_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcroute.Route",
	HandlerType: (*RouteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Route_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("route.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xca, 0x2f, 0x2d,
	0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x2a, 0x48, 0x06, 0xf3, 0x95, 0x94,
	0xb8, 0x78, 0x3c, 0x52, 0x73, 0x72, 0xf2, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84,
	0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25,
	0x4d, 0x2e, 0x5e, 0xa8, 0x9a, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x09, 0x2e, 0xf6, 0xdc,
	0xd4, 0xe2, 0xe2, 0xc4, 0x74, 0x98, 0x3a, 0x18, 0xd7, 0xc8, 0x8d, 0x8b, 0x35, 0x08, 0x64, 0xae,
	0x90, 0x2d, 0x17, 0x47, 0x70, 0x62, 0x25, 0x58, 0x9b, 0x90, 0x98, 0x1e, 0xcc, 0x3a, 0x3d, 0x64,
	0xbb, 0xa4, 0xc4, 0x31, 0xc4, 0x21, 0xe6, 0x2b, 0x31, 0x24, 0xb1, 0x81, 0xdd, 0x69, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x98, 0xfa, 0x6d, 0xf1, 0xb6, 0x00, 0x00, 0x00,
}
