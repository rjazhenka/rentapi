// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.11
// source: api/rent.proto

package api

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

// RentServiceClient is the client API for RentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RentServiceClient interface {
	CreateRent(ctx context.Context, in *CreateRentRequest, opts ...grpc.CallOption) (*CreateRentResponse, error)
	MarkAsSent(ctx context.Context, in *MarkAsSentRequest, opts ...grpc.CallOption) (*MarkAsSentResponse, error)
	GetRentToSend(ctx context.Context, in *GetRentToSendRequest, opts ...grpc.CallOption) (*GetRentToSendResponse, error)
	CheckIfExist(ctx context.Context, in *CheckIfExistRequest, opts ...grpc.CallOption) (*CheckIfExistResponse, error)
	ModifySearch(ctx context.Context, in *ModifySearchRequest, opts ...grpc.CallOption) (*ModifySearchResponse, error)
	CreateSearch(ctx context.Context, in *CreateSearchRequest, opts ...grpc.CallOption) (*CreateSearchResponse, error)
	GetSearch(ctx context.Context, in *GetSearchRequest, opts ...grpc.CallOption) (*GetSearchResponse, error)
	GetSearchToSend(ctx context.Context, in *GetSearchToSendRequest, opts ...grpc.CallOption) (*GetSearchToSendResponse, error)
	MarkSearchAsSent(ctx context.Context, in *MarkSearchAsSentRequest, opts ...grpc.CallOption) (*MarkSearchAsSentResponse, error)
	GetTownsByCity(ctx context.Context, in *GetTownsByCityRequest, opts ...grpc.CallOption) (*GetTownsByCityResponse, error)
	GetQuartersByTowns(ctx context.Context, in *GetQuartersByTownsRequest, opts ...grpc.CallOption) (*GetQuartersByTownsResponse, error)
}

type rentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRentServiceClient(cc grpc.ClientConnInterface) RentServiceClient {
	return &rentServiceClient{cc}
}

