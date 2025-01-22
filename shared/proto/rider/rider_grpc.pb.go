// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: rider.proto

package rider

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RiderServiceClient is the client API for RiderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RiderServiceClient interface {
	GetNearbyRiders(ctx context.Context, in *RiderRequest, opts ...grpc.CallOption) (*RiderResponse, error)
}

type riderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRiderServiceClient(cc grpc.ClientConnInterface) RiderServiceClient {
	return &riderServiceClient{cc}
}

func (c *riderServiceClient) GetNearbyRiders(ctx context.Context, in *RiderRequest, opts ...grpc.CallOption) (*RiderResponse, error) {
	out := new(RiderResponse)
	err := c.cc.Invoke(ctx, "/rider.RiderService/GetNearbyRiders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RiderServiceServer is the server API for RiderService service.
// All implementations must embed UnimplementedRiderServiceServer
// for forward compatibility
type RiderServiceServer interface {
	GetNearbyRiders(context.Context, *RiderRequest) (*RiderResponse, error)
	mustEmbedUnimplementedRiderServiceServer()
}

// UnimplementedRiderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRiderServiceServer struct {
}

func (UnimplementedRiderServiceServer) GetNearbyRiders(context.Context, *RiderRequest) (*RiderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNearbyRiders not implemented")
}
func (UnimplementedRiderServiceServer) mustEmbedUnimplementedRiderServiceServer() {}

// UnsafeRiderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RiderServiceServer will
// result in compilation errors.
type UnsafeRiderServiceServer interface {
	mustEmbedUnimplementedRiderServiceServer()
}

func RegisterRiderServiceServer(s grpc.ServiceRegistrar, srv RiderServiceServer) {
	s.RegisterService(&RiderService_ServiceDesc, srv)
}

func _RiderService_GetNearbyRiders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RiderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiderServiceServer).GetNearbyRiders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rider.RiderService/GetNearbyRiders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiderServiceServer).GetNearbyRiders(ctx, req.(*RiderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RiderService_ServiceDesc is the grpc.ServiceDesc for RiderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RiderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rider.RiderService",
	HandlerType: (*RiderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNearbyRiders",
			Handler:    _RiderService_GetNearbyRiders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rider.proto",
}
