// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: dagger.proto

package daggerProto

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
	Remove_RemKey_FullMethodName = "/dagger.Remove/RemKey"
)

// RemoveClient is the client API for Remove service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemoveClient interface {
	RemKey(ctx context.Context, in *DelKey, opts ...grpc.CallOption) (*ResponseCode, error)
}

type removeClient struct {
	cc grpc.ClientConnInterface
}

func NewRemoveClient(cc grpc.ClientConnInterface) RemoveClient {
	return &removeClient{cc}
}

func (c *removeClient) RemKey(ctx context.Context, in *DelKey, opts ...grpc.CallOption) (*ResponseCode, error) {
	out := new(ResponseCode)
	err := c.cc.Invoke(ctx, Remove_RemKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoveServer is the server API for Remove service.
// All implementations must embed UnimplementedRemoveServer
// for forward compatibility
type RemoveServer interface {
	RemKey(context.Context, *DelKey) (*ResponseCode, error)
	mustEmbedUnimplementedRemoveServer()
}

// UnimplementedRemoveServer must be embedded to have forward compatible implementations.
type UnimplementedRemoveServer struct {
}

func (UnimplementedRemoveServer) RemKey(context.Context, *DelKey) (*ResponseCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemKey not implemented")
}
func (UnimplementedRemoveServer) mustEmbedUnimplementedRemoveServer() {}

// UnsafeRemoveServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemoveServer will
// result in compilation errors.
type UnsafeRemoveServer interface {
	mustEmbedUnimplementedRemoveServer()
}

func RegisterRemoveServer(s grpc.ServiceRegistrar, srv RemoveServer) {
	s.RegisterService(&Remove_ServiceDesc, srv)
}

func _Remove_RemKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoveServer).RemKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Remove_RemKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoveServer).RemKey(ctx, req.(*DelKey))
	}
	return interceptor(ctx, in, info, handler)
}

// Remove_ServiceDesc is the grpc.ServiceDesc for Remove service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Remove_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dagger.Remove",
	HandlerType: (*RemoveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RemKey",
			Handler:    _Remove_RemKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dagger.proto",
}

const (
	GetAll_GetAll_FullMethodName = "/dagger.GetAll/GetAll"
)

// GetAllClient is the client API for GetAll service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetAllClient interface {
	GetAll(ctx context.Context, in *GetKey, opts ...grpc.CallOption) (*DbContents, error)
}

type getAllClient struct {
	cc grpc.ClientConnInterface
}

func NewGetAllClient(cc grpc.ClientConnInterface) GetAllClient {
	return &getAllClient{cc}
}

func (c *getAllClient) GetAll(ctx context.Context, in *GetKey, opts ...grpc.CallOption) (*DbContents, error) {
	out := new(DbContents)
	err := c.cc.Invoke(ctx, GetAll_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetAllServer is the server API for GetAll service.
// All implementations must embed UnimplementedGetAllServer
// for forward compatibility
type GetAllServer interface {
	GetAll(context.Context, *GetKey) (*DbContents, error)
	mustEmbedUnimplementedGetAllServer()
}

// UnimplementedGetAllServer must be embedded to have forward compatible implementations.
type UnimplementedGetAllServer struct {
}

func (UnimplementedGetAllServer) GetAll(context.Context, *GetKey) (*DbContents, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedGetAllServer) mustEmbedUnimplementedGetAllServer() {}

// UnsafeGetAllServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetAllServer will
// result in compilation errors.
type UnsafeGetAllServer interface {
	mustEmbedUnimplementedGetAllServer()
}

func RegisterGetAllServer(s grpc.ServiceRegistrar, srv GetAllServer) {
	s.RegisterService(&GetAll_ServiceDesc, srv)
}

func _GetAll_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetAllServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GetAll_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetAllServer).GetAll(ctx, req.(*GetKey))
	}
	return interceptor(ctx, in, info, handler)
}

// GetAll_ServiceDesc is the grpc.ServiceDesc for GetAll service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetAll_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dagger.GetAll",
	HandlerType: (*GetAllServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _GetAll_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dagger.proto",
}

const (
	HgetRecord_Hget_FullMethodName = "/dagger.hgetRecord/Hget"
)

// HgetRecordClient is the client API for HgetRecord service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HgetRecordClient interface {
	Hget(ctx context.Context, in *GetUUID, opts ...grpc.CallOption) (*UpdateObject, error)
}

