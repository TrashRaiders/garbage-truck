// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

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

// ShopServiceClient is the client API for ShopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShopServiceClient interface {
	CreateShop(ctx context.Context, in *Shop, opts ...grpc.CallOption) (*Reply, error)
	GetShop(ctx context.Context, in *ShopRequest, opts ...grpc.CallOption) (*Shop, error)
}

type shopServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShopServiceClient(cc grpc.ClientConnInterface) ShopServiceClient {
	return &shopServiceClient{cc}
}

func (c *shopServiceClient) CreateShop(ctx context.Context, in *Shop, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/protos.ShopService/createShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) GetShop(ctx context.Context, in *ShopRequest, opts ...grpc.CallOption) (*Shop, error) {
	out := new(Shop)
	err := c.cc.Invoke(ctx, "/protos.ShopService/getShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopServiceServer is the server API for ShopService service.
// All implementations must embed UnimplementedShopServiceServer
// for forward compatibility
type ShopServiceServer interface {
	CreateShop(context.Context, *Shop) (*Reply, error)
	GetShop(context.Context, *ShopRequest) (*Shop, error)
	mustEmbedUnimplementedShopServiceServer()
}

// UnimplementedShopServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShopServiceServer struct {
}

func (UnimplementedShopServiceServer) CreateShop(context.Context, *Shop) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShop not implemented")
}
func (UnimplementedShopServiceServer) GetShop(context.Context, *ShopRequest) (*Shop, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShop not implemented")
}
func (UnimplementedShopServiceServer) mustEmbedUnimplementedShopServiceServer() {}

// UnsafeShopServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShopServiceServer will
// result in compilation errors.
type UnsafeShopServiceServer interface {
	mustEmbedUnimplementedShopServiceServer()
}

func RegisterShopServiceServer(s grpc.ServiceRegistrar, srv ShopServiceServer) {
	s.RegisterService(&ShopService_ServiceDesc, srv)
}

func _ShopService_CreateShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Shop)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).CreateShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ShopService/createShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).CreateShop(ctx, req.(*Shop))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_GetShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ShopService/getShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetShop(ctx, req.(*ShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShopService_ServiceDesc is the grpc.ServiceDesc for ShopService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShopService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ShopService",
	HandlerType: (*ShopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createShop",
			Handler:    _ShopService_CreateShop_Handler,
		},
		{
			MethodName: "getShop",
			Handler:    _ShopService_GetShop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
