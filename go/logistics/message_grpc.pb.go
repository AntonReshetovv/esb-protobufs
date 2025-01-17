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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MessagesClient is the client API for Messages service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagesClient interface {
	Create(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessageId, error)
	Get(ctx context.Context, in *MessageId, opts ...grpc.CallOption) (*Message, error)
	List(ctx context.Context, in *ListMessagesRequest, opts ...grpc.CallOption) (*ListMessagesResponse, error)
	Update(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Upsert(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Delete(ctx context.Context, in *MessageId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddTransportCompany(ctx context.Context, in *MessageToTransportCompany, opts ...grpc.CallOption) (*MessageToTransportCompany, error)
	DeleteTransportCompany(ctx context.Context, in *MessageToTransportCompany, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type messagesClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagesClient(cc grpc.ClientConnInterface) MessagesClient {
	return &messagesClient{cc}
}

func (c *messagesClient) Create(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessageId, error) {
	out := new(MessageId)
	err := c.cc.Invoke(ctx, "/logistics.Messages/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagesClient) Get(ctx context.Context, in *MessageId, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/logistics.Messages/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagesClient) List(ctx context.Context, in *ListMessagesRequest, opts ...grpc.CallOption) (*ListMessagesResponse, error) {
	out := new(ListMessagesResponse)
	err := c.cc.Invoke(ctx, "/logistics.Messages/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagesClient) Update(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/logistics.Messages/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagesClient) Upsert(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/logistics.Messages/Upsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagesClient) Delete(ctx context.Context, in *MessageId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/logistics.Messages/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagesClient) AddTransportCompany(ctx context.Context, in *MessageToTransportCompany, opts ...grpc.CallOption) (*MessageToTransportCompany, error) {
	out := new(MessageToTransportCompany)
	err := c.cc.Invoke(ctx, "/logistics.Messages/AddTransportCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagesClient) DeleteTransportCompany(ctx context.Context, in *MessageToTransportCompany, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/logistics.Messages/DeleteTransportCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessagesServer is the server API for Messages service.
// All implementations should embed UnimplementedMessagesServer
// for forward compatibility
type MessagesServer interface {
	Create(context.Context, *Message) (*MessageId, error)
	Get(context.Context, *MessageId) (*Message, error)
	List(context.Context, *ListMessagesRequest) (*ListMessagesResponse, error)
	Update(context.Context, *Message) (*Message, error)
	Upsert(context.Context, *Message) (*Message, error)
	Delete(context.Context, *MessageId) (*emptypb.Empty, error)
	AddTransportCompany(context.Context, *MessageToTransportCompany) (*MessageToTransportCompany, error)
	DeleteTransportCompany(context.Context, *MessageToTransportCompany) (*emptypb.Empty, error)
}

// UnimplementedMessagesServer should be embedded to have forward compatible implementations.
type UnimplementedMessagesServer struct {
}

func (UnimplementedMessagesServer) Create(context.Context, *Message) (*MessageId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMessagesServer) Get(context.Context, *MessageId) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedMessagesServer) List(context.Context, *ListMessagesRequest) (*ListMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedMessagesServer) Update(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMessagesServer) Upsert(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upsert not implemented")
}
func (UnimplementedMessagesServer) Delete(context.Context, *MessageId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedMessagesServer) AddTransportCompany(context.Context, *MessageToTransportCompany) (*MessageToTransportCompany, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTransportCompany not implemented")
}
func (UnimplementedMessagesServer) DeleteTransportCompany(context.Context, *MessageToTransportCompany) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTransportCompany not implemented")
}

// UnsafeMessagesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagesServer will
// result in compilation errors.
type UnsafeMessagesServer interface {
	mustEmbedUnimplementedMessagesServer()
}

func RegisterMessagesServer(s grpc.ServiceRegistrar, srv MessagesServer) {
	s.RegisterService(&Messages_ServiceDesc, srv)
}

func _Messages_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.Messages/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServer).Create(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messages_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.Messages/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServer).Get(ctx, req.(*MessageId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messages_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.Messages/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServer).List(ctx, req.(*ListMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messages_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.Messages/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServer).Update(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messages_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServer).Upsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.Messages/Upsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServer).Upsert(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messages_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.Messages/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServer).Delete(ctx, req.(*MessageId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messages_AddTransportCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageToTransportCompany)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServer).AddTransportCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.Messages/AddTransportCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServer).AddTransportCompany(ctx, req.(*MessageToTransportCompany))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messages_DeleteTransportCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageToTransportCompany)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServer).DeleteTransportCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.Messages/DeleteTransportCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServer).DeleteTransportCompany(ctx, req.(*MessageToTransportCompany))
	}
	return interceptor(ctx, in, info, handler)
}

// Messages_ServiceDesc is the grpc.ServiceDesc for Messages service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Messages_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logistics.Messages",
	HandlerType: (*MessagesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Messages_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Messages_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Messages_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Messages_Update_Handler,
		},
		{
			MethodName: "Upsert",
			Handler:    _Messages_Upsert_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Messages_Delete_Handler,
		},
		{
			MethodName: "AddTransportCompany",
			Handler:    _Messages_AddTransportCompany_Handler,
		},
		{
			MethodName: "DeleteTransportCompany",
			Handler:    _Messages_DeleteTransportCompany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/logistics/message.proto",
}
