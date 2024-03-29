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

// DelegateServiceClient is the client API for DelegateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DelegateServiceClient interface {
	// BucketInfo provides voting bucket detail information for candidates within a range of epochs
	BucketInfo(ctx context.Context, in *BucketInfoRequest, opts ...grpc.CallOption) (*BucketInfoResponse, error)
	// BookKeeping gives delegates an overview of the reward distributions to their voters within a range of epochs
	BookKeeping(ctx context.Context, in *BookKeepingRequest, opts ...grpc.CallOption) (*BookKeepingResponse, error)
	// Productivity gives block productivity of producers within a range of epochs
	Productivity(ctx context.Context, in *ProductivityRequest, opts ...grpc.CallOption) (*ProductivityResponse, error)
	// Rewards provides reward detail information for candidates within a range of epochs
	Reward(ctx context.Context, in *RewardRequest, opts ...grpc.CallOption) (*RewardResponse, error)
	// Staking provides staking information for candidates within a range of epochs
	Staking(ctx context.Context, in *StakingRequest, opts ...grpc.CallOption) (*StakingResponse, error)
	// ProbationHistoricalRate provides the rate of probation for a given delegate
	ProbationHistoricalRate(ctx context.Context, in *ProbationHistoricalRateRequest, opts ...grpc.CallOption) (*ProbationHistoricalRateResponse, error)
	// PaidToDelegates provides the amount of rewards paid to delegates
	PaidToDelegates(ctx context.Context, in *PaidToDelegatesRequest, opts ...grpc.CallOption) (*PaidToDelegatesResponse, error)
}

type delegateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDelegateServiceClient(cc grpc.ClientConnInterface) DelegateServiceClient {
	return &delegateServiceClient{cc}
}

