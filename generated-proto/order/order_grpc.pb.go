// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package order

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrderManagementClient is the client API for OrderManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderManagementClient interface {
	//unary
	GetOrder(ctx context.Context, in *ProductType, opts ...grpc.CallOption) (*Order, error)
	// server side streaming
	// when i am in my local pc
	GetOrderItems(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (OrderManagement_GetOrderItemsClient, error)
	// client side streaming
	AddOrder(ctx context.Context, opts ...grpc.CallOption) (OrderManagement_AddOrderClient, error)
	GetAllOrders(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Orders, error)
	// bidirectional streaming
	AddOrderAndRec(ctx context.Context, opts ...grpc.CallOption) (OrderManagement_AddOrderAndRecClient, error)
}

type orderManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderManagementClient(cc grpc.ClientConnInterface) OrderManagementClient {
	return &orderManagementClient{cc}
}

func (c *orderManagementClient) GetOrder(ctx context.Context, in *ProductType, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/order.OrderManagement/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderManagementClient) GetOrderItems(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (OrderManagement_GetOrderItemsClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderManagement_ServiceDesc.Streams[0], "/order.OrderManagement/GetOrderItems", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderManagementGetOrderItemsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderManagement_GetOrderItemsClient interface {
	Recv() (*Order, error)
	grpc.ClientStream
}

type orderManagementGetOrderItemsClient struct {
	grpc.ClientStream
}

func (x *orderManagementGetOrderItemsClient) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderManagementClient) AddOrder(ctx context.Context, opts ...grpc.CallOption) (OrderManagement_AddOrderClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderManagement_ServiceDesc.Streams[1], "/order.OrderManagement/AddOrder", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderManagementAddOrderClient{stream}
	return x, nil
}

type OrderManagement_AddOrderClient interface {
	Send(*MapDb) error
	CloseAndRecv() (*wrapperspb.StringValue, error)
	grpc.ClientStream
}

type orderManagementAddOrderClient struct {
	grpc.ClientStream
}

func (x *orderManagementAddOrderClient) Send(m *MapDb) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderManagementAddOrderClient) CloseAndRecv() (*wrapperspb.StringValue, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(wrapperspb.StringValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderManagementClient) GetAllOrders(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Orders, error) {
	out := new(Orders)
	err := c.cc.Invoke(ctx, "/order.OrderManagement/GetAllOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderManagementClient) AddOrderAndRec(ctx context.Context, opts ...grpc.CallOption) (OrderManagement_AddOrderAndRecClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderManagement_ServiceDesc.Streams[2], "/order.OrderManagement/AddOrderAndRec", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderManagementAddOrderAndRecClient{stream}
	return x, nil
}

type OrderManagement_AddOrderAndRecClient interface {
	Send(*MapDb) error
	Recv() (*Order, error)
	grpc.ClientStream
}

type orderManagementAddOrderAndRecClient struct {
	grpc.ClientStream
}

func (x *orderManagementAddOrderAndRecClient) Send(m *MapDb) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderManagementAddOrderAndRecClient) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderManagementServer is the server API for OrderManagement service.
// All implementations must embed UnimplementedOrderManagementServer
// for forward compatibility
type OrderManagementServer interface {
	//unary
	GetOrder(context.Context, *ProductType) (*Order, error)
	// server side streaming
	// when i am in my local pc
	GetOrderItems(*wrapperspb.StringValue, OrderManagement_GetOrderItemsServer) error
	// client side streaming
	AddOrder(OrderManagement_AddOrderServer) error
	GetAllOrders(context.Context, *emptypb.Empty) (*Orders, error)
	// bidirectional streaming
	AddOrderAndRec(OrderManagement_AddOrderAndRecServer) error
	//mustEmbedUnimplementedOrderManagementServer()
}

// UnimplementedOrderManagementServer must be embedded to have forward compatible implementations.
type UnimplementedOrderManagementServer struct {
}

func (UnimplementedOrderManagementServer) GetOrder(context.Context, *ProductType) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrderManagementServer) GetOrderItems(*wrapperspb.StringValue, OrderManagement_GetOrderItemsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetOrderItems not implemented")
}
func (UnimplementedOrderManagementServer) AddOrder(OrderManagement_AddOrderServer) error {
	return status.Errorf(codes.Unimplemented, "method AddOrder not implemented")
}
func (UnimplementedOrderManagementServer) GetAllOrders(context.Context, *emptypb.Empty) (*Orders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllOrders not implemented")
}
func (UnimplementedOrderManagementServer) AddOrderAndRec(OrderManagement_AddOrderAndRecServer) error {
	return status.Errorf(codes.Unimplemented, "method AddOrderAndRec not implemented")
}
func (UnimplementedOrderManagementServer) mustEmbedUnimplementedOrderManagementServer() {}

// UnsafeOrderManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderManagementServer will
// result in compilation errors.
type UnsafeOrderManagementServer interface {
	mustEmbedUnimplementedOrderManagementServer()
}

func RegisterOrderManagementServer(s grpc.ServiceRegistrar, srv OrderManagementServer) {
	s.RegisterService(&OrderManagement_ServiceDesc, srv)
}

func _OrderManagement_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderManagement/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServer).GetOrder(ctx, req.(*ProductType))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderManagement_GetOrderItems_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(wrapperspb.StringValue)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderManagementServer).GetOrderItems(m, &orderManagementGetOrderItemsServer{stream})
}

type OrderManagement_GetOrderItemsServer interface {
	Send(*Order) error
	grpc.ServerStream
}

type orderManagementGetOrderItemsServer struct {
	grpc.ServerStream
}

func (x *orderManagementGetOrderItemsServer) Send(m *Order) error {
	return x.ServerStream.SendMsg(m)
}

func _OrderManagement_AddOrder_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderManagementServer).AddOrder(&orderManagementAddOrderServer{stream})
}

type OrderManagement_AddOrderServer interface {
	SendAndClose(*wrapperspb.StringValue) error
	Recv() (*MapDb, error)
	grpc.ServerStream
}

type orderManagementAddOrderServer struct {
	grpc.ServerStream
}

func (x *orderManagementAddOrderServer) SendAndClose(m *wrapperspb.StringValue) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderManagementAddOrderServer) Recv() (*MapDb, error) {
	m := new(MapDb)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _OrderManagement_GetAllOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServer).GetAllOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderManagement/GetAllOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServer).GetAllOrders(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderManagement_AddOrderAndRec_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderManagementServer).AddOrderAndRec(&orderManagementAddOrderAndRecServer{stream})
}

type OrderManagement_AddOrderAndRecServer interface {
	Send(*Order) error
	Recv() (*MapDb, error)
	grpc.ServerStream
}

type orderManagementAddOrderAndRecServer struct {
	grpc.ServerStream
}

func (x *orderManagementAddOrderAndRecServer) Send(m *Order) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderManagementAddOrderAndRecServer) Recv() (*MapDb, error) {
	m := new(MapDb)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderManagement_ServiceDesc is the grpc.ServiceDesc for OrderManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.OrderManagement",
	HandlerType: (*OrderManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrder",
			Handler:    _OrderManagement_GetOrder_Handler,
		},
		{
			MethodName: "GetAllOrders",
			Handler:    _OrderManagement_GetAllOrders_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetOrderItems",
			Handler:       _OrderManagement_GetOrderItems_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AddOrder",
			Handler:       _OrderManagement_AddOrder_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "AddOrderAndRec",
			Handler:       _OrderManagement_AddOrderAndRec_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/order.proto",
}
