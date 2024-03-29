// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pbsample

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SampleClient is the client API for Sample service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SampleClient interface {
	Ping(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error)
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Test, error)
}

type sampleClient struct {
	cc grpc.ClientConnInterface
}

func NewSampleClient(cc grpc.ClientConnInterface) SampleClient {
	return &sampleClient{cc}
}

func (c *sampleClient) Ping(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error) {
	out := new(Test)
	err := c.cc.Invoke(ctx, "/pbsample.Sample/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Test, error) {
	out := new(Test)
	err := c.cc.Invoke(ctx, "/pbsample.Sample/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SampleServer is the server API for Sample service.
// All implementations must embed UnimplementedSampleServer
// for forward compatibility
type SampleServer interface {
	Ping(context.Context, *Test) (*Test, error)
	Version(context.Context, *emptypb.Empty) (*Test, error)
	mustEmbedUnimplementedSampleServer()
}

// UnimplementedSampleServer must be embedded to have forward compatible implementations.
type UnimplementedSampleServer struct {
}

func (UnimplementedSampleServer) Ping(context.Context, *Test) (*Test, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedSampleServer) Version(context.Context, *emptypb.Empty) (*Test, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedSampleServer) mustEmbedUnimplementedSampleServer() {}

// UnsafeSampleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SampleServer will
// result in compilation errors.
type UnsafeSampleServer interface {
	mustEmbedUnimplementedSampleServer()
}

func RegisterSampleServer(s grpc.ServiceRegistrar, srv SampleServer) {
	s.RegisterService(&Sample_ServiceDesc, srv)
}

func _Sample_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Test)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbsample.Sample/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServer).Ping(ctx, req.(*Test))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sample_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbsample.Sample/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Sample_ServiceDesc is the grpc.ServiceDesc for Sample service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sample_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbsample.Sample",
	HandlerType: (*SampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Sample_Ping_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _Sample_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pbsample/sample.proto",
}
