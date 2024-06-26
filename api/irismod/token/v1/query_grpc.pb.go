// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: irismod/token/v1/query.proto

package tokenv1

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
	Query_Tokens_FullMethodName    = "/irismod.token.v1.Query/Tokens"
	Query_Token_FullMethodName     = "/irismod.token.v1.Query/Token"
	Query_Fees_FullMethodName      = "/irismod.token.v1.Query/Fees"
	Query_Params_FullMethodName    = "/irismod.token.v1.Query/Params"
	Query_TotalBurn_FullMethodName = "/irismod.token.v1.Query/TotalBurn"
	Query_Balances_FullMethodName  = "/irismod.token.v1.Query/Balances"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Tokens returns the token list
	Tokens(ctx context.Context, in *QueryTokensRequest, opts ...grpc.CallOption) (*QueryTokensResponse, error)
	// Token returns token with token name
	Token(ctx context.Context, in *QueryTokenRequest, opts ...grpc.CallOption) (*QueryTokenResponse, error)
	// Fees returns the fees to issue or mint a token
	Fees(ctx context.Context, in *QueryFeesRequest, opts ...grpc.CallOption) (*QueryFeesResponse, error)
	// Params queries the token parameters
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// TotalBurn queries all the burnt coins
	TotalBurn(ctx context.Context, in *QueryTotalBurnRequest, opts ...grpc.CallOption) (*QueryTotalBurnResponse, error)
	// Balances queries the balance of the specified token (including erc20
	// balance)
	Balances(ctx context.Context, in *QueryBalancesRequest, opts ...grpc.CallOption) (*QueryBalancesResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Tokens(ctx context.Context, in *QueryTokensRequest, opts ...grpc.CallOption) (*QueryTokensResponse, error) {
	out := new(QueryTokensResponse)
	err := c.cc.Invoke(ctx, Query_Tokens_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Token(ctx context.Context, in *QueryTokenRequest, opts ...grpc.CallOption) (*QueryTokenResponse, error) {
	out := new(QueryTokenResponse)
	err := c.cc.Invoke(ctx, Query_Token_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Fees(ctx context.Context, in *QueryFeesRequest, opts ...grpc.CallOption) (*QueryFeesResponse, error) {
	out := new(QueryFeesResponse)
	err := c.cc.Invoke(ctx, Query_Fees_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TotalBurn(ctx context.Context, in *QueryTotalBurnRequest, opts ...grpc.CallOption) (*QueryTotalBurnResponse, error) {
	out := new(QueryTotalBurnResponse)
	err := c.cc.Invoke(ctx, Query_TotalBurn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Balances(ctx context.Context, in *QueryBalancesRequest, opts ...grpc.CallOption) (*QueryBalancesResponse, error) {
	out := new(QueryBalancesResponse)
	err := c.cc.Invoke(ctx, Query_Balances_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Tokens returns the token list
	Tokens(context.Context, *QueryTokensRequest) (*QueryTokensResponse, error)
	// Token returns token with token name
	Token(context.Context, *QueryTokenRequest) (*QueryTokenResponse, error)
	// Fees returns the fees to issue or mint a token
	Fees(context.Context, *QueryFeesRequest) (*QueryFeesResponse, error)
	// Params queries the token parameters
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// TotalBurn queries all the burnt coins
	TotalBurn(context.Context, *QueryTotalBurnRequest) (*QueryTotalBurnResponse, error)
	// Balances queries the balance of the specified token (including erc20
	// balance)
	Balances(context.Context, *QueryBalancesRequest) (*QueryBalancesResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Tokens(context.Context, *QueryTokensRequest) (*QueryTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Tokens not implemented")
}
func (UnimplementedQueryServer) Token(context.Context, *QueryTokenRequest) (*QueryTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Token not implemented")
}
func (UnimplementedQueryServer) Fees(context.Context, *QueryFeesRequest) (*QueryFeesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fees not implemented")
}
func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) TotalBurn(context.Context, *QueryTotalBurnRequest) (*QueryTotalBurnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalBurn not implemented")
}
func (UnimplementedQueryServer) Balances(context.Context, *QueryBalancesRequest) (*QueryBalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balances not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Tokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Tokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Tokens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Tokens(ctx, req.(*QueryTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Token_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Token(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Token_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Token(ctx, req.(*QueryTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Fees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFeesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Fees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Fees_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Fees(ctx, req.(*QueryFeesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TotalBurn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTotalBurnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TotalBurn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_TotalBurn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TotalBurn(ctx, req.(*QueryTotalBurnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Balances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBalancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Balances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Balances_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Balances(ctx, req.(*QueryBalancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "irismod.token.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Tokens",
			Handler:    _Query_Tokens_Handler,
		},
		{
			MethodName: "Token",
			Handler:    _Query_Token_Handler,
		},
		{
			MethodName: "Fees",
			Handler:    _Query_Fees_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "TotalBurn",
			Handler:    _Query_TotalBurn_Handler,
		},
		{
			MethodName: "Balances",
			Handler:    _Query_Balances_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "irismod/token/v1/query.proto",
}
