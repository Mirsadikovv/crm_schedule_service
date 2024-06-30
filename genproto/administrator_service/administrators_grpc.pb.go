// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: administrators.proto

package administrator_service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AdministratorServiceClient is the client API for AdministratorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdministratorServiceClient interface {
	Create(ctx context.Context, in *CreateAdministrator, opts ...grpc.CallOption) (*GetAdministrator, error)
	GetByID(ctx context.Context, in *AdministratorPrimaryKey, opts ...grpc.CallOption) (*GetAdministrator, error)
	GetList(ctx context.Context, in *GetListAdministratorRequest, opts ...grpc.CallOption) (*GetListAdministratorResponse, error)
	Update(ctx context.Context, in *UpdateAdministrator, opts ...grpc.CallOption) (*GetAdministrator, error)
	Delete(ctx context.Context, in *AdministratorPrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error)
	Login(ctx context.Context, in *AdministratorLoginRequest, opts ...grpc.CallOption) (*AdministratorLoginResponse, error)
	Register(ctx context.Context, in *AdministratorRegisterRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	RegisterConfirm(ctx context.Context, in *AdministratorRegisterConfRequest, opts ...grpc.CallOption) (*AdministratorLoginResponse, error)
	ChangePassword(ctx context.Context, in *AdministratorChangePassword, opts ...grpc.CallOption) (*AdministratorChangePasswordResp, error)
}

type administratorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdministratorServiceClient(cc grpc.ClientConnInterface) AdministratorServiceClient {
	return &administratorServiceClient{cc}
}

func (c *administratorServiceClient) Create(ctx context.Context, in *CreateAdministrator, opts ...grpc.CallOption) (*GetAdministrator, error) {
	out := new(GetAdministrator)
	err := c.cc.Invoke(ctx, "/administrator_service_go.AdministratorService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *administratorServiceClient) GetByID(ctx context.Context, in *AdministratorPrimaryKey, opts ...grpc.CallOption) (*GetAdministrator, error) {
	out := new(GetAdministrator)
	err := c.cc.Invoke(ctx, "/administrator_service_go.AdministratorService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *administratorServiceClient) GetList(ctx context.Context, in *GetListAdministratorRequest, opts ...grpc.CallOption) (*GetListAdministratorResponse, error) {
	out := new(GetListAdministratorResponse)
	err := c.cc.Invoke(ctx, "/administrator_service_go.AdministratorService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *administratorServiceClient) Update(ctx context.Context, in *UpdateAdministrator, opts ...grpc.CallOption) (*GetAdministrator, error) {
	out := new(GetAdministrator)
	err := c.cc.Invoke(ctx, "/administrator_service_go.AdministratorService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *administratorServiceClient) Delete(ctx context.Context, in *AdministratorPrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/administrator_service_go.AdministratorService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *administratorServiceClient) Login(ctx context.Context, in *AdministratorLoginRequest, opts ...grpc.CallOption) (*AdministratorLoginResponse, error) {
	out := new(AdministratorLoginResponse)
	err := c.cc.Invoke(ctx, "/administrator_service_go.AdministratorService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *administratorServiceClient) Register(ctx context.Context, in *AdministratorRegisterRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/administrator_service_go.AdministratorService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *administratorServiceClient) RegisterConfirm(ctx context.Context, in *AdministratorRegisterConfRequest, opts ...grpc.CallOption) (*AdministratorLoginResponse, error) {
	out := new(AdministratorLoginResponse)
	err := c.cc.Invoke(ctx, "/administrator_service_go.AdministratorService/RegisterConfirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *administratorServiceClient) ChangePassword(ctx context.Context, in *AdministratorChangePassword, opts ...grpc.CallOption) (*AdministratorChangePasswordResp, error) {
	out := new(AdministratorChangePasswordResp)
	err := c.cc.Invoke(ctx, "/administrator_service_go.AdministratorService/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdministratorServiceServer is the server API for AdministratorService service.
// All implementations should embed UnimplementedAdministratorServiceServer
// for forward compatibility
type AdministratorServiceServer interface {
	Create(context.Context, *CreateAdministrator) (*GetAdministrator, error)
	GetByID(context.Context, *AdministratorPrimaryKey) (*GetAdministrator, error)
	GetList(context.Context, *GetListAdministratorRequest) (*GetListAdministratorResponse, error)
	Update(context.Context, *UpdateAdministrator) (*GetAdministrator, error)
	Delete(context.Context, *AdministratorPrimaryKey) (*empty.Empty, error)
	Login(context.Context, *AdministratorLoginRequest) (*AdministratorLoginResponse, error)
	Register(context.Context, *AdministratorRegisterRequest) (*empty.Empty, error)
	RegisterConfirm(context.Context, *AdministratorRegisterConfRequest) (*AdministratorLoginResponse, error)
	ChangePassword(context.Context, *AdministratorChangePassword) (*AdministratorChangePasswordResp, error)
}

// UnimplementedAdministratorServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAdministratorServiceServer struct {
}

func (UnimplementedAdministratorServiceServer) Create(context.Context, *CreateAdministrator) (*GetAdministrator, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAdministratorServiceServer) GetByID(context.Context, *AdministratorPrimaryKey) (*GetAdministrator, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedAdministratorServiceServer) GetList(context.Context, *GetListAdministratorRequest) (*GetListAdministratorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedAdministratorServiceServer) Update(context.Context, *UpdateAdministrator) (*GetAdministrator, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAdministratorServiceServer) Delete(context.Context, *AdministratorPrimaryKey) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAdministratorServiceServer) Login(context.Context, *AdministratorLoginRequest) (*AdministratorLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAdministratorServiceServer) Register(context.Context, *AdministratorRegisterRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAdministratorServiceServer) RegisterConfirm(context.Context, *AdministratorRegisterConfRequest) (*AdministratorLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterConfirm not implemented")
}
func (UnimplementedAdministratorServiceServer) ChangePassword(context.Context, *AdministratorChangePassword) (*AdministratorChangePasswordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}

// UnsafeAdministratorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdministratorServiceServer will
// result in compilation errors.
type UnsafeAdministratorServiceServer interface {
	mustEmbedUnimplementedAdministratorServiceServer()
}

func RegisterAdministratorServiceServer(s grpc.ServiceRegistrar, srv AdministratorServiceServer) {
	s.RegisterService(&AdministratorService_ServiceDesc, srv)
}

func _AdministratorService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAdministrator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministratorServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/administrator_service_go.AdministratorService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministratorServiceServer).Create(ctx, req.(*CreateAdministrator))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdministratorService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdministratorPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministratorServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/administrator_service_go.AdministratorService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministratorServiceServer).GetByID(ctx, req.(*AdministratorPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdministratorService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListAdministratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministratorServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/administrator_service_go.AdministratorService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministratorServiceServer).GetList(ctx, req.(*GetListAdministratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdministratorService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAdministrator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministratorServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/administrator_service_go.AdministratorService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministratorServiceServer).Update(ctx, req.(*UpdateAdministrator))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdministratorService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdministratorPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministratorServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/administrator_service_go.AdministratorService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministratorServiceServer).Delete(ctx, req.(*AdministratorPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdministratorService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdministratorLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministratorServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/administrator_service_go.AdministratorService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministratorServiceServer).Login(ctx, req.(*AdministratorLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdministratorService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdministratorRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministratorServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/administrator_service_go.AdministratorService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministratorServiceServer).Register(ctx, req.(*AdministratorRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdministratorService_RegisterConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdministratorRegisterConfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministratorServiceServer).RegisterConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/administrator_service_go.AdministratorService/RegisterConfirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministratorServiceServer).RegisterConfirm(ctx, req.(*AdministratorRegisterConfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdministratorService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdministratorChangePassword)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministratorServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/administrator_service_go.AdministratorService/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministratorServiceServer).ChangePassword(ctx, req.(*AdministratorChangePassword))
	}
	return interceptor(ctx, in, info, handler)
}

// AdministratorService_ServiceDesc is the grpc.ServiceDesc for AdministratorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdministratorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "administrator_service_go.AdministratorService",
	HandlerType: (*AdministratorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AdministratorService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _AdministratorService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _AdministratorService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AdministratorService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AdministratorService_Delete_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AdministratorService_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _AdministratorService_Register_Handler,
		},
		{
			MethodName: "RegisterConfirm",
			Handler:    _AdministratorService_RegisterConfirm_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _AdministratorService_ChangePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "administrators.proto",
}
