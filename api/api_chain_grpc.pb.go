// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// ChainServiceClient is the client API for ChainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChainServiceClient interface {
	Chain(ctx context.Context, in *ChainRequest, opts ...grpc.CallOption) (*ChainResponse, error)
	// MostRecentTPS gives the latest transactions per second
	MostRecentTPS(ctx context.Context, in *MostRecentTPSRequest, opts ...grpc.CallOption) (*MostRecentTPSResponse, error)
	// NumberOfActions gives the number of actions
	NumberOfActions(ctx context.Context, in *NumberOfActionsRequest, opts ...grpc.CallOption) (*NumberOfActionsResponse, error)
	// TotalTransferredTokens gives the amount of tokens transferred within a time frame
	TotalTransferredTokens(ctx context.Context, in *TotalTransferredTokensRequest, opts ...grpc.CallOption) (*TotalTransferredTokensResponse, error)
	// ChartSync gives the chart sync status
	ChartSync(ctx context.Context, in *ChartSyncRequest, opts ...grpc.CallOption) (*ChartSyncResponse, error)
}

type chainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChainServiceClient(cc grpc.ClientConnInterface) ChainServiceClient {
	return &chainServiceClient{cc}
}

func (c *chainServiceClient) Chain(ctx context.Context, in *ChainRequest, opts ...grpc.CallOption) (*ChainResponse, error) {
	out := new(ChainResponse)
	err := c.cc.Invoke(ctx, "/api.ChainService/Chain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainServiceClient) MostRecentTPS(ctx context.Context, in *MostRecentTPSRequest, opts ...grpc.CallOption) (*MostRecentTPSResponse, error) {
	out := new(MostRecentTPSResponse)
	err := c.cc.Invoke(ctx, "/api.ChainService/MostRecentTPS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainServiceClient) NumberOfActions(ctx context.Context, in *NumberOfActionsRequest, opts ...grpc.CallOption) (*NumberOfActionsResponse, error) {
	out := new(NumberOfActionsResponse)
	err := c.cc.Invoke(ctx, "/api.ChainService/NumberOfActions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainServiceClient) TotalTransferredTokens(ctx context.Context, in *TotalTransferredTokensRequest, opts ...grpc.CallOption) (*TotalTransferredTokensResponse, error) {
	out := new(TotalTransferredTokensResponse)
	err := c.cc.Invoke(ctx, "/api.ChainService/TotalTransferredTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainServiceClient) ChartSync(ctx context.Context, in *ChartSyncRequest, opts ...grpc.CallOption) (*ChartSyncResponse, error) {
	out := new(ChartSyncResponse)
	err := c.cc.Invoke(ctx, "/api.ChainService/ChartSync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChainServiceServer is the server API for ChainService service.
// All implementations must embed UnimplementedChainServiceServer
// for forward compatibility
type ChainServiceServer interface {
	Chain(context.Context, *ChainRequest) (*ChainResponse, error)
	// MostRecentTPS gives the latest transactions per second
	MostRecentTPS(context.Context, *MostRecentTPSRequest) (*MostRecentTPSResponse, error)
	// NumberOfActions gives the number of actions
	NumberOfActions(context.Context, *NumberOfActionsRequest) (*NumberOfActionsResponse, error)
	// TotalTransferredTokens gives the amount of tokens transferred within a time frame
	TotalTransferredTokens(context.Context, *TotalTransferredTokensRequest) (*TotalTransferredTokensResponse, error)
	// ChartSync gives the chart sync status
	ChartSync(context.Context, *ChartSyncRequest) (*ChartSyncResponse, error)
	mustEmbedUnimplementedChainServiceServer()
}

// UnimplementedChainServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChainServiceServer struct {
}

func (UnimplementedChainServiceServer) Chain(context.Context, *ChainRequest) (*ChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Chain not implemented")
}
func (UnimplementedChainServiceServer) MostRecentTPS(context.Context, *MostRecentTPSRequest) (*MostRecentTPSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MostRecentTPS not implemented")
}
func (UnimplementedChainServiceServer) NumberOfActions(context.Context, *NumberOfActionsRequest) (*NumberOfActionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NumberOfActions not implemented")
}
func (UnimplementedChainServiceServer) TotalTransferredTokens(context.Context, *TotalTransferredTokensRequest) (*TotalTransferredTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalTransferredTokens not implemented")
}
func (UnimplementedChainServiceServer) ChartSync(context.Context, *ChartSyncRequest) (*ChartSyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChartSync not implemented")
}
func (UnimplementedChainServiceServer) mustEmbedUnimplementedChainServiceServer() {}

// UnsafeChainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChainServiceServer will
// result in compilation errors.
type UnsafeChainServiceServer interface {
	mustEmbedUnimplementedChainServiceServer()
}

func RegisterChainServiceServer(s grpc.ServiceRegistrar, srv ChainServiceServer) {
	s.RegisterService(&ChainService_ServiceDesc, srv)
}

func _ChainService_Chain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).Chain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ChainService/Chain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).Chain(ctx, req.(*ChainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainService_MostRecentTPS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MostRecentTPSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).MostRecentTPS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ChainService/MostRecentTPS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).MostRecentTPS(ctx, req.(*MostRecentTPSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainService_NumberOfActions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NumberOfActionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).NumberOfActions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ChainService/NumberOfActions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).NumberOfActions(ctx, req.(*NumberOfActionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainService_TotalTransferredTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TotalTransferredTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).TotalTransferredTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ChainService/TotalTransferredTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).TotalTransferredTokens(ctx, req.(*TotalTransferredTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainService_ChartSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChartSyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).ChartSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ChainService/ChartSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).ChartSync(ctx, req.(*ChartSyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChainService_ServiceDesc is the grpc.ServiceDesc for ChainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ChainService",
	HandlerType: (*ChainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Chain",
			Handler:    _ChainService_Chain_Handler,
		},
		{
			MethodName: "MostRecentTPS",
			Handler:    _ChainService_MostRecentTPS_Handler,
		},
		{
			MethodName: "NumberOfActions",
			Handler:    _ChainService_NumberOfActions_Handler,
		},
		{
			MethodName: "TotalTransferredTokens",
			Handler:    _ChainService_TotalTransferredTokens_Handler,
		},
		{
			MethodName: "ChartSync",
			Handler:    _ChainService_ChartSync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api_chain.proto",
}
