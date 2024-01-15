// for syntax highlighting we use proto3 version

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: proto/assignment.proto

// package name for our proto file

package proto

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
	Assignment_SayHello_FullMethodName       = "/assignment.Assignment/SayHello"
	Assignment_BookTicket_FullMethodName     = "/assignment.Assignment/BookTicket"
	Assignment_ShowAllTickets_FullMethodName = "/assignment.Assignment/ShowAllTickets"
	Assignment_TicketsByType_FullMethodName  = "/assignment.Assignment/TicketsByType"
	Assignment_RemoveUser_FullMethodName     = "/assignment.Assignment/RemoveUser"
	Assignment_ModifyUser_FullMethodName     = "/assignment.Assignment/ModifyUser"
)

// AssignmentClient is the client API for Assignment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssignmentClient interface {
	// simple RPC
	SayHello(ctx context.Context, in *HelloParam, opts ...grpc.CallOption) (*StringResponse, error)
	BookTicket(ctx context.Context, in *Ticket, opts ...grpc.CallOption) (*Ticket, error)
	ShowAllTickets(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (Assignment_ShowAllTicketsClient, error)
	TicketsByType(ctx context.Context, in *SeatSectionRequest, opts ...grpc.CallOption) (Assignment_TicketsByTypeClient, error)
	RemoveUser(ctx context.Context, in *UserIDRequest, opts ...grpc.CallOption) (*StringResponse, error)
	ModifyUser(ctx context.Context, in *UserIDRequest, opts ...grpc.CallOption) (*StringResponse, error)
}

type assignmentClient struct {
	cc grpc.ClientConnInterface
}

func NewAssignmentClient(cc grpc.ClientConnInterface) AssignmentClient {
	return &assignmentClient{cc}
}

