// Copyright 2021-2023 the Kubeapps contributors.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: kubeappsapis/core/packages/v1alpha1/packages.proto

package v1alpha1

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
	PackagesService_GetAvailablePackageSummaries_FullMethodName    = "/kubeappsapis.core.packages.v1alpha1.PackagesService/GetAvailablePackageSummaries"
	PackagesService_GetAvailablePackageDetail_FullMethodName       = "/kubeappsapis.core.packages.v1alpha1.PackagesService/GetAvailablePackageDetail"
	PackagesService_GetAvailablePackageVersions_FullMethodName     = "/kubeappsapis.core.packages.v1alpha1.PackagesService/GetAvailablePackageVersions"
	PackagesService_GetInstalledPackageSummaries_FullMethodName    = "/kubeappsapis.core.packages.v1alpha1.PackagesService/GetInstalledPackageSummaries"
	PackagesService_GetInstalledPackageDetail_FullMethodName       = "/kubeappsapis.core.packages.v1alpha1.PackagesService/GetInstalledPackageDetail"
	PackagesService_CreateInstalledPackage_FullMethodName          = "/kubeappsapis.core.packages.v1alpha1.PackagesService/CreateInstalledPackage"
	PackagesService_UpdateInstalledPackage_FullMethodName          = "/kubeappsapis.core.packages.v1alpha1.PackagesService/UpdateInstalledPackage"
	PackagesService_DeleteInstalledPackage_FullMethodName          = "/kubeappsapis.core.packages.v1alpha1.PackagesService/DeleteInstalledPackage"
	PackagesService_GetInstalledPackageResourceRefs_FullMethodName = "/kubeappsapis.core.packages.v1alpha1.PackagesService/GetInstalledPackageResourceRefs"
)

// PackagesServiceClient is the client API for PackagesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PackagesServiceClient interface {
	GetAvailablePackageSummaries(ctx context.Context, in *GetAvailablePackageSummariesRequest, opts ...grpc.CallOption) (*GetAvailablePackageSummariesResponse, error)
	GetAvailablePackageDetail(ctx context.Context, in *GetAvailablePackageDetailRequest, opts ...grpc.CallOption) (*GetAvailablePackageDetailResponse, error)
	GetAvailablePackageVersions(ctx context.Context, in *GetAvailablePackageVersionsRequest, opts ...grpc.CallOption) (*GetAvailablePackageVersionsResponse, error)
	GetInstalledPackageSummaries(ctx context.Context, in *GetInstalledPackageSummariesRequest, opts ...grpc.CallOption) (*GetInstalledPackageSummariesResponse, error)
	GetInstalledPackageDetail(ctx context.Context, in *GetInstalledPackageDetailRequest, opts ...grpc.CallOption) (*GetInstalledPackageDetailResponse, error)
	CreateInstalledPackage(ctx context.Context, in *CreateInstalledPackageRequest, opts ...grpc.CallOption) (*CreateInstalledPackageResponse, error)
	UpdateInstalledPackage(ctx context.Context, in *UpdateInstalledPackageRequest, opts ...grpc.CallOption) (*UpdateInstalledPackageResponse, error)
	DeleteInstalledPackage(ctx context.Context, in *DeleteInstalledPackageRequest, opts ...grpc.CallOption) (*DeleteInstalledPackageResponse, error)
	GetInstalledPackageResourceRefs(ctx context.Context, in *GetInstalledPackageResourceRefsRequest, opts ...grpc.CallOption) (*GetInstalledPackageResourceRefsResponse, error)
}

type packagesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPackagesServiceClient(cc grpc.ClientConnInterface) PackagesServiceClient {
	return &packagesServiceClient{cc}
}

