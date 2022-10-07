// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package account

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

// DepositServiceClient is the client API for DepositService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DepositServiceClient interface {
	Deposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*DepositResponse, error)
}

type depositServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDepositServiceClient(cc grpc.ClientConnInterface) DepositServiceClient {
	return &depositServiceClient{cc}
}

func (c *depositServiceClient) Deposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*DepositResponse, error) {
	out := new(DepositResponse)
	err := c.cc.Invoke(ctx, "/account.DepositService/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DepositServiceServer is the server API for DepositService service.
// All implementations must embed UnimplementedDepositServiceServer
// for forward compatibility
type DepositServiceServer interface {
	Deposit(context.Context, *DepositRequest) (*DepositResponse, error)
	mustEmbedUnimplementedDepositServiceServer()
}

// UnimplementedDepositServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDepositServiceServer struct {
}

func (UnimplementedDepositServiceServer) Deposit(context.Context, *DepositRequest) (*DepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (UnimplementedDepositServiceServer) mustEmbedUnimplementedDepositServiceServer() {}

// UnsafeDepositServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DepositServiceServer will
// result in compilation errors.
type UnsafeDepositServiceServer interface {
	mustEmbedUnimplementedDepositServiceServer()
}

func RegisterDepositServiceServer(s grpc.ServiceRegistrar, srv DepositServiceServer) {
	s.RegisterService(&DepositService_ServiceDesc, srv)
}

func _DepositService_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepositServiceServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.DepositService/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepositServiceServer).Deposit(ctx, req.(*DepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DepositService_ServiceDesc is the grpc.ServiceDesc for DepositService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DepositService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.DepositService",
	HandlerType: (*DepositServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Deposit",
			Handler:    _DepositService_Deposit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/deposit.proto",
}
