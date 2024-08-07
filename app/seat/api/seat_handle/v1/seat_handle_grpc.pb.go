// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: seat_handle/v1/seat_handle.proto

package v1

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

// SeatHandleClient is the client API for SeatHandle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SeatHandleClient interface {
	Reserve(ctx context.Context, in *ReserveReq, opts ...grpc.CallOption) (*ReserveRes, error)
}

type seatHandleClient struct {
	cc grpc.ClientConnInterface
}

func NewSeatHandleClient(cc grpc.ClientConnInterface) SeatHandleClient {
	return &seatHandleClient{cc}
}

func (c *seatHandleClient) Reserve(ctx context.Context, in *ReserveReq, opts ...grpc.CallOption) (*ReserveRes, error) {
	out := new(ReserveRes)
	err := c.cc.Invoke(ctx, "/seat_handle.v1.SeatHandle/Reserve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SeatHandleServer is the server API for SeatHandle service.
// All implementations must embed UnimplementedSeatHandleServer
// for forward compatibility
type SeatHandleServer interface {
	Reserve(context.Context, *ReserveReq) (*ReserveRes, error)
	mustEmbedUnimplementedSeatHandleServer()
}

// UnimplementedSeatHandleServer must be embedded to have forward compatible implementations.
type UnimplementedSeatHandleServer struct {
}

func (UnimplementedSeatHandleServer) Reserve(context.Context, *ReserveReq) (*ReserveRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reserve not implemented")
}
func (UnimplementedSeatHandleServer) mustEmbedUnimplementedSeatHandleServer() {}

// UnsafeSeatHandleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SeatHandleServer will
// result in compilation errors.
type UnsafeSeatHandleServer interface {
	mustEmbedUnimplementedSeatHandleServer()
}

func RegisterSeatHandleServer(s grpc.ServiceRegistrar, srv SeatHandleServer) {
	s.RegisterService(&SeatHandle_ServiceDesc, srv)
}

func _SeatHandle_Reserve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReserveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeatHandleServer).Reserve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seat_handle.v1.SeatHandle/Reserve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeatHandleServer).Reserve(ctx, req.(*ReserveReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SeatHandle_ServiceDesc is the grpc.ServiceDesc for SeatHandle service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SeatHandle_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "seat_handle.v1.SeatHandle",
	HandlerType: (*SeatHandleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reserve",
			Handler:    _SeatHandle_Reserve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "seat_handle/v1/seat_handle.proto",
}