func (c *rentServiceClient) CreateRent(ctx context.Context, in *CreateRentRequest, opts ...grpc.CallOption) (*CreateRentResponse, error) {
	out := new(CreateRentResponse)
	err := c.cc.Invoke(ctx, "/realty.rent.api.v1.RentService/CreateRent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentServiceClient) MarkAsSent(ctx context.Context, in *MarkAsSentRequest, opts ...grpc.CallOption) (*MarkAsSentResponse, error) {
	out := new(MarkAsSentResponse)
	err := c.cc.Invoke(ctx, "/realty.rent.api.v1.RentService/MarkAsSent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentServiceClient) GetRentToSend(ctx context.Context, in *GetRentToSendRequest, opts ...grpc.CallOption) (*GetRentToSendResponse, error) {
	out := new(GetRentToSendResponse)
	err := c.cc.Invoke(ctx, "/realty.rent.api.v1.RentService/GetRentToSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentServiceClient) CheckIfExist(ctx context.Context, in *CheckIfExistRequest, opts ...grpc.CallOption) (*CheckIfExistResponse, error) {
	out := new(CheckIfExistResponse)
	err := c.cc.Invoke(ctx, "/realty.rent.api.v1.RentService/CheckIfExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentServiceClient) ModifySearch(ctx context.Context, in *ModifySearchRequest, opts ...grpc.CallOption) (*ModifySearchResponse, error) {
	out := new(ModifySearchResponse)
	err := c.cc.Invoke(ctx, "/realty.rent.api.v1.RentService/ModifySearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentServiceClient) CreateSearch(ctx context.Context, in *CreateSearchRequest, opts ...grpc.CallOption) (*CreateSearchResponse, error) {
	out := new(CreateSearchResponse)
	err := c.cc.Invoke(ctx, "/realty.rent.api.v1.RentService/CreateSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentServiceClient) GetSearch(ctx context.Context, in *GetSearchRequest, opts ...grpc.CallOption) (*GetSearchResponse, error) {
	out := new(GetSearchResponse)
	err := c.cc.Invoke(ctx, "/realty.rent.api.v1.RentService/GetSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentServiceClient) GetSearchToSend(ctx context.Context, in *GetSearchToSendRequest, opts ...grpc.CallOption) (*GetSearchToSendResponse, error) {
	out := new(GetSearchToSendResponse)
	err := c.cc.Invoke(ctx, "/realty.rent.api.v1.RentService/GetSearchToSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentServiceClient) MarkSearchAsSent(ctx context.Context, in *MarkSearchAsSentRequest, opts ...grpc.CallOption) (*MarkSearchAsSentResponse, error) {
	out := new(MarkSearchAsSentResponse)
	err := c.cc.Invoke(ctx, "/realty.rent.api.v1.RentService/MarkSearchAsSent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentServiceClient) GetTownsByCity(ctx context.Context, in *GetTownsByCityRequest, opts ...grpc.CallOption) (*GetTownsByCityResponse, error) {
	out := new(GetTownsByCityResponse)
	err := c.cc.Invoke(ctx, "/realty.rent.api.v1.RentService/GetTownsByCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentServiceClient) GetQuartersByTowns(ctx context.Context, in *GetQuartersByTownsRequest, opts ...grpc.CallOption) (*GetQuartersByTownsResponse, error) {
	out := new(GetQuartersByTownsResponse)
	err := c.cc.Invoke(ctx, "/realty.rent.api.v1.RentService/GetQuartersByTowns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RentServiceServer is the server API for RentService service.
// All implementations must embed UnimplementedRentServiceServer
// for forward compatibility
type RentServiceServer interface {
	CreateRent(context.Context, *CreateRentRequest) (*CreateRentResponse, error)
	MarkAsSent(context.Context, *MarkAsSentRequest) (*MarkAsSentResponse, error)
	GetRentToSend(context.Context, *GetRentToSendRequest) (*GetRentToSendResponse, error)
	CheckIfExist(context.Context, *CheckIfExistRequest) (*CheckIfExistResponse, error)
	ModifySearch(context.Context, *ModifySearchRequest) (*ModifySearchResponse, error)
	CreateSearch(context.Context, *CreateSearchRequest) (*CreateSearchResponse, error)
	GetSearch(context.Context, *GetSearchRequest) (*GetSearchResponse, error)
	GetSearchToSend(context.Context, *GetSearchToSendRequest) (*GetSearchToSendResponse, error)
	MarkSearchAsSent(context.Context, *MarkSearchAsSentRequest) (*MarkSearchAsSentResponse, error)
	GetTownsByCity(context.Context, *GetTownsByCityRequest) (*GetTownsByCityResponse, error)
	GetQuartersByTowns(context.Context, *GetQuartersByTownsRequest) (*GetQuartersByTownsResponse, error)
	mustEmbedUnimplementedRentServiceServer()
}

// UnimplementedRentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRentServiceServer struct {
}

func (UnimplementedRentServiceServer) CreateRent(context.Context, *CreateRentRequest) (*CreateRentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRent not implemented")
}
func (UnimplementedRentServiceServer) MarkAsSent(context.Context, *MarkAsSentRequest) (*MarkAsSentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAsSent not implemented")
}
func (UnimplementedRentServiceServer) GetRentToSend(context.Context, *GetRentToSendRequest) (*GetRentToSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRentToSend not implemented")
}
func (UnimplementedRentServiceServer) CheckIfExist(context.Context, *CheckIfExistRequest) (*CheckIfExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIfExist not implemented")
}
func (UnimplementedRentServiceServer) ModifySearch(context.Context, *ModifySearchRequest) (*ModifySearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySearch not implemented")
}
func (UnimplementedRentServiceServer) CreateSearch(context.Context, *CreateSearchRequest) (*CreateSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSearch not implemented")
}
func (UnimplementedRentServiceServer) GetSearch(context.Context, *GetSearchRequest) (*GetSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSearch not implemented")
}
func (UnimplementedRentServiceServer) GetSearchToSend(context.Context, *GetSearchToSendRequest) (*GetSearchToSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSearchToSend not implemented")
}
func (UnimplementedRentServiceServer) MarkSearchAsSent(context.Context, *MarkSearchAsSentRequest) (*MarkSearchAsSentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkSearchAsSent not implemented")
}
func (UnimplementedRentServiceServer) GetTownsByCity(context.Context, *GetTownsByCityRequest) (*GetTownsByCityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTownsByCity not implemented")
}
func (UnimplementedRentServiceServer) GetQuartersByTowns(context.Context, *GetQuartersByTownsRequest) (*GetQuartersByTownsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuartersByTowns not implemented")
}
func (UnimplementedRentServiceServer) mustEmbedUnimplementedRentServiceServer() {}

// UnsafeRentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RentServiceServer will
// result in compilation errors.
type UnsafeRentServiceServer interface {
	mustEmbedUnimplementedRentServiceServer()
}

func RegisterRentServiceServer(s grpc.ServiceRegistrar, srv RentServiceServer) {
	s.RegisterService(&RentService_ServiceDesc, srv)
}

func _RentService_CreateRent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentServiceServer).CreateRent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realty.rent.api.v1.RentService/CreateRent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentServiceServer).CreateRent(ctx, req.(*CreateRentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RentService_MarkAsSent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkAsSentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentServiceServer).MarkAsSent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realty.rent.api.v1.RentService/MarkAsSent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentServiceServer).MarkAsSent(ctx, req.(*MarkAsSentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RentService_GetRentToSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRentToSendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentServiceServer).GetRentToSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realty.rent.api.v1.RentService/GetRentToSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentServiceServer).GetRentToSend(ctx, req.(*GetRentToSendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RentService_CheckIfExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckIfExistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentServiceServer).CheckIfExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realty.rent.api.v1.RentService/CheckIfExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentServiceServer).CheckIfExist(ctx, req.(*CheckIfExistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RentService_ModifySearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentServiceServer).ModifySearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realty.rent.api.v1.RentService/ModifySearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentServiceServer).ModifySearch(ctx, req.(*ModifySearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RentService_CreateSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentServiceServer).CreateSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realty.rent.api.v1.RentService/CreateSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentServiceServer).CreateSearch(ctx, req.(*CreateSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RentService_GetSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentServiceServer).GetSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realty.rent.api.v1.RentService/GetSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentServiceServer).GetSearch(ctx, req.(*GetSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RentService_GetSearchToSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSearchToSendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentServiceServer).GetSearchToSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realty.rent.api.v1.RentService/GetSearchToSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentServiceServer).GetSearchToSend(ctx, req.(*GetSearchToSendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RentService_MarkSearchAsSent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkSearchAsSentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentServiceServer).MarkSearchAsSent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realty.rent.api.v1.RentService/MarkSearchAsSent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentServiceServer).MarkSearchAsSent(ctx, req.(*MarkSearchAsSentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RentService_GetTownsByCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTownsByCityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentServiceServer).GetTownsByCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realty.rent.api.v1.RentService/GetTownsByCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentServiceServer).GetTownsByCity(ctx, req.(*GetTownsByCityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RentService_GetQuartersByTowns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuartersByTownsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentServiceServer).GetQuartersByTowns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realty.rent.api.v1.RentService/GetQuartersByTowns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentServiceServer).GetQuartersByTowns(ctx, req.(*GetQuartersByTownsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RentService_ServiceDesc is the grpc.ServiceDesc for RentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "realty.rent.api.v1.RentService",
	HandlerType: (*RentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRent",
			Handler:    _RentService_CreateRent_Handler,
		},
		{
			MethodName: "MarkAsSent",
			Handler:    _RentService_MarkAsSent_Handler,
		},
		{
			MethodName: "GetRentToSend",
			Handler:    _RentService_GetRentToSend_Handler,
		},
		{
			MethodName: "CheckIfExist",
			Handler:    _RentService_CheckIfExist_Handler,
		},
		{
			MethodName: "ModifySearch",
			Handler:    _RentService_ModifySearch_Handler,
		},
		{
			MethodName: "CreateSearch",
			Handler:    _RentService_CreateSearch_Handler,
		},
		{
			MethodName: "GetSearch",
			Handler:    _RentService_GetSearch_Handler,
		},
		{
			MethodName: "GetSearchToSend",
			Handler:    _RentService_GetSearchToSend_Handler,
		},
		{
			MethodName: "MarkSearchAsSent",
			Handler:    _RentService_MarkSearchAsSent_Handler,
		},
		{
			MethodName: "GetTownsByCity",
			Handler:    _RentService_GetTownsByCity_Handler,
		},
		{
			MethodName: "GetQuartersByTowns",
			Handler:    _RentService_GetQuartersByTowns_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/rent.proto",
}