type hgetRecordClient struct {
	cc grpc.ClientConnInterface
}

func NewHgetRecordClient(cc grpc.ClientConnInterface) HgetRecordClient {
	return &hgetRecordClient{cc}
}

func (c *hgetRecordClient) Hget(ctx context.Context, in *GetUUID, opts ...grpc.CallOption) (*UpdateObject, error) {
	out := new(UpdateObject)
	err := c.cc.Invoke(ctx, HgetRecord_Hget_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HgetRecordServer is the server API for HgetRecord service.
// All implementations must embed UnimplementedHgetRecordServer
// for forward compatibility
type HgetRecordServer interface {
	Hget(context.Context, *GetUUID) (*UpdateObject, error)
	mustEmbedUnimplementedHgetRecordServer()
}

// UnimplementedHgetRecordServer must be embedded to have forward compatible implementations.
type UnimplementedHgetRecordServer struct {
}

func (UnimplementedHgetRecordServer) Hget(context.Context, *GetUUID) (*UpdateObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hget not implemented")
}
func (UnimplementedHgetRecordServer) mustEmbedUnimplementedHgetRecordServer() {}

// UnsafeHgetRecordServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HgetRecordServer will
// result in compilation errors.
type UnsafeHgetRecordServer interface {
	mustEmbedUnimplementedHgetRecordServer()
}

func RegisterHgetRecordServer(s grpc.ServiceRegistrar, srv HgetRecordServer) {
	s.RegisterService(&HgetRecord_ServiceDesc, srv)
}

func _HgetRecord_Hget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HgetRecordServer).Hget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HgetRecord_Hget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HgetRecordServer).Hget(ctx, req.(*GetUUID))
	}
	return interceptor(ctx, in, info, handler)
}

// HgetRecord_ServiceDesc is the grpc.ServiceDesc for HgetRecord service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HgetRecord_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dagger.hgetRecord",
	HandlerType: (*HgetRecordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hget",
			Handler:    _HgetRecord_Hget_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dagger.proto",
}

const (
	GetRecord_Get_FullMethodName = "/dagger.getRecord/get"
)

// GetRecordClient is the client API for GetRecord service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetRecordClient interface {
	Get(ctx context.Context, in *GetUUID, opts ...grpc.CallOption) (*UpdateObject, error)
}

type getRecordClient struct {
	cc grpc.ClientConnInterface
}

func NewGetRecordClient(cc grpc.ClientConnInterface) GetRecordClient {
	return &getRecordClient{cc}
}

func (c *getRecordClient) Get(ctx context.Context, in *GetUUID, opts ...grpc.CallOption) (*UpdateObject, error) {
	out := new(UpdateObject)
	err := c.cc.Invoke(ctx, GetRecord_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetRecordServer is the server API for GetRecord service.
// All implementations must embed UnimplementedGetRecordServer
// for forward compatibility
type GetRecordServer interface {
	Get(context.Context, *GetUUID) (*UpdateObject, error)
	mustEmbedUnimplementedGetRecordServer()
}

// UnimplementedGetRecordServer must be embedded to have forward compatible implementations.
type UnimplementedGetRecordServer struct {
}

func (UnimplementedGetRecordServer) Get(context.Context, *GetUUID) (*UpdateObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedGetRecordServer) mustEmbedUnimplementedGetRecordServer() {}

// UnsafeGetRecordServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetRecordServer will
// result in compilation errors.
type UnsafeGetRecordServer interface {
	mustEmbedUnimplementedGetRecordServer()
}

func RegisterGetRecordServer(s grpc.ServiceRegistrar, srv GetRecordServer) {
	s.RegisterService(&GetRecord_ServiceDesc, srv)
}

func _GetRecord_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetRecordServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GetRecord_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetRecordServer).Get(ctx, req.(*GetUUID))
	}
	return interceptor(ctx, in, info, handler)
}

// GetRecord_ServiceDesc is the grpc.ServiceDesc for GetRecord service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetRecord_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dagger.getRecord",
	HandlerType: (*GetRecordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _GetRecord_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dagger.proto",
}

const (
	Builder_StartBuilding_FullMethodName = "/dagger.Builder/StartBuilding"
)

// BuilderClient is the client API for Builder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BuilderClient interface {
	StartBuilding(ctx context.Context, in *BuildRoutine, opts ...grpc.CallOption) (*ResponseCode, error)
}

