// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	RegisterClientPortOnPrimaryServer(ctx context.Context, in *RequestPort, opts ...grpc.CallOption) (*ReturnPort, error)
	Bid(ctx context.Context, in *RequestBid, opts ...grpc.CallOption) (*Ack, error)
	Result(ctx context.Context, in *RequestStatus, opts ...grpc.CallOption) (*AuctionResults, error)
	PublishReplicate(ctx context.Context, in *Replica, opts ...grpc.CallOption) (*Ack, error)
	ReplicateHighestBid(ctx context.Context, in *HighestBidInfo, opts ...grpc.CallOption) (*Ack, error)
	Broadcast(ctx context.Context, in *RequestBid, opts ...grpc.CallOption) (*Ack, error)
	GetName(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) RegisterClientPortOnPrimaryServer(ctx context.Context, in *RequestPort, opts ...grpc.CallOption) (*ReturnPort, error) {
	out := new(ReturnPort)
	err := c.cc.Invoke(ctx, "/service.Service/RegisterClientPortOnPrimaryServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Bid(ctx context.Context, in *RequestBid, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/service.Service/Bid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Result(ctx context.Context, in *RequestStatus, opts ...grpc.CallOption) (*AuctionResults, error) {
	out := new(AuctionResults)
	err := c.cc.Invoke(ctx, "/service.Service/Result", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) PublishReplicate(ctx context.Context, in *Replica, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/service.Service/PublishReplicate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ReplicateHighestBid(ctx context.Context, in *HighestBidInfo, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/service.Service/ReplicateHighestBid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Broadcast(ctx context.Context, in *RequestBid, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/service.Service/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetName(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/service.Service/GetName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	RegisterClientPortOnPrimaryServer(context.Context, *RequestPort) (*ReturnPort, error)
	Bid(context.Context, *RequestBid) (*Ack, error)
	Result(context.Context, *RequestStatus) (*AuctionResults, error)
	PublishReplicate(context.Context, *Replica) (*Ack, error)
	ReplicateHighestBid(context.Context, *HighestBidInfo) (*Ack, error)
	Broadcast(context.Context, *RequestBid) (*Ack, error)
	GetName(context.Context, *InfoRequest) (*InfoResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) RegisterClientPortOnPrimaryServer(context.Context, *RequestPort) (*ReturnPort, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterClientPortOnPrimaryServer not implemented")
}
func (UnimplementedServiceServer) Bid(context.Context, *RequestBid) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bid not implemented")
}
func (UnimplementedServiceServer) Result(context.Context, *RequestStatus) (*AuctionResults, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Result not implemented")
}
func (UnimplementedServiceServer) PublishReplicate(context.Context, *Replica) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishReplicate not implemented")
}
func (UnimplementedServiceServer) ReplicateHighestBid(context.Context, *HighestBidInfo) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplicateHighestBid not implemented")
}
func (UnimplementedServiceServer) Broadcast(context.Context, *RequestBid) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedServiceServer) GetName(context.Context, *InfoRequest) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetName not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_RegisterClientPortOnPrimaryServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPort)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).RegisterClientPortOnPrimaryServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/RegisterClientPortOnPrimaryServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).RegisterClientPortOnPrimaryServer(ctx, req.(*RequestPort))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Bid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestBid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Bid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/Bid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Bid(ctx, req.(*RequestBid))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Result_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Result(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/Result",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Result(ctx, req.(*RequestStatus))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_PublishReplicate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Replica)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).PublishReplicate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/PublishReplicate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).PublishReplicate(ctx, req.(*Replica))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ReplicateHighestBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HighestBidInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ReplicateHighestBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/ReplicateHighestBid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ReplicateHighestBid(ctx, req.(*HighestBidInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestBid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Broadcast(ctx, req.(*RequestBid))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/GetName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetName(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterClientPortOnPrimaryServer",
			Handler:    _Service_RegisterClientPortOnPrimaryServer_Handler,
		},
		{
			MethodName: "Bid",
			Handler:    _Service_Bid_Handler,
		},
		{
			MethodName: "Result",
			Handler:    _Service_Result_Handler,
		},
		{
			MethodName: "PublishReplicate",
			Handler:    _Service_PublishReplicate_Handler,
		},
		{
			MethodName: "ReplicateHighestBid",
			Handler:    _Service_ReplicateHighestBid_Handler,
		},
		{
			MethodName: "Broadcast",
			Handler:    _Service_Broadcast_Handler,
		},
		{
			MethodName: "GetName",
			Handler:    _Service_GetName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/service.proto",
}
