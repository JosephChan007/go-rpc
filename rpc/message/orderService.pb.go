// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: orderService.proto

package message

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_orderService_proto protoreflect.FileDescriptor

var file_orderService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa2, 0x02,
	0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x76,
	0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x7b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x7d, 0x12, 0x3a, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x12, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x69, 0x7a, 0x65, 0x1a, 0x16, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x36, 0x0a,
	0x08, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x16, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_orderService_proto_goTypes = []interface{}{
	(*OrderRequest)(nil),       // 0: message.OrderRequest
	(*OrderSize)(nil),          // 1: message.OrderSize
	(*OrderData)(nil),          // 2: message.OrderData
	(*OrderStatusRequest)(nil), // 3: message.OrderStatusRequest
	(*OrderInfo)(nil),          // 4: message.OrderInfo
	(*OrderInfoList)(nil),      // 5: message.OrderInfoList
	(*OrderResponse)(nil),      // 6: message.OrderResponse
}
var file_orderService_proto_depIdxs = []int32{
	0, // 0: message.OrderService.GetOrder:input_type -> message.OrderRequest
	1, // 1: message.OrderService.GetOrderList:input_type -> message.OrderSize
	2, // 2: message.OrderService.NewOrder:input_type -> message.OrderData
	3, // 3: message.OrderService.GetStatusOrderList:input_type -> message.OrderStatusRequest
	4, // 4: message.OrderService.GetOrder:output_type -> message.OrderInfo
	5, // 5: message.OrderService.GetOrderList:output_type -> message.OrderInfoList
	6, // 6: message.OrderService.NewOrder:output_type -> message.OrderResponse
	5, // 7: message.OrderService.GetStatusOrderList:output_type -> message.OrderInfoList
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_orderService_proto_init() }
func file_orderService_proto_init() {
	if File_orderService_proto != nil {
		return
	}
	file_orderModel_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orderService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orderService_proto_goTypes,
		DependencyIndexes: file_orderService_proto_depIdxs,
	}.Build()
	File_orderService_proto = out.File
	file_orderService_proto_rawDesc = nil
	file_orderService_proto_goTypes = nil
	file_orderService_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderServiceClient interface {
	GetOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderInfo, error)
	GetOrderList(ctx context.Context, in *OrderSize, opts ...grpc.CallOption) (*OrderInfoList, error)
	NewOrder(ctx context.Context, in *OrderData, opts ...grpc.CallOption) (*OrderResponse, error)
	GetStatusOrderList(ctx context.Context, in *OrderStatusRequest, opts ...grpc.CallOption) (OrderService_GetStatusOrderListClient, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) GetOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderInfo, error) {
	out := new(OrderInfo)
	err := c.cc.Invoke(ctx, "/message.OrderService/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrderList(ctx context.Context, in *OrderSize, opts ...grpc.CallOption) (*OrderInfoList, error) {
	out := new(OrderInfoList)
	err := c.cc.Invoke(ctx, "/message.OrderService/GetOrderList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) NewOrder(ctx context.Context, in *OrderData, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/message.OrderService/NewOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetStatusOrderList(ctx context.Context, in *OrderStatusRequest, opts ...grpc.CallOption) (OrderService_GetStatusOrderListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OrderService_serviceDesc.Streams[0], "/message.OrderService/GetStatusOrderList", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderServiceGetStatusOrderListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderService_GetStatusOrderListClient interface {
	Recv() (*OrderInfoList, error)
	grpc.ClientStream
}

type orderServiceGetStatusOrderListClient struct {
	grpc.ClientStream
}

func (x *orderServiceGetStatusOrderListClient) Recv() (*OrderInfoList, error) {
	m := new(OrderInfoList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderServiceServer is the server API for OrderService service.
type OrderServiceServer interface {
	GetOrder(context.Context, *OrderRequest) (*OrderInfo, error)
	GetOrderList(context.Context, *OrderSize) (*OrderInfoList, error)
	NewOrder(context.Context, *OrderData) (*OrderResponse, error)
	GetStatusOrderList(*OrderStatusRequest, OrderService_GetStatusOrderListServer) error
}

// UnimplementedOrderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (*UnimplementedOrderServiceServer) GetOrder(context.Context, *OrderRequest) (*OrderInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (*UnimplementedOrderServiceServer) GetOrderList(context.Context, *OrderSize) (*OrderInfoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderList not implemented")
}
func (*UnimplementedOrderServiceServer) NewOrder(context.Context, *OrderData) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewOrder not implemented")
}
func (*UnimplementedOrderServiceServer) GetStatusOrderList(*OrderStatusRequest, OrderService_GetStatusOrderListServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStatusOrderList not implemented")
}

func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer) {
	s.RegisterService(&_OrderService_serviceDesc, srv)
}

func _OrderService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.OrderService/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrder(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderSize)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.OrderService/GetOrderList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrderList(ctx, req.(*OrderSize))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_NewOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).NewOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.OrderService/NewOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).NewOrder(ctx, req.(*OrderData))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetStatusOrderList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OrderStatusRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderServiceServer).GetStatusOrderList(m, &orderServiceGetStatusOrderListServer{stream})
}

type OrderService_GetStatusOrderListServer interface {
	Send(*OrderInfoList) error
	grpc.ServerStream
}

type orderServiceGetStatusOrderListServer struct {
	grpc.ServerStream
}

func (x *orderServiceGetStatusOrderListServer) Send(m *OrderInfoList) error {
	return x.ServerStream.SendMsg(m)
}

var _OrderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrder",
			Handler:    _OrderService_GetOrder_Handler,
		},
		{
			MethodName: "GetOrderList",
			Handler:    _OrderService_GetOrderList_Handler,
		},
		{
			MethodName: "NewOrder",
			Handler:    _OrderService_NewOrder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStatusOrderList",
			Handler:       _OrderService_GetStatusOrderList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "orderService.proto",
}