func (c *delegateServiceClient) BucketInfo(ctx context.Context, in *BucketInfoRequest, opts ...grpc.CallOption) (*BucketInfoResponse, error) {
	out := new(BucketInfoResponse)
	err := c.cc.Invoke(ctx, "/api.DelegateService/BucketInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *delegateServiceClient) BookKeeping(ctx context.Context, in *BookKeepingRequest, opts ...grpc.CallOption) (*BookKeepingResponse, error) {
	out := new(BookKeepingResponse)
	err := c.cc.Invoke(ctx, "/api.DelegateService/BookKeeping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *delegateServiceClient) Productivity(ctx context.Context, in *ProductivityRequest, opts ...grpc.CallOption) (*ProductivityResponse, error) {
	out := new(ProductivityResponse)
	err := c.cc.Invoke(ctx, "/api.DelegateService/Productivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *delegateServiceClient) Reward(ctx context.Context, in *RewardRequest, opts ...grpc.CallOption) (*RewardResponse, error) {
	out := new(RewardResponse)
	err := c.cc.Invoke(ctx, "/api.DelegateService/Reward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *delegateServiceClient) Staking(ctx context.Context, in *StakingRequest, opts ...grpc.CallOption) (*StakingResponse, error) {
	out := new(StakingResponse)
	err := c.cc.Invoke(ctx, "/api.DelegateService/Staking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *delegateServiceClient) ProbationHistoricalRate(ctx context.Context, in *ProbationHistoricalRateRequest, opts ...grpc.CallOption) (*ProbationHistoricalRateResponse, error) {
	out := new(ProbationHistoricalRateResponse)
	err := c.cc.Invoke(ctx, "/api.DelegateService/ProbationHistoricalRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *delegateServiceClient) PaidToDelegates(ctx context.Context, in *PaidToDelegatesRequest, opts ...grpc.CallOption) (*PaidToDelegatesResponse, error) {
	out := new(PaidToDelegatesResponse)
	err := c.cc.Invoke(ctx, "/api.DelegateService/PaidToDelegates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DelegateServiceServer is the server API for DelegateService service.
// All implementations must embed UnimplementedDelegateServiceServer
// for forward compatibility
type DelegateServiceServer interface {
	// BucketInfo provides voting bucket detail information for candidates within a range of epochs
	BucketInfo(context.Context, *BucketInfoRequest) (*BucketInfoResponse, error)
	// BookKeeping gives delegates an overview of the reward distributions to their voters within a range of epochs
	BookKeeping(context.Context, *BookKeepingRequest) (*BookKeepingResponse, error)
	// Productivity gives block productivity of producers within a range of epochs
	Productivity(context.Context, *ProductivityRequest) (*ProductivityResponse, error)
	// Rewards provides reward detail information for candidates within a range of epochs
	Reward(context.Context, *RewardRequest) (*RewardResponse, error)
	// Staking provides staking information for candidates within a range of epochs
	Staking(context.Context, *StakingRequest) (*StakingResponse, error)
	// ProbationHistoricalRate provides the rate of probation for a given delegate
	ProbationHistoricalRate(context.Context, *ProbationHistoricalRateRequest) (*ProbationHistoricalRateResponse, error)
	// PaidToDelegates provides the amount of rewards paid to delegates
	PaidToDelegates(context.Context, *PaidToDelegatesRequest) (*PaidToDelegatesResponse, error)
	mustEmbedUnimplementedDelegateServiceServer()
}

// UnimplementedDelegateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDelegateServiceServer struct {
}

func (UnimplementedDelegateServiceServer) BucketInfo(context.Context, *BucketInfoRequest) (*BucketInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BucketInfo not implemented")
}
func (UnimplementedDelegateServiceServer) BookKeeping(context.Context, *BookKeepingRequest) (*BookKeepingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookKeeping not implemented")
}
func (UnimplementedDelegateServiceServer) Productivity(context.Context, *ProductivityRequest) (*ProductivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Productivity not implemented")
}
func (UnimplementedDelegateServiceServer) Reward(context.Context, *RewardRequest) (*RewardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reward not implemented")
}
func (UnimplementedDelegateServiceServer) Staking(context.Context, *StakingRequest) (*StakingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Staking not implemented")
}
func (UnimplementedDelegateServiceServer) ProbationHistoricalRate(context.Context, *ProbationHistoricalRateRequest) (*ProbationHistoricalRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProbationHistoricalRate not implemented")
}
func (UnimplementedDelegateServiceServer) PaidToDelegates(context.Context, *PaidToDelegatesRequest) (*PaidToDelegatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PaidToDelegates not implemented")
}
func (UnimplementedDelegateServiceServer) mustEmbedUnimplementedDelegateServiceServer() {}

// UnsafeDelegateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DelegateServiceServer will
// result in compilation errors.
type UnsafeDelegateServiceServer interface {
	mustEmbedUnimplementedDelegateServiceServer()
}

func RegisterDelegateServiceServer(s grpc.ServiceRegistrar, srv DelegateServiceServer) {
	s.RegisterService(&DelegateService_ServiceDesc, srv)
}

func _DelegateService_BucketInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BucketInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DelegateServiceServer).BucketInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DelegateService/BucketInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DelegateServiceServer).BucketInfo(ctx, req.(*BucketInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DelegateService_BookKeeping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookKeepingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DelegateServiceServer).BookKeeping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DelegateService/BookKeeping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DelegateServiceServer).BookKeeping(ctx, req.(*BookKeepingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DelegateService_Productivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DelegateServiceServer).Productivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DelegateService/Productivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DelegateServiceServer).Productivity(ctx, req.(*ProductivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DelegateService_Reward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RewardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DelegateServiceServer).Reward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DelegateService/Reward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DelegateServiceServer).Reward(ctx, req.(*RewardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DelegateService_Staking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StakingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DelegateServiceServer).Staking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DelegateService/Staking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DelegateServiceServer).Staking(ctx, req.(*StakingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DelegateService_ProbationHistoricalRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProbationHistoricalRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DelegateServiceServer).ProbationHistoricalRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DelegateService/ProbationHistoricalRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DelegateServiceServer).ProbationHistoricalRate(ctx, req.(*ProbationHistoricalRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DelegateService_PaidToDelegates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaidToDelegatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DelegateServiceServer).PaidToDelegates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DelegateService/PaidToDelegates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DelegateServiceServer).PaidToDelegates(ctx, req.(*PaidToDelegatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DelegateService_ServiceDesc is the grpc.ServiceDesc for DelegateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DelegateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.DelegateService",
	HandlerType: (*DelegateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BucketInfo",
			Handler:    _DelegateService_BucketInfo_Handler,
		},
		{
			MethodName: "BookKeeping",
			Handler:    _DelegateService_BookKeeping_Handler,
		},
		{
			MethodName: "Productivity",
			Handler:    _DelegateService_Productivity_Handler,
		},
		{
			MethodName: "Reward",
			Handler:    _DelegateService_Reward_Handler,
		},
		{
			MethodName: "Staking",
			Handler:    _DelegateService_Staking_Handler,
		},
		{
			MethodName: "ProbationHistoricalRate",
			Handler:    _DelegateService_ProbationHistoricalRate_Handler,
		},
		{
			MethodName: "PaidToDelegates",
			Handler:    _DelegateService_PaidToDelegates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api_delegate.proto",
}
