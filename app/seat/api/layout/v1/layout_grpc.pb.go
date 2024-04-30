// protoc --go_out=plugins=grpc:. *.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: layout/v1/layout.proto

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

const (
	Layout_Create_FullMethodName  = "/layout.v1.Layout/Create"
	Layout_GetOne_FullMethodName  = "/layout.v1.Layout/GetOne"
	Layout_GetList_FullMethodName = "/layout.v1.Layout/GetList"
	Layout_Update_FullMethodName  = "/layout.v1.Layout/Update"
	Layout_Delete_FullMethodName  = "/layout.v1.Layout/Delete"
)

// LayoutClient is the client API for Layout service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LayoutClient interface {
	Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateRes, error)
	GetOne(ctx context.Context, in *GetOneReq, opts ...grpc.CallOption) (*GetOneRes, error)
	GetList(ctx context.Context, in *GetListReq, opts ...grpc.CallOption) (*GetListRes, error)
	Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateRes, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteRes, error)
}

type layoutClient struct {
	cc grpc.ClientConnInterface
}

func NewLayoutClient(cc grpc.ClientConnInterface) LayoutClient {
	return &layoutClient{cc}
}

func (c *layoutClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateRes, error) {
	out := new(CreateRes)
	err := c.cc.Invoke(ctx, Layout_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *layoutClient) GetOne(ctx context.Context, in *GetOneReq, opts ...grpc.CallOption) (*GetOneRes, error) {
	out := new(GetOneRes)
	err := c.cc.Invoke(ctx, Layout_GetOne_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *layoutClient) GetList(ctx context.Context, in *GetListReq, opts ...grpc.CallOption) (*GetListRes, error) {
	out := new(GetListRes)
	err := c.cc.Invoke(ctx, Layout_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *layoutClient) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateRes, error) {
	out := new(UpdateRes)
	err := c.cc.Invoke(ctx, Layout_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *layoutClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteRes, error) {
	out := new(DeleteRes)
	err := c.cc.Invoke(ctx, Layout_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LayoutServer is the server API for Layout service.
// All implementations must embed UnimplementedLayoutServer
// for forward compatibility
type LayoutServer interface {
	Create(context.Context, *CreateReq) (*CreateRes, error)
	GetOne(context.Context, *GetOneReq) (*GetOneRes, error)
	GetList(context.Context, *GetListReq) (*GetListRes, error)
	Update(context.Context, *UpdateReq) (*UpdateRes, error)
	Delete(context.Context, *DeleteReq) (*DeleteRes, error)
	mustEmbedUnimplementedLayoutServer()
}

// UnimplementedLayoutServer must be embedded to have forward compatible implementations.
type UnimplementedLayoutServer struct {
}

func (UnimplementedLayoutServer) Create(context.Context, *CreateReq) (*CreateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedLayoutServer) GetOne(context.Context, *GetOneReq) (*GetOneRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (UnimplementedLayoutServer) GetList(context.Context, *GetListReq) (*GetListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedLayoutServer) Update(context.Context, *UpdateReq) (*UpdateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedLayoutServer) Delete(context.Context, *DeleteReq) (*DeleteRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedLayoutServer) mustEmbedUnimplementedLayoutServer() {}

// UnsafeLayoutServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LayoutServer will
// result in compilation errors.
type UnsafeLayoutServer interface {
	mustEmbedUnimplementedLayoutServer()
}

func RegisterLayoutServer(s grpc.ServiceRegistrar, srv LayoutServer) {
	s.RegisterService(&Layout_ServiceDesc, srv)
}

func _Layout_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LayoutServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Layout_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LayoutServer).Create(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Layout_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LayoutServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Layout_GetOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LayoutServer).GetOne(ctx, req.(*GetOneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Layout_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LayoutServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Layout_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LayoutServer).GetList(ctx, req.(*GetListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Layout_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LayoutServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Layout_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LayoutServer).Update(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Layout_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LayoutServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Layout_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LayoutServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Layout_ServiceDesc is the grpc.ServiceDesc for Layout service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Layout_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "layout.v1.Layout",
	HandlerType: (*LayoutServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Layout_Create_Handler,
		},
		{
			MethodName: "GetOne",
			Handler:    _Layout_GetOne_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _Layout_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Layout_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Layout_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "layout/v1/layout.proto",
}