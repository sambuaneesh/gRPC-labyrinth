// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: labyrinth.proto

package labyrinth

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LabyrinthGame_GetLabyrinthInfo_FullMethodName = "/LabyrinthGame/GetLabyrinthInfo"
	LabyrinthGame_GetPlayerStatus_FullMethodName  = "/LabyrinthGame/GetPlayerStatus"
	LabyrinthGame_RegisterMove_FullMethodName     = "/LabyrinthGame/RegisterMove"
	LabyrinthGame_Revelio_FullMethodName          = "/LabyrinthGame/Revelio"
	LabyrinthGame_Bombarda_FullMethodName         = "/LabyrinthGame/Bombarda"
)

// LabyrinthGameClient is the client API for LabyrinthGame service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// -------------------------------------------------------------------------------------------------
// # Services
type LabyrinthGameClient interface {
	GetLabyrinthInfo(ctx context.Context, in *GetLabyrinthInfoRequest, opts ...grpc.CallOption) (*GetLabyrinthInfoResponse, error)
	GetPlayerStatus(ctx context.Context, in *GetPlayerStatusRequest, opts ...grpc.CallOption) (*GetPlayerStatusResponse, error)
	RegisterMove(ctx context.Context, in *RegisterMoveRequest, opts ...grpc.CallOption) (*RegisterMoveResponse, error)
	Revelio(ctx context.Context, in *RevelioRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[RevelioResponse], error)
	Bombarda(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[BombardaRequest, BombardaResponse], error)
}

type labyrinthGameClient struct {
	cc grpc.ClientConnInterface
}

func NewLabyrinthGameClient(cc grpc.ClientConnInterface) LabyrinthGameClient {
	return &labyrinthGameClient{cc}
}

