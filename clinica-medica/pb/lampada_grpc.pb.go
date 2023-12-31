// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: proto/lampada.proto

package pb

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

// LampadaClient is the client API for Lampada service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LampadaClient interface {
	LigarLampada(ctx context.Context, in *LampadaRequest, opts ...grpc.CallOption) (*Response, error)
	DesligarLampada(ctx context.Context, in *LampadaRequest, opts ...grpc.CallOption) (*Response, error)
}

type lampadaClient struct {
	cc grpc.ClientConnInterface
}

func NewLampadaClient(cc grpc.ClientConnInterface) LampadaClient {
	return &lampadaClient{cc}
}

func (c *lampadaClient) LigarLampada(ctx context.Context, in *LampadaRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Lampada/LigarLampada", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lampadaClient) DesligarLampada(ctx context.Context, in *LampadaRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Lampada/DesligarLampada", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LampadaServer is the server API for Lampada service.
// All implementations must embed UnimplementedLampadaServer
// for forward compatibility
type LampadaServer interface {
	LigarLampada(context.Context, *LampadaRequest) (*Response, error)
	DesligarLampada(context.Context, *LampadaRequest) (*Response, error)
	mustEmbedUnimplementedLampadaServer()
}

// UnimplementedLampadaServer must be embedded to have forward compatible implementations.
type UnimplementedLampadaServer struct {
}

func (UnimplementedLampadaServer) LigarLampada(context.Context, *LampadaRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LigarLampada not implemented")
}
func (UnimplementedLampadaServer) DesligarLampada(context.Context, *LampadaRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DesligarLampada not implemented")
}
func (UnimplementedLampadaServer) mustEmbedUnimplementedLampadaServer() {}

// UnsafeLampadaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LampadaServer will
// result in compilation errors.
type UnsafeLampadaServer interface {
	mustEmbedUnimplementedLampadaServer()
}

func RegisterLampadaServer(s grpc.ServiceRegistrar, srv LampadaServer) {
	s.RegisterService(&Lampada_ServiceDesc, srv)
}

func _Lampada_LigarLampada_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LampadaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LampadaServer).LigarLampada(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Lampada/LigarLampada",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LampadaServer).LigarLampada(ctx, req.(*LampadaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lampada_DesligarLampada_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LampadaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LampadaServer).DesligarLampada(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Lampada/DesligarLampada",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LampadaServer).DesligarLampada(ctx, req.(*LampadaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Lampada_ServiceDesc is the grpc.ServiceDesc for Lampada service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Lampada_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Lampada",
	HandlerType: (*LampadaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LigarLampada",
			Handler:    _Lampada_LigarLampada_Handler,
		},
		{
			MethodName: "DesligarLampada",
			Handler:    _Lampada_DesligarLampada_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/lampada.proto",
}

// ArCondicionadoClient is the client API for ArCondicionado service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArCondicionadoClient interface {
	LigarArCondicionado(ctx context.Context, in *ArCondicionadoRequest, opts ...grpc.CallOption) (*Response, error)
	DesligarArCondicionado(ctx context.Context, in *ArCondicionadoRequest, opts ...grpc.CallOption) (*Response, error)
}

type arCondicionadoClient struct {
	cc grpc.ClientConnInterface
}

func NewArCondicionadoClient(cc grpc.ClientConnInterface) ArCondicionadoClient {
	return &arCondicionadoClient{cc}
}

func (c *arCondicionadoClient) LigarArCondicionado(ctx context.Context, in *ArCondicionadoRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ArCondicionado/LigarArCondicionado", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arCondicionadoClient) DesligarArCondicionado(ctx context.Context, in *ArCondicionadoRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ArCondicionado/DesligarArCondicionado", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArCondicionadoServer is the server API for ArCondicionado service.
// All implementations must embed UnimplementedArCondicionadoServer
// for forward compatibility
type ArCondicionadoServer interface {
	LigarArCondicionado(context.Context, *ArCondicionadoRequest) (*Response, error)
	DesligarArCondicionado(context.Context, *ArCondicionadoRequest) (*Response, error)
	mustEmbedUnimplementedArCondicionadoServer()
}

// UnimplementedArCondicionadoServer must be embedded to have forward compatible implementations.
type UnimplementedArCondicionadoServer struct {
}

func (UnimplementedArCondicionadoServer) LigarArCondicionado(context.Context, *ArCondicionadoRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LigarArCondicionado not implemented")
}
func (UnimplementedArCondicionadoServer) DesligarArCondicionado(context.Context, *ArCondicionadoRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DesligarArCondicionado not implemented")
}
func (UnimplementedArCondicionadoServer) mustEmbedUnimplementedArCondicionadoServer() {}

// UnsafeArCondicionadoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArCondicionadoServer will
// result in compilation errors.
type UnsafeArCondicionadoServer interface {
	mustEmbedUnimplementedArCondicionadoServer()
}

func RegisterArCondicionadoServer(s grpc.ServiceRegistrar, srv ArCondicionadoServer) {
	s.RegisterService(&ArCondicionado_ServiceDesc, srv)
}

func _ArCondicionado_LigarArCondicionado_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArCondicionadoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArCondicionadoServer).LigarArCondicionado(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ArCondicionado/LigarArCondicionado",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArCondicionadoServer).LigarArCondicionado(ctx, req.(*ArCondicionadoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArCondicionado_DesligarArCondicionado_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArCondicionadoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArCondicionadoServer).DesligarArCondicionado(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ArCondicionado/DesligarArCondicionado",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArCondicionadoServer).DesligarArCondicionado(ctx, req.(*ArCondicionadoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ArCondicionado_ServiceDesc is the grpc.ServiceDesc for ArCondicionado service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArCondicionado_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ArCondicionado",
	HandlerType: (*ArCondicionadoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LigarArCondicionado",
			Handler:    _ArCondicionado_LigarArCondicionado_Handler,
		},
		{
			MethodName: "DesligarArCondicionado",
			Handler:    _ArCondicionado_DesligarArCondicionado_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/lampada.proto",
}

// ControleDeIncendioClient is the client API for ControleDeIncendio service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ControleDeIncendioClient interface {
	LigarControleDeIncendio(ctx context.Context, in *ControleDeIncendioRequest, opts ...grpc.CallOption) (*Response, error)
	DesligarControleDeIncendio(ctx context.Context, in *ControleDeIncendioRequest, opts ...grpc.CallOption) (*Response, error)
}

type controleDeIncendioClient struct {
	cc grpc.ClientConnInterface
}

func NewControleDeIncendioClient(cc grpc.ClientConnInterface) ControleDeIncendioClient {
	return &controleDeIncendioClient{cc}
}

func (c *controleDeIncendioClient) LigarControleDeIncendio(ctx context.Context, in *ControleDeIncendioRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ControleDeIncendio/LigarControleDeIncendio", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controleDeIncendioClient) DesligarControleDeIncendio(ctx context.Context, in *ControleDeIncendioRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ControleDeIncendio/DesligarControleDeIncendio", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControleDeIncendioServer is the server API for ControleDeIncendio service.
// All implementations must embed UnimplementedControleDeIncendioServer
// for forward compatibility
type ControleDeIncendioServer interface {
	LigarControleDeIncendio(context.Context, *ControleDeIncendioRequest) (*Response, error)
	DesligarControleDeIncendio(context.Context, *ControleDeIncendioRequest) (*Response, error)
	mustEmbedUnimplementedControleDeIncendioServer()
}

// UnimplementedControleDeIncendioServer must be embedded to have forward compatible implementations.
type UnimplementedControleDeIncendioServer struct {
}

func (UnimplementedControleDeIncendioServer) LigarControleDeIncendio(context.Context, *ControleDeIncendioRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LigarControleDeIncendio not implemented")
}
func (UnimplementedControleDeIncendioServer) DesligarControleDeIncendio(context.Context, *ControleDeIncendioRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DesligarControleDeIncendio not implemented")
}
func (UnimplementedControleDeIncendioServer) mustEmbedUnimplementedControleDeIncendioServer() {}

// UnsafeControleDeIncendioServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ControleDeIncendioServer will
// result in compilation errors.
type UnsafeControleDeIncendioServer interface {
	mustEmbedUnimplementedControleDeIncendioServer()
}

func RegisterControleDeIncendioServer(s grpc.ServiceRegistrar, srv ControleDeIncendioServer) {
	s.RegisterService(&ControleDeIncendio_ServiceDesc, srv)
}

func _ControleDeIncendio_LigarControleDeIncendio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControleDeIncendioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControleDeIncendioServer).LigarControleDeIncendio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ControleDeIncendio/LigarControleDeIncendio",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControleDeIncendioServer).LigarControleDeIncendio(ctx, req.(*ControleDeIncendioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControleDeIncendio_DesligarControleDeIncendio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControleDeIncendioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControleDeIncendioServer).DesligarControleDeIncendio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ControleDeIncendio/DesligarControleDeIncendio",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControleDeIncendioServer).DesligarControleDeIncendio(ctx, req.(*ControleDeIncendioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ControleDeIncendio_ServiceDesc is the grpc.ServiceDesc for ControleDeIncendio service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ControleDeIncendio_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ControleDeIncendio",
	HandlerType: (*ControleDeIncendioServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LigarControleDeIncendio",
			Handler:    _ControleDeIncendio_LigarControleDeIncendio_Handler,
		},
		{
			MethodName: "DesligarControleDeIncendio",
			Handler:    _ControleDeIncendio_DesligarControleDeIncendio_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/lampada.proto",
}
