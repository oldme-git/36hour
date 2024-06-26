// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: lib/v1/lib.proto

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

// LibClient is the client API for Lib service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LibClient interface {
	Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateRes, error)
	GetOne(ctx context.Context, in *GetOneReq, opts ...grpc.CallOption) (*GetOneRes, error)
	GetList(ctx context.Context, in *GetListReq, opts ...grpc.CallOption) (*GetListRes, error)
	Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateRes, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteRes, error)
}

type libClient struct {
	cc grpc.ClientConnInterface
}

func NewLibClient(cc grpc.ClientConnInterface) LibClient {
	return &libClient{cc}
}

func (c *libClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateRes, error) {
	out := new(CreateRes)
	err := c.cc.Invoke(ctx, "/lib.v1.Lib/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libClient) GetOne(ctx context.Context, in *GetOneReq, opts ...grpc.CallOption) (*GetOneRes, error) {
	out := new(GetOneRes)
	err := c.cc.Invoke(ctx, "/lib.v1.Lib/GetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libClient) GetList(ctx context.Context, in *GetListReq, opts ...grpc.CallOption) (*GetListRes, error) {
	out := new(GetListRes)
	err := c.cc.Invoke(ctx, "/lib.v1.Lib/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libClient) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateRes, error) {
	out := new(UpdateRes)
	err := c.cc.Invoke(ctx, "/lib.v1.Lib/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteRes, error) {
	out := new(DeleteRes)
	err := c.cc.Invoke(ctx, "/lib.v1.Lib/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LibServer is the server API for Lib service.
// All implementations must embed UnimplementedLibServer
// for forward compatibility
type LibServer interface {
	Create(context.Context, *CreateReq) (*CreateRes, error)
	GetOne(context.Context, *GetOneReq) (*GetOneRes, error)
	GetList(context.Context, *GetListReq) (*GetListRes, error)
	Update(context.Context, *UpdateReq) (*UpdateRes, error)
	Delete(context.Context, *DeleteReq) (*DeleteRes, error)
	mustEmbedUnimplementedLibServer()
}

// UnimplementedLibServer must be embedded to have forward compatible implementations.
type UnimplementedLibServer struct {
}

func (UnimplementedLibServer) Create(context.Context, *CreateReq) (*CreateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedLibServer) GetOne(context.Context, *GetOneReq) (*GetOneRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (UnimplementedLibServer) GetList(context.Context, *GetListReq) (*GetListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedLibServer) Update(context.Context, *UpdateReq) (*UpdateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedLibServer) Delete(context.Context, *DeleteReq) (*DeleteRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedLibServer) mustEmbedUnimplementedLibServer() {}

// UnsafeLibServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LibServer will
// result in compilation errors.
type UnsafeLibServer interface {
	mustEmbedUnimplementedLibServer()
}

func RegisterLibServer(s grpc.ServiceRegistrar, srv LibServer) {
	s.RegisterService(&Lib_ServiceDesc, srv)
}

func _Lib_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lib.v1.Lib/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibServer).Create(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lib_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lib.v1.Lib/GetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibServer).GetOne(ctx, req.(*GetOneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lib_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lib.v1.Lib/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibServer).GetList(ctx, req.(*GetListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lib_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lib.v1.Lib/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibServer).Update(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lib_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lib.v1.Lib/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Lib_ServiceDesc is the grpc.ServiceDesc for Lib service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Lib_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lib.v1.Lib",
	HandlerType: (*LibServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Lib_Create_Handler,
		},
		{
			MethodName: "GetOne",
			Handler:    _Lib_GetOne_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _Lib_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Lib_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Lib_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lib/v1/lib.proto",
}
