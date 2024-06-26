// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: policy_prepare/v1/policy_prepare.proto

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

// PolicyPrepareClient is the client API for PolicyPrepare service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PolicyPrepareClient interface {
	Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateRes, error)
	GetOne(ctx context.Context, in *GetOneReq, opts ...grpc.CallOption) (*GetOneRes, error)
	GetList(ctx context.Context, in *GetListReq, opts ...grpc.CallOption) (*GetListRes, error)
	Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateRes, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteRes, error)
}

type policyPrepareClient struct {
	cc grpc.ClientConnInterface
}

func NewPolicyPrepareClient(cc grpc.ClientConnInterface) PolicyPrepareClient {
	return &policyPrepareClient{cc}
}

func (c *policyPrepareClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateRes, error) {
	out := new(CreateRes)
	err := c.cc.Invoke(ctx, "/policy_prepare.v1.PolicyPrepare/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyPrepareClient) GetOne(ctx context.Context, in *GetOneReq, opts ...grpc.CallOption) (*GetOneRes, error) {
	out := new(GetOneRes)
	err := c.cc.Invoke(ctx, "/policy_prepare.v1.PolicyPrepare/GetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyPrepareClient) GetList(ctx context.Context, in *GetListReq, opts ...grpc.CallOption) (*GetListRes, error) {
	out := new(GetListRes)
	err := c.cc.Invoke(ctx, "/policy_prepare.v1.PolicyPrepare/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyPrepareClient) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateRes, error) {
	out := new(UpdateRes)
	err := c.cc.Invoke(ctx, "/policy_prepare.v1.PolicyPrepare/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyPrepareClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteRes, error) {
	out := new(DeleteRes)
	err := c.cc.Invoke(ctx, "/policy_prepare.v1.PolicyPrepare/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolicyPrepareServer is the server API for PolicyPrepare service.
// All implementations must embed UnimplementedPolicyPrepareServer
// for forward compatibility
type PolicyPrepareServer interface {
	Create(context.Context, *CreateReq) (*CreateRes, error)
	GetOne(context.Context, *GetOneReq) (*GetOneRes, error)
	GetList(context.Context, *GetListReq) (*GetListRes, error)
	Update(context.Context, *UpdateReq) (*UpdateRes, error)
	Delete(context.Context, *DeleteReq) (*DeleteRes, error)
	mustEmbedUnimplementedPolicyPrepareServer()
}

// UnimplementedPolicyPrepareServer must be embedded to have forward compatible implementations.
type UnimplementedPolicyPrepareServer struct {
}

func (UnimplementedPolicyPrepareServer) Create(context.Context, *CreateReq) (*CreateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPolicyPrepareServer) GetOne(context.Context, *GetOneReq) (*GetOneRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (UnimplementedPolicyPrepareServer) GetList(context.Context, *GetListReq) (*GetListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedPolicyPrepareServer) Update(context.Context, *UpdateReq) (*UpdateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPolicyPrepareServer) Delete(context.Context, *DeleteReq) (*DeleteRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPolicyPrepareServer) mustEmbedUnimplementedPolicyPrepareServer() {}

// UnsafePolicyPrepareServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PolicyPrepareServer will
// result in compilation errors.
type UnsafePolicyPrepareServer interface {
	mustEmbedUnimplementedPolicyPrepareServer()
}

func RegisterPolicyPrepareServer(s grpc.ServiceRegistrar, srv PolicyPrepareServer) {
	s.RegisterService(&PolicyPrepare_ServiceDesc, srv)
}

func _PolicyPrepare_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyPrepareServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/policy_prepare.v1.PolicyPrepare/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyPrepareServer).Create(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyPrepare_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyPrepareServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/policy_prepare.v1.PolicyPrepare/GetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyPrepareServer).GetOne(ctx, req.(*GetOneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyPrepare_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyPrepareServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/policy_prepare.v1.PolicyPrepare/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyPrepareServer).GetList(ctx, req.(*GetListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyPrepare_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyPrepareServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/policy_prepare.v1.PolicyPrepare/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyPrepareServer).Update(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyPrepare_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyPrepareServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/policy_prepare.v1.PolicyPrepare/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyPrepareServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PolicyPrepare_ServiceDesc is the grpc.ServiceDesc for PolicyPrepare service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PolicyPrepare_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "policy_prepare.v1.PolicyPrepare",
	HandlerType: (*PolicyPrepareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PolicyPrepare_Create_Handler,
		},
		{
			MethodName: "GetOne",
			Handler:    _PolicyPrepare_GetOne_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _PolicyPrepare_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PolicyPrepare_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PolicyPrepare_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "policy_prepare/v1/policy_prepare.proto",
}
