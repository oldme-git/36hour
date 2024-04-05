// protoc --go_out=plugins=grpc:. *.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: hall/v1/hall.proto

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
	Hall_Create_FullMethodName  = "/user.Hall/Create"
	Hall_GetOne_FullMethodName  = "/user.Hall/GetOne"
	Hall_GetList_FullMethodName = "/user.Hall/GetList"
	Hall_Update_FullMethodName  = "/user.Hall/Update"
	Hall_Delete_FullMethodName  = "/user.Hall/Delete"
)

// HallClient is the client API for Hall service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HallClient interface {
	Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateRes, error)
	GetOne(ctx context.Context, in *GetOneReq, opts ...grpc.CallOption) (*GetOneRes, error)
	GetList(ctx context.Context, in *GetListReq, opts ...grpc.CallOption) (*GetListRes, error)
	Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateRes, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteRes, error)
}

type hallClient struct {
	cc grpc.ClientConnInterface
}

func NewHallClient(cc grpc.ClientConnInterface) HallClient {
	return &hallClient{cc}
}

func (c *hallClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateRes, error) {
	out := new(CreateRes)
	err := c.cc.Invoke(ctx, Hall_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hallClient) GetOne(ctx context.Context, in *GetOneReq, opts ...grpc.CallOption) (*GetOneRes, error) {
	out := new(GetOneRes)
	err := c.cc.Invoke(ctx, Hall_GetOne_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hallClient) GetList(ctx context.Context, in *GetListReq, opts ...grpc.CallOption) (*GetListRes, error) {
	out := new(GetListRes)
	err := c.cc.Invoke(ctx, Hall_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hallClient) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateRes, error) {
	out := new(UpdateRes)
	err := c.cc.Invoke(ctx, Hall_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hallClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteRes, error) {
	out := new(DeleteRes)
	err := c.cc.Invoke(ctx, Hall_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HallServer is the server API for Hall service.
// All implementations must embed UnimplementedHallServer
// for forward compatibility
type HallServer interface {
	Create(context.Context, *CreateReq) (*CreateRes, error)
	GetOne(context.Context, *GetOneReq) (*GetOneRes, error)
	GetList(context.Context, *GetListReq) (*GetListRes, error)
	Update(context.Context, *UpdateReq) (*UpdateRes, error)
	Delete(context.Context, *DeleteReq) (*DeleteRes, error)
	mustEmbedUnimplementedHallServer()
}

// UnimplementedHallServer must be embedded to have forward compatible implementations.
type UnimplementedHallServer struct {
}

func (UnimplementedHallServer) Create(context.Context, *CreateReq) (*CreateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedHallServer) GetOne(context.Context, *GetOneReq) (*GetOneRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (UnimplementedHallServer) GetList(context.Context, *GetListReq) (*GetListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedHallServer) Update(context.Context, *UpdateReq) (*UpdateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedHallServer) Delete(context.Context, *DeleteReq) (*DeleteRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedHallServer) mustEmbedUnimplementedHallServer() {}

// UnsafeHallServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HallServer will
// result in compilation errors.
type UnsafeHallServer interface {
	mustEmbedUnimplementedHallServer()
}

func RegisterHallServer(s grpc.ServiceRegistrar, srv HallServer) {
	s.RegisterService(&Hall_ServiceDesc, srv)
}

func _Hall_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HallServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hall_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HallServer).Create(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hall_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HallServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hall_GetOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HallServer).GetOne(ctx, req.(*GetOneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hall_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HallServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hall_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HallServer).GetList(ctx, req.(*GetListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hall_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HallServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hall_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HallServer).Update(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hall_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HallServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hall_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HallServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Hall_ServiceDesc is the grpc.ServiceDesc for Hall service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hall_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.Hall",
	HandlerType: (*HallServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Hall_Create_Handler,
		},
		{
			MethodName: "GetOne",
			Handler:    _Hall_GetOne_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _Hall_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Hall_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Hall_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hall/v1/hall.proto",
}