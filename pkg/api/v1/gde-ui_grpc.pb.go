// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// GdeUIClient is the client API for GdeUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GdeUIClient interface {
	// ListAnalysisSpecs returns a list of Genetic Analysis that can be started through the UI.
	ListAnalysisSpecs(ctx context.Context, in *ListAnalysisSpecsRequest, opts ...grpc.CallOption) (GdeUI_ListAnalysisSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type gdeUIClient struct {
	cc grpc.ClientConnInterface
}

func NewGdeUIClient(cc grpc.ClientConnInterface) GdeUIClient {
	return &gdeUIClient{cc}
}

func (c *gdeUIClient) ListAnalysisSpecs(ctx context.Context, in *ListAnalysisSpecsRequest, opts ...grpc.CallOption) (GdeUI_ListAnalysisSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &GdeUI_ServiceDesc.Streams[0], "/v1.GdeUI/ListAnalysisSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &gdeUIListAnalysisSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GdeUI_ListAnalysisSpecsClient interface {
	Recv() (*ListAnalysisSpecsResponse, error)
	grpc.ClientStream
}

type gdeUIListAnalysisSpecsClient struct {
	grpc.ClientStream
}

func (x *gdeUIListAnalysisSpecsClient) Recv() (*ListAnalysisSpecsResponse, error) {
	m := new(ListAnalysisSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gdeUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.GdeUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GdeUIServer is the server API for GdeUI service.
// All implementations must embed UnimplementedGdeUIServer
// for forward compatibility
type GdeUIServer interface {
	// ListAnalysisSpecs returns a list of Genetic Analysis that can be started through the UI.
	ListAnalysisSpecs(*ListAnalysisSpecsRequest, GdeUI_ListAnalysisSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedGdeUIServer()
}

// UnimplementedGdeUIServer must be embedded to have forward compatible implementations.
type UnimplementedGdeUIServer struct {
}

func (UnimplementedGdeUIServer) ListAnalysisSpecs(*ListAnalysisSpecsRequest, GdeUI_ListAnalysisSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAnalysisSpecs not implemented")
}
func (UnimplementedGdeUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedGdeUIServer) mustEmbedUnimplementedGdeUIServer() {}

// UnsafeGdeUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GdeUIServer will
// result in compilation errors.
type UnsafeGdeUIServer interface {
	mustEmbedUnimplementedGdeUIServer()
}

func RegisterGdeUIServer(s grpc.ServiceRegistrar, srv GdeUIServer) {
	s.RegisterService(&GdeUI_ServiceDesc, srv)
}

func _GdeUI_ListAnalysisSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListAnalysisSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GdeUIServer).ListAnalysisSpecs(m, &gdeUIListAnalysisSpecsServer{stream})
}

type GdeUI_ListAnalysisSpecsServer interface {
	Send(*ListAnalysisSpecsResponse) error
	grpc.ServerStream
}

type gdeUIListAnalysisSpecsServer struct {
	grpc.ServerStream
}

func (x *gdeUIListAnalysisSpecsServer) Send(m *ListAnalysisSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GdeUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GdeUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.GdeUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GdeUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GdeUI_ServiceDesc is the grpc.ServiceDesc for GdeUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GdeUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.GdeUI",
	HandlerType: (*GdeUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _GdeUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListAnalysisSpecs",
			Handler:       _GdeUI_ListAnalysisSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gde-ui.proto",
}