type builderClient struct {
	cc grpc.ClientConnInterface
}

func NewBuilderClient(cc grpc.ClientConnInterface) BuilderClient {
	return &builderClient{cc}
}

func (c *builderClient) StartBuilding(ctx context.Context, in *BuildRoutine, opts ...grpc.CallOption) (*ResponseCode, error) {
	out := new(ResponseCode)
	err := c.cc.Invoke(ctx, Builder_StartBuilding_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuilderServer is the server API for Builder service.
// All implementations must embed UnimplementedBuilderServer
// for forward compatibility
type BuilderServer interface {
	StartBuilding(context.Context, *BuildRoutine) (*ResponseCode, error)
	mustEmbedUnimplementedBuilderServer()
}

// UnimplementedBuilderServer must be embedded to have forward compatible implementations.
type UnimplementedBuilderServer struct {
}

func (UnimplementedBuilderServer) StartBuilding(context.Context, *BuildRoutine) (*ResponseCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartBuilding not implemented")
}
func (UnimplementedBuilderServer) mustEmbedUnimplementedBuilderServer() {}

// UnsafeBuilderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BuilderServer will
// result in compilation errors.
type UnsafeBuilderServer interface {
	mustEmbedUnimplementedBuilderServer()
}

func RegisterBuilderServer(s grpc.ServiceRegistrar, srv BuilderServer) {
	s.RegisterService(&Builder_ServiceDesc, srv)
}

func _Builder_StartBuilding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildRoutine)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServer).StartBuilding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Builder_StartBuilding_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServer).StartBuilding(ctx, req.(*BuildRoutine))
	}
	return interceptor(ctx, in, info, handler)
}

// Builder_ServiceDesc is the grpc.ServiceDesc for Builder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Builder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dagger.Builder",
	HandlerType: (*BuilderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartBuilding",
			Handler:    _Builder_StartBuilding_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dagger.proto",
}

const (
	UpdateRecord_SendUpdate_FullMethodName = "/dagger.UpdateRecord/SendUpdate"
)

// UpdateRecordClient is the client API for UpdateRecord service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UpdateRecordClient interface {
	SendUpdate(ctx context.Context, in *UpdateObject, opts ...grpc.CallOption) (*ResponseCode, error)
}

type updateRecordClient struct {
	cc grpc.ClientConnInterface
}

func NewUpdateRecordClient(cc grpc.ClientConnInterface) UpdateRecordClient {
	return &updateRecordClient{cc}
}

func (c *updateRecordClient) SendUpdate(ctx context.Context, in *UpdateObject, opts ...grpc.CallOption) (*ResponseCode, error) {
	out := new(ResponseCode)
	err := c.cc.Invoke(ctx, UpdateRecord_SendUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateRecordServer is the server API for UpdateRecord service.
// All implementations must embed UnimplementedUpdateRecordServer
// for forward compatibility
type UpdateRecordServer interface {
	SendUpdate(context.Context, *UpdateObject) (*ResponseCode, error)
	mustEmbedUnimplementedUpdateRecordServer()
}

// UnimplementedUpdateRecordServer must be embedded to have forward compatible implementations.
type UnimplementedUpdateRecordServer struct {
}

func (UnimplementedUpdateRecordServer) SendUpdate(context.Context, *UpdateObject) (*ResponseCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUpdate not implemented")
}
func (UnimplementedUpdateRecordServer) mustEmbedUnimplementedUpdateRecordServer() {}

// UnsafeUpdateRecordServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UpdateRecordServer will
// result in compilation errors.
type UnsafeUpdateRecordServer interface {
	mustEmbedUnimplementedUpdateRecordServer()
}

func RegisterUpdateRecordServer(s grpc.ServiceRegistrar, srv UpdateRecordServer) {
	s.RegisterService(&UpdateRecord_ServiceDesc, srv)
}

func _UpdateRecord_SendUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateRecordServer).SendUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UpdateRecord_SendUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateRecordServer).SendUpdate(ctx, req.(*UpdateObject))
	}
	return interceptor(ctx, in, info, handler)
}

// UpdateRecord_ServiceDesc is the grpc.ServiceDesc for UpdateRecord service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UpdateRecord_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dagger.UpdateRecord",
	HandlerType: (*UpdateRecordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendUpdate",
			Handler:    _UpdateRecord_SendUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dagger.proto",
}
