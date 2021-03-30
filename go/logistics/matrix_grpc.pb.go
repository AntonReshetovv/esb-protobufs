// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package logistics

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// MatricesClient is the client API for Matrices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MatricesClient interface {
	Get(ctx context.Context, in *GetMatrixRequest, opts ...grpc.CallOption) (*Matrix, error)
	List(ctx context.Context, in *ListMatrixRequest, opts ...grpc.CallOption) (*ListMatrixResponse, error)
}

type matricesClient struct {
	cc grpc.ClientConnInterface
}

func NewMatricesClient(cc grpc.ClientConnInterface) MatricesClient {
	return &matricesClient{cc}
}

func (c *matricesClient) Get(ctx context.Context, in *GetMatrixRequest, opts ...grpc.CallOption) (*Matrix, error) {
	out := new(Matrix)
	err := c.cc.Invoke(ctx, "/logistics.Matrices/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matricesClient) List(ctx context.Context, in *ListMatrixRequest, opts ...grpc.CallOption) (*ListMatrixResponse, error) {
	out := new(ListMatrixResponse)
	err := c.cc.Invoke(ctx, "/logistics.Matrices/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatricesServer is the server API for Matrices service.
// All implementations must embed UnimplementedMatricesServer
// for forward compatibility
type MatricesServer interface {
	Get(context.Context, *GetMatrixRequest) (*Matrix, error)
	List(context.Context, *ListMatrixRequest) (*ListMatrixResponse, error)
	mustEmbedUnimplementedMatricesServer()
}

// UnimplementedMatricesServer must be embedded to have forward compatible implementations.
type UnimplementedMatricesServer struct {
}

func (UnimplementedMatricesServer) Get(context.Context, *GetMatrixRequest) (*Matrix, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedMatricesServer) List(context.Context, *ListMatrixRequest) (*ListMatrixResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedMatricesServer) mustEmbedUnimplementedMatricesServer() {}

// UnsafeMatricesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MatricesServer will
// result in compilation errors.
type UnsafeMatricesServer interface {
	mustEmbedUnimplementedMatricesServer()
}

func RegisterMatricesServer(s grpc.ServiceRegistrar, srv MatricesServer) {
	s.RegisterService(&_Matrices_serviceDesc, srv)
}

func _Matrices_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMatrixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatricesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.Matrices/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatricesServer).Get(ctx, req.(*GetMatrixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Matrices_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMatrixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatricesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.Matrices/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatricesServer).List(ctx, req.(*ListMatrixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Matrices_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logistics.Matrices",
	HandlerType: (*MatricesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Matrices_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Matrices_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/logistics/matrix.proto",
}