func (c *labyrinthGameClient) GetLabyrinthInfo(ctx context.Context, in *GetLabyrinthInfoRequest, opts ...grpc.CallOption) (*GetLabyrinthInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLabyrinthInfoResponse)
	err := c.cc.Invoke(ctx, LabyrinthGame_GetLabyrinthInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labyrinthGameClient) GetPlayerStatus(ctx context.Context, in *GetPlayerStatusRequest, opts ...grpc.CallOption) (*GetPlayerStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPlayerStatusResponse)
	err := c.cc.Invoke(ctx, LabyrinthGame_GetPlayerStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labyrinthGameClient) RegisterMove(ctx context.Context, in *RegisterMoveRequest, opts ...grpc.CallOption) (*RegisterMoveResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterMoveResponse)
	err := c.cc.Invoke(ctx, LabyrinthGame_RegisterMove_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labyrinthGameClient) Revelio(ctx context.Context, in *RevelioRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[RevelioResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &LabyrinthGame_ServiceDesc.Streams[0], LabyrinthGame_Revelio_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[RevelioRequest, RevelioResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type LabyrinthGame_RevelioClient = grpc.ServerStreamingClient[RevelioResponse]

func (c *labyrinthGameClient) Bombarda(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[BombardaRequest, BombardaResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &LabyrinthGame_ServiceDesc.Streams[1], LabyrinthGame_Bombarda_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[BombardaRequest, BombardaResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type LabyrinthGame_BombardaClient = grpc.BidiStreamingClient[BombardaRequest, BombardaResponse]

// LabyrinthGameServer is the server API for LabyrinthGame service.
// All implementations must embed UnimplementedLabyrinthGameServer
// for forward compatibility.
//
// -------------------------------------------------------------------------------------------------
// # Services
type LabyrinthGameServer interface {
	GetLabyrinthInfo(context.Context, *GetLabyrinthInfoRequest) (*GetLabyrinthInfoResponse, error)
	GetPlayerStatus(context.Context, *GetPlayerStatusRequest) (*GetPlayerStatusResponse, error)
	RegisterMove(context.Context, *RegisterMoveRequest) (*RegisterMoveResponse, error)
	Revelio(*RevelioRequest, grpc.ServerStreamingServer[RevelioResponse]) error
	Bombarda(grpc.BidiStreamingServer[BombardaRequest, BombardaResponse]) error
	mustEmbedUnimplementedLabyrinthGameServer()
}

// UnimplementedLabyrinthGameServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLabyrinthGameServer struct{}

func (UnimplementedLabyrinthGameServer) GetLabyrinthInfo(context.Context, *GetLabyrinthInfoRequest) (*GetLabyrinthInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLabyrinthInfo not implemented")
}
func (UnimplementedLabyrinthGameServer) GetPlayerStatus(context.Context, *GetPlayerStatusRequest) (*GetPlayerStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerStatus not implemented")
}
func (UnimplementedLabyrinthGameServer) RegisterMove(context.Context, *RegisterMoveRequest) (*RegisterMoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterMove not implemented")
}
func (UnimplementedLabyrinthGameServer) Revelio(*RevelioRequest, grpc.ServerStreamingServer[RevelioResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Revelio not implemented")
}
func (UnimplementedLabyrinthGameServer) Bombarda(grpc.BidiStreamingServer[BombardaRequest, BombardaResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Bombarda not implemented")
}
func (UnimplementedLabyrinthGameServer) mustEmbedUnimplementedLabyrinthGameServer() {}
func (UnimplementedLabyrinthGameServer) testEmbeddedByValue()                       {}

// UnsafeLabyrinthGameServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LabyrinthGameServer will
// result in compilation errors.
type UnsafeLabyrinthGameServer interface {
	mustEmbedUnimplementedLabyrinthGameServer()
}

func RegisterLabyrinthGameServer(s grpc.ServiceRegistrar, srv LabyrinthGameServer) {
	// If the following call pancis, it indicates UnimplementedLabyrinthGameServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LabyrinthGame_ServiceDesc, srv)
}

func _LabyrinthGame_GetLabyrinthInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabyrinthInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabyrinthGameServer).GetLabyrinthInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabyrinthGame_GetLabyrinthInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabyrinthGameServer).GetLabyrinthInfo(ctx, req.(*GetLabyrinthInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabyrinthGame_GetPlayerStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabyrinthGameServer).GetPlayerStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabyrinthGame_GetPlayerStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabyrinthGameServer).GetPlayerStatus(ctx, req.(*GetPlayerStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabyrinthGame_RegisterMove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterMoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabyrinthGameServer).RegisterMove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabyrinthGame_RegisterMove_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabyrinthGameServer).RegisterMove(ctx, req.(*RegisterMoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabyrinthGame_Revelio_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RevelioRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LabyrinthGameServer).Revelio(m, &grpc.GenericServerStream[RevelioRequest, RevelioResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type LabyrinthGame_RevelioServer = grpc.ServerStreamingServer[RevelioResponse]

func _LabyrinthGame_Bombarda_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LabyrinthGameServer).Bombarda(&grpc.GenericServerStream[BombardaRequest, BombardaResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type LabyrinthGame_BombardaServer = grpc.BidiStreamingServer[BombardaRequest, BombardaResponse]

// LabyrinthGame_ServiceDesc is the grpc.ServiceDesc for LabyrinthGame service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LabyrinthGame_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "LabyrinthGame",
	HandlerType: (*LabyrinthGameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLabyrinthInfo",
			Handler:    _LabyrinthGame_GetLabyrinthInfo_Handler,
		},
		{
			MethodName: "GetPlayerStatus",
			Handler:    _LabyrinthGame_GetPlayerStatus_Handler,
		},
		{
			MethodName: "RegisterMove",
			Handler:    _LabyrinthGame_RegisterMove_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Revelio",
			Handler:       _LabyrinthGame_Revelio_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Bombarda",
			Handler:       _LabyrinthGame_Bombarda_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "labyrinth.proto",
}
