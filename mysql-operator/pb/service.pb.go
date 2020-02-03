// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/service.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/service.proto
	pb/datatype.proto

It has these top-level messages:
	CrdReqResp
	CrdRecipient
	StorageRecipient
*/
package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CrdReqResp struct {
	Recipe       *CrdRecipient `protobuf:"bytes,1,opt,name=recipe" json:"recipe,omitempty"`
	StateCode    int32         `protobuf:"varint,2,opt,name=state_code,json=stateCode,proto3" json:"state_code,omitempty"`
	StateMessage string        `protobuf:"bytes,3,opt,name=state_message,json=stateMessage,proto3" json:"state_message,omitempty"`
}

func (m *CrdReqResp) Reset()                    { *m = CrdReqResp{} }
func (m *CrdReqResp) String() string            { return proto.CompactTextString(m) }
func (*CrdReqResp) ProtoMessage()               {}
func (*CrdReqResp) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{0} }

func (m *CrdReqResp) GetRecipe() *CrdRecipient {
	if m != nil {
		return m.Recipe
	}
	return nil
}

func (m *CrdReqResp) GetStateCode() int32 {
	if m != nil {
		return m.StateCode
	}
	return 0
}

func (m *CrdReqResp) GetStateMessage() string {
	if m != nil {
		return m.StateMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*CrdReqResp)(nil), "pb.CrdReqResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SimpleGRpcService service

type SimpleGRpcServiceClient interface {
	CreateCrd(ctx context.Context, in *CrdReqResp, opts ...grpc.CallOption) (*CrdReqResp, error)
	ReapCrd(ctx context.Context, in *CrdReqResp, opts ...grpc.CallOption) (*CrdReqResp, error)
}

type simpleGRpcServiceClient struct {
	cc *grpc.ClientConn
}

func NewSimpleGRpcServiceClient(cc *grpc.ClientConn) SimpleGRpcServiceClient {
	return &simpleGRpcServiceClient{cc}
}

func (c *simpleGRpcServiceClient) CreateCrd(ctx context.Context, in *CrdReqResp, opts ...grpc.CallOption) (*CrdReqResp, error) {
	out := new(CrdReqResp)
	err := grpc.Invoke(ctx, "/pb.SimpleGRpcService/CreateCrd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleGRpcServiceClient) ReapCrd(ctx context.Context, in *CrdReqResp, opts ...grpc.CallOption) (*CrdReqResp, error) {
	out := new(CrdReqResp)
	err := grpc.Invoke(ctx, "/pb.SimpleGRpcService/ReapCrd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SimpleGRpcService service

type SimpleGRpcServiceServer interface {
	CreateCrd(context.Context, *CrdReqResp) (*CrdReqResp, error)
	ReapCrd(context.Context, *CrdReqResp) (*CrdReqResp, error)
}

func RegisterSimpleGRpcServiceServer(s *grpc.Server, srv SimpleGRpcServiceServer) {
	s.RegisterService(&_SimpleGRpcService_serviceDesc, srv)
}

func _SimpleGRpcService_CreateCrd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CrdReqResp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleGRpcServiceServer).CreateCrd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SimpleGRpcService/CreateCrd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleGRpcServiceServer).CreateCrd(ctx, req.(*CrdReqResp))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleGRpcService_ReapCrd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CrdReqResp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleGRpcServiceServer).ReapCrd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SimpleGRpcService/ReapCrd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleGRpcServiceServer).ReapCrd(ctx, req.(*CrdReqResp))
	}
	return interceptor(ctx, in, info, handler)
}

var _SimpleGRpcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SimpleGRpcService",
	HandlerType: (*SimpleGRpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCrd",
			Handler:    _SimpleGRpcService_CreateCrd_Handler,
		},
		{
			MethodName: "ReapCrd",
			Handler:    _SimpleGRpcService_ReapCrd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/service.proto",
}

func init() { proto.RegisterFile("pb/service.proto", fileDescriptorService) }

var fileDescriptorService = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xb1, 0x4a, 0xc4, 0x40,
	0x10, 0x86, 0xd9, 0x88, 0x27, 0x99, 0xf3, 0x24, 0xb7, 0x82, 0x84, 0xa0, 0x10, 0x62, 0x13, 0x2c,
	0x12, 0x3c, 0x3b, 0x0b, 0x9b, 0x14, 0x56, 0x36, 0x7b, 0x0f, 0x20, 0x9b, 0xdd, 0x21, 0x2c, 0xdc,
	0x65, 0xc7, 0xdd, 0xe5, 0x40, 0x4b, 0x5f, 0xc1, 0xc6, 0xf7, 0xf2, 0x15, 0x7c, 0x10, 0x71, 0x73,
	0x88, 0x58, 0x59, 0xce, 0xff, 0xcf, 0x7c, 0xfc, 0xff, 0x40, 0x46, 0x7d, 0xeb, 0xd1, 0xed, 0x8c,
	0xc2, 0x86, 0x9c, 0x0d, 0x96, 0x27, 0xd4, 0x17, 0xe7, 0x83, 0xb5, 0xc3, 0x06, 0x5b, 0x49, 0xa6,
	0x95, 0xe3, 0x68, 0x83, 0x0c, 0xc6, 0x8e, 0x7e, 0xda, 0x28, 0x96, 0xd4, 0xb7, 0x5a, 0x06, 0x19,
	0x9e, 0x69, 0x7f, 0x54, 0xbd, 0x00, 0x74, 0x4e, 0x0b, 0x7c, 0x12, 0xe8, 0x89, 0xd7, 0x30, 0x73,
	0xa8, 0x0c, 0x61, 0xce, 0x4a, 0x56, 0xcf, 0x57, 0x59, 0x43, 0x7d, 0x13, 0x7d, 0x65, 0xc8, 0xe0,
	0x18, 0xc4, 0xde, 0xe7, 0x17, 0x00, 0x3e, 0xc8, 0x80, 0x8f, 0xca, 0x6a, 0xcc, 0x93, 0x92, 0xd5,
	0x87, 0x22, 0x8d, 0x4a, 0x67, 0x35, 0xf2, 0x4b, 0x58, 0x4c, 0xf6, 0x16, 0xbd, 0x97, 0x03, 0xe6,
	0x07, 0x25, 0xab, 0x53, 0x71, 0x1c, 0xc5, 0x87, 0x49, 0x5b, 0xbd, 0x33, 0x58, 0xae, 0xcd, 0x96,
	0x36, 0x78, 0x2f, 0x48, 0xad, 0xa7, 0x32, 0xbc, 0x83, 0xb4, 0x73, 0xf8, 0x0d, 0x72, 0x9a, 0x9f,
	0xfc, 0x04, 0x88, 0x01, 0x8b, 0x3f, 0x73, 0x75, 0xf6, 0xfa, 0xf1, 0xf9, 0x96, 0x64, 0xd5, 0x3c,
	0x36, 0xde, 0x5d, 0xb7, 0xca, 0xe9, 0x5b, 0x76, 0xc5, 0xef, 0xe0, 0x48, 0xa0, 0xa4, 0xff, 0x20,
	0x4e, 0x23, 0x62, 0xc1, 0x7f, 0x23, 0xfa, 0x59, 0xfc, 0xce, 0xcd, 0x57, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x31, 0xb3, 0xb0, 0xbb, 0x66, 0x01, 0x00, 0x00,
}