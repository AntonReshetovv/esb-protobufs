// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package logistics

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DeliveriesManualPrioritiesClient is the client API for DeliveriesManualPriorities service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeliveriesManualPrioritiesClient interface {
	Create(ctx context.Context, in *DeliveryManualPriority, opts ...grpc.CallOption) (*DeliveryManualPriorityId, error)
	Get(ctx context.Context, in *DeliveryManualPriorityId, opts ...grpc.CallOption) (*DeliveryManualPriority, error)
	List(ctx context.Context, in *ListDeliveriesManualPrioritiesRequest, opts ...grpc.CallOption) (*ListDeliveriesManualPrioritiesResponse, error)
	Update(ctx context.Context, in *DeliveryManualPriority, opts ...grpc.CallOption) (*DeliveryManualPriority, error)
	Delete(ctx context.Context, in *DeliveryManualPriorityId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddTransportCompany(ctx context.Context, in *DeliveryManualToTransportCompany, opts ...grpc.CallOption) (*DeliveryManualToTransportCompany, error)
	DeleteTransportCompany(ctx context.Context, in *DeliveryManualToTransportCompany, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type deliveriesManualPrioritiesClient struct {
	cc grpc.ClientConnInterface
}

func NewDeliveriesManualPrioritiesClient(cc grpc.ClientConnInterface) DeliveriesManualPrioritiesClient {
	return &deliveriesManualPrioritiesClient{cc}
}

func (c *deliveriesManualPrioritiesClient) Create(ctx context.Context, in *DeliveryManualPriority, opts ...grpc.CallOption) (*DeliveryManualPriorityId, error) {
	out := new(DeliveryManualPriorityId)
	err := c.cc.Invoke(ctx, "/logistics.DeliveriesManualPriorities/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveriesManualPrioritiesClient) Get(ctx context.Context, in *DeliveryManualPriorityId, opts ...grpc.CallOption) (*DeliveryManualPriority, error) {
	out := new(DeliveryManualPriority)
	err := c.cc.Invoke(ctx, "/logistics.DeliveriesManualPriorities/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveriesManualPrioritiesClient) List(ctx context.Context, in *ListDeliveriesManualPrioritiesRequest, opts ...grpc.CallOption) (*ListDeliveriesManualPrioritiesResponse, error) {
	out := new(ListDeliveriesManualPrioritiesResponse)
	err := c.cc.Invoke(ctx, "/logistics.DeliveriesManualPriorities/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveriesManualPrioritiesClient) Update(ctx context.Context, in *DeliveryManualPriority, opts ...grpc.CallOption) (*DeliveryManualPriority, error) {
	out := new(DeliveryManualPriority)
	err := c.cc.Invoke(ctx, "/logistics.DeliveriesManualPriorities/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveriesManualPrioritiesClient) Delete(ctx context.Context, in *DeliveryManualPriorityId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/logistics.DeliveriesManualPriorities/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveriesManualPrioritiesClient) AddTransportCompany(ctx context.Context, in *DeliveryManualToTransportCompany, opts ...grpc.CallOption) (*DeliveryManualToTransportCompany, error) {
	out := new(DeliveryManualToTransportCompany)
	err := c.cc.Invoke(ctx, "/logistics.DeliveriesManualPriorities/AddTransportCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveriesManualPrioritiesClient) DeleteTransportCompany(ctx context.Context, in *DeliveryManualToTransportCompany, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/logistics.DeliveriesManualPriorities/DeleteTransportCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeliveriesManualPrioritiesServer is the server API for DeliveriesManualPriorities service.
// All implementations must embed UnimplementedDeliveriesManualPrioritiesServer
// for forward compatibility
type DeliveriesManualPrioritiesServer interface {
	Create(context.Context, *DeliveryManualPriority) (*DeliveryManualPriorityId, error)
	Get(context.Context, *DeliveryManualPriorityId) (*DeliveryManualPriority, error)
	List(context.Context, *ListDeliveriesManualPrioritiesRequest) (*ListDeliveriesManualPrioritiesResponse, error)
	Update(context.Context, *DeliveryManualPriority) (*DeliveryManualPriority, error)
	Delete(context.Context, *DeliveryManualPriorityId) (*emptypb.Empty, error)
	AddTransportCompany(context.Context, *DeliveryManualToTransportCompany) (*DeliveryManualToTransportCompany, error)
	DeleteTransportCompany(context.Context, *DeliveryManualToTransportCompany) (*emptypb.Empty, error)
	mustEmbedUnimplementedDeliveriesManualPrioritiesServer()
}

// UnimplementedDeliveriesManualPrioritiesServer must be embedded to have forward compatible implementations.
type UnimplementedDeliveriesManualPrioritiesServer struct {
}

func (UnimplementedDeliveriesManualPrioritiesServer) Create(context.Context, *DeliveryManualPriority) (*DeliveryManualPriorityId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDeliveriesManualPrioritiesServer) Get(context.Context, *DeliveryManualPriorityId) (*DeliveryManualPriority, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDeliveriesManualPrioritiesServer) List(context.Context, *ListDeliveriesManualPrioritiesRequest) (*ListDeliveriesManualPrioritiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDeliveriesManualPrioritiesServer) Update(context.Context, *DeliveryManualPriority) (*DeliveryManualPriority, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDeliveriesManualPrioritiesServer) Delete(context.Context, *DeliveryManualPriorityId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDeliveriesManualPrioritiesServer) AddTransportCompany(context.Context, *DeliveryManualToTransportCompany) (*DeliveryManualToTransportCompany, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTransportCompany not implemented")
}
func (UnimplementedDeliveriesManualPrioritiesServer) DeleteTransportCompany(context.Context, *DeliveryManualToTransportCompany) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTransportCompany not implemented")
}
func (UnimplementedDeliveriesManualPrioritiesServer) mustEmbedUnimplementedDeliveriesManualPrioritiesServer() {
}

// UnsafeDeliveriesManualPrioritiesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeliveriesManualPrioritiesServer will
// result in compilation errors.
type UnsafeDeliveriesManualPrioritiesServer interface {
	mustEmbedUnimplementedDeliveriesManualPrioritiesServer()
}

func RegisterDeliveriesManualPrioritiesServer(s grpc.ServiceRegistrar, srv DeliveriesManualPrioritiesServer) {
	s.RegisterService(&_DeliveriesManualPriorities_serviceDesc, srv)
}

func _DeliveriesManualPriorities_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliveryManualPriority)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveriesManualPrioritiesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.DeliveriesManualPriorities/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveriesManualPrioritiesServer).Create(ctx, req.(*DeliveryManualPriority))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveriesManualPriorities_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliveryManualPriorityId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveriesManualPrioritiesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.DeliveriesManualPriorities/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveriesManualPrioritiesServer).Get(ctx, req.(*DeliveryManualPriorityId))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveriesManualPriorities_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeliveriesManualPrioritiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveriesManualPrioritiesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.DeliveriesManualPriorities/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveriesManualPrioritiesServer).List(ctx, req.(*ListDeliveriesManualPrioritiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveriesManualPriorities_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliveryManualPriority)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveriesManualPrioritiesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.DeliveriesManualPriorities/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveriesManualPrioritiesServer).Update(ctx, req.(*DeliveryManualPriority))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveriesManualPriorities_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliveryManualPriorityId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveriesManualPrioritiesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.DeliveriesManualPriorities/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveriesManualPrioritiesServer).Delete(ctx, req.(*DeliveryManualPriorityId))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveriesManualPriorities_AddTransportCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliveryManualToTransportCompany)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveriesManualPrioritiesServer).AddTransportCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.DeliveriesManualPriorities/AddTransportCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveriesManualPrioritiesServer).AddTransportCompany(ctx, req.(*DeliveryManualToTransportCompany))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveriesManualPriorities_DeleteTransportCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliveryManualToTransportCompany)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveriesManualPrioritiesServer).DeleteTransportCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.DeliveriesManualPriorities/DeleteTransportCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveriesManualPrioritiesServer).DeleteTransportCompany(ctx, req.(*DeliveryManualToTransportCompany))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeliveriesManualPriorities_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logistics.DeliveriesManualPriorities",
	HandlerType: (*DeliveriesManualPrioritiesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DeliveriesManualPriorities_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DeliveriesManualPriorities_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DeliveriesManualPriorities_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DeliveriesManualPriorities_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DeliveriesManualPriorities_Delete_Handler,
		},
		{
			MethodName: "AddTransportCompany",
			Handler:    _DeliveriesManualPriorities_AddTransportCompany_Handler,
		},
		{
			MethodName: "DeleteTransportCompany",
			Handler:    _DeliveriesManualPriorities_DeleteTransportCompany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/logistics/delivery_manual_priority.proto",
}