func (c *packagesServiceClient) GetAvailablePackageSummaries(ctx context.Context, in *GetAvailablePackageSummariesRequest, opts ...grpc.CallOption) (*GetAvailablePackageSummariesResponse, error) {
	out := new(GetAvailablePackageSummariesResponse)
	err := c.cc.Invoke(ctx, PackagesService_GetAvailablePackageSummaries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packagesServiceClient) GetAvailablePackageDetail(ctx context.Context, in *GetAvailablePackageDetailRequest, opts ...grpc.CallOption) (*GetAvailablePackageDetailResponse, error) {
	out := new(GetAvailablePackageDetailResponse)
	err := c.cc.Invoke(ctx, PackagesService_GetAvailablePackageDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packagesServiceClient) GetAvailablePackageVersions(ctx context.Context, in *GetAvailablePackageVersionsRequest, opts ...grpc.CallOption) (*GetAvailablePackageVersionsResponse, error) {
	out := new(GetAvailablePackageVersionsResponse)
	err := c.cc.Invoke(ctx, PackagesService_GetAvailablePackageVersions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packagesServiceClient) GetInstalledPackageSummaries(ctx context.Context, in *GetInstalledPackageSummariesRequest, opts ...grpc.CallOption) (*GetInstalledPackageSummariesResponse, error) {
	out := new(GetInstalledPackageSummariesResponse)
	err := c.cc.Invoke(ctx, PackagesService_GetInstalledPackageSummaries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packagesServiceClient) GetInstalledPackageDetail(ctx context.Context, in *GetInstalledPackageDetailRequest, opts ...grpc.CallOption) (*GetInstalledPackageDetailResponse, error) {
	out := new(GetInstalledPackageDetailResponse)
	err := c.cc.Invoke(ctx, PackagesService_GetInstalledPackageDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packagesServiceClient) CreateInstalledPackage(ctx context.Context, in *CreateInstalledPackageRequest, opts ...grpc.CallOption) (*CreateInstalledPackageResponse, error) {
	out := new(CreateInstalledPackageResponse)
	err := c.cc.Invoke(ctx, PackagesService_CreateInstalledPackage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packagesServiceClient) UpdateInstalledPackage(ctx context.Context, in *UpdateInstalledPackageRequest, opts ...grpc.CallOption) (*UpdateInstalledPackageResponse, error) {
	out := new(UpdateInstalledPackageResponse)
	err := c.cc.Invoke(ctx, PackagesService_UpdateInstalledPackage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packagesServiceClient) DeleteInstalledPackage(ctx context.Context, in *DeleteInstalledPackageRequest, opts ...grpc.CallOption) (*DeleteInstalledPackageResponse, error) {
	out := new(DeleteInstalledPackageResponse)
	err := c.cc.Invoke(ctx, PackagesService_DeleteInstalledPackage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packagesServiceClient) GetInstalledPackageResourceRefs(ctx context.Context, in *GetInstalledPackageResourceRefsRequest, opts ...grpc.CallOption) (*GetInstalledPackageResourceRefsResponse, error) {
	out := new(GetInstalledPackageResourceRefsResponse)
	err := c.cc.Invoke(ctx, PackagesService_GetInstalledPackageResourceRefs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PackagesServiceServer is the server API for PackagesService service.
// All implementations should embed UnimplementedPackagesServiceServer
// for forward compatibility
type PackagesServiceServer interface {
	GetAvailablePackageSummaries(context.Context, *GetAvailablePackageSummariesRequest) (*GetAvailablePackageSummariesResponse, error)
	GetAvailablePackageDetail(context.Context, *GetAvailablePackageDetailRequest) (*GetAvailablePackageDetailResponse, error)
	GetAvailablePackageVersions(context.Context, *GetAvailablePackageVersionsRequest) (*GetAvailablePackageVersionsResponse, error)
	GetInstalledPackageSummaries(context.Context, *GetInstalledPackageSummariesRequest) (*GetInstalledPackageSummariesResponse, error)
	GetInstalledPackageDetail(context.Context, *GetInstalledPackageDetailRequest) (*GetInstalledPackageDetailResponse, error)
	CreateInstalledPackage(context.Context, *CreateInstalledPackageRequest) (*CreateInstalledPackageResponse, error)
	UpdateInstalledPackage(context.Context, *UpdateInstalledPackageRequest) (*UpdateInstalledPackageResponse, error)
	DeleteInstalledPackage(context.Context, *DeleteInstalledPackageRequest) (*DeleteInstalledPackageResponse, error)
	GetInstalledPackageResourceRefs(context.Context, *GetInstalledPackageResourceRefsRequest) (*GetInstalledPackageResourceRefsResponse, error)
}

// UnimplementedPackagesServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPackagesServiceServer struct {
}

func (UnimplementedPackagesServiceServer) GetAvailablePackageSummaries(context.Context, *GetAvailablePackageSummariesRequest) (*GetAvailablePackageSummariesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailablePackageSummaries not implemented")
}
func (UnimplementedPackagesServiceServer) GetAvailablePackageDetail(context.Context, *GetAvailablePackageDetailRequest) (*GetAvailablePackageDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailablePackageDetail not implemented")
}
func (UnimplementedPackagesServiceServer) GetAvailablePackageVersions(context.Context, *GetAvailablePackageVersionsRequest) (*GetAvailablePackageVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailablePackageVersions not implemented")
}
func (UnimplementedPackagesServiceServer) GetInstalledPackageSummaries(context.Context, *GetInstalledPackageSummariesRequest) (*GetInstalledPackageSummariesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstalledPackageSummaries not implemented")
}
func (UnimplementedPackagesServiceServer) GetInstalledPackageDetail(context.Context, *GetInstalledPackageDetailRequest) (*GetInstalledPackageDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstalledPackageDetail not implemented")
}
func (UnimplementedPackagesServiceServer) CreateInstalledPackage(context.Context, *CreateInstalledPackageRequest) (*CreateInstalledPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInstalledPackage not implemented")
}
func (UnimplementedPackagesServiceServer) UpdateInstalledPackage(context.Context, *UpdateInstalledPackageRequest) (*UpdateInstalledPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInstalledPackage not implemented")
}
func (UnimplementedPackagesServiceServer) DeleteInstalledPackage(context.Context, *DeleteInstalledPackageRequest) (*DeleteInstalledPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInstalledPackage not implemented")
}
func (UnimplementedPackagesServiceServer) GetInstalledPackageResourceRefs(context.Context, *GetInstalledPackageResourceRefsRequest) (*GetInstalledPackageResourceRefsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstalledPackageResourceRefs not implemented")
}

// UnsafePackagesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PackagesServiceServer will
// result in compilation errors.
type UnsafePackagesServiceServer interface {
	mustEmbedUnimplementedPackagesServiceServer()
}

func RegisterPackagesServiceServer(s grpc.ServiceRegistrar, srv PackagesServiceServer) {
	s.RegisterService(&PackagesService_ServiceDesc, srv)
}

func _PackagesService_GetAvailablePackageSummaries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvailablePackageSummariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackagesServiceServer).GetAvailablePackageSummaries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PackagesService_GetAvailablePackageSummaries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackagesServiceServer).GetAvailablePackageSummaries(ctx, req.(*GetAvailablePackageSummariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackagesService_GetAvailablePackageDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvailablePackageDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackagesServiceServer).GetAvailablePackageDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PackagesService_GetAvailablePackageDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackagesServiceServer).GetAvailablePackageDetail(ctx, req.(*GetAvailablePackageDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackagesService_GetAvailablePackageVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvailablePackageVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackagesServiceServer).GetAvailablePackageVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PackagesService_GetAvailablePackageVersions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackagesServiceServer).GetAvailablePackageVersions(ctx, req.(*GetAvailablePackageVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackagesService_GetInstalledPackageSummaries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstalledPackageSummariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackagesServiceServer).GetInstalledPackageSummaries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PackagesService_GetInstalledPackageSummaries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackagesServiceServer).GetInstalledPackageSummaries(ctx, req.(*GetInstalledPackageSummariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackagesService_GetInstalledPackageDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstalledPackageDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackagesServiceServer).GetInstalledPackageDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PackagesService_GetInstalledPackageDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackagesServiceServer).GetInstalledPackageDetail(ctx, req.(*GetInstalledPackageDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackagesService_CreateInstalledPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInstalledPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackagesServiceServer).CreateInstalledPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PackagesService_CreateInstalledPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackagesServiceServer).CreateInstalledPackage(ctx, req.(*CreateInstalledPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackagesService_UpdateInstalledPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInstalledPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackagesServiceServer).UpdateInstalledPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PackagesService_UpdateInstalledPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackagesServiceServer).UpdateInstalledPackage(ctx, req.(*UpdateInstalledPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackagesService_DeleteInstalledPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInstalledPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackagesServiceServer).DeleteInstalledPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PackagesService_DeleteInstalledPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackagesServiceServer).DeleteInstalledPackage(ctx, req.(*DeleteInstalledPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackagesService_GetInstalledPackageResourceRefs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstalledPackageResourceRefsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackagesServiceServer).GetInstalledPackageResourceRefs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PackagesService_GetInstalledPackageResourceRefs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackagesServiceServer).GetInstalledPackageResourceRefs(ctx, req.(*GetInstalledPackageResourceRefsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PackagesService_ServiceDesc is the grpc.ServiceDesc for PackagesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PackagesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kubeappsapis.core.packages.v1alpha1.PackagesService",
	HandlerType: (*PackagesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAvailablePackageSummaries",
			Handler:    _PackagesService_GetAvailablePackageSummaries_Handler,
		},
		{
			MethodName: "GetAvailablePackageDetail",
			Handler:    _PackagesService_GetAvailablePackageDetail_Handler,
		},
		{
			MethodName: "GetAvailablePackageVersions",
			Handler:    _PackagesService_GetAvailablePackageVersions_Handler,
		},
		{
			MethodName: "GetInstalledPackageSummaries",
			Handler:    _PackagesService_GetInstalledPackageSummaries_Handler,
		},
		{
			MethodName: "GetInstalledPackageDetail",
			Handler:    _PackagesService_GetInstalledPackageDetail_Handler,
		},
		{
			MethodName: "CreateInstalledPackage",
			Handler:    _PackagesService_CreateInstalledPackage_Handler,
		},
		{
			MethodName: "UpdateInstalledPackage",
			Handler:    _PackagesService_UpdateInstalledPackage_Handler,
		},
		{
			MethodName: "DeleteInstalledPackage",
			Handler:    _PackagesService_DeleteInstalledPackage_Handler,
		},
		{
			MethodName: "GetInstalledPackageResourceRefs",
			Handler:    _PackagesService_GetInstalledPackageResourceRefs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kubeappsapis/core/packages/v1alpha1/packages.proto",
}