func (c *assignmentClient) SayHello(ctx context.Context, in *HelloParam, opts ...grpc.CallOption) (*StringResponse, error) {
	out := new(StringResponse)
	err := c.cc.Invoke(ctx, Assignment_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assignmentClient) BookTicket(ctx context.Context, in *Ticket, opts ...grpc.CallOption) (*Ticket, error) {
	out := new(Ticket)
	err := c.cc.Invoke(ctx, Assignment_BookTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assignmentClient) ShowAllTickets(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (Assignment_ShowAllTicketsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Assignment_ServiceDesc.Streams[0], Assignment_ShowAllTickets_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &assignmentShowAllTicketsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Assignment_ShowAllTicketsClient interface {
	Recv() (*Ticket, error)
	grpc.ClientStream
}

type assignmentShowAllTicketsClient struct {
	grpc.ClientStream
}

func (x *assignmentShowAllTicketsClient) Recv() (*Ticket, error) {
	m := new(Ticket)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *assignmentClient) TicketsByType(ctx context.Context, in *SeatSectionRequest, opts ...grpc.CallOption) (Assignment_TicketsByTypeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Assignment_ServiceDesc.Streams[1], Assignment_TicketsByType_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &assignmentTicketsByTypeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Assignment_TicketsByTypeClient interface {
	Recv() (*Ticket, error)
	grpc.ClientStream
}

type assignmentTicketsByTypeClient struct {
	grpc.ClientStream
}

func (x *assignmentTicketsByTypeClient) Recv() (*Ticket, error) {
	m := new(Ticket)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *assignmentClient) RemoveUser(ctx context.Context, in *UserIDRequest, opts ...grpc.CallOption) (*StringResponse, error) {
	out := new(StringResponse)
	err := c.cc.Invoke(ctx, Assignment_RemoveUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assignmentClient) ModifyUser(ctx context.Context, in *UserIDRequest, opts ...grpc.CallOption) (*StringResponse, error) {
	out := new(StringResponse)
	err := c.cc.Invoke(ctx, Assignment_ModifyUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssignmentServer is the server API for Assignment service.
// All implementations must embed UnimplementedAssignmentServer
// for forward compatibility
type AssignmentServer interface {
	// simple RPC
	SayHello(context.Context, *HelloParam) (*StringResponse, error)
	BookTicket(context.Context, *Ticket) (*Ticket, error)
	ShowAllTickets(*NoParam, Assignment_ShowAllTicketsServer) error
	TicketsByType(*SeatSectionRequest, Assignment_TicketsByTypeServer) error
	RemoveUser(context.Context, *UserIDRequest) (*StringResponse, error)
	ModifyUser(context.Context, *UserIDRequest) (*StringResponse, error)
	mustEmbedUnimplementedAssignmentServer()
}

// UnimplementedAssignmentServer must be embedded to have forward compatible implementations.
type UnimplementedAssignmentServer struct {
}

func (UnimplementedAssignmentServer) SayHello(context.Context, *HelloParam) (*StringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedAssignmentServer) BookTicket(context.Context, *Ticket) (*Ticket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookTicket not implemented")
}
func (UnimplementedAssignmentServer) ShowAllTickets(*NoParam, Assignment_ShowAllTicketsServer) error {
	return status.Errorf(codes.Unimplemented, "method ShowAllTickets not implemented")
}
func (UnimplementedAssignmentServer) TicketsByType(*SeatSectionRequest, Assignment_TicketsByTypeServer) error {
	return status.Errorf(codes.Unimplemented, "method TicketsByType not implemented")
}
func (UnimplementedAssignmentServer) RemoveUser(context.Context, *UserIDRequest) (*StringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedAssignmentServer) ModifyUser(context.Context, *UserIDRequest) (*StringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyUser not implemented")
}
func (UnimplementedAssignmentServer) mustEmbedUnimplementedAssignmentServer() {}

// UnsafeAssignmentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssignmentServer will
// result in compilation errors.
type UnsafeAssignmentServer interface {
	mustEmbedUnimplementedAssignmentServer()
}

func RegisterAssignmentServer(s grpc.ServiceRegistrar, srv AssignmentServer) {
	s.RegisterService(&Assignment_ServiceDesc, srv)
}

func _Assignment_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignmentServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Assignment_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignmentServer).SayHello(ctx, req.(*HelloParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Assignment_BookTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ticket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignmentServer).BookTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Assignment_BookTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignmentServer).BookTicket(ctx, req.(*Ticket))
	}
	return interceptor(ctx, in, info, handler)
}

func _Assignment_ShowAllTickets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NoParam)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AssignmentServer).ShowAllTickets(m, &assignmentShowAllTicketsServer{stream})
}

type Assignment_ShowAllTicketsServer interface {
	Send(*Ticket) error
	grpc.ServerStream
}

type assignmentShowAllTicketsServer struct {
	grpc.ServerStream
}

func (x *assignmentShowAllTicketsServer) Send(m *Ticket) error {
	return x.ServerStream.SendMsg(m)
}

func _Assignment_TicketsByType_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SeatSectionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AssignmentServer).TicketsByType(m, &assignmentTicketsByTypeServer{stream})
}

type Assignment_TicketsByTypeServer interface {
	Send(*Ticket) error
	grpc.ServerStream
}

type assignmentTicketsByTypeServer struct {
	grpc.ServerStream
}

func (x *assignmentTicketsByTypeServer) Send(m *Ticket) error {
	return x.ServerStream.SendMsg(m)
}

func _Assignment_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignmentServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Assignment_RemoveUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignmentServer).RemoveUser(ctx, req.(*UserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Assignment_ModifyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignmentServer).ModifyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Assignment_ModifyUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignmentServer).ModifyUser(ctx, req.(*UserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Assignment_ServiceDesc is the grpc.ServiceDesc for Assignment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Assignment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "assignment.Assignment",
	HandlerType: (*AssignmentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Assignment_SayHello_Handler,
		},
		{
			MethodName: "BookTicket",
			Handler:    _Assignment_BookTicket_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _Assignment_RemoveUser_Handler,
		},
		{
			MethodName: "ModifyUser",
			Handler:    _Assignment_ModifyUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ShowAllTickets",
			Handler:       _Assignment_ShowAllTickets_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "TicketsByType",
			Handler:       _Assignment_TicketsByType_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/assignment.proto",
}