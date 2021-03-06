// Code generated by protoc-gen-go.
// source: log_it.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	log_it.proto

It has these top-level messages:
	LogData
	LogSuccess
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type LogData struct {
	Severity int32  `protobuf:"varint,1,opt,name=severity" json:"severity,omitempty"`
	Data     string `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *LogData) Reset()                    { *m = LogData{} }
func (m *LogData) String() string            { return proto1.CompactTextString(m) }
func (*LogData) ProtoMessage()               {}
func (*LogData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LogData) GetSeverity() int32 {
	if m != nil {
		return m.Severity
	}
	return 0
}

func (m *LogData) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type LogSuccess struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *LogSuccess) Reset()                    { *m = LogSuccess{} }
func (m *LogSuccess) String() string            { return proto1.CompactTextString(m) }
func (*LogSuccess) ProtoMessage()               {}
func (*LogSuccess) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LogSuccess) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *LogSuccess) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto1.RegisterType((*LogData)(nil), "proto.LogData")
	proto1.RegisterType((*LogSuccess)(nil), "proto.LogSuccess")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LogIt service

type LogItClient interface {
	// Log message, and swap log file
	SwapLogFile(ctx context.Context, in *LogData, opts ...grpc.CallOption) (*LogSuccess, error)
	// Repond to message - no loggin -just responde with success
	IAmAlive(ctx context.Context, in *LogData, opts ...grpc.CallOption) (*LogSuccess, error)
	// just log some data
	LogMessage(ctx context.Context, in *LogData, opts ...grpc.CallOption) (*LogSuccess, error)
}

type logItClient struct {
	cc *grpc.ClientConn
}

func NewLogItClient(cc *grpc.ClientConn) LogItClient {
	return &logItClient{cc}
}

func (c *logItClient) SwapLogFile(ctx context.Context, in *LogData, opts ...grpc.CallOption) (*LogSuccess, error) {
	out := new(LogSuccess)
	err := grpc.Invoke(ctx, "/proto.LogIt/SwapLogFile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logItClient) IAmAlive(ctx context.Context, in *LogData, opts ...grpc.CallOption) (*LogSuccess, error) {
	out := new(LogSuccess)
	err := grpc.Invoke(ctx, "/proto.LogIt/IAmAlive", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logItClient) LogMessage(ctx context.Context, in *LogData, opts ...grpc.CallOption) (*LogSuccess, error) {
	out := new(LogSuccess)
	err := grpc.Invoke(ctx, "/proto.LogIt/LogMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LogIt service

type LogItServer interface {
	// Log message, and swap log file
	SwapLogFile(context.Context, *LogData) (*LogSuccess, error)
	// Repond to message - no loggin -just responde with success
	IAmAlive(context.Context, *LogData) (*LogSuccess, error)
	// just log some data
	LogMessage(context.Context, *LogData) (*LogSuccess, error)
}

func RegisterLogItServer(s *grpc.Server, srv LogItServer) {
	s.RegisterService(&_LogIt_serviceDesc, srv)
}

func _LogIt_SwapLogFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogItServer).SwapLogFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LogIt/SwapLogFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogItServer).SwapLogFile(ctx, req.(*LogData))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogIt_IAmAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogItServer).IAmAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LogIt/IAmAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogItServer).IAmAlive(ctx, req.(*LogData))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogIt_LogMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogItServer).LogMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LogIt/LogMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogItServer).LogMessage(ctx, req.(*LogData))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogIt_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LogIt",
	HandlerType: (*LogItServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SwapLogFile",
			Handler:    _LogIt_SwapLogFile_Handler,
		},
		{
			MethodName: "IAmAlive",
			Handler:    _LogIt_IAmAlive_Handler,
		},
		{
			MethodName: "LogMessage",
			Handler:    _LogIt_LogMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "log_it.proto",
}

func init() { proto1.RegisterFile("log_it.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x8e, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x5d, 0x35, 0xb5, 0x1d, 0x45, 0x74, 0x0e, 0xa5, 0xf4, 0x54, 0x72, 0x90, 0x9e, 0x56,
	0xac, 0x20, 0x78, 0x6c, 0x11, 0xa1, 0x10, 0xa1, 0xa4, 0x0f, 0x20, 0x6b, 0x5c, 0x86, 0x85, 0x0d,
	0x13, 0x32, 0xd3, 0xaa, 0x4f, 0xe3, 0xab, 0x4a, 0xd7, 0xa0, 0xd7, 0x9c, 0xe6, 0xff, 0x0e, 0x1f,
	0xf3, 0xc1, 0x45, 0x64, 0x7a, 0x0d, 0x6a, 0x9b, 0x96, 0x95, 0x31, 0x4b, 0x27, 0x7f, 0x84, 0xb3,
	0x82, 0xe9, 0xc9, 0xa9, 0xc3, 0x29, 0x0c, 0xc5, 0xef, 0x7d, 0x1b, 0xf4, 0x6b, 0x62, 0x66, 0x66,
	0x9e, 0x95, 0x7f, 0x8c, 0x08, 0xa7, 0xef, 0x4e, 0xdd, 0xe4, 0x78, 0x66, 0xe6, 0xa3, 0x32, 0xed,
	0xfc, 0x01, 0xa0, 0x60, 0xda, 0xee, 0xaa, 0xca, 0x8b, 0xe0, 0x18, 0x06, 0xa2, 0x4e, 0x77, 0x92,
	0xdc, 0x51, 0xd9, 0x11, 0x5e, 0xc1, 0x49, 0x2d, 0xd4, 0x89, 0x87, 0xb9, 0xf8, 0x36, 0x90, 0x15,
	0x4c, 0x6b, 0xc5, 0x05, 0x9c, 0x6f, 0x3f, 0x5c, 0x53, 0x30, 0x3d, 0x87, 0xe8, 0xf1, 0xf2, 0x37,
	0xcd, 0x76, 0x41, 0xd3, 0xeb, 0x7f, 0xee, 0xbe, 0xe4, 0x47, 0x78, 0x0b, 0xc3, 0xf5, 0xb2, 0x5e,
	0xc6, 0xb0, 0xef, 0x29, 0xdc, 0xa5, 0xcc, 0x17, 0x2f, 0xe2, 0xa8, 0x9f, 0xb2, 0xba, 0x81, 0x71,
	0x60, 0x4b, 0x6d, 0x53, 0x59, 0xff, 0xe9, 0xea, 0x26, 0x7a, 0xb1, 0x91, 0x29, 0xe8, 0x0a, 0x52,
	0xf8, 0xe6, 0xa0, 0x6c, 0xcc, 0xdb, 0x20, 0xb9, 0xf7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfb,
	0x2e, 0xb0, 0xb6, 0x5a, 0x01, 0x00, 0x00,
}
